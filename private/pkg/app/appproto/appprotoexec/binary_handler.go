// Copyright 2020-2023 Buf Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package appprotoexec

import (
	"bytes"
	"context"
	"io"
	"path/filepath"

	"github.com/bufbuild/buf/private/pkg/app"
	"github.com/bufbuild/buf/private/pkg/app/appproto"
	"github.com/bufbuild/buf/private/pkg/command"
	"github.com/bufbuild/buf/private/pkg/ioextended"
	"github.com/bufbuild/buf/private/pkg/protoencoding"
	"go.opencensus.io/trace"
	"google.golang.org/protobuf/types/pluginpb"
)

type binaryHandler struct {
	runner     command.Runner
	pluginPath string
}

func newBinaryHandler(
	runner command.Runner,
	pluginPath string,
) *binaryHandler {
	return &binaryHandler{
		runner:     runner,
		pluginPath: pluginPath,
	}
}

func (h *binaryHandler) Handle(
	ctx context.Context,
	container app.EnvStderrContainer,
	responseWriter appproto.ResponseBuilder,
	request *pluginpb.CodeGeneratorRequest,
) error {
	ctx, span := trace.StartSpan(ctx, "plugin_proxy")
	span.AddAttributes(trace.StringAttribute("plugin", filepath.Base(h.pluginPath)))
	defer span.End()
	requestData, err := protoencoding.NewWireMarshaler().Marshal(request)
	if err != nil {
		return err
	}
	responseBuffer := bytes.NewBuffer(nil)
	stderrWriteCloser := newStderrWriteCloser(container.Stderr(), h.pluginPath)
	if err := h.runner.Run(
		ctx,
		h.pluginPath,
		command.RunWithEnv(app.EnvironMap(container)),
		command.RunWithStdin(bytes.NewReader(requestData)),
		command.RunWithStdout(responseBuffer),
		command.RunWithStderr(stderrWriteCloser),
	); err != nil {
		return err
	}
	response := &pluginpb.CodeGeneratorResponse{}
	if err := protoencoding.NewWireUnmarshaler(nil).Unmarshal(responseBuffer.Bytes(), response); err != nil {
		return err
	}
	response, err = normalizeCodeGeneratorResponse(response)
	if err != nil {
		return err
	}
	if response.GetSupportedFeatures()&uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL) != 0 {
		responseWriter.SetFeatureProto3Optional()
	}
	for _, file := range response.File {
		if err := responseWriter.AddFile(file); err != nil {
			return err
		}
	}
	// plugin.proto specifies that only non-empty errors are considered errors.
	// This is also consistent with protoc's behavior.
	// Ref: https://github.com/protocolbuffers/protobuf/blob/069f989b483e63005f87ab309de130677718bbec/src/google/protobuf/compiler/plugin.proto#L100-L108.
	if response.GetError() != "" {
		responseWriter.AddError(response.GetError())
	}
	return nil
}

func newStderrWriteCloser(delegate io.Writer, pluginPath string) io.WriteCloser {
	switch filepath.Base(pluginPath) {
	case "protoc-gen-swift":
		// https://github.com/bufbuild/buf/issues/1736
		// Swallowing specific stderr message for protoc-gen-swift as protoc-gen-swift, see issue.
		// This is all disgusting code but it's simple and it works.
		// We did not document if pluginPath is normalized or not, so
		return newProtocGenSwiftStderrWriteCloser(delegate)
	default:
		return ioextended.NopWriteCloser(delegate)
	}
}

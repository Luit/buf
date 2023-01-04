// Copyright 2020-2022 Buf Technologies, Inc.
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

package generator

import (
	"bytes"
	"fmt"
	"html"
	"io"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strings"

	"github.com/spf13/cobra"
)

func printOptions(buf *bytes.Buffer, cmd *cobra.Command, name string) error {
	flags := cmd.NonInheritedFlags()
	flags.SetOutput(buf)
	if flags.HasAvailableFlags() {
		buf.WriteString("### Flags\n\n```\n")
		flags.PrintDefaults()
		buf.WriteString("```\n\n")
	}

	parentFlags := cmd.InheritedFlags()
	parentFlags.SetOutput(buf)
	if parentFlags.HasAvailableFlags() {
		buf.WriteString("### Flags inherited from parent commands\n\n```\n")
		parentFlags.PrintDefaults()
		buf.WriteString("```\n\n")
	}
	return nil
}

// GenMarkdownCustom creates custom markdown output.
func GenMarkdownCustom(cmd *cobra.Command, w io.Writer, childLinkHandler, parentLinkHandler func(*cobra.Command) string) error {
	cmd.InitDefaultHelpCmd()
	cmd.InitDefaultHelpFlag()
	buf := new(bytes.Buffer)
	name := cmd.CommandPath()
	if cmd.Version != "" {
		buf.WriteString("version " + cmd.Version + "\n\n")
	}
	buf.WriteString(cmd.Short + "\n\n")
	if len(cmd.Long) > 0 {
		buf.WriteString("### Synopsis\n\n")
		buf.WriteString(html.EscapeString(cmd.Long) + "\n\n")
	}
	if cmd.Runnable() {
		buf.WriteString(fmt.Sprintf("```\n%s\n```\n\n", cmd.UseLine()))
	}
	if len(cmd.Example) > 0 {
		buf.WriteString("### Examples\n\n")
		buf.WriteString(fmt.Sprintf("```\n%s\n```\n\n", html.EscapeString(cmd.Example)))
	}
	if err := printOptions(buf, cmd, name); err != nil {
		return err
	}
	if hasSubCommands(cmd) {
		buf.WriteString("### Subcommands\n\n")
		children := cmd.Commands()
		sort.Slice(children, func(i, j int) bool {
			return children[i].Name() < children[j].Name()
		})

		for _, child := range children {
			if !child.IsAvailableCommand() || child.IsAdditionalHelpTopicCommand() {
				continue
			}
			cname := name + " " + child.Name()
			link := cname + ".md"
			link = strings.ReplaceAll(link, " ", "-")

			childLink := childLinkHandler(child)

			buf.WriteString(fmt.Sprintf("* [%s](%s)\t - %s\n", cname, childLink, child.Short))
		}
		buf.WriteString("\n")
	}

	if cmd.HasParent() {
		buf.WriteString("### Parent Command\n\n")
		if cmd.HasParent() {
			parent := cmd.Parent()
			pname := parent.CommandPath()
			link := pname + ".md"
			link = strings.ReplaceAll(link, " ", "-")

			parentLink := parentLinkHandler(parent)

			buf.WriteString(fmt.Sprintf("* [%s](%s)\t - %s\n", pname, parentLink, parent.Short))
			cmd.VisitParents(func(c *cobra.Command) {
				if c.DisableAutoGenTag {
					cmd.DisableAutoGenTag = c.DisableAutoGenTag
				}
			})
		}
	}
	_, err := buf.WriteTo(w)
	return err
}

// GenMarkdownTree will generate a markdown page for this command and all
// descendants in the directory given. The header may be nil.
// This function may not work correctly if your command names have `-` in them.
// If you have `cmd` with two subcmds, `sub` and `sub-third`,
// and `sub` has a subcommand called `third`, it is undefined which
// help output will be in the file `cmd-sub-third.1`.
func GenMarkdownTree(cmd *cobra.Command, dir string) error {
	childLinkHandler := func(cmd *cobra.Command) string {
		return cmd.Parent().Name() + "/" + cmd.Name() + ".md"
	}
	parentLinkHandler := func(cmd *cobra.Command) string {
		return "../" + cmd.Name() + ".md"
	}
	var count int
	frontMatter := func(s string) string {
		count--
		s = strings.TrimSuffix(path.Base(s), ".md")
		return fmt.Sprintf(`---
id: %s
title: %s
sidebar_position: %d
---
`, s, strings.ReplaceAll(s, "-", " "), count)
	}
	pathGenerator := func(cmd *cobra.Command) string {
		cmdPath := cmd.CommandPath()
		cmdAll := strings.ReplaceAll(cmdPath, " ", "/")
		if cmd.HasSubCommands() {
			return path.Join(cmdAll, "index.md")
		}
		return cmdAll + ".md"
	}
	return GenMarkdownTreeCustom(cmd, dir, frontMatter, pathGenerator, childLinkHandler, parentLinkHandler)
}

// GenMarkdownTreeCustom is the the same as GenMarkdownTree, but
// with custom filePrepender and linkHandler.
func GenMarkdownTreeCustom(cmd *cobra.Command, dir string, filePrepender func(string) string, pathGenerator, childLinkHandler, parentLinkHandler func(*cobra.Command) string) error {
	for _, c := range cmd.Commands() {
		if !c.IsAvailableCommand() || c.IsAdditionalHelpTopicCommand() {
			continue
		}
		if err := GenMarkdownTreeCustom(c, dir, filePrepender, pathGenerator, childLinkHandler, parentLinkHandler); err != nil {
			return err
		}
	}

	basename := pathGenerator(cmd)
	filename := filepath.Join(dir, basename)
	if err := os.MkdirAll(path.Dir(filename), os.ModePerm); err != nil {
		return err
	}
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err := io.WriteString(f, filePrepender(filename)); err != nil {
		return err
	}
	if err := GenMarkdownCustom(cmd, f, childLinkHandler, parentLinkHandler); err != nil {
		return err
	}
	return nil
}

func hasSubCommands(cmd *cobra.Command) bool {
	for _, c := range cmd.Commands() {
		if !c.IsAvailableCommand() || c.IsAdditionalHelpTopicCommand() {
			continue
		}
		return true
	}
	return false
}

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
	"strings"

	"github.com/spf13/cobra"
)

func GenMarkdownTree(cmd *cobra.Command, dir string, num int) error {
	childNum := 0
	for _, c := range cmd.Commands() {
		childNum += 1
		if err := GenMarkdownTree(c, dir, childNum); err != nil {
			return err
		}
	}
	if !cmd.IsAvailableCommand() || cmd.IsAdditionalHelpTopicCommand() {
		return nil
	}
	num--
	var basename string
	cmdAll := strings.ReplaceAll(cmd.CommandPath(), " ", "/")
	if cmd.HasSubCommands() {
		basename = path.Join(cmdAll, "index.md")
	} else {
		basename = cmdAll + ".md"
	}
	filename := filepath.Join(dir, basename)
	if err := os.MkdirAll(path.Dir(filename), os.ModePerm); err != nil {
		return err
	}
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	frontMatter := fmt.Sprintf(`---
id: %s
title: %s
sidebar_position: %d
---
`, strings.TrimSuffix(path.Base(filename), ".md"), cmd.Name(), num)
	if _, err := io.WriteString(f, frontMatter); err != nil {
		return err
	}
	if err := GenMarkdownPage(cmd, f); err != nil {
		return err
	}
	return nil
}

// GenMarkdownPage creates custom markdown output.
func GenMarkdownPage(cmd *cobra.Command, w io.Writer) error {
	cmd.InitDefaultHelpCmd()
	cmd.InitDefaultHelpFlag()
	name := cmd.CommandPath()
	buf := new(bytes.Buffer)
	// Name + Version
	if cmd.Version != "" {
		buf.WriteString("version " + cmd.Version + "\n\n")
	}
	buf.WriteString(cmd.Short + "\n\n")
	// Synopsis
	if len(cmd.Long) > 0 {
		buf.WriteString("### Synopsis\n\n")
		buf.WriteString(html.EscapeString(cmd.Long) + "\n\n")
	}
	if cmd.Runnable() {
		buf.WriteString(fmt.Sprintf("```\n%s\n```\n\n", cmd.UseLine()))
	}
	// Examples
	if len(cmd.Example) > 0 {
		buf.WriteString("### Examples\n\n")
		buf.WriteString(fmt.Sprintf("```\n%s\n```\n\n", html.EscapeString(cmd.Example)))
	}
	// Flags
	flags := cmd.NonInheritedFlags()
	flags.SetOutput(buf)
	if flags.HasAvailableFlags() {
		buf.WriteString("### Flags\n\n```\n")
		flags.PrintDefaults()
		buf.WriteString("```\n\n")
	}
	// Parent Flags
	parentFlags := cmd.InheritedFlags()
	parentFlags.SetOutput(buf)
	if parentFlags.HasAvailableFlags() {
		buf.WriteString("### Flags inherited from parent commands\n\n```\n")
		parentFlags.PrintDefaults()
		buf.WriteString("```\n\n")
	}
	// Subcommands
	if hasSubCommands(cmd) {
		buf.WriteString("### Subcommands\n\n")
		children := cmd.Commands()
		//sort.Slice(children, func(i, j int) bool {
		//	return children[i].Name() < children[j].Name()
		//})
		for _, child := range children {
			if !child.IsAvailableCommand() || child.IsAdditionalHelpTopicCommand() {
				continue
			}
			commandName := name + " " + child.Name()
			link := commandName + ".md"
			link = strings.ReplaceAll(link, " ", "-")
			childLink := child.Name() + ".md"
			if child.HasSubCommands() {
				childLink = child.Name() + "/index.md"
			}
			buf.WriteString(fmt.Sprintf("* [%s](%s)\t - %s\n", commandName, childLink, child.Short))
		}
		buf.WriteString("\n")
	}
	// Parent Command
	if cmd.HasParent() {
		buf.WriteString("### Parent Command\n\n")
		if cmd.HasParent() {
			parent := cmd.Parent()
			parentName := parent.CommandPath()
			link := "index.md"
			if cmd.HasSubCommands() {
				link = "../" + link
			}
			buf.WriteString(fmt.Sprintf("* [%s](%s)\t - %s\n", parentName, "index.md", parent.Short))
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

func hasSubCommands(cmd *cobra.Command) bool {
	for _, c := range cmd.Commands() {
		if !c.IsAvailableCommand() || c.IsAdditionalHelpTopicCommand() {
			continue
		}
		return true
	}
	return false
}

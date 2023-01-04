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
	"fmt"
	"html"
	"io"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

func GenMarkdownTree(cmd *cobra.Command, dir string) error {
	for _, c := range cmd.Commands() {
		if err := GenMarkdownTree(c, dir); err != nil {
			return err
		}
	}
	if !cmd.IsAvailableCommand() || cmd.IsAdditionalHelpTopicCommand() {
		return nil
	}
	cmdPath := strings.ReplaceAll(cmd.CommandPath(), " ", "/")
	if cmd.HasSubCommands() {
		cmdPath = path.Join(cmdPath, "index.md")
	} else {
		cmdPath += ".md"
	}
	filename := filepath.Join(dir, cmdPath)
	if err := os.MkdirAll(path.Dir(filename), os.ModePerm); err != nil {
		return err
	}
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	if err := GenMarkdownPage(cmd, f); err != nil {
		return err
	}
	return nil
}

// GenMarkdownPage creates custom markdown output.
func GenMarkdownPage(cmd *cobra.Command, w io.Writer) error {
	var err error
	p := func(format string, a ...any) {
		_, err = w.Write([]byte(fmt.Sprintf(format, a...)))
	}
	p("---\n")
	p("id: %s\n", pageId(cmd))
	p("title: %s\n", cmd.Name())
	p("sidebar_position: %d\n", order(cmd))
	p("---\n")
	cmd.InitDefaultHelpCmd()
	cmd.InitDefaultHelpFlag()
	name := cmd.CommandPath()
	// Name + Version
	if cmd.Version != "" {
		p("version %s\n\n", cmd.Version)
	}
	p(cmd.Short)
	p("\n\n")
	// Synopsis
	if len(cmd.Long) > 0 {
		p("### Synopsis\n\n")
		p("%s \n\n", html.EscapeString(cmd.Long))
	}
	if cmd.Runnable() {
		p("```\n%s\n```\n\n", cmd.UseLine())
	}
	// Examples
	if len(cmd.Example) > 0 {
		p("### Examples\n\n")
		p("```\n%s\n```\n\n", html.EscapeString(cmd.Example))
	}
	// Flags
	flags := cmd.NonInheritedFlags()
	flags.SetOutput(w)
	if flags.HasAvailableFlags() {
		p("### Flags\n\n")
		p("```\n")
		flags.PrintDefaults()
		p("```\n\n")
	}
	// Parent Flags
	parentFlags := cmd.InheritedFlags()
	parentFlags.SetOutput(w)
	if parentFlags.HasAvailableFlags() {
		p("### Flags inherited from parent commands\n\n```\n")
		parentFlags.PrintDefaults()
		p("```\n\n")
	}
	// Subcommands
	if hasSubCommands(cmd) {
		p("### Subcommands\n\n")
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
			p("* [%s](%s)\t - %s\n", commandName, childLink, child.Short)
		}
		p("\n")
	}
	// Parent Command
	if cmd.HasParent() {
		p("### Parent Command\n\n")
		if cmd.HasParent() {
			parent := cmd.Parent()
			parentName := parent.CommandPath()
			link := "index.md"
			if cmd.HasSubCommands() {
				link = "../" + link
			}
			p("* [%s](%s)\t - %s\n", parentName, "index.md", parent.Short)
			cmd.VisitParents(func(c *cobra.Command) {
				if c.DisableAutoGenTag {
					cmd.DisableAutoGenTag = c.DisableAutoGenTag
				}
			})
		}
	}
	return err
}

func order(cmd *cobra.Command) int {
	var i int
	if !cmd.HasParent() {
		return 0
	}
	for _, sibling := range cmd.Parent().Commands() {
		if !cmd.IsAvailableCommand() || cmd.IsAdditionalHelpTopicCommand() {
			continue
		}
		i++
		if sibling.CommandPath() == cmd.CommandPath() {
			return i
		}
	}
	return -1
}

func pageId(cmd *cobra.Command) string {
	if cmd.HasSubCommands() {
		return "index"
	}
	return cmd.Name()
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

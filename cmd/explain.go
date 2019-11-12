/*
Copyright © 2019 Mark Freriks <m.freriks@avisi.nl>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/djmarkoz/kubesafe/pkg/kubesafe"

	"github.com/spf13/cobra"
)

// explainCmd represents the explain command
var explainCmd = &cobra.Command{
	Use:   "explain",
	Short: "Explain which directory expects which cluster",
	Run: func(cmd *cobra.Command, args []string) {
		definitions, err := kubesafe.ExpectedClusterList()
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "failed to determine expected cluster list: %v\n", err)
			os.Exit(1)
		}

		max := 0
		for _, definition := range definitions {
			if len(definition.Dir) > max {
				max = len(definition.Dir)
			}
		}

		previousDir := ""
		for i, definition := range definitions {
			if i == 0 {
				_, err := fmt.Println("/")
				if err != nil {
					os.Exit(1)
				}
			} else {
				line := ""
				for j := 0; j < i; j++ {
					if j+1 == i {
						line = fmt.Sprintf("%s└── ", line)
					} else {
						line = fmt.Sprintf("%s    ", line)
					}
				}
				line = fmt.Sprintf("%s%s", line, strings.Replace(definition.Dir, previousDir, "", 1))
				diff := max - len(line)
				for j := 0; j < diff; j++ {
					line = fmt.Sprintf("%s ", line)
				}

				_, err = fmt.Fprintf(os.Stdout, "%s %s\n", line, definition.ExpectedCluster)
				if err != nil {
					_, _ = fmt.Fprintf(os.Stderr, "failed to print cluster definition: %v\n", err)
					os.Exit(1)
				}
			}
			previousDir = definition.Dir
		}
	},
}

func init() {
	rootCmd.AddCommand(explainCmd)
}

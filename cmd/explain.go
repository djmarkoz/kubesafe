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

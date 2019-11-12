package cmd

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/djmarkoz/kubesafe/pkg/kubesafe"

	"github.com/spf13/cobra"
)

// setCmd represents the set command
var setCmd = &cobra.Command{
	Use:   "set",
	Short: fmt.Sprintf("Create a '%s' file in the current directory specifying the currently active cluster", kubesafe.ExpectedClusterFilename),
	Run: func(cmd *cobra.Command, args []string) {
		currentCluster, err := kubesafe.CurrentCluster()
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "failed to determine current cluster\n", err)
			os.Exit(1)
		}

		if currentCluster == "" {
			_, _ = fmt.Fprintf(os.Stderr, "no active cluster to set\n")
			os.Exit(1)
		} else {
			err = ioutil.WriteFile(kubesafe.ExpectedClusterFilename, []byte(currentCluster+"\n"), 0644)
			if err != nil {
				_, _ = fmt.Fprintf(os.Stderr, "failed to write '%s': %v\n", kubesafe.ExpectedClusterFilename, err)
				os.Exit(1)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(setCmd)
}

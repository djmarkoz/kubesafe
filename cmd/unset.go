package cmd

import (
	"fmt"
	"os"

	"github.com/djmarkoz/kubesafe/pkg/kubesafe"

	"github.com/spf13/cobra"
)

// unsetCmd represents the unset command
var unsetCmd = &cobra.Command{
	Use:   "unset",
	Short: fmt.Sprintf("Remove the '%s' file in the current directory", kubesafe.ExpectedClusterFilename),
	Run: func(cmd *cobra.Command, args []string) {
		err := os.Remove(kubesafe.ExpectedClusterFilename)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(unsetCmd)
}

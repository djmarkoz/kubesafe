package cmd

import (
	"fmt"
	"os"

	"github.com/djmarkoz/kubesafe/pkg/kubesafe"

	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Returns the currently expected cluster",
	Run: func(cmd *cobra.Command, args []string) {
		expectedCluster, err := kubesafe.ExpectedCluster()
		if err != nil {
			if err == kubesafe.NotFound {
				// nothing
			} else {
				_, _ = fmt.Fprintf(os.Stderr, "failed to determine expected cluster: %v", err)
				os.Exit(1)
			}
		}
		fmt.Printf("%s\n", expectedCluster)
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}

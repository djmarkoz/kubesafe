package cmd

import (
	"fmt"
	"os"

	"github.com/djmarkoz/kubesafe/pkg/kubesafe"

	"github.com/spf13/cobra"
)

// verifyCmd represents the verify command
var verifyCmd = &cobra.Command{
	Use:   "verify [flags] -- COMMAND [args..] [options]",
	Short: "Verify if the current cluster is the expected cluster and optionally run a command",
	Run: func(cmd *cobra.Command, args []string) {
		expectedCluster, err := kubesafe.ExpectedCluster()
		if err != nil {
			if err == kubesafe.NotFound {
				_, _ = fmt.Fprintf(os.Stderr, "kubesafe: no validation of expected cluster\n")
			} else {
				_, _ = fmt.Fprintf(os.Stderr, "kubesafe: failed to determine expected cluster: %v\n", err)
				os.Exit(1)
			}
		} else {
			currentCluster, err := kubesafe.CurrentCluster()
			if err != nil {
				_, _ = fmt.Fprintf(os.Stderr, "kubesafe: failed to determine current cluster: %v\n", err)
				os.Exit(1)
			}

			if currentCluster != expectedCluster {
				_, _ = fmt.Fprintf(os.Stderr, "kubesafe: current cluster '%s' does not match expected cluster '%s'\n", currentCluster, expectedCluster)
				os.Exit(1)
			}
		}

		if len(args) > 0 {
			executable := args[0]
			err = kubesafe.Exec(executable, args[1:])
			if err != nil {
				_, _ = fmt.Fprintf(os.Stderr, "kubesafe: failed to execute kubectl command: %v\n", err)
				os.Exit(1)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(verifyCmd)
}

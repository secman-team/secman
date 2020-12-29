package cmdutil

import (
	"os"

	"github.com/abdfnx/secman/v3/api/repox"
	"github.com/spf13/cobra"
)

func EnableRepoOverride(cmd *cobra.Command, f *Factory) {
	cmd.PersistentFlags().StringP("repo", "R", "", "Select another repository using the `[HOST/]OWNER/REPO` format")

	cmd.PersistentPreRun = func(cmd *cobra.Command, args []string) {
		repoOverride, _ := cmd.Flags().GetString("repo")
		if repoFromEnv := os.Getenv("GH_REPO"); repoOverride == "" && repoFromEnv != "" {
			repoOverride = repoFromEnv
		}
		if repoOverride != "" {
			// NOTE: this mutates the factory
			f.BaseRepo = func() (repox.Interface, error) {
				return repox.FromFullName(repoOverride)
			}
		}
	}
}

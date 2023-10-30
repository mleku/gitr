package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// cloneCmd represents the clone command
var cloneCmd = &cobra.Command{
	Use:   "clone <repository> [<directory>]",
	Short: "fetch a git repository from a another location and store it locally",
	Long: `clones a repository from a given path, creates a remote tracking branch for each branch in the repository, and checks out an initial branch that is forked from the cloned repository's currently active branch.

optionally, it will be cloned into the given directory, otherwise the current working directory will be used.
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("clone called")

	},
}

func init() {
	cloneCmd.Flags().BoolP("bare", "b", false, "bare clone - does not create a working tree, but creates a remote tracking for the current (default) branch")
	cloneCmd.Flags().BoolP("mirror", "m", false, "mirror - the same as bare except without any remote tracking added")
	rootCmd.AddCommand(cloneCmd)
}

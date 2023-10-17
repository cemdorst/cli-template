package cmd

import (
	"fmt"
	"os"

	"github.com/cemdorst/cli-template/cmd/corecmd"
	"github.com/cemdorst/cli-template/pkg/version"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:           "cli-template",
	SilenceUsage:  true,
	SilenceErrors: true,
	Short:         "CLI Template",
	Long: `CLI template - Boilerplate code.

Modify as you wish ;)`,
	RunE: corecmd.RequireSubcommandE,
}

var RootCmd = rootCmd

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Version info",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%s\n", version.VERSION)
		fmt.Printf("build: %s\n", version.VERSIONDATETIME)
		fmt.Printf("commit: %s\n", version.VERSIONCOMMIT)
	},
}

func init() {
	rootCmd.PersistentFlags().BoolVar(&corecmd.DebugMode, "debug", false, "debug output")
	rootCmd.AddCommand(versionCmd)
}

func Run() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
	os.Exit(0)
}

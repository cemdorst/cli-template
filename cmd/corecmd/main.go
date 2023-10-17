package corecmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var DebugMode bool

func Debug(a ...interface{}) {
	var f *os.File
	if DebugMode {
		f = os.Stderr
	}
	fmt.Fprint(f, a...)
}

func Debugf(format string, a ...interface{}) {
	var f *os.File
	if DebugMode {
		f = os.Stderr
	}
	fmt.Fprintf(f, format, a...)
}

func RequireSubcommandE(cmd *cobra.Command, args []string) error {
	if len(args) < 1 {
		cmd.Help()
		os.Exit(1)
	}
	return fmt.Errorf("unknown subcommand: %s", args[0])
}

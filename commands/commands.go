package commands

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	// RootCmd sets task command config
	RootCmd = &cobra.Command{
		Use: "fortnite",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Usage()
		},
	}
)

// Exit finishs requests
func Exit(err error, codes ...int) {
	var code int
	if len(codes) > 0 {
		code = codes[0]
	} else {
		code = 2
	}
	fmt.Println(err)
	os.Exit(code)
}

// Run runs CLI action
func Run() {
	if err := RootCmd.Execute(); err != nil {
		Exit(fmt.Errorf("failed to run: %s", err.Error()), 1)
	}
}

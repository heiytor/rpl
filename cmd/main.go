package main

import (
	"os"

	"github.com/spf13/cobra"
)

var rplCmd = &cobra.Command{
	Version: "dev-0.0.1",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Println("Hello World!")
	},
}

func main() {
	if err := rplCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

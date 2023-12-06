package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

func parseArgs(args []string) (pattern string, target string, files []string, ok bool) {
	if len(args) < 3 {
		return pattern, target, files, ok
	}

	pattern = args[0]
	target = args[1]
	files = args[2:]
	ok = true

	return pattern, target, files, ok
}

var rplCmd = &cobra.Command{
	Version: "dev-0.0.1",
	Run: func(cmd *cobra.Command, args []string) {
		pattern, target, files, ok := parseArgs(args)
		if !ok {
			fmt.Println("Invalid args")
			os.Exit(1)
		}

		fmt.Println(pattern)
		fmt.Println(target)
		fmt.Println(files)

		for _, f := range files {
			fpath := filepath.Clean(f)

			if _, err := os.Stat(fpath); os.IsNotExist(err) {
				fmt.Printf("File %s does not exists.\n", fpath)
				os.Exit(1)
			}

			content, err := os.ReadFile(fpath)
			if err != nil {
				fmt.Printf("Error reading file %s: %v\n", fpath, err)
				os.Exit(1)
			}

			newContent := strings.Replace(string(content), pattern, target, -1)

			err = os.WriteFile(fpath, []byte(newContent), 0644)
			if err != nil {
				fmt.Printf("Error writing file %s: %v\n", fpath, err)
				os.Exit(1)
			}
		}
	},
}

func main() {
	if err := rplCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

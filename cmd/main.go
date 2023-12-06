package main

import (
	"fmt"
	"io/fs"
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

func replaceString(path, pattern, target string) {
	content, err := os.ReadFile(path)
	if err != nil {
		fmt.Printf("Error reading file %s: %v\n", path, err)
		os.Exit(1)
	}

	newContent := strings.Replace(string(content), pattern, target, -1)

	err = os.WriteFile(path, []byte(newContent), 0644)
	if err != nil {
		fmt.Printf("Error writing file %s: %v\n", path, err)
		os.Exit(1)
	}
}

var flagRecursive bool

var rplCmd = &cobra.Command{
	Version: "dev-0.0.1",
	Run: func(cmd *cobra.Command, args []string) {
		pattern, target, files, ok := parseArgs(args)
		if !ok {
			fmt.Println("Invalid args")
			os.Exit(1)
		}

		fmt.Println("******************")
		fmt.Printf("pattern: %s\n", pattern)
		fmt.Printf("target: %s\n", target)
		fmt.Printf("files: %s\n", files)
		fmt.Println("******************")

		for _, f := range files {
			path := filepath.Clean(f)

			info, err := os.Stat(path)
			if os.IsNotExist(err) {
				fmt.Printf("File %s does not exists.\n", path)
				os.Exit(1)
			}

			if info.IsDir() {
				if !flagRecursive {
					fmt.Println("Trying to replace an entire directory without recursive flag. Try using -r or --recursive.")
					os.Exit(1)
				}

				filepath.WalkDir(path, func(path string, d fs.DirEntry, _ error) error {
					info, err := d.Info()
					if err != nil {
						panic(err)
					}

					if info.IsDir() {
						return nil
					}

					replaceString(path, pattern, target)

					return nil
				})
			} else {
				replaceString(path, pattern, target)
			}
		}
	},
}

func main() {
	rplCmd.PersistentFlags().BoolVarP(&flagRecursive, "recursive", "r", false, "...")

	if err := rplCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

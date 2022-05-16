package main

import (
	"github.com/hnlq715/struct2interface"
	"github.com/spf13/cobra"
)

func main() {
	var (
		dir string
	)

	root := &cobra.Command{
		Use: "struct2interface",
		RunE: func(cmd *cobra.Command, args []string) error {
			return struct2interface.MakeDir(dir)
		},
	}

	root.Flags().StringVarP(&dir, "dir", "d", ".", "Go source file dir to read")
	if err := root.Execute(); err != nil {
		panic(err)
	}
}

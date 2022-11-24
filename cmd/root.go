package cmd

import (
	"fmt"
	"os"

	"github.com/osborngh/break/pkg"
	"github.com/spf13/cobra"
)

var directory string

var version = `0.1.0
Copyright 2022 Osborn`

var rootCmd = &cobra.Command{
	Use: "break",
	Version: version,
	Short: "break - a simple build tool.",
	Long: `break - the 'make' alternative you never needed.`,
	Run: func(cmd *cobra.Command, args []string) {
		pkg.HandleRoot(directory, args)
	},
}

func Execute() {
	rootCmd.Flags().StringVarP(&directory, "directory", "d", ".", "Specify Breakfile Path")
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "[ERROR]: An Error Occurred!\n")
		os.Exit(1)
	}
}

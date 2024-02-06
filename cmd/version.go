/*
Copyright © 2024 Rudolf Farkas @rudifa rudi.farkas@gmail.com
*/

// Package cmd implements the cuedo commands
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	cueCommitSHA    string
	cueCommitDate   string
	cuedoCommitSHA  string
	cuedoCommitDate string
	DateNow         string
	User            string
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version data",
	Long:  `Print the version data.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("cuedo built %s by %s with:\n", DateNow, User)
		fmt.Printf("        cue %s %s\n", cueCommitDate, cueCommitSHA)
		fmt.Printf("      cuedo %s %s\n", cuedoCommitDate, cuedoCommitSHA)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

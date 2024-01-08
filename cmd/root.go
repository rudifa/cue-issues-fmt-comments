/*
Copyright © 2024 Rudolf Farkas @rudifa rudi.farkas@gmail.com
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var verbose bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cuedo-fmt",
	Short: "Tools for investigation of CUE issues related to formatting.",
	Long: `Tools for investigation of CUE issues related to formatting.

	This evolving set of tools requires a compatible set of patches to the CUE source code (*).

	As of 7 Jan 2024, the CUE patch extras are turned on by setting these environment variables to non-empty values:

	CUEDO_AST_SPEW - turns on the AST spew mode
	CUEDO_AST_TREE - turns on the AST tree mode
	CUEDO_AST_TYPE - turns on the AST type mode
	CUEDO_FBB_KLUDGE - turns on a kludge in formatter to anticipate a fix for the issue #2567
	CUEDO_FORMATTER_HEXDUMP - turns on hex dump of the formatter's internal buffer, as each fragment is printed
	CUEDO_FORMATTER_STACKTRACE - turns on dump of a stack trace in the formatter, as each fragment is printed
	CUEDO_PARSER_COMMENTS_POS - adds to the parser.Trace mode data about the comment positions and texts
	CUEDO_PARSER_DEBUG_STR - turns on the parser debug string mode
	CUEDO_PARSER_TRACE - turns on the parser.Trace mode

	Note that these patches are not part of the official CUE source code and are not supported by the CUE team.

	(*) branch https://github.com/rudifa/cue/tree/change-1173870-cuedo
	`,
	// Run: func(cmd *cobra.Command, args []string) {
	// },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.EnablePrefixMatching = true                  // allow typing partial command names
	rootCmd.CompletionOptions.DisableDefaultCmd = true // disable default completion

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Enable verbose output")
}

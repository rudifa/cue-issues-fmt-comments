/*
Copyright © 2024 Rudolf Farkas @rudifa rudi.farkas@gmail.com
*/
package cmd

import (
	"log"
	"os"
	"strings"

	"github.com/rudifa/cuedo-fmt/runner"
	"github.com/spf13/cobra"
)

// formatCmd represents the format command
var formatCmd = &cobra.Command{
	Use:   "format",
	Short: "Parse and format a CUE file, optionally displaying the parser and formatter inner data.",
	Long: `Parse and format a CUE file, optionally displaying the parser and formatter inner data.

	In the absence of any flags, the command will parse and format the CUE file and print the input and the result.

	Flags -x and -s are best used together for debugging the formatter.

	Add flags in any combination to display the inner data.
	`,
	Aliases: []string{"fmt"},
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		// Get the parser_trace flag value and set the environment variables
		setEnvironmentVariables(cmd)

		// Get the filename from the command arguments
		filename := args[0]

		// log.Println("Filename:", filename)

		runner.RunParseAndFormat(filename)
	},
}

func init() {
	rootCmd.AddCommand(formatCmd)

	formatCmd.Flags().BoolP("fbb_kludge", "k", false, "formatting: enable a kludge in formatter to anticipate a fix for the issue #2567")
	formatCmd.Flags().BoolP("formatter_hexdump", "x", false, "formatting: print a hex dump of the formatter's internal buffer, for each fragment when printed")
	formatCmd.Flags().BoolP("formatter_stacktrace", "s", false, "formatting: print the stack trace for each fragment when printed")
	formatCmd.Flags().BoolP("full_monty", "f", false, "turn on everything")
	formatCmd.Flags().BoolP("parser_ast_spew", "w", false, "parser: deep pretty print the AST tree, when a node is closed")
	formatCmd.Flags().BoolP("parser_ast_tree", "a", false, "parser: print the final file AST tree")
	formatCmd.Flags().BoolP("parser_ast_node_type", "n", false, "parser: print the AST node type, when a node is closed")
	formatCmd.Flags().BoolP("parser_comments_pos", "c", false, "parser: print the comment position and text while creating nodes")
	formatCmd.Flags().BoolP("parser_debug_str", "d", false, "parser: print the final file AST debug string")
	formatCmd.Flags().BoolP("parser_trace", "t", false, "parser: print the parser.Trace while creating nodes")

	// customize the usage message _after_ flags are defined
	defaultUsage := formatCmd.UsageString()
	customUsage := strings.Replace(defaultUsage, "cuedo-fmt format", "cuedo-fmt format <inputfile>.cue", -1)
	formatCmd.SetUsageTemplate(customUsage)
}

func setEnvironmentVariables(cmd *cobra.Command) {

	fbb_kludge, _ := cmd.Flags().GetBool("fbb_kludge")
	formatter_hexdump, _ := cmd.Flags().GetBool("formatter_hexdump")
	formatter_stacktrace, _ := cmd.Flags().GetBool("formatter_stacktrace")
	full_monty, _ := cmd.Flags().GetBool("full_monty")
	parser_ast_spew, _ := cmd.Flags().GetBool("parser_ast_spew")
	parser_ast_tree, _ := cmd.Flags().GetBool("parser_ast_tree")
	parser_ast_node_type, _ := cmd.Flags().GetBool("parser_ast_node_type")
	parser_comments_pos, _ := cmd.Flags().GetBool("parser_comments_pos")
	parser_debug_str, _ := cmd.Flags().GetBool("parser_debug_str")
	parser_trace, _ := cmd.Flags().GetBool("parser_trace")

	if full_monty {
		fbb_kludge = true
		formatter_hexdump = true
		formatter_stacktrace = true
		parser_ast_spew = true
		parser_ast_tree = true
		parser_ast_node_type = true
		parser_comments_pos = true
		parser_debug_str = true
		parser_trace = true
	}

	if fbb_kludge {
		os.Setenv("CUEDO_FBB_KLUDGE", "1")
	}
	if formatter_hexdump {
		os.Setenv("CUEDO_FORMATTER_HEXDUMP", "1")
	}
	if formatter_stacktrace {
		os.Setenv("CUEDO_FORMATTER_STACKTRACE", "1")
	}
	if parser_ast_spew {
		os.Setenv("CUEDO_AST_NODE_SPEW", "1")
	}
	if parser_ast_tree {
		os.Setenv("CUEDO_AST_TREE", "1")
	}
	if parser_ast_node_type {
		os.Setenv("CUEDO_AST_NODE_TYPE", "1")
	}
	if parser_comments_pos {
		os.Setenv("CUEDO_PARSER_COMMENTS_POS", "1")
		parser_trace = true
	}
	if parser_debug_str {
		os.Setenv("CUEDO_PARSER_DEBUG_STR", "1")
	}
	if parser_trace {
		os.Setenv("CUEDO_PARSER_TRACE", "1")
	}

	if verbose {

		// print flags that are set
		{
			if fbb_kludge {
				log.Println("fbb_kludge:", fbb_kludge)
			}
			if formatter_hexdump {
				log.Println("formatter_hexdump:", formatter_hexdump)
			}
			if formatter_stacktrace {
				log.Println("formatter_stacktrace:", formatter_stacktrace)
			}
			if full_monty {
				log.Println("full_monty:", full_monty)
			}
			if parser_ast_spew {
				log.Println("parser_ast_spew:", parser_ast_spew)
			}
			if parser_ast_tree {
				log.Println("parser_ast_tree:", parser_ast_tree)
			}
			if parser_ast_node_type {
				log.Println("parser_ast_node_type:", parser_ast_node_type)
			}
			if parser_comments_pos {
				log.Println("parser_comments_pos:", parser_comments_pos)
			}
			if parser_debug_str {
				log.Println("parser_debug_str:", parser_debug_str)
			}
			if parser_trace {
				log.Println("parser_trace:", parser_trace)
			}
		}
		printEnvVars()
	}
}

// printEnvVars prints envvars that are set
func printEnvVars() {
	envvars := []string{
		"CUEDO_FBB_KLUDGE",
		"CUEDO_FORMATTER_STACKTRACE",
		"CUEDO_FORMATTER_HEXDUMP",
		"CUEDO_PARSER_TRACE",
		"CUEDO_PARSER_DEBUG_STR",
		"CUEDO_AST_TREE",
		"CUEDO_AST_NODE_TYPE",
		"CUEDO_PARSER_COMMENTS_POS",
		"CUEDO_AST_NODE_SPEW",
	}
	for _, envvar := range envvars {
		if os.Getenv(envvar) != "" {
			log.Printf("%s is set\n", envvar)
		}
	}
}

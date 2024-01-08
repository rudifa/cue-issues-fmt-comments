package cmd

/*
Copyright © 2024 Rudolf Farkas @rudifa rudi.farkas@gmail.com
*/

import (
	"log"
	"os"
	"strings"

	"github.com/rudifa/cuedo/runner"
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
	formatCmd.Flags().BoolP("parser_ast_node_type_and_comments", "n", false, "parser: print the AST node type and comments, when a node is closed")
	formatCmd.Flags().BoolP("parser_comments_pos", "c", false, "parser: print the comment position and text while creating nodes")
	formatCmd.Flags().BoolP("parser_debug_str", "d", false, "parser: print the final file AST debug string")
	formatCmd.Flags().BoolP("parser_trace", "t", false, "parser: print the parser.Trace while creating nodes")

	// customize the usage message _after_ flags are defined
	defaultUsage := formatCmd.UsageString()
	customUsage := strings.Replace(defaultUsage, "cuedo format", "cuedo format <inputfile>.cue", -1)
	formatCmd.SetUsageTemplate(customUsage)
}

func setEnvironmentVariables(cmd *cobra.Command) {

	fbbKludge, _ := cmd.Flags().GetBool("fbb_kludge")
	formatterHexdump, _ := cmd.Flags().GetBool("formatter_hexdump")
	formatterStacktrace, _ := cmd.Flags().GetBool("formatter_stacktrace")
	fullMonty, _ := cmd.Flags().GetBool("full_monty")
	parserAstSpew, _ := cmd.Flags().GetBool("parser_ast_spew")
	parserAstTree, _ := cmd.Flags().GetBool("parser_ast_tree")
	parserAstNodeTypeAndSomments, _ := cmd.Flags().GetBool("parser_ast_node_type_and_comments")
	parserCommentsPos, _ := cmd.Flags().GetBool("parser_comments_pos")
	parserDebugStr, _ := cmd.Flags().GetBool("parser_debug_str")
	parserTrace, _ := cmd.Flags().GetBool("parser_trace")

	if fullMonty {
		fbbKludge = true
		formatterHexdump = true
		formatterStacktrace = true
		parserAstSpew = true
		parserAstTree = true
		parserAstNodeTypeAndSomments = true
		parserCommentsPos = true
		parserDebugStr = true
		parserTrace = true
	}

	if fbbKludge {
		os.Setenv("CUEDO_FBB_KLUDGE", "1")
	}
	if formatterHexdump {
		os.Setenv("CUEDO_FORMATTER_HEXDUMP", "1")
	}
	if formatterStacktrace {
		os.Setenv("CUEDO_FORMATTER_STACKTRACE", "1")
	}
	if parserAstSpew {
		os.Setenv("CUEDO_AST_NODE_SPEW", "1")
	}
	if parserAstTree {
		os.Setenv("CUEDO_AST_TREE", "1")
	}
	if parserAstNodeTypeAndSomments {
		os.Setenv("CUEDO_AST_NODE_TYPE_AND_COMMENTS", "1")
	}
	if parserCommentsPos {
		os.Setenv("CUEDO_PARSER_COMMENTS_POS", "1")
		parserTrace = true
	}
	if parserDebugStr {
		os.Setenv("CUEDO_PARSER_DEBUG_STR", "1")
	}
	if parserTrace {
		os.Setenv("CUEDO_PARSER_TRACE", "1")
	}

	if verbose {

		// print flags that are set
		{
			if fbbKludge {
				log.Println("fbb_kludge:", fbbKludge)
			}
			if formatterHexdump {
				log.Println("formatter_hexdump:", formatterHexdump)
			}
			if formatterStacktrace {
				log.Println("formatter_stacktrace:", formatterStacktrace)
			}
			if fullMonty {
				log.Println("full_monty:", fullMonty)
			}
			if parserAstSpew {
				log.Println("parser_ast_spew:", parserAstSpew)
			}
			if parserAstTree {
				log.Println("parser_ast_tree:", parserAstTree)
			}
			if parserAstNodeTypeAndSomments {
				log.Println("parser_ast_node_type_and_comments:", parserAstNodeTypeAndSomments)
			}
			if parserCommentsPos {
				log.Println("parser_comments_pos:", parserCommentsPos)
			}
			if parserDebugStr {
				log.Println("parser_debug_str:", parserDebugStr)
			}
			if parserTrace {
				log.Println("parser_trace:", parserTrace)
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
		"CUEDO_AST_NODE_TYPE_AND_COMMENTS",
		"CUEDO_PARSER_COMMENTS_POS",
		"CUEDO_AST_NODE_SPEW",
	}
	for _, envvar := range envvars {
		if os.Getenv(envvar) != "" {
			log.Printf("%s is set\n", envvar)
		}
	}
}

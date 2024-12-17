package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/ezep02/gemini-chat-cli/pkg/gemini"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "hello",
	Short: "This is the first command",
	Long: `A longer description 
	for the first command`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("This is the first cobra example")
	},
}

var AskGemmini = &cobra.Command{
	Use:     "Gemmini",
	Short:   "Ask to Gemmini",
	Aliases: []string{"Gem"},
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		resp, err := gemini.Gemini(args[0])
		if err != nil {
			log.Println("[ERROR]", err.Error())
		}

		for _, cand := range resp.Candidates {
			if cand.Content != nil {
				for _, part := range cand.Content.Parts {
					fmt.Println(part)
				}
			}
		}
		fmt.Println("---")
	},
}

func Init() {
	rootCmd.AddCommand(AskGemmini)
}

func Execute() {
	Init()

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}

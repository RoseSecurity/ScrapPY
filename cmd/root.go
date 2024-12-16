package cmd

import (
	"fmt"
	"log"

	tui "github.com/RoseSecurity/ScrapNGo/internal/tui/utils"
	"github.com/RoseSecurity/ScrapNGo/pkg/utils"
	"github.com/spf13/cobra"
)

var (
	blue  = "\033[34m"
	red   = "\033[91m"
	green = "\033[92m"
	norm  = "\x1b[0m"
	tag   = "@RoseSecurity"
)

var rootCmd = &cobra.Command{
	Use:   "scrapNGo",
	Short: "ScrapNGo enumerates documents, manuals, and sensitive PDFs for key phrases and words that can be utilized in dictionary and brute force attacks.",
	Long: `ScrapNGo enumerates documents, manuals, and sensitive PDFs for key phrases and words 
that can be utilized in dictionary and brute force attacks. These keywords are outputted 
to a text file (ScrapNGo.txt in the directory which the tool was run from) that can be read 
by tools such as Hydra, Dirb, and other offensive security tools for initial access and 
lateral movement.`,
	Run: func(cmd *cobra.Command, args []string) {
		file, _ := cmd.Flags().GetString("file")
		mode, _ := cmd.Flags().GetString("mode")
		outputFile, _ := cmd.Flags().GetString("output-file")

		// Print help if no file is provided
		if file == "" {
			cmd.Help()
			return
		}

		// Extract text from PDF
		fileContent, err := utils.ExtractTextFromPDF(file)
		if err != nil {
			utils.LogErrorAndExit(err)
		}

		wordList := utils.RemoveCommonWords(fileContent)

		if len(fileContent) == 0 {
			log.Fatalf("No content found in the PDF file: %s", file)
		}

		// Process PDF content based on mode
		var keywords []string
		switch mode {
		case "word-frequency":
			keywords = utils.WordFrequency(wordList, 100)
		case "metadata":
			utils.PrintMetadata(file)
			if err != nil {
				utils.LogErrorAndExit(err)
			}
			return
		case "entropy":
			// keywords = utils.CalculateEntropy(wordList, 100)
			// default:
			// 	keywords = utils.ExtractKeywords(fileContent)
		}

		// Write output to file
		if err := utils.WriteToFile(outputFile, keywords); err != nil {
			utils.LogErrorAndExit(err)
		}

		fmt.Printf("%s has been created!\n", outputFile)
	},
}

func init() {
	// Custom help menu to display banner
	rootCmd.SetHelpFunc(func(cmd *cobra.Command, args []string) {
		fmt.Println()
		tui.PrintStyledText("SCRAPPY")
		fmt.Println(cmd.UsageString())
	})
	// Docs and Version commands
	rootCmd.AddCommand(docsCmd)
	rootCmd.AddCommand(versionCmd)
	rootCmd.Flags().StringP("file", "f", "", "PDF input file")
	rootCmd.Flags().StringP("mode", "m", "full", "Modes of operation: full, word-frequency, metadata, entropy")
	rootCmd.Flags().StringP("output-file", "o", "ScrapNGo.txt", "Output file name")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		utils.LogErrorAndExit(err)
	}
}

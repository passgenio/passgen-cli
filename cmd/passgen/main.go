package main

import (
	"fmt"

	"github.com/amirrezaask/passgen"
	"github.com/spf13/cobra"
)

var masterPassword string
var algo string

var rootCmd = &cobra.Command{
	Use: `passgen`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var genCmd = &cobra.Command{
	Use:   `gen`,
	Short: "generate a new password for website.",
	Long:  `generate a new password for website.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			panic(args)
		}
		p := passgen.NewPassGen([]byte(masterPassword), passgen.NewAlgorithm(algo))

		generated, err := p.GenFor(args[0])
		if err != nil {
			panic(err)
		}
		fmt.Printf("%x\n", generated)
	},
}

func init() {
	rootCmd.PersistentFlags().StringVar(&masterPassword, "master", "", "Master password")
	rootCmd.PersistentFlags().StringVar(&algo, "algo", "aes", "Algorithm to use.")
	genCmd.Aliases = []string{"generate", "g"}
	rootCmd.AddCommand(genCmd)

}

func main() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}

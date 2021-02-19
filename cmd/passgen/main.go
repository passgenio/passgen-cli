package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/amirrezaask/passgen"
	"github.com/spf13/cobra"
)

type config struct {
	Master string `json:"master"`
}

var configPath string
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

		f, err := os.OpenFile(filepath.Clean(configPath), os.O_CREATE|os.O_RDONLY, 0644)
		if err != nil {
			panic(err)
		}
		var c config
		err = json.NewDecoder(f).Decode(&c)
		if err != nil {
			panic(err)
		}
		p := passgen.NewPassGen([]byte(c.Master), passgen.NewAlgorithm(algo))
		generated, err := p.GenFor(args[0])
		if err != nil {
			panic(err)
		}
		fmt.Printf("%x\n", generated)
	},
}

func init() {
	rootCmd.PersistentFlags().StringVar(&configPath, "config", fmt.Sprintf("%s/.passgen.json", os.Getenv("HOME")), "config path that contains your configuration")
	rootCmd.PersistentFlags().StringVar(&algo, "algo", "aes", "Algorithm to use.")
	genCmd.Aliases = []string{"generate", "g"}
	rootCmd.AddCommand(genCmd)

}

func main() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}

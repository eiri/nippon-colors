package main

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

var (
	input, output string

	rootCmd = &cobra.Command{
		Use:   "convert",
		Short: "A cli helper for nippon-colors package",
	}
)

func init() {
	log.SetFlags(0)
	log.SetOutput(os.Stdout)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

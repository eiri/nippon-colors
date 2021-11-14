package main

import (
	"github.com/spf13/cobra"

	color "github.com/eiri/nippon-colors"
)

var (
	demoCmd = &cobra.Command{
		Use:   "demo",
		Short: "Ouput colors demo",
		Run:   runDemo,
	}
)

func init() {
	rootCmd.AddCommand(demoCmd)
}

func runDemo(cmd *cobra.Command, args []string) {
	color.Shironezumi.Println("Colors for dark background")
	for _, c := range color.Dark.Series() {
		c.Print("▀")
	}
	color.Shironezumi.Println("")

	color.Shironezumi.Println("Colors for bright background")
	for _, c := range color.Bright.Series() {
		c.Print("▀")
	}
	color.Shironezumi.Println("")
}

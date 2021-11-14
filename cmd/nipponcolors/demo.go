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

	color.Shironezumi.Println("Reds")
	for _, c := range color.Red.Series() {
		c.Print("▀")
	}
	color.Shironezumi.Println("")

	color.Shironezumi.Println("Blacks")
	for _, c := range color.Black.Series() {
		c.Print("▀")
	}
	color.Shironezumi.Println("")

	color.Shironezumi.Println("Whites")
	for _, c := range color.White.Series() {
		c.Print("▀")
	}
	color.Shironezumi.Println("")

	color.Shironezumi.Println("Purples")
	for _, c := range color.Purple.Series() {
		c.Print("▀")
	}
	color.Shironezumi.Println("")

	color.Shironezumi.Println("Cyans")
	for _, c := range color.Cyan.Series() {
		c.Print("▀")
	}
	color.Shironezumi.Println("")

	color.Shironezumi.Println("Greens")
	for _, c := range color.Green.Series() {
		c.Print("▀")
	}
	color.Shironezumi.Println("")

	color.Shironezumi.Println("Yellows")
	for _, c := range color.Yellow.Series() {
		c.Print("▀")
	}
	color.Shironezumi.Println("")
}

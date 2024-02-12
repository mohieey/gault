/*
Copyright © 2024
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Clear all the secrets(delete the secrets file)",
	Run: func(cmd *cobra.Command, args []string) {
		err := os.Remove(secretsPath())
		if err != nil {
			fmt.Printf("err: %v\n", err)
			os.Exit(1)
		}

		fmt.Println("secrets cleared successfully")
	},
}

func init() {
	rootCmd.AddCommand(clearCmd)
}

/*
Copyright © 2024
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/mohieey/gault/vault"
	"github.com/spf13/cobra"
)

var setCmd = &cobra.Command{
	Use:   "set",
	Short: "setting a secret with a key",
	Run: func(cmd *cobra.Command, args []string) {
		v := vault.File(encodingKey, secretsPath())
		key, value := args[0], args[1]
		err := v.Set(key, value)
		if err != nil {
			fmt.Printf("err: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("value: %v has been set successfully\n", value)
	},
}

func init() {
	rootCmd.AddCommand(setCmd)
}

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

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "getting a secret with a key",
	Run: func(cmd *cobra.Command, args []string) {
		v := vault.File(encodingKey, secretsPath())
		key := args[0]
		value, err := v.Get(key)
		if err != nil {
			fmt.Printf("err: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("value: %v\n", value)
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}

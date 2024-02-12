/*
Copyright © 2024
*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
)

var encodingKey string

var rootCmd = &cobra.Command{
	Use:   "gault",
	Short: "A light copy of vault by hashicorp",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.PersistentFlags().StringVarP(&encodingKey, "key", "k", "", "The ke used to encrypt the secrets file")
}

func secretsPath() string {
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return filepath.Join(home, ".gault")
}

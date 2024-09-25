package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var pathToConf string

func init() {
	rootCmd.PersistentFlags().StringVar(&pathToConf, "conf", "config.yaml", "path to config")
}

var rootCmd = &cobra.Command{
	Use: "kashein",
}

func Execute() error {
	fmt.Printf("kashein\n")

	return nil
}

/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// printYamlCmd represents the printYaml command
var printYamlCmd = &cobra.Command{
	Use:   "printYaml",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		for i, key := range viper.AllKeys() {
			fmt.Printf("%d: %s = %s\n", i+1, key, viper.Get(key))
		}
	},
}

func init() {
	rootCmd.AddCommand(printYamlCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// printYamlCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// printYamlCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"sort"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// printEnvCmd represents the printEnv command
var printEnvCmd = &cobra.Command{
	Use:   "printEnv",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		keys := viper.AllKeys()
		sort.SliceStable(keys, func(i, j int) bool {
			return keys[i] < keys[j]
		})
		for _, key := range keys {
			k := strings.ToUpper(strings.Replace(key, ".", "_", -1))
			v := viper.Get(key)
			switch t := v.(type) {
			case int, int32, int64:
				fmt.Printf("%s\t%d\n", k, t)
			case float32, float64:
				fmt.Printf("%s\t%f\n", k, t)
			case bool:
				f := "false"
				if t {
					f = "true"
				}
				fmt.Printf("%s\t%s\n", k, f)
			default:
				fmt.Printf("%s\t%s\n", k, v)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(printEnvCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// printEnvCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// printEnvCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

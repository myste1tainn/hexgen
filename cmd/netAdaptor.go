/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"
	"text/template"

	"github.com/myste1tainn/hexgen/core"
	"github.com/myste1tainn/hexgen/templates/adaptor"
	"github.com/myste1tainn/hexgen/util/fs"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// netAdaptorCmd represents the netAdaptor command
var netAdaptorCmd = &cobra.Command{
	Use:   "netAdaptor",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fs.CreateAllDirectories()

		var config core.AdaptorConfig
		configKey := args[0]
		key := fmt.Sprintf("integration.%s", configKey)
		if err := viper.UnmarshalKey(key, &config); err != nil {
			panic(err)
		}

		tmpl, err := template.New("Adatpor").Parse(adaptor.AdaptorTmpl)
		// Capture any error
		if err != nil {
			panic(err)
		}

		configTmpl, err := template.New("AdaptorConfig").Parse(adaptor.AdaptorConfigTmpl)
		// Capture any error
		if err != nil {
			panic(err)
		}

		// Print out the tmpl to std
		var fns []string
		for _, v := range config.Apis {
			if v.Name == "" {
				panic("function name cannot be empty")
			}
			fns = append(fns, v.Name)
		}

		if config.Name == "" {
			panic("adaptor name cannot be empty")
		}

		value := core.AdaptorTemplateValues{
			Module:  core.GetModule(),
			Name:    config.Name,
			KeyPath: "integration." + configKey,
			Fns:     fns,
		}

		targetDir := fmt.Sprintf("./internal/adaptor/repo")
		fp := fmt.Sprintf("%s/%sRepo.go", targetDir, value.Name)
		f, err := os.Create(fp)
		defer f.Close()
		if err != nil {
			panic(err)
		}
		fmt.Printf("Creating adaptor at %s\n", fp)
		if err := tmpl.Execute(f, value); err != nil {
			panic(err)
		}

		fp = fmt.Sprintf("%s/%sRepoConfig.go", targetDir, value.Name)
		configF, err := os.Create(fp)
		defer configF.Close()
		if err != nil {
			panic(err)
		}
		fmt.Printf("Creating adaptor at %s\n", fp)
		if err := configTmpl.Execute(configF, value); err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(netAdaptorCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// netAdaptorCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// netAdaptorCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

package cmd

import (
	"github.com/myste1tainn/hexgen/util/fs"
	"github.com/spf13/cobra"
)

// netAdaptorsCmd represents the netAdaptors command
var netsCmd = &cobra.Command{
	Use:   "nets",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fs.CreateAllDirectories()

		netAdaptorsCmd.Run(cmd, args)
		netPortsCmd.Run(cmd, args)
	},
}

func init() {
	rootCmd.AddCommand(netsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// netsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// netsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

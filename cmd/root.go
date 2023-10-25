package cmd

import (
	"log"

	"github.com/rezzamaqfiro/wallet/dep"
	"github.com/spf13/cobra"
)

var configFile string
var di *dep.DI

var rootCmd = &cobra.Command{
	Use:   "wallet",
	Short: "Wallet backend application",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		var err error
		di, err = dep.InitDI(configFile)
		if err != nil {
			log.Fatal(err)
		}

	},
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func Execute() error {
	rootCmd.AddCommand(&cobra.Command{Use: "completion", Hidden: true})
	rootCmd.Flags().StringVarP(&configFile, "config", "c", "config.yaml", "Config file to use")
	return rootCmd.Execute()
}

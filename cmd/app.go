package cmd

import (
	"github.com/neutrinoapp/neutrino-client"
	"github.com/spf13/cobra"
)

var appCmd = &cobra.Command{
	Use:   "app",
	Short: "App",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Println("Specify app id")
			return
		}

		conf := Config()
		conf.App = args[0]
		err := conf.Save()
		if err != nil {
			cmd.Println(err)
			return
		}

		cmd.Println("AppId set to:", args[0])
	},
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		c := neutrino.NewApiClientClean()
		apps, err := c.GetApps()
		if err != nil {
			cmd.Println(err)
			return
		}

		cmd.Println(apps)
	},
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Println("Specify app name")
			return
		}

		c := neutrino.NewApiClientClean()
		id, err := c.CreateApp(args[0])
		if err != nil {
			cmd.Println(err)
			return
		}

		cmd.Println("Created app with id:", id)
	},
}

func init() {
	RootCmd.AddCommand(appCmd)
	appCmd.AddCommand(listCmd, createCmd)
}

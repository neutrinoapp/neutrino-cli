package cmd

import "github.com/spf13/cobra"

var collectionCmd = &cobra.Command{
	Use:   "collection",
	Short: "Collection",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var collectionListCmd = &cobra.Command{
	Use:   "list",
	Short: "collection",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		collections, err := Client().GetCollections()
		if err != nil {
			cmd.Println(err)
			return
		}

		cmd.Println(collections)
	},
}

func init() {
	RootCmd.AddCommand(collectionCmd)
	collectionCmd.AddCommand(collectionListCmd)
}

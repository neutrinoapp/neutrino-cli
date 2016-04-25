package cmd

import (
	"bufio"
	"os"

	"github.com/neutrinoapp/neutrino-client"
	"github.com/spf13/cobra"
)

// registerCmd represents the register command
var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "Register to neutrino",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Println("Specify a username")
		}

		u := args[0]
		c := neutrino.NewApiClientClean()
		reader := bufio.NewReader(os.Stdin)
		cmd.Print("Enter password: ")
		p, _ := reader.ReadString('\n')

		err := c.Register(u, p)
		if err != nil {
			cmd.Println(err)
		} else {
			cmd.Printf("Registered, you can now login\r\n")
		}
	},
}

func init() {
	RootCmd.AddCommand(registerCmd)
}

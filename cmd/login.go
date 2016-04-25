package cmd

import (
	"bufio"
	"os"

	"github.com/neutrinoapp/neutrino-cli/cli"
	"github.com/neutrinoapp/neutrino-client"
	"github.com/spf13/cobra"
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login to neutrino",
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

		token, err := c.Login(u, p)
		if err != nil {
			cmd.Println(err)
			return
		} else if token == "" {
			cmd.Println("Invalid credentials")
			return
		}

		cmd.Printf("Logged in, your token is: %s\r\n", token)
		conf := &cli.Config{
			Username: u,
			Token:    token,
		}
		saveErr := conf.Save()
		if saveErr != nil {
			cmd.Println(saveErr)
			return
		}
	},
}

func init() {
	RootCmd.AddCommand(loginCmd)
}

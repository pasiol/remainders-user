package cmd

import (
	"fmt"

	"github.com/pasiol/remainders-user/internal"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new user",
	Long: `Create a new user and marking it approved. For example:
	remainders-user create --username "Matti Möttönen" --password "qwerty" --mongouri "mongodb+srv://<username>:<password>@<host.domain.name>/<database>?retryWrites=true&w=majority"
	Look more how to generate mongodb uri from: https://docs.mongodb.com/manual/reference/connection-string/`,
	Run: func(cmd *cobra.Command, args []string) {

		username, _ := cmd.Flags().GetString("username")
		password, _ := cmd.Flags().GetString("password")
		mongouri, _ := cmd.Flags().GetString("mongouri")

		err := internal.CreateNewUser(username, password, mongouri)
		if err != nil {
			fmt.Printf("Creating user failed: %s", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	createCmd.Flags().String("username", "", "Login name for user")
	createCmd.Flags().String("password", "", "Password for user")
	createCmd.Flags().String("mongouri", "", "Mongouri for example: mongodb+srv://<username>:<password>@<host.domain.name>/<database>?retryWrites=true&w=majority")

}

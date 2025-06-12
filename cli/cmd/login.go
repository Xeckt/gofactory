package cmd

import (
	"fmt"
	"strings"

	"github.com/alchemicalkube/gofactory/api"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

var privilegeFlag string

var validPrivilegeFlags = []string{
	api.NOT_AUTHENTICATED_PRIVILEGE,
	api.CLIENT_PRIVILEGE,
	api.ADMINISTRATOR_PRIVILEGE,
	api.INITIAL_ADMIN_PRIVILEGE,
	api.API_TOKEN_PRIVILEGE,
}

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "command to specify the type of login you wish to do",
	Run: func(cmd *cobra.Command, args []string) {
		switch strings.ToLower(privilegeFlag) {
		case "notauthenticated":
			privilegeFlag = api.NOT_AUTHENTICATED_PRIVILEGE
		case "client":
			privilegeFlag = api.CLIENT_PRIVILEGE
		case "administrator":
			privilegeFlag = api.ADMINISTRATOR_PRIVILEGE
		case "initialadmin":
			privilegeFlag = api.INITIAL_ADMIN_PRIVILEGE
		case "apitoken":
			privilegeFlag = api.API_TOKEN_PRIVILEGE
		default:
			Logger.Fatal("Unknown privilege type", Logger.Args(
				"specified:", privilegeFlag,
				"expected:", validPrivilegeFlags,
			))
		}
	},
}

var passwordlessSubCmd = &cobra.Command{
	Use:   "passwordless",
	Short: "authenticate with the dedicated server without password",
	Long:  "returns the new token retrieved from the passwordless privilege to stdout",
	Run: func(cmd *cobra.Command, args []string) {
		Logger.Trace("passwordless", Logger.Args(
			"client object", client,
			"client pointer", &client,
			"client privilege", privilegeFlag,
		))
		passwordlessLogin(privilegeFlag)
	},
}

func passwordlessLogin(privilege string) {
	err := client.PasswordlessLogin(ctx, privilege)
	if err != nil {
		Logger.Fatal("error with passwordless login", Logger.Args("error", err))
	}

	if len(client.Token) == 0 {
		Logger.Fatal("api returned an empty token. are you sure it is not claimed or no client protection password is enabled?")
	}

	Logger.AppendKeyStyle("warning", *pterm.NewStyle(pterm.BgYellow))

	Logger.Info("server response success", Logger.Args(
		"privilege", privilegeFlag,
		"new token", client.Token,
		"warning", pterm.NewStyle(pterm.FgWhite, pterm.BgYellow).
			Sprint(fmt.Sprintf("if you wish to use this with gofactory-cli, replace your %s environment variable", ENV_GF_TOKEN)),
	))
}

var passwordFlag string

var passwordSubCmd = &cobra.Command{
	Use:   "password",
	Short: "authenticate with the dedicated server with a password",
	Run: func(cmd *cobra.Command, args []string) {
		if passwordFlag == "" {
			Logger.Fatal("You must provide a password")
		}
		passwordLogin(passwordFlag, privilegeFlag)
	},
}

func passwordLogin(password string, privilege string) {
	Logger.Trace("password", Logger.Args(
		"client object", client,
		"client pointer", &client,
		"client privilege", privilegeFlag,
		"password", passwordFlag,
	))

	err := client.PasswordLogin(ctx, privilege, password)
	if err != nil {
		Logger.Fatal("password command error", Logger.Args("error", err))
	}

	if client.Token == "" {
		Logger.Fatal("api returned an empty token. is your password correct?")
	}

	Logger.AppendKeyStyle("warning", *pterm.NewStyle(pterm.BgYellow))

	Logger.Info("server response success", Logger.Args(
		"privilege", privilegeFlag,
		"new token", client.Token,
		"warning", pterm.NewStyle(pterm.FgWhite, pterm.BgYellow).
			Sprint(fmt.Sprintf("if you wish to use this with gofactory-cli, replace your %s environment variable", ENV_GF_TOKEN)),
	))
}

func init() {
	Root.AddCommand(loginCmd)

	loginCmd.PersistentFlags().StringVarP(&privilegeFlag, "privilege", "p", "", "privilege to use")
	loginCmd.MarkPersistentFlagRequired("privilege")

	passwordSubCmd.Flags().StringVarP(&passwordFlag, "password", "s", "", "password to authenticate with")
	passwordSubCmd.MarkFlagRequired("password")

	loginCmd.AddCommand(passwordlessSubCmd)
	loginCmd.AddCommand(passwordSubCmd)
}

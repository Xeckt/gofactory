package cmd

import (
	"strings"

	"github.com/alchemicalkube/gofactory/api"
	"github.com/rs/zerolog/log"
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
}

var passwordlessSubCmd = &cobra.Command{
	Use:   "passwordless",
	Short: "authenticate with the dedicated server without password",
	Long:  "returns the new token retrieved from the passwordless privilege to stdout",
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
			log.Fatal().Msgf("Unknown privilege %s, please use one of %v", privilegeFlag, validPrivilegeFlags)
		}

		err := client.PasswordlessLogin(ctx, privilegeFlag)
		if err != nil {
			log.Fatal().Err(err)
		}

		if len(client.Token) == 0 {
			log.Fatal().Msg("token returned is empty. Are you sure it is not claimed or no client protection password is enabled?")
		}

		log.Info().Msgf("Successfully authenticated with privilege: %s", privilegeFlag)
		log.Info().Msgf("Token returned: %s", client.Token)
		log.Warn().Msgf("If you wish to use this token, make sure to replace your %s environment variable!", ENV_TOKEN)
	},
}

var passwordFlag string

var passwordSubCmd = &cobra.Command{
	Use:   "password",
	Short: "authenticate with the dedicated server with a password",
	Run: func(cmd *cobra.Command, args []string) {
		if passwordFlag == "" {
			log.Fatal().Msgf("You must provide a password")
		}
		err := client.PasswordLogin(ctx, privilegeFlag, passwordFlag)
		if err != nil {
			log.Fatal().Err(err)
		}
		if client.Token == "" {
			log.Fatal().Msgf("token returned is empty. are you sure the information you provided was correct?")
		}
		log.Info().Msgf("Successfully authenticated with privilege: %s", privilegeFlag)
		log.Info().Msgf("Token returned: %s", client.Token)
		log.Warn().Msgf("If you wish to use this token, make sure to replace your %s environment variable!", ENV_TOKEN)
	},
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

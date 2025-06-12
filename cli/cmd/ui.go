package cmd

import (
	"github.com/alchemicalkube/gofactory/api"
	"github.com/pterm/pterm"
)

var (
	selectMenu = pterm.InteractiveSelectPrinter{
		TextStyle:     pterm.NewStyle(pterm.FgLightCyan),
		DefaultText:   "select a command to run",
		Options:       []string{"query server", "health check", "login"},
		OptionStyle:   pterm.NewStyle(pterm.FgCyan),
		DefaultOption: "",
		MaxHeight:     5,
		Selector:      "➜",
		SelectorStyle: pterm.NewStyle(pterm.FgCyan),
		Filter:        true,
	}

	loginMenu = pterm.InteractiveSelectPrinter{
		TextStyle:     pterm.NewStyle(pterm.FgLightCyan),
		DefaultText:   "select the type of login to try",
		Options:       []string{"password", "passwordless"},
		OptionStyle:   pterm.NewStyle(pterm.FgCyan),
		DefaultOption: "",
		MaxHeight:     5,
		Selector:      "➜",
		SelectorStyle: pterm.NewStyle(pterm.FgCyan),
		Filter:        true,
	}

	privilegeMenu = pterm.InteractiveSelectPrinter{
		TextStyle:   pterm.NewStyle(pterm.FgLightCyan),
		DefaultText: "select a privilege",
		Options: []string{
			api.NOT_AUTHENTICATED_PRIVILEGE,
			api.INITIAL_ADMIN_PRIVILEGE,
			api.CLIENT_PRIVILEGE,
			api.ADMINISTRATOR_PRIVILEGE,
			api.API_TOKEN_PRIVILEGE,
		},
		OptionStyle:   pterm.NewStyle(pterm.FgCyan),
		DefaultOption: "",
		MaxHeight:     5,
		Selector:      "➜",
		SelectorStyle: pterm.NewStyle(pterm.FgCyan),
		Filter:        true,
	}
)

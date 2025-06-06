module github.com/alchemicalkube/gofactory/cli

go 1.24

require (
	github.com/alchemicalkube/gofactory/api v1.0.0
	github.com/rs/zerolog v1.34.0
	github.com/spf13/cobra v1.9.1
)

require (
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/mattn/go-colorable v0.1.14 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/spf13/pflag v1.0.6 // indirect
	golang.org/x/sys v0.33.0 // indirect
)

replace github.com/alchemicalkube/gofactory/api => ../api

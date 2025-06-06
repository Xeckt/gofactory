module github.com/alchemicalkube/gofactory/cli

go 1.24

require (
	github.com/alchemicalkube/gofactory/api v1.0.0
	github.com/alecthomas/kong v1.11.0
	github.com/rs/zerolog v1.34.0
)

require (
	github.com/goccy/go-yaml v1.18.0 // indirect
	github.com/mattn/go-colorable v0.1.14 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	golang.org/x/sys v0.33.0 // indirect
)

replace github.com/alchemicalkube/gofactory/api => ../api

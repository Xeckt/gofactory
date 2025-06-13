# GoFactory

<img src="./logo.png" width=250 height =250 alt="_">

GoFactory is a cross-platform CLI tool bundled with 
an API client library to interact with your Satisfactory dedicated server HTTPS and Lightquery API.

## Current state

### API Client Library
- [x] Full integration with the Satisfactory v1 HTTPS API
- [ ] Full integration with the Satisfactory Lightquery API (In progress, almost finished)
- [x] Goroutine compatible (contexts)
- [x] Convenient helper functions
- [x] Strongly typed

The v1.0.0 API library has just been released. See: [API MD file](./API.md)

### CLI
- [ ] Full TUI 
- [ ] Support for all API operations
- [ ] Server specific features
  - Non-API export/import saves
  - export/import blueprints
  - export/import configs
  - log reading

The cli tool will not manage mods because the [ficsit-cli](https://github.com/satisfactorymodding/ficsit-cli) tool is already available.

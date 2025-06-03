# GoFactory

<img src="./logo.png" width=250 height =250 alt="_">

GoFactory is a (planned) cross-platform CLI tool to interact with your Satisfactory dedicated server, locally or remotely via API.

The project will be modularised into two parts - the API client library, and the CLI tool itself. 

So, if you wish to take the API library and use it for yourself, you can do so.

## Roadmap

### CLI

- [ ] Full CLI GUI
- [ ] Support for all API operations
- [ ] Various support for handling secrets

### API Library 
- [x] Full integration with the Satisfactory v1 API
- [x] Goroutine compatible (contexts)
- [x] Convenient helper functions
- [ ] Rate limiting (To be discussed)
- [x] Strongly typed

## Current state

The API library is ready and usable, though it is still in development and hasn't released in a major version yet so
use at your own risk:
```bash
go get github.com/xeckt/gofactory/api
```
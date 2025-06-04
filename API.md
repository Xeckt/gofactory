# GoFactory API Client Library

## Downloading

First download the latest version of the module into your project:
```bash
go get github.com/alchemicalkube/gofactory/api
```
---

## Usage

There are various ways to use the library, and you can utilise a bunch of functionality since there is full integration
with the Satisfactory v1 API. Let's say we want to claim an unclaimed Satisfactory dedicated server, we must first
obtain an `InitialAdmin` privileged token:

```go
func main() {
    const token = ""
    const url = "https://dedicatedserver.co.uk:7777"
	
    client := api.NewGoFactoryClient(url, token, true)
    err := client.PasswordlessLogin(context.Background(), api.INITIAL_ADMIN_PRIVILEGE)
    if err != nil {
        log.Fatal(err)
    }
	
    fmt.Println(client.Token)
}
```
> [!WARNING]
> There is a caveat to pay attention to here. The library operates directly on the pointed client struct when it comes to important
information in the struct such as the `Token`. So once this is called, it will update accordingly directly on your struct
object.

Let's proceed by now claiming the server. Claiming also provides a new, permanent token, so the old `InitialAdmin` permissive
token is discarded and replaced with the new, privileged `Administrator` token in the defined object:

```go
func main() {
    const token = ""
    const url = "https://dedicatedserver.co.uk:7777"

    client := api.NewGoFactoryClient(url, token, true)
	
    err := client.PasswordlessLogin(context.Background(), api.INITIAL_ADMIN_PRIVILEGE)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(client.Token)

    err = client.ClaimServer(context.Background(), api.ClaimRequestData{
        ServerName:    "DedicatedServerName",
        AdminPassword: "YourNewAdminPassword",
    })

    fmt.Println(client.Token)

    if err != nil {
        log.Fatal(err)
    }
}
```

The privilege check is done in the background by the `ClaimServer` function, since the `PasswordlessLogin` function is called
before, it automatically updated the unexported `currentPrivilege` field and then checked by `ClaimServer` to ensure it was
done correctly.

> [!IMPORTANT]
> Although combining `PasswordlessLogin` and `ClaimServer` into a single function for convenience is possible,
GoFactory's primary goal is to provide developers with maximum access and flexibility, without limitations for
additional capabilities.


All functions are exported, so if you are not a fond of this behaviour you can take advantage of the helper functions:

https://github.com/alchemicalkube/gofactory/blob/982206d8f8305bc9bea67a0e12347455350a43bc/api/client.go#L35-L112


See here for an example of how they are used inside the library

https://github.com/alchemicalkube/gofactory/blob/982206d8f8305bc9bea67a0e12347455350a43bc/api/rename.go#L1-L41


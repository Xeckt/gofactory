# Contributing to [GoFactory](github.com/xeckt/gofactory)

Thank you for your interest in contributing! 

[GoFactory](github.com/xeckt/gofactory) has two components to it's source. One is the 
[API Client Library](github.com/xeckt/gofactory/api) and the other is the [CLI tool](github.com/xeckt/gofactory/cli)

You will need to clone the whole repo even if you plan to only work on one, this is to consolidate everything in one place
and make it as available as possible for users & developers alike.

---

## Getting Started

1. **Fork the repository**

2. **Clone your fork**
   ```bash
   git clone https://github.com/YOUR_USERNAME/REPO_NAME.git
   cd REPO_NAME
   ```
3. **Create a new branch**


## Code Guidelines
### Formatting
Use `go fmt` to ensure consistent formatting.

### Lint your code
Use `golangci-lint` or `go vet` to catch potential issues

### Code style
Code style here is not too strict, I like to allow as much room for anything here.

Keep code focused. Function names can be long if the name communicates necessarily
and there are no other alternatives

Add tests if you’re adding or changing features, this is a must to ensure tests always pass.

### Comitting
Prefix your commits with anything related, such as:

Features:
- `feature:`
- `feat`

Fixes:
- `fix:`


```
feat: add new API client for handling rate limits
fix: correct typo in README.md
docs: update contribution guidelines
```

More guidelines may be added here in the future so be sure to check back.
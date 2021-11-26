# td [![Go Reference](https://img.shields.io/badge/go-pkg-00ADD8)](https://pkg.go.dev/github.com/gotd/td#section-documentation) [![codecov](https://img.shields.io/codecov/c/github/gotd/td?label=cover)](https://codecov.io/gh/gotd/td) [![alpha](https://img.shields.io/badge/-alpha-orange)](https://go-faster.org/docs/projects/status#alpha)

A fast Telegram client in pure Go.

* [Examples](examples)
* [Security policy](.github/SECURITY.md)
* [User support and dev chat](.github/SUPPORT.md)
* [Roadmap](ROADMAP.md)
* [Contributing](CONTRIBUTING.md)
* [Architecture](ARCHITECTURE.md)
* [Generated Go Documentation](https://ref.gotd.dev/pkg/github.com/gotd/td/tg.html)

Before using this library, read [How To Not Get Banned](.github/SUPPORT.md#how-to-not-get-banned) guide.
Due to limitations of `pkg.go.dev`, documentation for `tg` package is not shown, but there is [hosted version](https://ref.gotd.dev/pkg/github.com/gotd/td/tg.html).

## Usage

```console
go get github.com/gotd/td
```

```go
package main

import (
	"context"

	"github.com/gotd/td/telegram"
)

func main() {
	// https://core.telegram.org/api/obtaining_api_id
	client := telegram.NewClient(appID, appHash, telegram.Options{})
	if err := client.Run(context.Background(), func(ctx context.Context) error {
		// It is only valid to use client while this function is not returned
		// and ctx is not cancelled.
		api := client.API()

		// Now you can invoke MTProto RPC requests by calling the API.
		// ...

		// Return to close client connection and free up resources.
		return nil
	}); err != nil {
		panic(err)
	}
	// Client is closed.
}
```

See [examples](examples) for more info.

## Status

Work is still in progress (mostly helpers and convenience wrappers), but basic functionality were tested in production and works fine.
Only go1.16 is supported and no backward compatibility is provided for now.

The goal of this project is to implement a stable, performant and safe client for Telegram in pure Go while
having a simple and convenient API and a feature parity with TDLib.

This project is fully non-commercial and not affiliated with any commercial organization
(including Telegram LLC).

Also, see the [comparison](#difference-to-other-projects) with other Go Telegram clients.

## Features

* Low memory overhead, 150kb per idle client
* Full MTProto 2.0 implementation in Golang
* Code for Telegram types generated by `./cmd/gotdgen` (based on [gotd/tl](https://github.com/gotd/tl) parser) with embedded [official documentation](https://core.telegram.org/schema)
* Pluggable session storage
* Automatic re-connects with keepalive
* Vendored Telegram public keys that are kept up-to-date
* Rigorously tested
  * End-to-end with real Telegram server in CI
  * End-to-end with gotd Telegram server (in pure Go)
  * Lots of unit testing
  * Fuzzing
  * 24/7 canary bot in production that tests reconnects, update handling, memory leaks and performance
* No runtime reflection overhead
* [Conforms](https://github.com/gotd/td/issues/155) to [Security guidelines](https://core.telegram.org/mtproto/security_guidelines) for Telegram client software developers
  * Secure PRNG used for crypto
  * Replay attack protection
* 2FA support
* MTProxy support
* Various helpers that lighten the complexity of the Telegram API
  * [uploads](https://pkg.go.dev/github.com/gotd/td/telegram/uploader) for big and small files with multiple streams for single file and progress reporting
  * [downloads](https://pkg.go.dev/github.com/gotd/td/telegram/downloader) with CDN support, also multiple streams
  * [messages](https://pkg.go.dev/github.com/gotd/td/telegram/message) with various convenience builders and text styling support
  * [query](https://pkg.go.dev/github.com/gotd/td/telegram/query) with pagination helpers
  * [middleware](https://pkg.go.dev/github.com/gotd/td/telegram#Middleware) for [rate limiting](https://pkg.go.dev/github.com/gotd/contrib/middleware/ratelimit) and [FLOOD_WAIT handling](https://pkg.go.dev/github.com/gotd/contrib/middleware/floodwait)
* Connection pooling
* Automatic datacenter migration and redirects handling
* Graceful [request cancellation](https://core.telegram.org/mtproto/service_messages#cancellation-of-an-rpc-query) via context
* WebSocket transport support (works in WASM)

## Examples

See [examples](examples) directory.

Also take a look at

* [gotd/bot](https://github.com/gotd/bot) with updates recovery enabled, used as canary for stability testing
* [gotd/cli](https://github.com/gotd/cli), command line interface for subset of telegram methods.

### Auth

#### User

You can use `td/telegram/auth.Flow` to simplify user authentications.

```go
codePrompt := func(ctx context.Context, sentCode *tg.AuthSentCode) (string, error) {
    // NB: Use "golang.org/x/crypto/ssh/terminal" to prompt password.
    fmt.Print("Enter code: ")
    code, err := bufio.NewReader(os.Stdin).ReadString('\n')
    if err != nil {
        return "", err
    }
    return strings.TrimSpace(code), nil
}
// This will setup and perform authentication flow.
// If account does not require 2FA password, use telegram.CodeOnlyAuth
// instead of telegram.ConstantAuth.
if err := auth.NewFlow(
    auth.Constant(phone, password, auth.CodeAuthenticatorFunc(codePrompt)),
    auth.SendCodeOptions{},
).Run(ctx, client.Auth()); err != nil {
    panic(err)
}
```

#### Bot

Use bot token from [@BotFather](https://telegram.me/BotFather).

```go
if err := client.Auth().Bot(ctx, "token:12345"); err != nil {
    panic(err)
}
```

### Calling MTProto directly

You can use the generated `tg.Client` that allows calling any MTProto method
directly.

```go
// Grab these from https://my.telegram.org/apps.
// Never share it or hardcode!
client := telegram.NewClient(appID, appHash, telegram.Options{})
client.Run(ctx, func(ctx context.Context) error) {
  // Grab token from @BotFather.
  if err := client.Auth().Bot(ctx, "token:12345"); err != nil {
    return err
  }
  state, err := client.API().UpdatesGetState(ctx)
  if err != nil {
    return err
  }
  // Got state: &{Pts:197 Qts:0 Date:1606855030 Seq:1 UnreadCount:106}
  // This will close client and cleanup resources.
  return nil
}
```

### Generated code

The code output of `gotdgen` contains references to TL types, examples, URL to
official documentation and [extracted](https://github.com/gotd/getdoc) comments from it.

For example, the [auth.Authorization](https://core.telegram.org/type/auth.Authorization) type in `tg/tl_auth_authorization_gen.go`:

```go
// AuthAuthorizationClass represents auth.Authorization generic type.
//
// See https://core.telegram.org/type/auth.Authorization for reference.
//
// Example:
//  g, err := DecodeAuthAuthorization(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *AuthAuthorization: // auth.authorization#cd050916
//  case *AuthAuthorizationSignUpRequired: // auth.authorizationSignUpRequired#44747e9a
//  default: panic(v)
//  }
type AuthAuthorizationClass interface {
	bin.Encoder
	bin.Decoder
	construct() AuthAuthorizationClass
}
```

Also, the corresponding [auth.signIn](https://core.telegram.org/method/auth.signIn) method:

```go
// AuthSignIn invokes method auth.signIn#bcd51581 returning error if any.
// Signs in a user with a validated phone number.
//
// See https://core.telegram.org/method/auth.signIn for reference.
func (c *Client) AuthSignIn(ctx context.Context, request *AuthSignInRequest) (AuthAuthorizationClass, error) {}
```

The generated constructors contain detailed official documentation, including links:

```go
// FoldersDeleteFolderRequest represents TL type `folders.deleteFolder#1c295881`.
// Delete a peer folder¹
//
// Links:
//  1) https://core.telegram.org/api/folders#peer-folders
//
// See https://core.telegram.org/method/folders.deleteFolder for reference.
type FoldersDeleteFolderRequest struct {
    // Peer folder ID, for more info click here¹
    //
    // Links:
    //  1) https://core.telegram.org/api/folders#peer-folders
    FolderID int
}
```

## Contributions

Huge thanks to all contributors. Dealing with a project of this scale alone is impossible.

Special thanks:

* [tdakkota](https://github.com/tdakkota)
  * Two-factor authentication (SRP)
  * Proxy support
  * Update dispatcher
  * Complete transport support (abridged, padded intermediate and full)
  * Telegram server for end-to-end testing
  * Multiple major refactorings, including critical cryptographical scope reduction
  * Code generation improvements (vector support, multiple modes for pretty-print)
  * And many other cool things and performance improvements
* [xjem](https://github.com/xjem)
  * Background pings
  * Links in generated documentation
  * Message acknowledgements
  * Retries
  * RPC Engine
  * Gap (Updates) engine

## Reference

The MTProto protocol description is [hosted](https://core.telegram.org/mtproto#general-description) by Telegram.

Most important parts for client implementations:

* [Security guidelines](https://core.telegram.org/mtproto/security_guidelines) for client software developers

Current implementation [mostly conforms](https://github.com/gotd/td/issues/155) to security guidelines, but no
formal security audit were performed.

## Prior art

* [Lonami/grammers](https://github.com/Lonami/grammers) (Great Telegram client in Rust, many test vectors were used as reference)
* [sdidyk/mtproto](https://github.com/sdidyk/mtproto), [cjongseok/mtproto](https://github.com/cjongseok/mtproto), [xelaj/mtproto](https://github.com/xelaj/mtproto) (MTProto 1.0 in go)

## Difference to other projects

Status by 11.07.2021

<table>
  <tr>
   <th>Topic</th>
   <th colspan="2">gotd</th>
   <th colspan="2"><a href="https://github.com/xelaj/mtproto">xelaj/mtproto</a></th>
  <tr>
   <td align="center">🚧<br>Maintainers</td>
   <td colspan="2">At least 3 core contributors, 1.7K commits, actively maintained, <a href="https://github.com/gotd/td/pull/409">automatically updated to latest layer</a></td>
   <td colspan="2">Single core contributor, 150 commits, effectively not maintaned from March 2021, <a href="https://github.com/xelaj/mtproto/issues/88">failed to update</a> to 129 layer</td>
  <tr>
   <td align="center">🔒️<br>Security</td>
   <td colspan="2">MTProto v2, <a href="https://github.com/gotd/td/issues/155">conforms</a> to security guidelines</td>
   <td colspan="2">MTProto v1 that is insecure and <a href="https://core.telegram.org/mtproto_v1">deprecated</a> since 2017 (SHA-1 <a href="https://github.com/xelaj/mtproto/blob/216789b95a5d644ebbdd1acb7eb46ff61647960a/internal/aes_ige/ige_cipher.go#L194-L197">is used</a> instead of SHA-256), code was probably copied from <a href="https://github.com/sdidyk/mtproto">sdidyk/mtproto</a> or <a href="https://github.com/cjongseok/mtproto">cjongseok/mtproto</a> with <a href="https://github.com/xelaj/mtproto/blob/216789b95a5d644ebbdd1acb7eb46ff61647960a/internal/aes_ige/ige_cipher.go#L160">little understanding</a></td>
  <tr>
   <td align="center">🚀<br>Features</td>
   <td colspan="2">MTProxy support, WebSocket transport, helpers for uploads, downloads, messages, text styling, datacenter redirects handling and more, but you can still use the raw API</td>
   <td colspan="2">Raw MTProto API</td>
  <tr>
   <td align="center">🔨<br>Platforms</td>
   <td colspan="2">All platforms including WASM and special stability features for Windows that considers clock resolution</td>
   <td colspan="2">Limited support for Windows</td>
  <tr>
   <td align="center">⚡️<br>Performance</td>
   <td colspan="2">Multiple optimizations for zero allocations</td>
   <td colspan="2">Uses reflection in runtime</td>
  <tr>
   <td align="center">🧪<br>Stability</td>
   <td colspan="2">Lots of unit tests (237 tests, 943 including subtests), end-to-end tests with self-made Telegram server in Go, end-to-end tests with real test servers, fuzzing</td>
   <td colspan="2">12 unit tests, 41 including sub-tests</td>
  <tr>
   <td align="center">⭐<br>GitHub Stargazers</td>
   <td align="center" width="20%">Has only 192 stars</td>
   <td align="center" width="20%">
    <img alt="GitHub stars" src="https://img.shields.io/github/stars/gotd/td?style=for-the-badge">
    <br>
    <img alt="GitHub watchers" src="https://img.shields.io/github/watchers/gotd/td?style=for-the-badge">
   </td>
   <td align="center" width="20%">Has 643 stars, which is much higher</td>
   <td align="center" width="20%">
    <img alt="GitHub stars" src="https://img.shields.io/github/stars/xelaj/mtproto?style=for-the-badge">
    <br>
    <img alt="GitHub watchers" src="https://img.shields.io/github/watchers/xelaj/mtproto?style=for-the-badge">
   </td>
</table>

## Who is using gotd?

[<img src="https://user-images.githubusercontent.com/43930873/142855897-7091ced0-4fe8-4f8d-ad43-e9db2723bacc.png" width="150">](https://telq.org)

Drop a comment [here](https://github.com/gotd/td/issues/568) to add your project.

## License

MIT License

Created by Aleksandr (ernado) Razumov

2020

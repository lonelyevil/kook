# Kook

A low level library for interacting with Kook bot API. Currently, it is WIP.

[![Go Reference](https://pkg.go.dev/badge/github.com/lonelyevil/kook.svg)](https://pkg.go.dev/github.com/lonelyevil/kook)
[![Go Report Card](https://goreportcard.com/badge/github.com/lonelyevil/kook)](https://goreportcard.com/report/github.com/lonelyevil/kook)
[![Server Badger](https://img.shields.io/badge/kookapp-dev--chat-informational)](https://kaihei.co/r5s1WO)
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Flonelyevil%2Fkook.svg?type=shield)](https://app.fossa.com/projects/git%2Bgithub.com%2Flonelyevil%2Fkook?ref=badge_shield)

## Get Started

It is not recommended to use it in production until it releases `v1.0.0`.

### Installing

This assumes that you already have a working Go environment.

```go get github.com/lonelyevil/kook```

Other than the library itself, you need to install a logger and an adapter for logger.
Personally, I only implement the adapter for [phuslu/log](https://github.com/phuslu/log). So, in order to use the library, the following installation is necessary.

```go get github.com/lonelyevil/kook/log_adapter/plog```

### Usage

See the examples in `examples` folder.

## Features and Roadmap

See [features.md](features.md).

## Documentation

WIP.

For code that are not well commented, users could refer to [bwmarrin/discordgo](https://github.com/bwmarrin/discordgo) as this library is heavily influenced by it.

## Versioning

I could not guarantee a stable API until the release of `1.0.0`.
Before that, any break change would happen.
After that, It would follow [semantic versioning](https://semver.org/).

## Contributing

Currently, this repo does not accept any large PRs, as the API may be broken at any time when I want.
Bug reports and code suggestions are greatly welcomed.

## LICENSE

kook is a free and open source software library distributed under the terms of [ISC License](LICENSE).


[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Flonelyevil%2Fkook.svg?type=large)](https://app.fossa.com/projects/git%2Bgithub.com%2Flonelyevil%2Fkook?ref=badge_large)

## Special Thanks To:

<img alt="GoLand logo" src="https://resources.jetbrains.com/storage/products/company/brand/logos/GoLand.png" width="250">

Built using IntelliJ IDEA
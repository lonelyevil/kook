# KHL

A low level library for interacting with kaiheila(开黑啦) bot API. Currently, it is WIP.

[![Go Reference](https://pkg.go.dev/badge/github.com/lonelyevil/khl.svg)](https://pkg.go.dev/github.com/lonelyevil/khl)
[![Go Report Card](https://goreportcard.com/badge/github.com/lonelyevil/khl)](https://goreportcard.com/report/github.com/lonelyevil/khl)
[![Server Badger](https://img.shields.io/badge/kaiheila-dev--chat-informational)](https://kaihei.co/r5s1WO)

## Get Started

It is not recommended to use it in production until it releases `v1.0.0`.

### Installing

This assumes that you already have a working Go environment.

```go get github.com/lonelyevil/khl```

Other than the library itself, you need to install a logger and an adapter for logger.
Personally, I only implement the adapter for [phuslu/log](https://github.com/phuslu/log). So, in order to use the library, the following installation is necessary.

```go get github.com/lonelyevil/khl/log_adapter/plog```

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

KHL is a free and open source software library distributed under the terms of [ISC License](LICENSE).

## Special Thanks To:

<img alt="GoLand logo" src="https://resources.jetbrains.com/storage/products/company/brand/logos/GoLand.png" width="250">

Built using IntelliJ IDEA
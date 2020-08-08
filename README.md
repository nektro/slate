# Slate

![loc](https://sloc.xyz/github/nektro/slate)
[![license](https://img.shields.io/github/license/nektro/slate.svg)](https://github.com/nektro/slate/blob/master/LICENSE)
[![discord](https://img.shields.io/discord/551971034593755159.svg?logo=discord)](https://discord.gg/P6Y4zQC)
[![paypal](https://img.shields.io/badge/donate-paypal-009cdf?logo=paypal)](https://paypal.me/nektro)
[![circleci](https://circleci.com/gh/nektro/slate.svg?style=svg)](https://circleci.com/gh/nektro/slate)
[![release](https://img.shields.io/github/v/release/nektro/slate)](https://github.com/nektro/slate/releases/latest)
[![goreportcard](https://goreportcard.com/badge/github.com/nektro/slate)](https://goreportcard.com/report/github.com/nektro/slate)
[![codefactor](https://www.codefactor.io/repository/github/nektro/slate/badge)](https://www.codefactor.io/repository/github/nektro/slate)
[![downloads](https://img.shields.io/github/downloads/nektro/slate/total.svg)](https://github.com/nektro/slate/releases)
[![crowdin](https://badges.crowdin.net/slate/localized.svg)](https://crowdin.com/project/slate)
[![issuehunt](https://img.shields.io/badge/issuehunt-slate-38d39f)](https://issuehunt.io/r/nektro/slate)
[![docker_pulls](https://img.shields.io/docker/pulls/nektro/slate)](https://hub.docker.com/r/nektro/slate)
[![docker_stars](https://img.shields.io/docker/stars/nektro/slate)](https://hub.docker.com/r/nektro/slate)

A new programming langauge made by me. Used as a tool to learn about compilers.

## Using

Dependencies
- Golang 1.14
- LLVM 10

Running `./start_local.sh` will build and run Slate on the x86_64 Hello World. Feel free to edit either file.
```
go build
./slate -run {path/to/src_file.slate}
llvm-as-10 out.ll
lli-10 out.bc
```

## Inspirations
- https://github.com/golang/go
- https://github.com/ziglang/zig
- https://github.com/denoland/deno

## Contributing
Issues and Pull Requests welcome. As the project progresses I will make a list of focus areas.

## Contact
- hello@nektro.net
- https://twitter.com/nektro

## License
MIT

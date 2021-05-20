# Slate

![loc](https://sloc.xyz/github/nektro/slate)
[![license](https://img.shields.io/github/license/nektro/slate.svg)](https://github.com/nektro/slate/blob/master/LICENSE)
[![discord](https://img.shields.io/discord/551971034593755159.svg?logo=discord)](https://discord.gg/P6Y4zQC)
[![goreportcard](https://goreportcard.com/badge/github.com/nektro/slate)](https://goreportcard.com/report/github.com/nektro/slate)
[![issuehunt](https://img.shields.io/badge/issuehunt-slate-38d39f)](https://issuehunt.io/r/nektro/slate)

A new programming language made by me. Used as a tool to learn about compilers.

> Note: some code snippets and examples may only work on arm64 Linux since that's what I have.

## Built With

- Golang 1.14
- LLVM 11

## Usage

Running `./start_local.sh` will build and run Slate on the arm64 Linux Hello World. Feel free to edit either file.

```
go build
./slate -run {path/to/src_file.slate}
llvm-as-11 out.ll
lli-11 out.bc
```

More code examples can be found in the [`./tests/`](./tests/) directory.

## Inspirations

- https://github.com/golang/go
- https://github.com/ziglang/zig
- https://github.com/denoland/deno
- https://github.com/rust-lang/rust

## Contributing

Issues welcome. As the project progresses I will make a list of focus areas.

## Contact

- hello@nektro.net
- https://twitter.com/nektro

## License

ACSL v1.4. See [LICENSE](LICENSE) for more info.

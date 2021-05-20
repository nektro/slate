# Slate

![loc](https://sloc.xyz/github/nektro/slate)
[![discord](https://img.shields.io/discord/551971034593755159.svg?logo=discord)](https://discord.gg/P6Y4zQC)
[![goreportcard](https://goreportcard.com/badge/github.com/nektro/slate)](https://goreportcard.com/report/github.com/nektro/slate)

A new programming language made by me. Used as a tool to learn about compilers.

> Note: some code snippets and examples may only work on arm64 Linux since that's what I have.

## Built With

- Golang 1.14
- LLVM 11
- https://github.com/llir/llvm
- https://github.com/nektro/go-util

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
- https://www.haskell.org/
- https://www.roc-lang.org/

## Contributing

Issues welcome. As the project progresses I will make a list of focus areas.

## Contact

- hello@nektro.net
- https://twitter.com/nektro

## License

ACSL v1.4. See [LICENSE](LICENSE) for more info.

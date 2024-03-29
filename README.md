# Slate

![loc](https://sloc.xyz/github/nektro/slate)
[![discord](https://img.shields.io/discord/551971034593755159.svg?logo=discord)](https://discord.gg/P6Y4zQC)
[![goreportcard](https://goreportcard.com/badge/github.com/nektro/slate)](https://goreportcard.com/report/github.com/nektro/slate)

A new programming language made by me. Used as a tool to learn about compilers.

> Note: some code snippets and examples may only work on x64 Linux since that's what I have.

## Built With

- Golang 1.14
- LLVM 12
- https://github.com/llir/llvm
- https://github.com/nektro/go-util

## Usage

Running `./start_local.sh` will build and run Slate on the x64 Linux Hello World. Feel free to edit either file.

```
go build
./slate -run {path/to/src_file.slate}
llvm-as-12 out.ll
lli-12 out.bc
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

In regards to pull requests, I appreciate the enthusiasm but at this stage I will be more hesitant than normal about what gets merged from other people. This project is largely an experiment and learning environment for me, so please keep that in mind.

## Contact

- hello@nektro.net
- https://twitter.com/nektro

## License

ACSL v1.4. See [LICENSE](LICENSE) for more info.

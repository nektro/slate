# Slate

![loc](https://tokei.rs/b1/github/nektro/slate)

A new, strongly-typed, fast, and secure programming language made for [WebAssembly](https://webassembly.org/). This repository contains the compiler, standard library, tests, and documentation. Slate syntax is largely inspried by JavaScript with influences also from Java, Kotlin, and more.

> NOTE:
> > The compiler is super-alpha (like 0.0.0.0.0.0.0.1 alpha) and not ready to be used in any sort of regular capacity just yet. Heavy development is still underway and more features are coming soon!

----

With that disclaimer out of the way I'm going to show what a potential program might look like with the currently planned features.

## Ideas
- Type-safe and intelligent. Slate is a safe, strongly typed language that compiles to [WebAssembly](https://webassembly.org/).
- Explicit. Primitive data types are the only thing in the global scope. In order to keep binary size down, all libraries (as well as language included ones) must be explicitly imported to use.
- Modular. Objects and functions are passed around to different files using imports and exports. Values can only be imported if they are exported by the file being imported.
- Meta-programming made easy. Object oriented programming is used to represent the real or virtual world in code. The world is complex but that doesn't mean our code has to. Operator overloading, object extensions, etc, will be added very early on.

## Conributing
Found a bug? Theres a lot of 'em! Post them all over on the [bug tracker](https://github.com/nektro/slate/issues), or if you're brave and found a fix yourself, definitely submit a [pull request](https://github.com/nektro/slate/pulls)!

## License
This project is licensed under the Apache 2.0 License. See the [LICENSE](LICENSE) file for details.

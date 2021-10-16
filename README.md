# df-rs-gen &mdash; [`Dragonfly Rust Generator`](https://github.com/df-mc/dragonfly/)

### Using the powers of Rust, Go and Dragonfly to make a vanilla-like world generation.
![Showcase](https://elon.is-from.space/r/kusshugr69a.jpg "Showcase")

## How to use
Clone the repo.
> `git clone https://github.com/Lucaskyy/df-rs-gen.git`

Download the latest DLL required. (or build it yourself, see below)
> Download [here!](https://nightly.link/Lucaskyy/df-rs-gen/workflows/build/master/rustgen-lib.zip)

Build the application. (or use it as a library, see below)
> `go build .`

Run the application. Make sure the application and DLL are in the same path, next to each other. Otherwise you will get an error!
> `df-rs-gen.exe`

### How to build the DLL yourself
To build the DLL yourself, enter the `./lib/rustgen` directory and run:
> `cargo build --release`

You will now find the DLL in the `./lib/rustgen/target/release` directory.

### How to use as a library

To use this as a library, use `go get` to get df-rs-gen.
> `go get github.com/Lucaskyy/df-rs-gen`

Then, somewhere in your code, set the generator to `NewRustGen`<br>
See `main.go` for an example.
> ```go
> srv.World().Generator(NewRustGen())
> ```

### Credit
Thanks to everyone developing Dragonfly, and [hansihe's voxel_worldgen](https://github.com/hansihe/voxel_worldgen) for making this possible.

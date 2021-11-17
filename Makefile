buildings:
	GOOS=linux   GOARCH=amd64 go build -o bin/linux/buildings/barracks       cmd/buildings/barracks.go
	GOOS=windows GOARCH=amd64 go build -o bin/windows/buildings/barracks.exe cmd/buildings/barracks.go

units:
	GOOS=linux   GOARCH=amd64 go build                -o bin/linux/units/conscript                   cmd/units/conscript.go
	GOOS=windows GOARCH=amd64 go build                -o bin/windows/units/conscript.exe             cmd/units/conscript.go
	GOOS=windows GOARCH=amd64 go build -buildmode=pie -o bin/windows/backedin/conscript_backedin.exe cmd/backedin/conscript.go

all: units buildings

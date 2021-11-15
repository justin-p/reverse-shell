buildings:
	GOOS=linux   GOARCH=amd64 go build -o bin/buildings/barracks       construction_yard/buildings/barracks.go
	GOOS=windows GOARCH=amd64 go build -o bin/buildings/barracks.exe construction_yard/buildings/barracks.go

units:
	GOOS=linux   GOARCH=amd64 go build -o bin/units/conscript       construction_yard/units/conscript.go
	GOOS=windows GOARCH=amd64 go build -o bin/units/conscript.exe construction_yard/units/conscript.go

all: units buildings
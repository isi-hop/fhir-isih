package main

import (
	"bufio"
	"fmt"
	"io"
	"os/exec"
)

func main() {
	running("/bin/sh","./r.d")

}

/******************************
	Lancer une commande...
*******************************/
func running(commande string, arguments string) {
	cmd := exec.Command(commande, arguments)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		panic(err)
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		panic(err)
	}
	err = cmd.Start()
	if err != nil {
		panic(err)
	}

	go copyOutput(stdout)
	go copyOutput(stderr)
	cmd.Wait()
}

func copyOutput(r io.Reader) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

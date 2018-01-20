package main

import (
	"log"
	"os"
	"os/exec"
)

func getTerminalWidth() {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	println(out)
}

func printBoldLine() {

}

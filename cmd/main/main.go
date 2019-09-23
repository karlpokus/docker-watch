package main

import (
	"fmt"
	"os"

	"dockerw"
)

func main() {
	p, err := dockerw.Args(os.Args)
	if err != nil {
		fmt.Printf("err: %s\n", err)
		return
	}
	if p.Cmd == "noop" {
		fmt.Println("usage: <placeholder>")
		return
	}
	if p.Cmd == "metadata" {
		fmt.Println(dockerw.Metadata())
		return
	}
	if p.Cmd == "version" {
		fmt.Println(dockerw.Version)
		return
	}
	fmt.Printf("%+v\n", p) // debug
}

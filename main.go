// main.go
package main

import "fmt"

var (
	version = "default"
	commit  = "none"
	date    = ""
	builtBy = ""
)

func main() {
	fmt.Printf("Ba dum, tss!, version=%v, commit=%v date=%v builtBy=%v")
}

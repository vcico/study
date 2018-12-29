package main

import (
	"fmt"
	"os/exec"
)

func main() {

	cmd0 := exec.Command("dir")
	if err := cmd0.Start(); err != nil {
		fmt.Printf("Error: The command No.0 can not be startup %s\n", err)
		return
	}

}

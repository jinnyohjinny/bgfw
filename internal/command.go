package internal

import (
	"os/exec"
)

func Command(cmd string) {

	command := exec.Command("bash", "-c", cmd)

	err := command.Run()
	if err != nil {
		panic(err)
	}

}

package util

import (
	"fmt"
	"os/exec"
)

func cmd(string)string{
	cmd := exec.Command("ls", "-lah")
	err := cmd.Run()
	data, err2 := cmd.Output()
	if err != nil || err2 != nil {
		return fmt.Sprintf("failed to call cmd.Run(): %v", err)
	}
	return string(data)
}

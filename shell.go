package main

import (
	"github.com/abdfnx/gosh"
	"log"
	"runtime"
)

// `Run` executes the same command for shell and powershell
func Run(cmd string) {
	err, out, errout := gosh.ShellOutput("")

	if runtime.GOOS == "windows" {
		err, out, errout = gosh.PowershellOutput(cmd)
	} else {
		err, out, errout = gosh.ShellOutput(cmd)
	}

	if err != nil {
		log.Printf("error: %v\n", err)
		log.Printf("stderr: %s\n", errout)
	}
	log.Printf("error: %v\n", err)
	log.Printf("stdout: %s\n", out)
}

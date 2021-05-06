package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	// get args
	args := os.Args[1:]

	// validate
	if len(args) != 1 {
		log.Panic(errors.New("usage: ./new <problem id>"))
	}

	// get problem id
	problemId := args[0]

	// list solutions
	out, err := exec.Command("ls", "./solutions").Output()

	if err != nil {
		log.Panic(err)
	}

	// go thru each file
	for _, str := range strings.Split(string(out), "\n") {
		// check if this problem already exists
		if str == problemId {
			log.Panic(errors.New("problem already exists"))
		}
	}

	// create new folder under /solutions for this file
	exec.Command("mkdir", "-p", "./solutions/"+problemId).Run()

	// path to go source file
	path := fmt.Sprintf("./solutions/%s/main.go", problemId)

	// create file path
	f, err := os.Create(path)

	// validate
	if err != nil {
		log.Panic(err)
	}

	defer f.Close()

	// write boilerplate
	f.WriteString("package main")

	// stage + commit new code
	exec.Command("git", "add", "-A").Run()
	exec.Command("git", "commit", "-m", fmt.Sprintf("create problem %s", problemId)).Run()
}

package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	// get args
	args := os.Args[1:]

	// validate
	if len(args) != 1 {
		fmt.Println("err: usage: ./new <problem id>")
		return
	}

	// get problem id
	problemId := args[0]

	// create new folder under /solutions for this file
	exec.Command("mkdir", "-p", "./solutions/"+problemId).Run()

	// path to go source file
	path := fmt.Sprintf("./solutions/%s/main.go", problemId)

	// create file path
	f, err := os.Create(path)

	// validate
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer f.Close()

	// write boilerplate
	f.WriteString("package main")

	// stage + commit new code
	exec.Command("git", "add", "-A").Run()
	exec.Command("git", "commit", "-m", fmt.Sprintf("\"create problem %s\"", problemId)).Run()
}

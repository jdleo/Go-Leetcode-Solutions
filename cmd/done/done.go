package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	// get args
	args := os.Args[1:]

	// validate
	if len(args) != 1 {
		fmt.Println("err: usage: ./done <problem id>")
		return
	}

	// get problem id
	problemId := args[0]

	// path to go source file
	path := fmt.Sprintf("./solutions/%s/main.go", problemId)

	// create file path
	f, err := os.Create("README.md")

	// validate
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// close file
	defer f.Close()

	// get current problems solved
	out, err2 := exec.Command("ls", "-l", "./solutions").Output()

	// get count of solved problems
	count := strings.Count(string(out), "\n")

	// validate
	if err2 != nil {
		fmt.Println(err.Error())
		return
	}

	// boilerplate readme
	readme := `# Leetcode Solutions in Go

This is a personal repository for doing and tracking Leetcode problems in Golang. You can use this yourself, but be sure to delete the problems under the %s/solutions%s folder for yourself. Also run %srm -rf .git%s in the root of this project to disassociate my git remote.
		
To compile helper cli's (to quickly start/track new problems):
	
%s
make init
%s 

To start a new problem:	
%s
./new <problem id>
%s 

To mark a problem as complete:	
%s 

./done <problem id>
%s 


### # Of Problems Solved: %d
`

	readme = fmt.Sprintf(readme, "`", "`", "`", "`", "```", "```", "```", "```", "```", "```", count-1)
	f.WriteString(readme)

	// stage + commit new code
	exec.Command("git", "add", path).Run()
	exec.Command("git", "commit", "-m", fmt.Sprintf("\"completed problem %s\"", problemId)).Run()
}

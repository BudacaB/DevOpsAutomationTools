package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"os/exec"
	"runtime"
	"time"
)

func execute() {

	var executeWith string
	var commitMessage string
	var branchName string

	flag.StringVar(&commitMessage, "m", "", "Input commit message (Required)")
	flag.StringVar(&branchName, "b", "", "Input branch name - master or other (Required)")
	flag.Parse()

	if runtime.GOOS == "windows" {
		executeWith = "powershell"
	} else if runtime.GOOS == "linux" || runtime.GOOS == "darwin" {
		executeWith = "bash"
	}

	if commitMessage == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	if branchName == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	outStatus, err := exec.Command(executeWith, "git status").Output()
	outputStatus := string(outStatus[:])
	fmt.Println(outputStatus)

	time.Sleep(1 * time.Second)

	outAdd := exec.Command(executeWith, "git add .")
	err = outAdd.Run()

	time.Sleep(2 * time.Second)

	outCommit, err := exec.Command(executeWith, fmt.Sprintf("git commit -am \"%s\"", commitMessage)).Output()
	outputCommit := string(outCommit[:])
	fmt.Println(outputCommit)

	time.Sleep(2 * time.Second)

	mwriter := io.MultiWriter(os.Stdout)
	outPush := exec.Command(executeWith, fmt.Sprintf("git push origin \"%s\"", branchName))
	outPush.Stderr = mwriter
	outPush.Stdout = mwriter
	err = outPush.Run()
	fmt.Println(outPush.Stderr)

	if err != nil {
		fmt.Printf("%s", err)
	} else {
		fmt.Println("Command Successfully Executed")
	}

}

func main() {
	execute()
}

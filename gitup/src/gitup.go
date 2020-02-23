package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
	"os/exec"
	"time"
	"runtime"
)

func execute() {

	var commitMessage string
	var branchName string

	flag.StringVar(&commitMessage, "m", "", "Input commit message (Required)")
	flag.StringVar(&branchName, "b", "", "Input branch name - master or other (Required)")
	flag.Parse()

	if commitMessage == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	if branchName == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	if runtime.GOOS == "windows" {
		outStatus, err := exec.Command("powershell", "git status").Output()
		outputStatus := string(outStatus[:])
		fmt.Println(outputStatus)
	
		time.Sleep(1 * time.Second)
	
		outAdd := exec.Command("powershell", "git add .")
		err = outAdd.Run()
	
		time.Sleep(2 * time.Second)
	
		outCommit, err := exec.Command("powershell", fmt.Sprintf("git commit -am \"%s\"", commitMessage)).Output()
		outputCommit := string(outCommit[:])
		fmt.Println(outputCommit)
	
		time.Sleep(2 * time.Second)
	
		var stdBuffer bytes.Buffer
		mwriter := io.MultiWriter(os.Stdout, &stdBuffer)
		outPush := exec.Command("powershell", fmt.Sprintf("git push origin \"%s\"", branchName))
		outPush.Stderr = mwriter
		outPush.Stdout = mwriter
		err = outPush.Run()
		fmt.Println(stdBuffer.String())
	
		if err != nil {
			fmt.Printf("%s", err)
		} else {
			fmt.Println("Command Successfully Executed")
		}
	} else if runtime.GOOS == "linux" || runtime.GOOS == "darwin" {
		outStatus, err := exec.Command("git", "status").Output()
		outputStatus := string(outStatus[:])
		fmt.Println(outputStatus)

		time.Sleep(1 * time.Second)

		outAdd := exec.Command("git", "add", ".")
		err = outAdd.Run()

		time.Sleep(2 * time.Second)

		outCommit, err := exec.Command("git", "commit", "-am", fmt.Sprintf("%s", commitMessage)).Output()
		outputCommit := string(outCommit[:])
		fmt.Println(outputCommit)

		time.Sleep(2 * time.Second)

		var stdBuffer bytes.Buffer
		mwriter := io.MultiWriter(os.Stdout, &stdBuffer)
		outPush := exec.Command("git", "push", "origin", fmt.Sprintf("%s", branchName))
		outPush.Stderr = mwriter
		outPush.Stdout = mwriter
		err = outPush.Run()
		fmt.Println(stdBuffer.String())

		if err != nil {
			fmt.Printf("%s", err)
		} else {
			fmt.Println("Command Successfully Executed")
		}
	}
}

func main() {
	execute()
}

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	iterate(cwd)
}

func iterate(path string) {
	filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatal(path)
		}

		if !info.IsDir() {
			return nil
		}

		files, err := ioutil.ReadDir(path)
		if err != nil {
			log.Fatal(path)
		}
		for _, file := range files {
			isVersionControlled := file.Name() == ".git"
			if isVersionControlled {
				fmt.Printf("%s\n", path)
				cmd := exec.Command("git", "status", "-s")
				cmd.Dir = path
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr
				err := cmd.Run()
				//out, err := cmd.CombinedOutput()
				if err != nil {
					log.Fatalf("cmd.Run() failed with %s\n", err)
				}
				//fmt.Printf("combined out:\n%s\n", string(out))
				//return filepath.SkipDir
			}
		}

		return nil
	})
}

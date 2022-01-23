package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

func main() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
	configBytes, err := os.ReadFile("./hriseob.conf")
	if errors.Is(err, os.ErrNotExist) {
		log.Fatal("Can't find file config (./hriseob.conf)")
		return
	}

	var config Config
	err = json.Unmarshal(configBytes, &config)
	if err != nil {
		log.Fatal("Can't unmarshal Config (./hriseob.conf): Please make sure to make it valid")
		return
	}
	args := os.Args[1:]
	if len(args) != 0 {
		Execute(args[0])
	} else {
		Execute(config.Home)
	}
}

func Execute(currentDir string) {
	filesString := ".\n..\n"
	files, _ := ioutil.ReadDir(currentDir)
	for _, file := range files {
		filesString += file.Name() + "\n"
	}
	fmt.Print(filesString)
	var filename string
	fmt.Scanln(&filename)
	file, _ := os.Stat(currentDir + "/" + filename)
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
	if file.IsDir() {
		Execute(currentDir + "/" + filename)
	} else {
		fileBytes, _ := os.ReadFile(currentDir + "/" + filename)
		fmt.Println(string(fileBytes))
		Execute(currentDir)
	}
}

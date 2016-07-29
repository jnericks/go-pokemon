package main

import (
	"io/ioutil"
	"log"
	"os"
	"path"
)

var pwd string

func init() {
	var err error
	pwd, err = os.Getwd()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

func getFullPath(fileName string) string {
	return path.Join(pwd, fileName)
}

func writeFile(fileName string, data []byte) error {
	return ioutil.WriteFile(getFullPath(fileName), data, 777)
}

func readFile(fileName string) ([]byte, error) {
	return ioutil.ReadFile(getFullPath(fileName))
}

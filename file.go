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

func recreateFile(fileName string) error {
	_ = os.Remove(getFullPath(fileName))
	_, err := os.Create(getFullPath(fileName))
	if err != nil {
		return err
	}

	return nil
}

func writeFile(fileName string, bytes []byte) error {
	// open file using READ & WRITE permission
	var file, err = os.OpenFile(getFullPath(fileName), os.O_RDWR, 0644)
	defer file.Close()
	if err != nil {
		return err
	}

	_, err = file.Write(bytes)
	if err != nil {
		return err
	}

	// save changes
	err = file.Sync()
	if err != nil {
		return err
	}

	return nil
}

func readFile(fileName string) ([]byte, error) {
	return ioutil.ReadFile(getFullPath(fileName))
}

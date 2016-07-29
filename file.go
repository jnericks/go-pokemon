package gopokemon

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
)

// the present working directory of the project
var pwd string

func init() {
	var err error
	pwd, err = os.Getwd()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

// GetDataPath will return the path to the /data folder
func GetDataPath() string {
	return pwd
}

// WriteFile will write to a file in the /data folder of the pwd
// If file does not already exist it will create the file
// If file already exists it will override the data already in the file
func WriteFile(fileName string, data []byte) error {
	return ioutil.WriteFile(getFullPath(fileName), data, 0666)
}

// AppendFile will append the data on a new line to a file in the /data folder of the pwd
func AppendFile(fileName string, data string) error {
	fullpath := getFullPath(fileName)

	_, fileNotExistsErr := os.Stat(fullpath)
	if os.IsNotExist(fileNotExistsErr) {
		return fileNotExistsErr
	}

	f, err1 := os.OpenFile(fullpath, os.O_APPEND|os.O_WRONLY, 0666)
	if err1 != nil {
		return err1
	}

	defer f.Close()

	_, err2 := f.WriteString(fmt.Sprintf("\n%s", data))
	if err2 != nil {
		return err2
	}

	return nil
}

// ReadFile will read all data from file in the /data folder of the pwd
func ReadFile(fileName string) ([]byte, error) {
	return ioutil.ReadFile(getFullPath(fileName))
}

func getFullPath(fileName string) string {
	return path.Join(pwd, fileName)
}

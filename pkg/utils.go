package pkg

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"time"
)

func DeleteFilesFromDir(dirPath string) error {
	dir, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return err
	}
	for _, d := range dir {
		err = os.RemoveAll(path.Join([]string{dirPath, d.Name()}...))
		if err != nil {
			return err
		}
	}

	return nil
}

func CreateFile(data []byte, dir string) {
	filename := fmt.Sprintf("%s/%v-data.json", dir, time.Now().Unix())
	fmt.Println("file data", string(data))
	os.WriteFile(filename, data, 0644)
}

func CreateDirectory(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		fmt.Println("I am creating a directory")
		err := os.Mkdir(dir, os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func ReadJsonFromFile(output *map[string]string, dir string) {
	var fileInfo os.FileInfo
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		if fileInfo == nil {
			fileInfo = f
			continue
		}
		if f.ModTime().After(fileInfo.ModTime()) {
			fileInfo = f
		}
	}

	if fileInfo != nil {
		data, err := ioutil.ReadFile(dir+ "/" + fileInfo.Name())
		if err != nil {
			log.Fatal(err)
		}
		jsonErr := json.Unmarshal(data, output)
		if jsonErr != nil {
			log.Fatal(jsonErr)
		}
	}
}
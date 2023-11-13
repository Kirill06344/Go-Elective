package main

import (
	"fmt"
	"github.com/Kirill06344/sequenceLib/pkg/substringSearch"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
	"os"
)

type FileInfo struct {
	Filename  string `yaml:"filename"`
	Substring string `yaml:"substring"`
}

func main() {

	filesYml, err := os.ReadFile("files.yaml")
	if err != nil {
		log.Error(err)
		return
	}

	var files []FileInfo
	err = yaml.Unmarshal(filesYml, &files)
	if err != nil {
		log.Error(err)
		return
	}

	for _, item := range files {
		contains, err := substringSearch.Contains(item.Filename, item.Substring)
		fmt.Printf("File %s: \n", item.Filename)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(contains)
		}
	}

}

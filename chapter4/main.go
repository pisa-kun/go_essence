package main

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

var content = `
{
	"groups": "放課後",
	"concept": "Girls",
	"informations": {
		"color": "orange",
		"member": 5
	}
}
`

type Informations struct {
	Color  string `json:"color"`
	Member int    `json:"member"`
}

type Data struct {
	Groups       string       `json:"groups"`
	Concept      string       `json:"concept"`
	Informations Informations `json:"informations"`
}

// カレントディレクトリ配下のファイル一覧を返す
func CurrentWalkDir() []string {
	files := []string{}

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	walkDir := func(path string, info fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		files = append(files, path)
		return nil
	}
	err = filepath.WalkDir(cwd, walkDir)
	if err != nil {
		log.Fatal(nil)
	}
	return files
}

func main() {
	// var data Data
	// err := json.Unmarshal([]byte(content), &data)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(data)

	f, err := os.Open("input.json")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var data Data
	err = json.NewDecoder(f).Decode(&data)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(data)

	fmt.Println("---CurrentDirectory under all files---")
	for _, dir := range CurrentWalkDir() {
		fmt.Println(dir)
	}
}

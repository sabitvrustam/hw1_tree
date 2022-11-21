package main

import (
	"fmt"
	"os"
)

type Object struct {
	name         []string
	fileOrFolder []bool
	amount       []int
	nesting      []int
	size         []int64
	number       []int
}

var object Object
var simbolDir string
var maxLevelNesting int
var stringPattern string

func main() {
	scanDir("testdata", 0)

	transitDir := "├"
	ultimateDir := "└"
	continuDir := "│"
	tab := "\t"
	openDir := "─────"

	for i := range object.nesting {
		if maxLevelNesting < object.nesting[i] {
			maxLevelNesting = object.nesting[i]
		}
	}
	lineDiagram := make([]bool, maxLevelNesting+1)

	for i := range object.fileOrFolder {
		stringPattern = ""
		for n := 0; n < object.nesting[i]; n++ {
			if !lineDiagram[n] {
				stringPattern += tab
			} else if lineDiagram[n] {
				stringPattern += continuDir + tab
			}
		}
		var fileSize string
		if !object.fileOrFolder[i] && object.size[i] != 0 {
			fileSize = fmt.Sprintf("(%db)", object.size[i])
		} else if !object.fileOrFolder[i] && object.size[i] == 0 {
			fileSize = "(empty)"
		}
		if object.number[i]+1 == object.amount[i] {
			simbolDir = ultimateDir
			lineDiagram[object.nesting[i]] = false

		} else {
			simbolDir = transitDir
			lineDiagram[object.nesting[i]] = true
		}

		fmt.Println(stringPattern+simbolDir+openDir, object.name[i], fileSize)
	}
}

func scanDir(firstDirectory string, numberDirectory int) {

	files, err := os.ReadDir(firstDirectory)
	if err != nil {
		fmt.Println("Каталога не существует", err)
	}

	for i := range files {
		file := files[i]
		info, _ := file.Info()
		object.number = append(object.number, i)
		object.size = append(object.size, info.Size())
		object.nesting = append(object.nesting, numberDirectory)
		object.amount = append(object.amount, len(files))
		object.name = append(object.name, file.Name())
		object.fileOrFolder = append(object.fileOrFolder, file.IsDir())

		if file.IsDir() {
			directory := file.Name()

			scanDir(firstDirectory+"/"+directory, numberDirectory+1)

		}

	}

}

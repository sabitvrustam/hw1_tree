package main

import (
	"fmt"
	"os"
)

type entry struct {
	name   string
	isDir  bool
	count  int
	height int
	size   int64
	number int
}

const (
	tSimbol     = "├"
	lSimbol     = "└"
	vertSimbol  = "│"
	tabSimbol   = "\t"
	horizSimbol = "─────"
)

func main() {
	entrys := scanDir("testdata", 0)
	var maxHeight int
	var simbolDir string

	for _, entry := range entrys {
		if maxHeight < entry.height {
			maxHeight = entry.height
		}
	}
	lineDiagram := make([]bool, maxHeight+1)

	for _, entry := range entrys {
		fileSize := ""
		stringPattern := ""
		for i := 0; i < entry.height; i++ {
			if !lineDiagram[i] {
				stringPattern += tabSimbol
			} else if lineDiagram[i] {
				stringPattern += vertSimbol + tabSimbol
			}
		}

		if !entry.isDir && entry.size != 0 {
			fileSize = fmt.Sprintf("(%db)", entry.size)
		} else if !entry.isDir && entry.size == 0 {
			fileSize = "(empty)"
		}

		if entry.number+1 == entry.count {
			simbolDir = lSimbol
			lineDiagram[entry.height] = false

		} else {
			simbolDir = tSimbol
			lineDiagram[entry.height] = true
		}

		fmt.Println(stringPattern+simbolDir+horizSimbol, entry.name, fileSize)
	}
}

func scanDir(firstDir string, height int) (entrys []entry) {
	var element entry

	files, err := os.ReadDir(firstDir)
	if err != nil {
		fmt.Println("Каталога не существует", err)
	}

	for i, file := range files {
		infoElement, _ := file.Info()
		element.name = file.Name()
		element.isDir = file.IsDir()
		element.count = len(files)
		element.number = i
		element.height = height
		element.size = infoElement.Size()

		entrys = append(entrys, element)

		if file.IsDir() {
			entrys = append(entrys, scanDir(firstDir+"/"+element.name, height+1)...)
		}

	}
	return entrys
}

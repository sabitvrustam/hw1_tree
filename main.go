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
	tSymbol     = "├"
	lSimbol     = "└"
	vertSymbol  = "│"
	tabSymbol   = "\t"
	horizSymbol = "─────"
)

func main() {
	entries := scanDir("testdata", 0)
	var maxHeight int
	var symbolDir string

	for _, entry := range entries {
		if maxHeight < entry.height {
			maxHeight = entry.height
		}
	}
	lineDiagram := make([]bool, maxHeight+1)

	for _, entry := range entries {
		fileSize := ""
		stringPattern := ""
		for i := 0; i < entry.height; i++ {
			if !lineDiagram[i] {
				stringPattern += tabSymbol
			} else if lineDiagram[i] {
				stringPattern += vertSymbol + tabSymbol
			}
		}

		if !entry.isDir {
			fileSize = "(empty)"
			if entry.size != 0 {
				fileSize = fmt.Sprintf("(%db)", entry.size)
			}
		}
		symbolDir = tSymbol
		lineDiagram[entry.height] = true

		if entry.number+1 == entry.count {
			symbolDir = lSimbol
			lineDiagram[entry.height] = false
		}

		fmt.Println(stringPattern+symbolDir+horizSymbol, entry.name, fileSize)
	}
}

func scanDir(firstDir string, height int) (entrys []entry) {
	var element entry

	files, err := os.ReadDir(firstDir)
	if err != nil {
		fmt.Println("Каталога не существует", err)
		panic(err)
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

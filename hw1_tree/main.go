package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	//out := os.Stdout
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}
	path := os.Args[1]
	fl := false

	if len(os.Args) > 2 {
		fl = bool(os.Args[2] == "-f")
	}

	//printFiles := len(os.Args) == 3 && os.Args[2] == "-f"

	/*
		outputDirRead, err := os.Open(path)
		if err != nil {
			panic(err)
		}

			fmt.Println("outputDirRead", outputDirRead, out, printFiles)
			outputDirFiles, _ := outputDirRead.Readdir(0)
			fmt.Println("outputDirFiles", outputDirFiles)

			// Loop over files.
			for outputIndex := range outputDirFiles {
				outputFileHere := outputDirFiles[outputIndex]

				// Get name of file.
				outputNameHere := outputFileHere.Name()

				// Print name.
				fmt.Println(outputNameHere)
			}
	*/
	//rez bytes.Buffer
	//_, err := dirTree(rez, path, true)

	var rez bytes.Buffer
	rez.Write([]byte(""))

	dirTree(&rez, path, fl)
	rez.WriteTo(os.Stdout)
	/*
		if err != nil {
			panic(err.Error())
		}
	*/
}

func dirTree(rez *bytes.Buffer, path string, fl bool) error {

	var e error
	out := ""
	printTreePath(&out, path, fl, "")
	//fmt.Println(out)
	rez.Write([]byte(out))

	return e
}

func printTreePath(fOut *string, fPath string, fFile bool, prefixLine1 string) {

	prefixLine := ""

	outputDirRead, err := os.Open(fPath) //получить доступ
	if err != nil {
		fmt.Println("Не смог получить доступ к каталогу:", fPath)
		panic(err)
	}

	outputDirFiles, _ := outputDirRead.Readdir(0)
	//fmt.Println("outputDirFiles", len(outputDirFiles))

	// Loop over files.
	if len(outputDirFiles) > 0 {

		var sliceDir [][]string //объявим слайс для хранения имен файлов и директорий
		for outputIndex := range outputDirFiles {
			if !outputDirFiles[outputIndex].IsDir() && !fFile { //нужно ли выводить файлы
				continue //выйдем если это файл
			}
			sliceDir = append(sliceDir, []string{outputDirFiles[outputIndex].Name(), strconv.FormatBool(outputDirFiles[outputIndex].IsDir())[0:1]}) //[0][0] имя;[0][1] признак каталога
		}
		outputDirRead.Close()
		sort.Slice(sliceDir[:], func(i, j int) bool { return sliceDir[i][0] < sliceDir[j][0] })

		//fmt.Println("sliceDir", sliceDir[0][0], "-", sliceDir[0][1])
		//fmt.Println(sliceDir)

		for i := 0; i < len(sliceDir); i++ {
			if i == (len(sliceDir) - 1) { //если последний
				prefixLine = prefixLine1 + "└───"
			} else { //если не последний
				prefixLine = prefixLine1 + "├───"
			}

			if sliceDir[i][1] != "t" {
				fp := fPath + "/" + sliceDir[i][0]
				fSize, e := os.Stat(fp)
				if e != nil {
					log.Fatal(err)
				}

				fs := "empty"
				if fSize.Size() != 0 {
					fs = strconv.FormatInt(fSize.Size(), 10) + "b"
				}
				*fOut = *fOut + prefixLine + sliceDir[i][0] + " (" + fs + ")\n" //Имя файла или директории; true - если директория
			} else {
				*fOut = *fOut + prefixLine + sliceDir[i][0] + "\n" //Имя файла или директории; true - если директория
			}

			//fmt.Println(len(sliceDir), i, level)
			//*fOut = *fOut + prefixLine + sliceDir[i][0] + "\n" //Имя файла или директории; true - если директория
			//fmt.Println(fOut)

			if sliceDir[i][1] == "t" {

				if i < (len(sliceDir) - 1) {
					printTreePath(fOut, fPath+"\\"+sliceDir[i][0], fFile, (prefixLine1 + "│\t"))
				} else {
					printTreePath(fOut, fPath+"\\"+sliceDir[i][0], fFile, (prefixLine1 + "\t"))
				}
			}
		}

	}
}

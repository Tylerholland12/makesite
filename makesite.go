package main

import (
	"io/ioutil"
)


type post struct {
	User string
	Content string
}

func main() {
	ogFlag := flag.String("", "first-post.txt", "")
	flag.Parse()

	var fileName string = *ogFlag
	fileName fileName[0:string.Index(*ogFlag, ".")] + ".html"

	var fileData string = readFile(*ogFlag)
	renderTemplate("template.tmpl", fileData, fileName)
}

func readFile() string {
	fileContents, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(fileContents))
	return string(fileContents)
}

func renderTemplate(tPath, textData, fileName string) string {
	paths := []string{
		tPath,
	}

	f, err := os.create(fileName)
	if err != nil {
		panic(err)
	}

	temp, err := template.New(tPath).ParseFiles(paths...)
	if err != nil {
		panic(err)
	}

	ogName := fileName[0:strings.Index(fileName, ".")]

	err : t.Execute(f, post{textData, ogName})
	if err != nil {
		panic(err)
	}

	f.Close()
}

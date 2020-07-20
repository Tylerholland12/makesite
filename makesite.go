package main

import (
	"flag"
	"io/ioutil"
	"strings"
	"text/template"

	"github.com/bregydoc/gtranslate"
	"golang.org/x/text/language"
	"gopkg.in/russross/blackfriday.v2"
)

type post struct {
	User    string
	Content string
}

func main() {
	ogFlag := flag.String("file", "first-post.txt", "define text")
	dirFlag := flag.String("directory", "none", "generates all .txt files in directory")
	outputDirFlag := flag.String("output", "templates/", "Generator output directory")
	flag.Parse()

	if *dirFlag == "none" {
		runFile(*ogFlag, "txt_dir/")
	} else {
		runDir(*dirFlag, *outputDirFlag)
	}
}

func runFile(fileFlag, directory string) {

	var fileName string = ogFlag

	if fileName[strings.Index(fileFlag, "."):len(fileFlag)] != ".txt" {
		return
	}

	if strings.Contains(strings.ToLower(fileFlag), ".md") {

		var data string = readFile(directory + fileFlag)
		tmpl := renderTemplate("template.tmpl", data, fileName)
		output := blackfriday.Run(tmpl)
		ioutil.WriteFile(output, ogFlag)

		return
	}

	fileName = fileName[0:strings.Index(fileFlag, ".")] + ".html"

	var data string = readFile(directory + fileFlag)
	renderTemplate("template.tmpl", data, fileName)
}

func readFile() string {
	fileContents, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	return string(fileContents)
}

func runDir(directory, output string) {

	if directory[len(directory)-1] != "/"[0] {
		directory += "/"
	}

	files, err := ioutil.ReadDir(directory)

	if err != nil {
		panic(err)
	}

	for _, file := range files {

		if file.IsDir() == false {
			runFile(file.Name(), directory)
		} else {
			runDir(directory+"/"+file.Name(), output)
		}
	}
}

func renderTemplate(tPath, textData, fileName string) {
	paths := []string{
		tPath,
	}

	f, err := os.create("templates/" + fileName)
	if err != nil {
		panic(err)
	}

	t, err := template.New(tPath).ParseFiles(paths...)
	if err != nil {
		panic(err)
	}

	ogName := fileName[0:strings.Index(fileName, ".")]

	txtTranslated := translateText(textData)

	err = t.Execute(f, post{txtTranslated, ogName})
	if err != nil {
		panic(err)
	}

	f.Close()
}

func translateText(txtData string) string {

	translated, err := gtranslate.Translate(txtData, language.English, language.French)
	if err != nil {
		panic(err)
	}

	return string(translated)
}

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

type TemplateData struct {
	AppName    string
	ModuleName string
	GoVersion  string
}

func main() {
	fmt.Println("Bootstrap")

	data := TemplateData{
		AppName:    "DemoApp",
		ModuleName: "github.com/username/demoapp",
		GoVersion:  "1.21",
	}

	generateFile("main.go", data)
	generateFile("Makefile", data)
	generateFile("go.mod", data)
}

func generateFile(outputFileName string, data TemplateData) {
	outputDir := "output"
	outputFilePath := filepath.Join(outputDir, outputFileName)
	templateDir := "template"
	templateFilePath := fmt.Sprintf("%s/%s.tpl", templateDir, outputFileName)

	tmpl, err := template.ParseFiles(templateFilePath)
	if err != nil {
		panic(err)
	}

	err = os.MkdirAll(outputDir, os.ModePerm)
	if err != nil {
		panic(err)
	}

	outputFile, err := os.Create(outputFilePath)
	if err != nil {
		panic(err)
	}

	defer outputFile.Close()

	err = tmpl.Execute(outputFile, data)
	if err != nil {
		panic(err)
	}

	fmt.Printf("'%s' generated\n", outputFilePath)
}

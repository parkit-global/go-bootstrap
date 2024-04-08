package main

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

type TemplateData struct {
	AppName string
}

func main() {
	fmt.Println("Bootstrap")

	outputFileName := "main.go"
	generateFile(
		outputFileName,
		TemplateData{
			AppName: "DemoApp",
		})
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

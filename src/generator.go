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

type Generator struct {
	OutputDir   string
	TemplateDir string
}

func (g *Generator) generateFile(outputFileName string, data TemplateData) {
	outputFilePath := filepath.Join(g.OutputDir, outputFileName)
	outputFileDir := filepath.Dir(outputFilePath)
	templateFilePath := fmt.Sprintf("%s/%s.tpl", g.TemplateDir, outputFileName)

	tmpl, err := template.ParseFiles(templateFilePath)
	if err != nil {
		panic(err)
	}

	err = os.MkdirAll(outputFileDir, os.ModePerm)
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

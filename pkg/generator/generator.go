package generator

import (
	"fmt"
	"io"
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

func (g *Generator) GenerateFile(outputFileName string, data TemplateData) error {
	outputFilePath := filepath.Join(g.OutputDir, outputFileName)
	outputFileDir := filepath.Dir(outputFilePath)
	templateFilePath := fmt.Sprintf("%s/%s.tpl", g.TemplateDir, outputFileName)

	tmpl, err := template.ParseFiles(templateFilePath)
	if err != nil {
		return err
	}

	err = os.MkdirAll(outputFileDir, os.ModePerm)
	if err != nil {
		return err
	}

	outputFile, err := os.Create(outputFilePath)
	if err != nil {
		return err
	}

	defer outputFile.Close()

	err = tmpl.Execute(outputFile, data)
	if err != nil {
		return err
	}

	fmt.Printf("'%s' generated\n", outputFilePath)

	return nil
}

func (g *Generator) GenerateFiles(outputFileNames []string, data TemplateData) error {
	for _, outputFileName := range outputFileNames {
		err := g.GenerateFile(outputFileName, data)
		if err != nil {
			return err
		}
	}
	return nil
}

func (g *Generator) CopyFile(outputFileName string) error {
	outputFilePath := filepath.Join(g.OutputDir, outputFileName)
	outputFileDir := filepath.Dir(outputFilePath)
	sourceFilePath := fmt.Sprintf("%s/%s", g.TemplateDir, outputFileName)

	err := os.MkdirAll(outputFileDir, os.ModePerm)
	if err != nil {
		return err
	}

	outputFile, err := os.Create(outputFilePath)
	if err != nil {
		return err
	}

	defer outputFile.Close()

	sourceFile, err := os.Open(sourceFilePath)
	if err != nil {
		return err
	}

	_, err = io.Copy(outputFile, sourceFile)
	if err != nil {
		return err
	}

	fmt.Printf("'%s' copied\n", outputFilePath)

	return nil
}

func (g *Generator) CopyFiles(outputFileNames []string) error {
	for _, outputFileName := range outputFileNames {
		err := g.CopyFile(outputFileName)
		if err != nil {
			return err
		}
	}
	return nil
}

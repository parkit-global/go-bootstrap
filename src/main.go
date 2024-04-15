package main

import (
	"github.com/spf13/cobra"
)

func main() {
	var outputDir string
	var templateDir string
	var cmd = &cobra.Command{
		Use: "go-bootstrap",
	}

	cmd.PersistentFlags().StringVar(&outputDir, "output", "output", "Output directory")
	cmd.PersistentFlags().StringVar(&templateDir, "template", "template", "Template directory")
	cmd.Execute()

	generator := Generator{
		OutputDir:   outputDir,
		TemplateDir: templateDir,
	}

	data := TemplateData{
		AppName:    "DemoApp",
		ModuleName: "github.com/username/demoapp",
		GoVersion:  "1.21",
	}

	err := generator.GenerateFiles(
		[]string{
			"src/main.go",
			"Makefile",
			"go.mod",
		},
		data)

	if err != nil {
		panic(err)
	}

	err = generator.CopyFiles(
		[]string{
			"application.yaml",
			"src/config.go",
			"src/endpoint.go",
		})

	if err != nil {
		panic(err)
	}
}

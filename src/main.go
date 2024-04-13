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

	generator.generateFile("src/main.go", data)
	generator.generateFile("src/config.go", data)
	generator.generateFile("Makefile", data)
	generator.generateFile("go.mod", data)
	generator.generateFile("application.yaml", data)
}

package main

import (
	"github.com/parkit-global/go-bootstrap/pkg/generator"
	"github.com/spf13/cobra"
)

func main() {
	var outputDir string
	var templateDir string
	var appName string
	var moduleName string
	var goVersion string
	var cmd = &cobra.Command{
		Use: "go-bootstrap",
	}

	cmd.PersistentFlags().StringVar(&outputDir, "output", "output", "Output directory")
	cmd.PersistentFlags().StringVar(&templateDir, "template", "template", "Template directory")
	cmd.PersistentFlags().StringVar(&appName, "app-name", "DemoApp", "Name of the application")
	cmd.PersistentFlags().StringVar(&moduleName, "module-name", "github.com/username/demoapp", "Name of the module")
	cmd.PersistentFlags().StringVar(&goVersion, "go-version", "1.21", "Go version")
	cmd.Execute()

	g := generator.Generator{
		OutputDir:   outputDir,
		TemplateDir: templateDir,
	}

	data := generator.TemplateData{
		AppName:    appName,
		ModuleName: moduleName,
		GoVersion:  goVersion,
	}

	err := g.GenerateFiles(
		[]string{
			"cmd/main.go",
			"Makefile",
			"go.mod",
		},
		data)

	if err != nil {
		panic(err)
	}

	err = g.CopyFiles(
		[]string{
			"application.yaml",
			"cmd/config.go",
			"cmd/endpoint.go",
		})

	if err != nil {
		panic(err)
	}
}

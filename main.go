package main

import (
	"fmt"
	"html/template"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const (
	srcDir    = "src/templates2"
	outputDir = "docs2"
)

func main() {
	// Define the source directory for templates and the output directory for rendered files
	// Ensure the output directory exists
	err := os.MkdirAll(outputDir, os.ModePerm)
	if err != nil {
		log.Fatalf("Error creating output directory: %v", err)
	}

	t, err := template.ParseGlob(srcDir + "/*")
	if err != nil {
		log.Fatal(err)
	}
	for _, tmplName := range t.Templates() {
		if tmplName.Name() != "about.html" {
			continue
		}
		// Print the name of the template
		if !strings.HasSuffix(tmplName.Name(), ".html") {
			continue
		}
		err = renderTemplate(tmplName)
		if err != nil {
			log.Fatal(err)
		}
	}
}
func renderTemplate(tmpl *template.Template) error {
	// Create an output file for each template
	outputFile, err := os.Create(filepath.Join(outputDir, tmpl.Name()))
	if err != nil {
		return fmt.Errorf("error creating output file: %v", err)
	}
	defer outputFile.Close()

	// Render the template (without any data, or you can pass data as needed)
	err = tmpl.Execute(outputFile, nil)
	if err != nil {
		return fmt.Errorf("error rendering template: %v", err)
	}

	fmt.Printf("Rendered template: %s\n", tmpl.Name())
	return nil
}

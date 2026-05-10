package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// Define HTML-related extensions
var htmlExtensions = map[string]bool{
	// --- Standard & Legacy ---
	".html": true, ".htm": true,
	".xhtml": true, ".xml": true,
	".asp": true, ".aspx": true, // Classic ASP / WebForms
	".cfm": true, ".cfml": true, // ColdFusion

	// --- Modern Frontend (JavaScript/TypeScript) ---
	".jsx": true, ".tsx": true, // React
	".vue":     true,                   // Vue
	".svelte":  true,                   // Svelte
	".astro":   true,                   // Astro
	".html.js": true, ".html.ts": true, // Lit / HTML-in-JS

	// --- PHP & Content Management ---
	".php": true, ".phtml": true,
	".twig":      true, // Symfony
	".blade.php": true, // Laravel
	".tpl":       true, // Smarty / Dwoo

	// --- Python & Ruby ---
	".jinja": true, ".jinja2": true, // Flask / Django
	".pyhtml": true,
	".erb":    true, // Ruby on Rails
	".rhtml":  true,

	// --- Go (Golang) Specific ---
	".gohtml": true, // Standard html/template
	".tmpl":   true, // Common generic template

	// --- Java & .NET ---
	".jsp": true, ".jspx": true, // Java Server Pages
	".tag":    true,
	".cshtml": true, ".razor": true, // ASP.NET Core / Blazor

	// --- Static Site Generators & Logicless ---
	".liquid":   true, // Shopify / Jekyll
	".mustache": true,
	".hbs":      true, ".handlebars": true, // Handlebars
	".ejs": true, // Embedded JS
	".njk": true, // Nunjucks
	".mdx": true, // Markdown + JSX

	// --- Abstraction Layers (HTML-like syntax) ---
	".pug": true, ".jade": true, // Pug
	".haml": true, // Haml
	".slim": true, // Slim
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: mycli <folder_path>")
		os.Exit(1)
	}

	targetDir := os.Args[1]
	results := make(map[string]string)

	// Regex handles multi-line content (?s)
	// re := regexp.MustCompile(`(?s)data-twomi18n=["']([^"']+)["'][^>]*>([^<]+)`)

	err := filepath.WalkDir(targetDir, func(path string, d os.DirEntry, err error) error {
		if err != nil || d.IsDir() {
			return nil
		}

		if !htmlExtensions[strings.ToLower(filepath.Ext(path))] {
			return nil
		}

		content, _ := os.ReadFile(path)
		fileResults := parseContent(content)

		for k, v := range fileResults {
			results[k] = v
		}

		return nil
	})

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "    ")
	enc.SetEscapeHTML(false)
	enc.Encode(results)
}

func scrubText(s string) string {
	s = strings.ReplaceAll(s, "\n", " ")
	s = strings.ReplaceAll(s, "\t", " ")
	s = strings.ReplaceAll(s, "\r", " ")

	spaceRegex := regexp.MustCompile(`\s+`)
	s = spaceRegex.ReplaceAllString(s, " ")

	return strings.TrimSpace(s)
}

// parseContent takes the raw file bytes and returns the key-value map.
func parseContent(content []byte) map[string]string {
	results := make(map[string]string)
	re := regexp.MustCompile(`(?s)data-twomi18n=["']([^"']+)["'][^>]*>([^<]+)`)

	matches := re.FindAllSubmatch(content, -1)
	for _, m := range matches {
		keysAttr := string(m[1])
		rawText := string(m[2])
		cleanVal := scrubText(rawText)

		keyList := strings.Fields(keysAttr)
		for _, fullKey := range keyList {
			if strings.Contains(fullKey, "[") {
				cleanKey := strings.Split(fullKey, "[")[0]
				results[cleanKey] = ""
			} else {
				results[fullKey] = cleanVal
			}
		}
	}
	return results
}

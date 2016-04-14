package main

import (
	"bytes"
	"errors"
	"html/template"
)

var templates = map[string]*template.Template{}

func getTemplateBytes(filename string) ([]byte, error) {
	tmplBytes, err := Asset(filename)
	if err != nil {
		return nil, errors.New(filename + " not found")
	}
	return tmplBytes, err
}

func generateStaticPath(p string) string {
	return baseURL + p
}

func parseFiles(filenames ...string) (*template.Template, error) {
	if len(filenames) == 0 {
		return nil, errors.New("template: no files named in call to parseFiles")
	}
	root := template.New("root")
	root.Funcs(template.FuncMap{
		"staticPath": generateStaticPath,
	})
	for _, filename := range filenames {
		tmplBytes, err := getTemplateBytes(filename)
		if err != nil {
			return nil, err
		}
		_, err = root.Parse(string(tmplBytes))
		if err != nil {
			return nil, err
		}
	}
	return root, nil
}

func cacheTemplates(sets []string) error {
	for _, set := range sets {
		root, err := parseFiles(
			"templates/navigation.tmpl",
			"templates/base.tmpl",
			"templates/footer.tmpl",
			set,
		)
		if err != nil {
			return err
		}
		templates[set] = root
	}
	return nil
}

func executeTemplate(path string, params map[string]interface{}) (string, error) {
	var buf bytes.Buffer
	root := templates[path]
	if root == nil {
		return "", errors.New("template: template not found")
	}
	root.ExecuteTemplate(&buf, "base", params)
	return buf.String(), nil
}

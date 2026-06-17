package handlers

import (
	"html/template"
	"net/http"
	"path/filepath"
	"strings"

	"alquimia-alimentar/internal/content"
)

// PageData is the data passed to every page template.
type PageData struct {
	Lang     string
	NavLinks []content.Link
	Content  *content.Content
}

// Home renders the landing page for the requested language.
func Home(w http.ResponseWriter, r *http.Request) {
	lang := languageFromPath(r.URL.Path)

	cont, err := content.Load(lang)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := PageData{
		Lang:     lang,
		NavLinks: cont.Header.Nav,
		Content:  cont,
	}

	tmpl, err := loadTemplates()
	if err != nil {
		http.Error(w, "Erro ao carregar templates: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.ExecuteTemplate(w, "base", data); err != nil {
		http.Error(w, "Erro ao renderizar página: "+err.Error(), http.StatusInternalServerError)
	}
}

// languageFromPath extracts the language code from the URL path.
func languageFromPath(path string) string {
	switch {
	case strings.HasPrefix(path, "/en/") || path == "/en":
		return "en"
	case strings.HasPrefix(path, "/nl/") || path == "/nl":
		return "nl"
	default:
		return "pt"
	}
}

// loadTemplates discovers and parses all HTML templates.
// base.html is parsed first so page templates can override its blocks.
func loadTemplates() (*template.Template, error) {
	funcMap := template.FuncMap{
		"add": func(a, b int) int { return a + b },
		"sub": func(a, b int) int { return a - b },
		"even": func(i int) bool { return i%2 == 0 },
		"rotateClass": func(r string) string {
			if strings.HasPrefix(r, "-") {
				return "md:-rotate-" + r[1:]
			}
			return "md:rotate-" + r
		},
	}

	baseFiles, err := filepath.Glob("web/templates/*.html")
	if err != nil {
		return nil, err
	}

	partialFiles, err := filepath.Glob("web/templates/partials/*.html")
	if err != nil {
		return nil, err
	}

	pageFiles, err := filepath.Glob("web/templates/pages/*.html")
	if err != nil {
		return nil, err
	}

	files := append(baseFiles, partialFiles...)
	files = append(files, pageFiles...)

	return template.New("").Funcs(funcMap).ParseFiles(files...)
}

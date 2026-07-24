// File: middle/template/main.go
// Level: Middle
// Topik: Templates (html/template & text/template)
//
// html/template: untuk web, otomatis escape HTML (anti XSS)
// text/template: untuk teks biasa, email, config files
//
// Fitur: variables, functions, pipelines, conditions, loops, template composition

package main

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
	"strings"
	"time"
)

// Data yang akan di-render ke template
type User struct {
	Name    string
	Age     int
	Admin   bool
	Skills  []string
	Profile Profile
}

type Profile struct {
	Bio   string
	Email string
}

// Custom function untuk template
var funcMap = template.FuncMap{
	"upper": strings.ToUpper,
	"lower": strings.ToLower,
	"greet": func(name string) string {
		return fmt.Sprintf("Halo %s, selamat datang!", name)
	},
	"since": func(t time.Time) string {
		return time.Since(t).Truncate(time.Second).String()
	},
}

func main() {
	// 1. TEXT TEMPLATE BASIC
	fmt.Println("=== Text Template Basic ===")
	tmplText := "Halo {{.Name}}! Umur {{.Age}} tahun.\n"

	tmpl, _ := template.New("greeting").Parse(tmplText)
	tmpl.Execute(os.Stdout, User{Name: "Anggi", Age: 20})
	fmt.Println()

	// 2. TEMPLATE WITH CONDITION
	fmt.Println("\n=== Template with Condition ===")
	tmplCond := `{{if .Admin}}Admin{{else}}User{{end}}: {{.Name}}` + "\n"

	tmpl2, _ := template.New("cond").Parse(tmplCond)
	tmpl2.Execute(os.Stdout, User{Name: "Anggi", Admin: true})
	tmpl2.Execute(os.Stdout, User{Name: "Budi", Admin: false})
	fmt.Println()

	// 3. TEMPLATE WITH LOOP
	fmt.Println("=== Template with Loop ===")
	tmplLoop := `Skills:
{{range .Skills}}- {{.}}
{{end}}` + "\n"

	tmpl3, _ := template.New("loop").Parse(tmplLoop)
	tmpl3.Execute(os.Stdout, User{
		Name:   "Anggi",
		Skills: []string{"Go", "React", "Docker"},
	})
	fmt.Println()

	// 4. TEMPLATE WITH PIPELINE & FUNCTION
	fmt.Println("=== Template with Functions ===")
	tmplFunc := `{{greet .Name}}
Bio: {{.Profile.Bio | upper}}
Email: {{.Profile.Email | lower}}
` + "\n"

	tmpl4 := template.Must(template.New("funcs").Funcs(funcMap).Parse(tmplFunc))
	tmpl4.Execute(os.Stdout, User{
		Name: "Anggi",
		Profile: Profile{
			Bio:   "Software Engineer",
			Email: "ANGGI@MAIL.COM",
		},
	})
	fmt.Println()

	// 5. HTML TEMPLATE (auto-escape)
	fmt.Println("=== HTML Template (auto-escape) ===")
	htmlTmpl := `<h1>{{.Name}}</h1>
<p>Bio: {{.Profile.Bio}}</p>
<p>Email: {{.Profile.Email}}</p>
`

	htmlT, _ := template.New("html").Parse(htmlTmpl)
	var buf bytes.Buffer
	htmlT.Execute(&buf, User{
		Name: "<script>alert('xss')</script>",
		Profile: Profile{
			Bio:   "Hacker",
			Email: "hacker@mail.com",
		},
	})
	fmt.Println("HTML Output (script escaped):")
	fmt.Println(buf.String())
	fmt.Println()

	// 6. NESTED TEMPLATE / TEMPLATE COMPOSITION
	fmt.Println("=== Nested Template ===")
	const base = `{{define "base"}}<html><body>{{template "content" .}}</body></html>{{end}}`
	const content = `{{define "content"}}<h2>{{.Name}}</h2><p>{{.Profile.Bio}}</p>{{end}}`

	nested := template.Must(template.New("base").Parse(base))
	template.Must(nested.New("content").Parse(content))

	nested.ExecuteTemplate(os.Stdout, "base", User{
		Name: "Anggi",
		Profile: Profile{Bio: "Go Developer"},
	})
	fmt.Println()

	// 7. MULTIPLE FILES
	fmt.Println("\n=== Multiple Template Files ===")
	// Simulasi parsing multiple files
	files := []string{"header.tmpl", "body.tmpl", "footer.tmpl"}
	// Dalam praktek: tmpl := template.Must(template.ParseFiles(files...))
	fmt.Println("Gunakan template.ParseFiles(files...) untuk multiple files")
}

/*
Cara Run:
go run main.go

Untuk production:
- Simpan template di file terpisah (*.tmpl atau *.html)
- Parse dengan template.ParseFiles() atau template.ParseGlob()
- Cache template dengan template.Must()
- Untuk web: gunakan http.ServeFile atau embed dengan //go:embed
*/
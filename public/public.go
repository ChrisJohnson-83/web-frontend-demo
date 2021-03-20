package public

import (
	"embed"
	"io"
	"text/template"
)

//go:embed *
var files embed.FS

type IndexParams struct {
	Title   string
	Message string
}

func Index(w io.Writer, p IndexParams) error {
	return parse("index.html").Execute(w, p)
}

type DashboardParams struct {
	Title   string
	Message string
}

func Dashboard(w io.Writer, p DashboardParams) error {
	return parse("dashboard.html").Execute(w, p)
}

type ProfileShowParams struct {
	Title   string
	Message string
}

func ProfileShow(w io.Writer, p ProfileShowParams) error {
	return parse("profile/show.html").Execute(w, p)
}

type ProfileEditParams struct {
	Title   string
	Message string
}

func ProfileEdit(w io.Writer, p ProfileEditParams) error {
	return parse("profile/edit.html").Execute(w, p)
}

func parse(file string) *template.Template {
	return template.Must(
		template.New("layout.html").ParseFS(files, "layout.html", file))
}

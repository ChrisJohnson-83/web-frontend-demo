package public

import (
	"embed"
	"io"
	"text/template"
)

//go:embed *
var files embed.FS

type Menu struct {
	Item string
}

type Menus []Menu

type DashboardParams struct {
	Menus   Menus
	Title   string
	Message string
}

func Dashboard(w io.Writer, p DashboardParams) error {
	return parse("layout.html", "dashboard.html").Execute(w, p)
}

type ProfileShowParams struct {
	Menus   Menus
	Title   string
	Message string
}

func ProfileShow(w io.Writer, p ProfileShowParams) error {
	return parse("layout2.html", "profile/show.html").Execute(w, p)
}

type ProfileEditParams struct {
	Menus   Menus
	Title   string
	Message string
}

func ProfileEdit(w io.Writer, p ProfileEditParams) error {
	return parse("layout.html", "profile/edit.html").Execute(w, p)
}

// func parse(layout string, file string) *template.Template {
// 	return template.Must(
// 		template.New(layout).ParseFS(files, layout, "menu.html", file),
// 	)
// }

// func parse(layout string, directory string) *template.Template {
// 	return template.Must(
// 		template.New(layout).ParseGlob(directory + "/*.html"),
// 	)
// }

// func parse(layout string, directory string) *template.Template {
// 	return template.Must(
// 		template.Must(template.Must(template.New(layout).ParseGlob("public/layout/*.html")).ParseGlob("public/partial/*.html")).ParseGlob("public/" + directory + "/*.html"),
// 	)
// }

func parse(layout string, file string) *template.Template {
	return template.Must(
		template.Must(template.Must(template.New(layout).ParseGlob("public/layout/*.html")).ParseGlob("public/partial/*.html")).ParseFS(files, file),
	)
}

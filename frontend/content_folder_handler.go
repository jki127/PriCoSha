package main

import (
	"html/template"
	"net/http"
	b "pricosha/backend"
)

func contentFolderHandler(w http.ResponseWriter, r *http.Request) {
	logged, username := getUserSessionInfo(r)

	if !logged {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	urlParams := r.URL.Query()
	folderName := urlParams["fn"][0]

	pageData := struct {
		Logged       bool
		Username     string
		ContentItems []*b.ContentItem
		FolderName   string
	}{
		logged,
		username,
		b.GetContentInFolder(folderName, username),
		folderName,
	}

	t := template.Must(template.New("").ParseFiles("../web/template/content_folder.html",
		"../web/template/base.html"))
	t.ExecuteTemplate(w, "base", pageData)
}

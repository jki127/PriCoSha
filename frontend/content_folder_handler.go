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

func newFolderHandler(w http.ResponseWriter, r *http.Request) {
	logged, username := getUserSessionInfo(r)
	errCookie, err := r.Cookie("folderErr")
	var errExists bool
	var errMsg string

	if err == nil {
		errExists = true
		errMsg = errCookie.Value
	} else {
		errExists = false
	}
	if !logged {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	pageData := struct {
		Logged   bool
		Username string
		ErrMsg   string
		IfErr    bool
	}{
		logged,
		username,
		errMsg,
		errExists,
	}

	t := template.Must(template.New("").ParseFiles("../web/template/new_folder.html",
		"../web/template/base.html"))
	t.ExecuteTemplate(w, "base", pageData)
}

func createFolderHandler(w http.ResponseWriter, r *http.Request) {
	clearCookie(&w, r, "folderErr")
	logged, username := getUserSessionInfo(r)
	if !logged {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	folderName := r.FormValue("folder_name")
	if folderName == "" {
		cookie := http.Cookie{Name: "folderErr", Value: "Cannot use empty folder name"}
		http.SetCookie(w, &cookie)
		http.Redirect(w, r, "/folder/new", http.StatusFound)
		return
	}

	err := b.CreateFolder(folderName, username)
	if err != nil {
		cookie := http.Cookie{Name: "folderErr", Value: "Cannot use duplicate folder name"}
		http.SetCookie(w, &cookie)
		http.Redirect(w, r, "/folder/new", http.StatusFound)
		return
	}

	http.Redirect(w, r, "/", http.StatusFound)
}

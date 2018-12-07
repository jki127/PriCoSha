package main

import (
	"log"
	"net/http"
	b "pricosha/backend"
	"strconv"
)

func addTagHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("username")
	if err != nil {
		log.Println(`frontend: addTagHandler(): User is not logged in and thus 
			cannot tag someone`)
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	uTagger := cookie.Value

	id, err := strconv.Atoi(r.URL.Query()["id"][0])
	if err != nil {
		log.Println("frontend: addTagHandler(): Failed to grab id from URL query")
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	uTagged := r.PostFormValue("username")

	log.Println("frontend: id:", id, "uTagger:", uTagger, "uTagged:", uTagged)

	err = b.InsertTag(id, uTagger, uTagged)

	switch {
	case err == nil:
		// do nothing
	case err.Error() == "noview":
		log.Println("received noview error from b.InsertTag()")
	case err.Error() == "failed":
		log.Println("received failed error from b.InsertTag()")
	}

	http.Redirect(w, r, r.Header.Get("referer"), http.StatusFound)
	return
}

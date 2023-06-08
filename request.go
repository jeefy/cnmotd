package main

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/feeds"
)

func handleRequest(w http.ResponseWriter, r *http.Request) {
	feed := feeds.Feed{
		Title:       "CNCF MOTD",
		Link:        &feeds.Link{Href: "https://motd.cncf.io"},
		Description: "CNCF MOTD",
		Author:      &feeds.Author{Name: "CNCF MOTD", Email: "projects@cncf.io"},
	}
	motdRequests.Inc()

	projects := strings.Split(r.URL.Query().Get("projects"), ",")
	level := r.URL.Query().Get("level")
	intVar, err := strconv.Atoi(level)
	if err != nil {
		intVar = 0
	}

	for _, entry := range fullFeed.Filter(projects, MOTDLevel(intVar)) {
		motdServed.Inc()
		feed.Add(&entry.Item)
	}

	atom, err := feed.ToAtom()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/atom+xml")
	w.Header().Set("Cache-Control", "max-age=300")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(atom))
}

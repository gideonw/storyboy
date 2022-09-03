package main

import "net/http"

type Control struct {
	r Repo
}

func NewControl(r Repo) Control {
	return Control{
		r: r,
	}
}

func (c Control) newEntry(rw http.ResponseWriter, r *http.Request) {
	resp := c.r.CreateEntry("NPC", "Geralt", "Human")

	rw.Write([]byte(resp))
}

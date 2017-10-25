package main

import (
	"html/template"
	"math/rand"
	"net/http"
	"time"
)

type data struct {
	Data string
}

var gifs = [...]string{"https://raw.githubusercontent.com/ecliptik/gifs/master/agent-cooper-coffee-01.gif",
	"https://raw.githubusercontent.com/ecliptik/gifs/master/agent-cooper-coffee-02.gif",
	"https://raw.githubusercontent.com/ecliptik/gifs/master/big-trouble-are-you-crazy.gif",
	"https://raw.githubusercontent.com/ecliptik/gifs/master/big-trouble-invincible.gif",
	"https://raw.githubusercontent.com/ecliptik/gifs/master/big-trouble-rub-the-wrong-way.gif",
	"https://raw.githubusercontent.com/ecliptik/gifs/master/dr-strangelove-chew-gum.gif",
	"https://raw.githubusercontent.com/ecliptik/gifs/master/dr-strangelove-slim-pickens.gif",
	"https://raw.githubusercontent.com/ecliptik/gifs/master/hackers-floppy-draw.gif"}

// Gif :
// Returns random gif
func Gif(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().Unix())
	t := template.New("Gif")
	t, _ = t.Parse("<html><body><img src=\"{{.Data}}\"></body></html>")

	val := data{Data: gifs[rand.Intn(len(gifs))]}
	t.Execute(w, val)
}

func main() {
	http.HandleFunc("/gif", Gif)
	http.ListenAndServe(":8080", nil)
}

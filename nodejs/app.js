const express = require('express')
const app = express()

// Gifs to choose from
var gifs = ["https://raw.githubusercontent.com/ecliptik/gifs/master/agent-cooper-coffee-01.gif",
        "https://raw.githubusercontent.com/ecliptik/gifs/master/agent-cooper-coffee-02.gif",
        "https://raw.githubusercontent.com/ecliptik/gifs/master/big-trouble-are-you-crazy.gif",
        "https://raw.githubusercontent.com/ecliptik/gifs/master/big-trouble-invincible.gif",
        "https://raw.githubusercontent.com/ecliptik/gifs/master/big-trouble-rub-the-wrong-way.gif",
        "https://raw.githubusercontent.com/ecliptik/gifs/master/dr-strangelove-chew-gum.gif",
        "https://raw.githubusercontent.com/ecliptik/gifs/master/dr-strangelove-slim-pickens.gif",
        "https://raw.githubusercontent.com/ecliptik/gifs/master/hackers-floppy-draw.gif"]

app.get('/gif', function (req, res) {
    var rnd = Math.floor(Math.random() * gifs.length)
    var gif = gifs[rnd]
  res.send("<img src=\"" + gif + "\">")
})

app.listen(3000, function () {
  console.log('Example app listening on port 3000!')
})

require "sinatra"

#Gifs to choose from
gifs = ["https://raw.githubusercontent.com/ecliptik/gifs/master/agent-cooper-coffee-01.gif",
        "https://raw.githubusercontent.com/ecliptik/gifs/master/agent-cooper-coffee-02.gif",
        "https://raw.githubusercontent.com/ecliptik/gifs/master/big-trouble-are-you-crazy.gif",
        "https://raw.githubusercontent.com/ecliptik/gifs/master/big-trouble-invincible.gif",
        "https://raw.githubusercontent.com/ecliptik/gifs/master/big-trouble-rub-the-wrong-way.gif",
        "https://raw.githubusercontent.com/ecliptik/gifs/master/dr-strangelove-chew-gum.gif",
        "https://raw.githubusercontent.com/ecliptik/gifs/master/dr-strangelove-slim-pickens.gif",
        "https://raw.githubusercontent.com/ecliptik/gifs/master/hackers-floppy-draw.gif"]

#Output random gif
get "/gif" do
  gif = gifs.sample
  "<img src=\"#{gif}\">"
end

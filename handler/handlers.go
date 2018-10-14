package handler

import (
	"bytes"
	"image/png"
	"log"
	"os"

	"github.com/valyala/fasthttp"
)

func MainHandler(ctx *fasthttp.RequestCtx) {
	ctx.SendFile("static/index.html")
}

func ConvertHandler(ctx *fasthttp.RequestCtx) {
	body := ctx.PostBody()

	r := bytes.NewReader(body)
	myImage, err := png.Decode(r)
	if err != nil {
		log.Panic(err)
	}

	f, err := os.Create("png/1.png")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	png.Encode(f, myImage)
}

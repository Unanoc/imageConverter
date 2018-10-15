package handler

import (
	"crypto/md5"
	"fmt"
	"image/jpeg"
	"image/png"
	"imageConverter/logger"
	"io/ioutil"
	"os"
	"time"

	"github.com/valyala/fasthttp"
)

const (
	imagesDirPath = "static/images/"
)

func ImgHandler(ctx *fasthttp.RequestCtx) {
	fasthttp.ServeFile(ctx, string(ctx.Path())[1:]) // removing "/" in the begining of the path "/static/images/{image_name}.jpg"
}

func MainHandler(ctx *fasthttp.RequestCtx) {
	ctx.SendFile("static/index.html")
}

func ConvertHandler(ctx *fasthttp.RequestCtx) {
	formData, err := ctx.FormFile("fileToUpload")
	if err != nil {
		logger.LoggerInfo(err.Error())
		ctx.SetBody([]byte(err.Error()))
		return
	}

	inputFile, err := formData.Open()
	if err != nil {
		logger.LoggerInfo(err.Error())
		ctx.SetBody([]byte(err.Error()))
		return
	}

	inputPNG, err := png.Decode(inputFile)
	if err != nil {
		logger.LoggerInfo(err.Error())
		ctx.SetBody([]byte(err.Error()))
		return
	}

	pngBytes, err := ioutil.ReadAll(inputFile)
	if err != nil {
		logger.LoggerInfo(err.Error())
		ctx.SetBody([]byte(err.Error()))
		return
	}

	if _, err := os.Stat(imagesDirPath); os.IsNotExist(err) {
		os.Mkdir(imagesDirPath, 0777)
	}
	_ = os.Mkdir(imagesDirPath, 0777)

	imagePath := imagesDirPath + fmt.Sprintf("%x", md5.Sum(pngBytes)) + ".jpg"
	jpgImgFile, err := os.Create(imagePath)
	if err != nil {
		logger.LoggerInfo(err.Error())
		ctx.SetBody([]byte(err.Error()))
		return
	}
	defer jpgImgFile.Close()

	err = jpeg.Encode(jpgImgFile, inputPNG, nil)
	if err != nil {
		logger.LoggerInfo(err.Error())
		ctx.SetBody([]byte(err.Error()))
		return
	}

	// remove image in 3 minutes
	go removeImageInTime(imagePath)
	fullImageUrl := string(ctx.Host()) + "/" + imagePath
	ctx.SetBody([]byte(fullImageUrl))
}

func removeImageInTime(pathToFile string) {
	time.Sleep(180 * time.Second)
	os.Remove(pathToFile)
}

package main

import (
	"github.com/disintegration/imaging"
	"github.com/gin-gonic/gin"
	"github.com/kolesa-team/go-webp/encoder"
	"github.com/kolesa-team/go-webp/webp"
	"log"
)

func main() {
	engine := gin.Default()
	engine.POST("/webp", func(ctx *gin.Context) {
		file, err := ctx.FormFile("file")
		if err != nil {
			log.Println(err.Error())
			Error(ctx, "param error")
			return
		}
		open, err := file.Open()
		if err != nil {
			log.Println(err.Error())
			Error(ctx, "network error")
			return
		}
		decode, err := imaging.Decode(open)
		if err != nil {
			log.Println(err.Error())
			Error(ctx, "decode image error")
			return
		}
		options, err := encoder.NewLossyEncoderOptions(encoder.PresetDefault, 75)
		options.LowMemory = true
		if err != nil {
			log.Println(err.Error())
			Error(ctx, "press param error")
			return
		}
		ctx.Writer.Header().Add("Content-Type", "image/webp")
		err = webp.Encode(ctx.Writer, decode, options)
		if err != nil {
			log.Println(err.Error())
			Error(ctx, "encode image error")
			return
		}
	})
	err := engine.Run(":8081")
	if err != nil {
		panic(err.Error())
	}
}

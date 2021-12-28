package main

import (
	"github.com/disintegration/imaging"
	"github.com/gin-gonic/gin"
	"github.com/kolesa-team/go-webp/encoder"
	"github.com/kolesa-team/go-webp/webp"
	"image"
	"log"
	"strconv"
	"strings"
)

func main() {
	engine := gin.Default()
	// /webp/s*size/q*quality
	engine.POST("/webp/*quality", func(ctx *gin.Context) {
		pressParam := strings.ToLower(ctx.Param("quality"))
		split := strings.Split(strings.TrimPrefix(pressParam, "/"), "/")
		var resizedWidth int
		var resizedHeight int
		var lossless bool
		var quality float32
		var level int
		for _, s := range split {
			if len(s) == 0 {
				continue
			}
			if s == "lossless" {
				lossless = true
				continue
			}
			cmd := s[0]
			switch cmd {
			case 's':
				p := s[1:]
				index := strings.Index(p, "*")
				if index == -1 {
					Error(ctx, "size param error")
					return
				}
				widthStr := p[:index]
				heightStr := p[index+1:]
				width, err := strconv.Atoi(widthStr)
				if err != nil {
					Error(ctx, "size param error")
					return
				}
				height, err := strconv.Atoi(heightStr)
				if err != nil {
					Error(ctx, "size param error")
					return
				}
				resizedWidth = width
				resizedHeight = height
			case 'q': //quality
				q, err := strconv.Atoi(s[1:])
				if err != nil {
					Error(ctx, "quality param error")
					return
				}
				quality = float32(q)
			case 'l': //level
				l, err := strconv.Atoi(s[1:])
				if err != nil {
					Error(ctx, "level param error")
					return
				}
				level = l
			}
		}
		var options *encoder.Options
		if lossless {
			encoderOptions, err := encoder.NewLosslessEncoderOptions(encoder.PresetDefault, level)
			if err != nil {
				Error(ctx, "param error")
				return
			}
			options = encoderOptions
		} else {
			encoderOptions, err := encoder.NewLossyEncoderOptions(encoder.PresetDefault, quality)
			if err != nil {
				log.Println(err.Error())
				Error(ctx, "param error")
				return
			}
			options = encoderOptions
		}
		options.LowMemory = true
		options.Quality = quality
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
		resized := resizeImg(decode, resizedWidth, resizedHeight)
		decode = nil
		if err != nil {
			log.Println(err.Error())
			Error(ctx, "decode image error")
			return
		}
		ctx.Writer.Header().Add("Content-Type", "image/webp")
		err = webp.Encode(ctx.Writer, resized, options)
		if err != nil {
			log.Println(err.Error())
			Error(ctx, "encode image error")
			return
		}
	})
	err := engine.Run(":8889")
	if err != nil {
		panic(err.Error())
	}
}

func resizeImg(img image.Image, width int, height int) image.Image {
	x := img.Bounds().Size().X
	y := img.Bounds().Size().Y
	if width == 0 || height == 0 {
		return img
	}
	var w = x
	var h = y
	if x >= y && x > width {
		w = width
		h = y * width / x
	} else if y >= x && y > height {
		h = height
		w = x * height / y
	}
	if w != x || h != y {
		return imaging.Resize(img, w, h, imaging.Lanczos)
	}
	return img
}

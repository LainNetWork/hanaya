package main

import "github.com/gin-gonic/gin"

type Result struct {
	IsOk bool        `json:"isOk"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func NoAuth(ctx *gin.Context, msg string) {
	ctx.JSON(401, Result{
		IsOk: false,
		Msg:  msg,
	})
}

func Success(ctx *gin.Context, msg string) {
	ctx.JSON(200, Result{
		IsOk: true,
		Msg:  msg,
	})
}

func Error(ctx *gin.Context, msg string) {
	ctx.JSON(200, Result{
		IsOk: false,
		Msg:  msg,
		Data: []string{},
	})
}

func SuccessWithData(ctx *gin.Context, data interface{}) {
	ctx.JSON(200, Result{
		IsOk: true,
		Data: data,
	})
}

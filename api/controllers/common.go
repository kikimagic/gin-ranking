package controllers

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/gin-gonic/gin"
)

type JsonInfo struct {
	Code  int         `json:"code"`
	Msg   interface{} `json:"msg"`
	Data  interface{} `json:"data"`
	Count int64       `json:"count"`
}

type JsonError struct {
	Code int         `json:"code"`
	Msg  interface{} `json:"msg"`
}

func ReturnSuccess(c *gin.Context, code int, msg interface{}, data interface{}, count int64) {
	json := &JsonInfo{Code: code, Msg: msg, Data: data, Count: count}
	c.JSON(200, json)
}

func ReturnError(c *gin.Context, code int, msg interface{}) {
	json := &JsonError{Code: code, Msg: msg}
	c.JSON(200, json)
}

func EncryptPassword(password string) string {
	ctx := md5.New()
	ctx.Write([]byte(password))
	return hex.EncodeToString(ctx.Sum(nil))
}

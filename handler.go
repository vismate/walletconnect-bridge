package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	version    = "v0.0.1-beta"
	name       = "zhcppy"
	repository = "https://github.com/DIMO-Network/walletconnect-bridge"
)

func HealthHandler(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}

func HelloHandler(ctx *gin.Context) {
	ctx.String(http.StatusOK, "Hello! This is DIMO WalletConnect %s.", version)
}

func InfoHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, map[string]string{
		"name":       name,
		"version":    version,
		"repository": repository,
	})
}

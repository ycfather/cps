package main

import (
	"gopkg.in/gin-gonic/gin.v1"
	"net/http"
	"strings"
)

func CheckLoginMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authcookie := c.Query("authcookie")
		if authcookie == "" {
			buf := []string{
				cpsConfig.Sso.Url,
				"/bacupurl=http://",
				c.Request.Host,
				c.Request.RequestURI,
				"&site=", cpsConfig.Sso.Site}
			redirect_url := strings.Join(buf, "")
			c.Redirect(http.StatusPermanentRedirect, redirect_url)

			return
		}

		c.Next()
	}
}

func HeaderMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")

		c.Next()
	}
}

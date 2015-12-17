package base

import (
	"net/http"
	"strings"

	"fmt"
	"github.com/gin-gonic/gin"
)

func Form(fm interface{}, fn func(*gin.Context, interface{}) (interface{}, error)) func(c *gin.Context) {
	return func(c *gin.Context) {
		if err := c.Bind(fm); err == nil {
			data, err := fn(c, fm)
			Json(c, data, err)
		} else {
			c.JSON(http.StatusOK, gin.H{"ok": false, "errors": strings.Split(err.Error(), "\n")})
		}
	}
}

func Message(c *gin.Context, ok bool, msg string) {
	c.Redirect(http.StatusFound, fmt.Sprintf("/?ok=%v&msg=%s#message", ok, msg))
}

func Json(c *gin.Context, data interface{}, err error) {
	if err == nil {
		Success(c, data)
	} else {
		Fail(c, err)
	}
}

func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{"ok": true, "data": data})
}

func Fail(c *gin.Context, err error) {
	c.JSON(http.StatusOK, gin.H{"ok": false, "errors": []string{err.Error()}})
}

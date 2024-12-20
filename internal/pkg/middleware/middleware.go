package middleware

import "github.com/gin-gonic/gin"

var GinMiddlewares = defaultMiddlewares()

func defaultMiddlewares() map[string]gin.HandlerFunc {
	return map[string]gin.HandlerFunc{}
}

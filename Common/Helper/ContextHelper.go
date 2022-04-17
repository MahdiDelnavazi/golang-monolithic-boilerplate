package Helper

import "github.com/gin-gonic/gin"

type contextKey int8

const key contextKey = 2

func GetValues(context gin.Context) {
	//value, ok := context.Value(key)
}

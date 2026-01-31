package main

import (
	"github.com/gin-gonic/gin"
	"medi-biz/api-gateway/router"
	_ "medi-biz/common/init"
)

func main() {
	r := gin.Default()
	router.Router(r)
	r.Run(":8000")
}

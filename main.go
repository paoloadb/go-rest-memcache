package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

var myStore = map[string]any{}

func setHandler(ctx *gin.Context) {
	holder := make(map[string]any)
	if err := ctx.BindJSON(&holder); err != nil {
		ctx.String(400, "bad request")
		return
	}

	ctx.JSON(200, gin.H(holder))
	hldrky := holder["key"].(string)
	hldrvl := holder["value"].(map[string]any)
	myStore[hldrky] = hldrvl
	fmt.Println(myStore)
}

func getHandler(ctx *gin.Context) {
	q := ctx.Param("id")
	if myStore[q] == nil {
		ctx.String(400, "bad request")
		return
	}
	ctx.JSON(200, gin.H(myStore[q].(map[string]any)))
}

func deleteHandler(ctx *gin.Context) {
	q := ctx.Param("id")
	if myStore[q] == nil {
		ctx.String(400, "bad request")
		return
	}
	delete(myStore, q)
	ctx.String(200, "deleted "+q)
}

func main() {
var port string

if len(os.Args) < 2 {
	port = ":6666"
} else {
	port = ":" + os.Args[1]
}

	router := gin.Default()

	router.PUT("/set", setHandler)
	router.POST("/set", setHandler)
	router.GET("/get/:id", getHandler)
	router.DELETE("/del/:id", deleteHandler)

	router.Run(port)
}
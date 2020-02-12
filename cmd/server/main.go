package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/FengGeSe/terminal/model"
	"github.com/FengGeSe/terminal/service/cmd"
	"github.com/FengGeSe/terminal/util"
)

var router = gin.Default()

func init() {
	// LoggerWithFormatter middleware will write the logs to gin.DefaultWriter
	// By default gin.DefaultWriter = os.Stdout
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		// your custom format
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	router.Use(gin.Recovery())
}

func main() {
	// Simple group: v1
	v1 := router.Group("/v1")
	{
		v1.POST("/execute", executeEndpoint)
	}
	router.StaticFS("/file", http.Dir("public"))
	router.Run(":8080") // listen and serve on 0.0.0.0:8080
}

func executeEndpoint(c *gin.Context) {
	// input
	args := c.PostFormArray("args")
	// process
	params, err := model.LoadExecuteParams(args)
	if err != nil {
		c.JSON(500, gin.H{
			"data": util.WrapRed("请求参数错误"),
		})
		return
	}

	ctx := context.Background()
	ctx, _ = context.WithTimeout(ctx, 3*time.Second)

	svc := cmd.NewCmdService()
	result, err := svc.Execute(ctx, params)
	if err != nil {
		c.JSON(500, gin.H{
			"data": err.Error(),
		})
		return
	}
	// output
	c.JSON(200, gin.H{
		"data": result.ToJson(),
	})
}

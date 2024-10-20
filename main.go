package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os/exec"
)

var page = `
  <form method="POST" action="/poweroff">
    <button type="submit">POWEROFF</button>
  </form>
`

func main() {
	g := gin.Default()

	g.GET("", func(ctx *gin.Context) {
		ctx.Header("Content-Type", "text/html")
		ctx.String(200, page)
	})

	g.POST("poweroff", func(ctx *gin.Context) {
		fmt.Println("got poweroff")
		cmd := exec.Command("poweroff")
		out, err := cmd.Output()
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s\n", out)
		ctx.String(200, "ok")
	})

	g.Run(":9876")
}

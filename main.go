package main

import (
	"net/http"

	SpringBoot "github.com/go-spring/spring-boot"

	SpringWeb "github.com/go-spring/spring-web"

	_ "github.com/go-spring/starter-gin"
)

func main() {
	SpringBoot.GetMapping("/", func(ctx SpringWeb.WebContext) {
		ctx.String(http.StatusOK, "hello world")
	})

	SpringBoot.RunApplication()
}

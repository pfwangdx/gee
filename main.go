package main

import (
	"engine"
)

func main() {
	eng := engine.New()

	eng.GET("/", func(context *engine.Context) {
		context.Html("<h1> hello gee <h1>")
	})

	// curl "http://localhost:9999/hello?name=geektutu"
	eng.GET("/hello", func(context *engine.Context) {
		context.String("hello %s\n", context.Query("name"))
	})

	// expect /hello/geektutu
	eng.GET("/hello/:name", func(context *engine.Context) {
		context.String("hello %s ", context.Params("name"))
	})

	// curl "http://localhost:9999/login" -X POST -d 'username=geektutu&password=1234'
	eng.POST("/login", func(context *engine.Context) {
		context.Json(map[string]string{
			"username": context.PostForm("username"),
			"password": context.PostForm("password"),
		})
	})

	eng.Run(":9999")
}

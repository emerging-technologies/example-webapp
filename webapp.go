// Author: Ian McLoughlin
// Date: October 6th, 2016
// Adapted from: https://go-macaron.com

package main

import "gopkg.in/macaron.v1"

func main() {
  m := macaron.Classic()
  m.Use(macaron.Renderer())
  
  m.Get("/macaron", func() string {
    return "Hello from Macaron!"
  })
  
  m.Get("/reverse/:name", func(ctx *macaron.Context) {
    // Adapted from: https://go-macaron.com/docs/middlewares/templating
    ctx.Data["Name"] = ctx.Params(":name")
    ctx.HTML(200, "hello")
  })
  
  m.Run()
}
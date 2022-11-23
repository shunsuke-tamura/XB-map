package main

import (
  "github.com/LoliGothic/lottery-map/controller"
)

func main() {
  router := controller.GetRouter()
  router.Run(":8080")
}
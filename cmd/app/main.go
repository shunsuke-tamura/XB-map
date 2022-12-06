package main

import (
  "github.com/LoliGothic/XB-map/controller"
)

func main() {
  router := controller.GetRouter()
  router.Run(":8080")
}
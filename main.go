package main
import (
    "github.com/zenazn/goji"
)


func main() {
     goji.Get("/employees/", employees)
     goji.Serve()
}


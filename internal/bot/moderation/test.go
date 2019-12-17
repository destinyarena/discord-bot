package main

import (
    "fmt"
    "regexp"
)

func main () {
    m, _ := regexp.Match(`^\d+$`, []byte("411323761116184578"))
    fmt.Println(m)
}

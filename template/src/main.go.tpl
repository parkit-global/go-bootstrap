package main

import "fmt"

func init() {
	InitConfig()
}

func main() {
    fmt.Println("{{.AppName}} Started")
}


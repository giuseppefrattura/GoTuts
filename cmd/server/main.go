package main

import "fmt"

func Run() error  {
    fmt.Println("Starting...")
    return nil
}

func main()  {
    fmt.Println("Hello world")
    if err := Run(); err!=nil{
        fmt.Println(err)
    }
}
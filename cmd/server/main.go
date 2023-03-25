package main

import (
    "context"
    "fmt"
    "github.com/TutorialEdge/go-rest-api-course/internal/db"
)

func Run() error  {
    fmt.Println("Starting...")

    db, err := db.NewDatabase()
    if err!=nil {
        fmt.Println("Failed to connect to database")
        return err
    }
    db.Ping(context.Background())
    return nil
}

func main()  {
    fmt.Println("Hello world")
    if err := Run(); err!=nil{
        fmt.Println(err)
    }
}
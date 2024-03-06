package main

import (
    "fmt"
    
    //"github.com/namoopsoo/orgroam2logseq"
    "testing"

)


func TestFoo(t *testing.T) {
    fmt.Println("hi")

    actual := MakeNewFileName("yo/cool")

    fmt.Printf("actual, %v", actual)
    if actual != "yo___cool.md" {
        t.Errorf("uhoh, %v, ", actual)
    }

    fmt.Println("bye")
    
}

func TestMigrate(t *testing.T) {
    fmt.Println("hi migrate test")
    sourceDir := "example"
    destinationDir := "temp"
    err := Migrate(sourceDir, destinationDir)

    if err != nil {
        t.Errorf("Oops: %v\n", err)
        }
    fmt.Println("bye")
}
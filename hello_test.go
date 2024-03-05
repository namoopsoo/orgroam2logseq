package main

import (
    "fmt"
    
    //"github.com/namoopsoo/orgroam2logseq"
    "testing"

)


func TestFoo(t *testing.T) {
    fmt.Println("hi")

    actual := main.MakeNewFileName("yo/cool")
    if actual != "yo___cool.md" {
        t.Errorf("uhoh, %v, ", actual)
    }
    
}
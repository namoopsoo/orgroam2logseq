package hello_main

import (
    "fmt"
    
    "github.com/namoopsoo/orgroam2logseq/hello"
    "testing"

)
import "testing"

func TestFoo(t *testing.T) {
    fmt.Println("hi")

    actual := hello.MakeNewFileName("yo/cool")
    if actual != "yo___cool.md" {
        t.Errorf("uhoh, %v, ", actual)
    }
    
}
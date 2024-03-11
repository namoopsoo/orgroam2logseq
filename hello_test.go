package main

import (
    "fmt"
    "testing"

    "github.com/namoopsoo/orgroam2logseq/okay"
)


func TestMakeNewFileName(t *testing.T) {
    fmt.Println("hi")

    actual := MakeNewFileName("yo/cool")

    fmt.Printf("actual, %v", actual)
    if actual != "yo___cool.md" {
        t.Errorf("uhoh, %v, ", actual)
    }

    fmt.Println("bye")
    
}

func TestMigrate(t *testing.T) {
    //fmt.Println("hi migrate test")

    // make temp
    //os.Mkdir("temp", os.FileMode(0777))
    //os.Mkdir("temp/journals", os.FileMode(0777))
    //os.Mkdir("temp/pages", os.FileMode(0777))

    err, _ := utils.ListDir(".")
    if err != nil {
        t.Errorf("listdir err %v", err)
    }
    // fmt.Printf("files %v\n", files)

    sourceDir := "example_org_roam"
    destinationDir := "temp"
    err = Migrate(sourceDir, destinationDir)

    if err != nil {
        t.Errorf("Oops: %v\n", err)
        }
    // fmt.Println("bye")
}

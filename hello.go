package main


import (
    "fmt"
    "os"
    "github.com/namoopsoo/orgroam2logseq/utils"
)

func Migrate(sourceDir string, destinationDir string) error {
    fmt.Print(sourceDir, "->", destinationDir)

    var lines []string
    lines, err = utils.ReadFileLines(sourceDir + "/example_org_roam/daily/2024-02-12.org")
    if err != nil {
        fmt.Println("error reading %v", err)
        return
    }
    fmt.Printf("lines \n\n%v\n\n", lines)

    return nil
}



func PrintHelp() {
    fmt.Print(`
Usage:
go run hello.go migrate sourceDir destinationDir

sourceDir: /path/to/org/roam/root
destinationDir: /path/to/clean/new/empty/logseq/graph/dir

`)
}


func main() {
    if len(os.Args) == 1 {
        PrintHelp()
        os.Exit(0)
    }
    switch os.Args[1] {
    case "migrate":
        sourceDir := os.Args[2]
        destinationDir := os.Args[3]
        err := Migrate(sourceDir, destinationDir)
        if err != nil {
            fmt.Fprintf(os.Stderr, "Oops: %v\n", err)
            os.Exit(1)
        }
    case "foo":
        fmt.Print("hi")
    }

}

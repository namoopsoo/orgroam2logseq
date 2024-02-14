package main


import (
    "fmt"
    "os"
    "github.com/namoopsoo/orgroam2logseq/okay"
    "regexp"
)

func BuildIdTitleMap() {
    var idMap map[string]string
    // for each file
    // FindIdTitle 
}

// find id, title for org file
func FindIdTitle(sourceDir string, destinationDir string) error {
    fmt.Print(sourceDir, "->", destinationDir)

    var lines []string
    lines, err := utils.ReadFileLines(sourceDir + "/daily/2024-02-12.org")

    if err != nil {
        fmt.Printf("error reading %v", err)
        return err
    }

    // ec22c32c-26b5-45a7-992-ff867494e7
    idRegexp := "(:ID:) ([a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{3}-[a-f0-9]{10})$"
    idRe := regexp.MustCompile(idRegexp)

    titleRegexp := "#+:title (.*)$"
    titleRe := regexp.MustCompile(titleRegexp)
    
    foundID := ""
    foundTitle := ""

    for i, line := range lines {
        matches := idRe.FindStringSubmatch(line)
        if len(matches) > 0 {
            fmt.Printf("%d, %v, match?\n", i, line)
            for j, m := range matches {
                fmt.Println(j, m)
            }
            foundID = matches[1]
        }

        // also match title        
        matches := titleRe.FindStringSubmatch(line)
        if len(matches) > 0 {
            fmt.Printf("%d, %v, match?\n", i, line)
            for j, m := range matches {
                fmt.Println(j, m)
            }
            foundTitle = matches[1]
        }

        // if find both id and title, break early
        if foundID != "" and foundTitle != "" {break}
    }

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

package main


import (
    "fmt"
    "os"
    "github.com/namoopsoo/orgroam2logseq/okay"
    "regexp"
    "net/url"
    "strings"
)

func BuildIdTitleMap(files []string) (map[string]string, error) {
    idMap := make(map[string]string)
    // for each file
    // FindIdTitle 
    for _, file := range files {
        // read, 
        id, title, err := FindIdTitle(file)
        if err != nil {
            return nil, fmt.Errorf("FindIdTitle: err %v", err)
        }

        fmt.Printf("file %v , id \"%v\", title \"%v\"\n\n", file, id, title)

        idMap[id] = title
    }
    return idMap, nil
}

// find id, title for org file
func FindIdTitle(filePath string) (string, string, error) {
    // fmt.Print(sourceDir, "->", destinationDir)

    var lines []string
    lines, err := utils.ReadFileLines(filePath)
    //sourceDir + "/daily/2024-02-12.org"

    if err != nil {
        fmt.Printf("error reading %v", err)
        return "", "", err
    }

    // ec22c32c-26b5-45a7-992-ff867494e7
    idRegexp := ":ID: ([a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12})$"
    idRe := regexp.MustCompile(idRegexp)

	titleRegexp := "[#][+]title: (.*)$"
    titleRe := regexp.MustCompile(titleRegexp)
    
    foundID := ""
    foundTitle := ""

    for i, line := range lines {
        //fmt.Printf("DEBUG: %v\n", line)
        matches := idRe.FindStringSubmatch(line)
        if len(matches) > 0 {
            fmt.Printf("%d, %v, match?\n", i, line)
            for j, m := range matches {
                fmt.Println(j, m)
            }
            foundID = matches[1]
        }

        // also match title        
        matches = titleRe.FindStringSubmatch(line)
        if len(matches) > 0 {
            fmt.Printf("%d, %v, match?\n", i, line)
            for j, m := range matches {
                fmt.Println(j, m)
            }
            foundTitle = matches[1]
        }

        // if find both id and title, break early
        if foundID != "" && foundTitle != "" {break}
    }
    return foundID, foundTitle, nil
    // return nil
}

func NewFileName(name string) string {
    // TODO error handling

    // lower
    s1 := strings.ToLower(name)

    // special characters -> percent encoded
    // https://www.urlencoder.io/golang/
    return url.QueryEscape(s1)
}

func Migrate(sourceDir string, destinationDir string) error {
    // copy/transform pages 

    // list all nonjournal files 
    err, files := utils.ListDir(sourceDir)
    if err != nil {
        return fmt.Errorf("listdir err %v", err)
    }
    

    var filePaths []string
    for _, fileName := range files {
        path := sourceDir + "/" + fileName
        fmt.Printf("path %v\n", path)
        filePaths = append(filePaths, path)
    }
    // return nil

    idMap, err := BuildIdTitleMap(filePaths)
    if err != nil {
        return fmt.Errorf("err %v", err)
    }
    
    fmt.Printf("id map, %v\n\n", idMap)

    // TODO and journal files too
    err, journalFiles := utils.ListDir(sourceDir + "/daily")
    if err != nil {
        return fmt.Errorf("listdir err %v", err)
    }
    fmt.Printf("journal files %v\n\n", journalFiles)

    // and transform !
    for _, fileName := range journalFiles {
        // new filename is %-encoded special characters and probably lower cased 
        newFileName := NewFileName(fileName)

        sourcePath := sourceDir + "/daily/" + fileName

        lines, err := utils.ReadFileLines(sourcePath)
        if err != nil {
            return fmt.Errorf("mmkay %v", err)
        }

        transformed := utils.TransformLines(lines, idMap)

        // write to new location 
        // pages dir
        newPath := destinationDir + "/" + newFileName
        err = utils.WriteLines(newPath, transformed)
        if err != nil {
            return fmt.Errorf("oops %v", err)
        }
    }

    // journals next 

    // assets next
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
    sourceDir := os.Args[2]
    destinationDir := os.Args[3]

    switch os.Args[1] {
    case "migrate":

        err := Migrate(sourceDir, destinationDir)
        if err != nil {
            fmt.Fprintf(os.Stderr, "Oops: %v\n", err)
            os.Exit(1)
        }
    case "foo":
        fmt.Print("hi")
    }

}

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
    idRegexp := `:ID:[\s]+([a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12})$`
    idRe := regexp.MustCompile(idRegexp)

	titleRegexp := `[#][+]title:[\s]+(.*)$`
    titleRe := regexp.MustCompile(titleRegexp)
    
    foundID := ""
    foundTitle := ""

    for _, line := range lines {
        //fmt.Printf("DEBUG: %v\n", line)
        matches := idRe.FindStringSubmatch(line)
        if len(matches) > 0 {
            // fmt.Printf("%d, %v, match?\n", i, line)
            for j, m := range matches {
                fmt.Println(j, m)
            }
            foundID = matches[1]
        }

        // also match title        
        matches = titleRe.FindStringSubmatch(line)
        if len(matches) > 0 {
            for j, m := range matches {
                fmt.Println(j, m)
            }
            foundTitle = matches[1]
        }

        // if find both id and title, break early
        if foundID != "" && foundTitle != "" {
            return foundID, foundTitle, nil
        }
    }

    return foundID, foundTitle, fmt.Errorf("Uh oh regex did not find id or title in %v", filePath)
    // return nil
}

func MakeNewFileName(name string) string {
    // TODO error handling

    // lower
    s1 := strings.ToLower(name)

    // replace / with ___
    s2 := strings.Replace(s1, "/", "___", -1)

    // special characters -> percent encoded
    // https://www.urlencoder.io/golang/
    s3 := url.QueryEscape(s2)

    // And spaces look like they dont need to be %20.
    return strings.Replace(s3, "+", " ", -1) + ".md"
}

func Migrate(sourceDir string, destinationDir string) error {
    // copy/transform pages 

    // list all nonjournal files 
    err, pageFiles := utils.ListDir(sourceDir)
    if err != nil {
        return fmt.Errorf("listdir err %v", err)
    }

    var filePaths []string
    for _, fileName := range pageFiles {
        path := sourceDir + "/" + fileName
        fmt.Printf("path %v\n", path)
        filePaths = append(filePaths, path)
    }

    var journalPaths []string
    err, journalFiles := utils.ListDir(sourceDir + "/daily")
    if err != nil {
        return fmt.Errorf("listdir err %v", err)
    }
    fmt.Printf("journal files %v\n\n", journalFiles)
    for _, x := range journalFiles {
        journalPaths = append(journalPaths, sourceDir + "/daily/" + x)
    }

    filePaths = append(filePaths, journalPaths...)

    idMap, err := BuildIdTitleMap(filePaths)
    if err != nil {
        return fmt.Errorf("err %v", err)
    }
    
    fmt.Printf("id map, %v\n\n", idMap)

    // and transform !
    for _, fileName := range journalFiles {

        sourcePath := sourceDir + "/daily/" + fileName

        // hmm although for journal files, date name
        // id, title, err := FindIdTitle(sourcePath)
        // find Id, title again?
        newFileName := strings.Replace(fileName, "-", "_", -1)
        newFileName = strings.Replace(fileName, ".org", ".md", 1)

        lines, err := utils.ReadFileLines(sourcePath)
        if err != nil {
            return fmt.Errorf("mmkay %v", err)
        }
        transformed := utils.TransformLines(lines, idMap)

        // write to new location 
        newPath := destinationDir + "/journals/" + newFileName
        err = utils.WriteLines(newPath, transformed)
        if err != nil {
            return fmt.Errorf("oops %v", err)
        }
    }

    // and transform pages too TODO dont copypasta
    for _, fileName := range pageFiles {
        newFileName := MakeNewFileName(fileName)

        sourcePath := sourceDir + "/" + fileName

        _, title, err := FindIdTitle(sourcePath)
        // find Id, title again?
        newFileName = MakeNewFileName(title)

        lines, err := utils.ReadFileLines(sourcePath)
        if err != nil {
            return fmt.Errorf("mmkay %v", err)
        }
        transformed := utils.TransformLines(lines, idMap)

        // write to new location 
        newPath := destinationDir + "/pages/" + newFileName
        err = utils.WriteLines(newPath, transformed)
        if err != nil {
            return fmt.Errorf("oops %v", err)
        }
    } 

    // assets next
    return nil
}

func FixLinksOneOff(workDir string) error {
    // list all  files 
    err, files := utils.ListDir(workDir)
    if err != nil {
        return fmt.Errorf("listdir err %v", err)
    }

    re := regexp.MustCompile(`\[\[([^\]]+)\]\[([^\]]+)\]\]`)

    replaceFn := func(m string) string {
        matches := re.FindStringSubmatch(m)
        if len(matches) == 3 {
            // matches[0] is the whole match, matches[1] is the id, matches[2] is the title
            // is it a url? 
            left := matches[1]
            right := matches[2]
            if strings.HasPrefix(left, "https://") {
                return fmt.Sprintf("[%s](%s)", right, left)
            }
        }
        return m // Return the original string if no replacement was made
    }

    for _, file := range files {
        // ok
        path := workDir + "/" + file
        // 
        //lines := readlines()
        lines, err := utils.ReadFileLines(sourcePath)
        if err != nil {
            return fmt.Errorf("mmkay %v", err)
        }

        var transformed []string
        for _, line := range lines {
            // 
            result := re.ReplaceAllStringFunc(line, replaceFn)
            transformed = append(transformed, result)

        }

        // write 
        err = utils.WriteLines(path, transformed)
        if err != nil {
            return fmt.Errorf("oops %v", err)
        }
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
    case "fixlinks":
        fmt.Print("hi")
        workDir := os.Args[2]
        FixLinksOneOff(workDir)
    }

}

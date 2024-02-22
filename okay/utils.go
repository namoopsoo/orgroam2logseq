package utils

import (
    "bufio"
    "fmt"
    "os"
    "strings"
    "regexp"
)

// ReadFileLines reads a file and returns its lines as a slice of strings.
func ReadFileLines(filePath string) ([]string, error) {
    // Check if the file exists
    if _, err := os.Stat(filePath); os.IsNotExist(err) {
        return nil, fmt.Errorf("file does not exist: %s", filePath)
    }

    // Open the file
    file, err := os.Open(filePath)
    if err != nil {
        return nil, fmt.Errorf("error opening file: %v", err)
    }
    defer file.Close()

    // Use bufio.Scanner to read the file line by line
    var lines []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }

    // Check for errors during scanning
    if err := scanner.Err(); err != nil {
        return nil, fmt.Errorf("error reading file lines: %v", err)
    }

    return lines, nil
}


func WriteLines(path string, lines []string) error {
    // open a file 
    file, error := os.Create(path)
    if error != nil {
        return fmt.Errorf("error opening %v", error)
    }
    defer file.Close()

    writer := bufio.NewWriter(file)

    for _, line := range lines {
        // fmt.Fprintf(writer, line + "\n")
        if _, err := writer.WriteString(line + "\n"); err != nil {
            return fmt.Errorf("error")
        }
    }

    // flush
    if err := writer.Flush(); err != nil {
        return fmt.Errorf("error flush %v", err)
    }

    return nil
}

func TransformLines(
    lines []string, idMap map[string]string,
) []string {

    // line := "Foo okay [[1259-aefe3-36def][Apple company]] okay and great [[473a-26faae-473d][intel.com]] ah nice"
    //idMap := map[string]string{
    //    "1259-aefe3-36def": "Apple.com",
    //    "473a-26faae-473d": "Intel",
    //}
    // Regex to find patterns like [[id][title]]
    re := regexp.MustCompile(`\[\[([^\]]+)\]\[([^\]]+)\]\]`)
    idRe := regexp.MustCompile(`id:(.*)$`)
    
    // Replacement function
    replaceFn := func(m string) string {
        matches := re.FindStringSubmatch(m)
        if len(matches) == 3 {
            // matches[0] is the whole match, matches[1] is the id, matches[2] is the title
            // is it a url? 
            left := matches[1]
            right := matches[2]
            if strings.HasPrefix(right, "https://") {
                return fmt.Sprintf("[%s](%s)", right, left)
            }
            if strings.HasPrefix(left, "id:"){
                matches := idRe.FindStringSubmatch(left)
                if len(matches) > 0 {
                    theId := matches[1]
                    if newName, ok := idMap[theId]; ok {
                        return fmt.Sprintf("[[%s]]", newName) // Use the new name from the map
                    }
                }
            }

        }
        return m // Return the original string if no replacement was made
    }

    var transformed []string
    for _, line := range lines {

        // Perform the replacement
        result := re.ReplaceAllStringFunc(line, replaceFn)
    
        if line != result {
            fmt.Println("\nDEBUG")
            fmt.Println("Original:", line)
            fmt.Println("Modified:", result)
        }
        transformed = append(transformed, result)
    }
    return transformed

}


func ListDir(folderPath string) (error, []string) {

    var paths []string

    //folderPath := "./path/to/your/folder"

    // Read the directory contents
    files, err := os.ReadDir(folderPath)
    if err != nil {
        fmt.Println("Error reading directory:", err)
        return err, nil
    }

    for _, file := range files {
        // Check if the directory entry is a file and not a directory
        info, err := file.Info()
        if err != nil {
            fmt.Println("Error getting file info:", err)
            continue
        }

        if !info.IsDir() {

            // only org files
            if ! strings.HasSuffix(file.Name(), ".org") {
                continue
            }
            paths = append(paths, file.Name())
        // fmt.Println(file.Name()) // Print the file name
        }

    }
    return nil, paths
}

//func main() {
//    // Example usage
//    filePath := "path/to/your/file.txt"
//    lines, err := ReadFileLines(filePath)
//    if err != nil {
//        fmt.Println(err)
//        return
//    }
//
//    // Process the lines
//    for i, line := range lines {
//        fmt.Printf("Line %d: %s\n", i+1, line)
//    }
//}


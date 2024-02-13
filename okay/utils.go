package utils

import (
    "bufio"
    "fmt"
    "os"
    // "strings"
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


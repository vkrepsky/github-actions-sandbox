package main

import (
    "fmt"
    "io"
    "os"
)

func main() {
    filePath := "dummy.txt"
    newDirPath := "dir"

    if err := os.Mkdir(newDirPath, os.ModePerm); err != nil {
        panic(err)
    }

    file, err := os.Open(filePath)
    if err != nil {
        panic(err)
    }
    defer file.Close()

    newFilePath := fmt.Sprintf("%s/%s", newDirPath, file.Name())
    newFile, err := os.Create(newFilePath)
    if err != nil {
        panic(err)
    }
    defer newFile.Close()

    _, err = io.Copy(newFile, file)
    if err != nil {
        panic(err)
    }

    fmt.Printf("File %s moved to directory %s\n", filePath, newDirPath)

    if err := os.Remove("dummy.txt"); err != nil {
        panic(err)
    }
}

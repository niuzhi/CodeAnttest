package main

import (
	"fmt"
	"os"
)

func readFileContent(filename string) (string, error) {
    data, err := os.ReadFile(filename)
    if err != nil {
        return "", err
    }
    return string(data), nil
}

func main() {
    content, err := readFileContent("nonexistent_file.txt")
    if err != nil {
        fmt.Println("读取文件失败：", err)
        return
    }
	fmt.Println("文件内容：", content)
}

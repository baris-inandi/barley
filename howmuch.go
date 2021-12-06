package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	gi "github.com/sabhiram/go-gitignore"
)

func parseTrackedFile(path string, name string, size int64) {
  // fmt.Println(path)
  // fmt.Println(name)
  // fmt.Println(size)
}

func main() {
  x := gi.CompileIgnoreLines(".gitignore")
  fmt.Println(x)
  err := filepath.Walk(".",
    func(path string, info os.FileInfo, err error) error {
    if err != nil {
        return err
    }
    parseTrackedFile(path, info.Name(), info.Size())
    return nil
})
if err != nil {
    log.Println(err)
}
}

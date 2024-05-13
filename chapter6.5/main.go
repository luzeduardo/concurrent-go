package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

func fileSearch(dir string, filename string, wg *sync.WaitGroup) {
	files, _ := os.ReadDir(dir)
	for _, file := range files {
		fpath := filepath.Join(dir, file.Name())
		if strings.Contains(file.Name(), filename) {
			fmt.Println(fpath)
		}
		if file.IsDir() {
			wg.Add(1)
			go fileSearch(fpath, filename, wg)
		}
	}
	wg.Done()
}

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	dir := os.Args[1]
	filename := os.Args[2]
	go fileSearch(dir, filename, &wg)
	wg.Wait()
}

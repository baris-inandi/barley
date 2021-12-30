package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
)

var langStore = map[string]int64{}
var langMap map[string]string = loadLanguagesJSON()
var totalSize int64 = 0

type file struct {
	name string
	size int64
}

func loadLanguagesJSON() map[string]string {
	jsonFile, err := os.Open("languages.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var result map[string]string
	json.Unmarshal([]byte(byteValue), &result)
	return result
}

func getLang(filename string) string {
	filenameRunes := []rune(filename)
	reverse(filenameRunes)
	exensionChars := []rune{}
	for i := range filenameRunes {
		currentChar := filenameRunes[i]
		if currentChar == '.' {
			reverse(exensionChars)
			break
		}
		exensionChars = append(exensionChars, currentChar)
	}
	return langMap[string(exensionChars)]
}

func reverse(s interface{}) {
	n := reflect.ValueOf(s).Len()
	swap := reflect.Swapper(s)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		swap(i, j)
	}
}

func trackFile(langName string, f file, detailed bool) {
	langStore[langName] += f.size
	totalSize += f.size
	if detailed {
		sizeStr := strconv.Itoa(int(f.size))
		fmt.Printf("\033[0;34m" + langName + "\033[0m" + strings.Repeat(" ", (30-len(langName))))
		fmt.Printf(f.name)
		fmt.Printf("\033[0;32m" + " (" + sizeStr + ")" + "\033[0m")
		fmt.Println()
	}
}

func forFiles(f file, detailed bool) {
	fileLanguage := getLang(f.name)
	if fileLanguage != "" {
		trackFile(fileLanguage, f, detailed)
	}
}

func visualizeStore() {
	fmt.Println("\nTOTAL:")
	for langName, size := range langStore {
		percentage := strconv.Itoa(int((size*100)/totalSize)) + "%%"
		langPreviewStr := "\033[0;34m" + langName + "\033[0m" + "\033[0;32m" + " (" + strconv.Itoa(int(size)) + ")" + "\033[0m"
		fmt.Printf(langPreviewStr)
		fmt.Printf(strings.Repeat(" ", (70 - len(langPreviewStr))))
		fmt.Printf(percentage)
		fmt.Println()
	}
}

func visualizeBars() {
	color := []string{
		"\033[0;31m",
		"\033[0;32m",
		"\033[0;33m",
		"\033[0;34m",
		"\033[0;35m",
		"\033[0;36m",
		"\033[0;37m",
	}
	colorReset := "\033[0m"
	colorBlack := "\033[0;30m"
	i := 0
	barView := []string{}
	summary := []string{}
	for langName, v := range langStore {
		fiftyNumerator := uint8((v * 50) / totalSize)
		if fiftyNumerator >= 1 {
			langSummary := color[i] + "\n" + langName + " => " + strconv.Itoa(int(fiftyNumerator)*2) + "%" + " (size: " + strconv.Itoa(int(v)) + ")"
			barView = append(barView, color[i]+strings.Repeat("█", int(fiftyNumerator))+"▌")
			summary = append(summary, langSummary)
			i++
		}
	}
	reverse(barView)
	reverse(summary)
	barView = append(barView, colorReset)
	barViewString := strings.Join(barView[:], "")
	summaryString := strings.Join(summary[:], "")

	// get parent directory name
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	projectName := filepath.Base(wd)
	headerFmt := "\n" + projectName + " " + colorBlack + wd + colorReset + "\n"

	// print out all the formatted data
	fmt.Println(headerFmt)
	fmt.Println(barViewString)
	fmt.Println(summaryString + "\n")
}

func main() {
	help := false
	detailed := false
	args := os.Args
	for i := range args {
		// command line arguments
		switch args[i] {
		case "--detailed", "-d":
			detailed = true
		case "--help", "-h":
			help = true
			fmt.Println(`
barley
barley analyzes the programming languages used in the current directory and produces a GitHub-like programming language usage bar
COMMANDS:
--detailed -d: Also provides file-specific information
--help -h: Shows this help message
			`)
		}
	}
	if !help {
		err := filepath.Walk(".",
			func(path string, info os.FileInfo, err error) error {
				if err != nil {
					return err
				}
				switch info.Name() {
				case ".git":
					if info.IsDir() {
						return filepath.SkipDir
					}
				case "node_modules":
					if info.IsDir() {
						return filepath.SkipDir
					}
				case "package.json":
					if !info.IsDir() {
						return filepath.SkipDir
					}
				case "package-lock.json":
					if !info.IsDir() {
						return filepath.SkipDir
					}
				}
				if info.IsDir() && info.Name() == ".git" {
					return filepath.SkipDir
				}
				// perform analysis on files, all values stored in langStore
				forFiles(file{name: info.Name(), size: info.Size()}, detailed)
				return nil
			})
		if err != nil {
			log.Println(err)
		}
		// visualize global language store
		if detailed {
			visualizeStore()
		}
		// visualize results
		visualizeBars()
	}
}

// read lang file code source: https://github.com/LeeReindeer/github-colors/blob/go/github-colors.go

package main

import (
	"fmt"
	"strings"
	"os"
	"io/ioutil"
	"io"
	"gopkg.in/yaml.v2"
	"net/http"
)

type Lang struct {
	// hex code
	color string
}

const langSrcFile = "languages.yml"
const langSrcURL = "https://raw.githubusercontent.com/github/linguist/master/lib/linguist/languages.yml"

func main() {
	GetGithubColors()
}

func GetGithubColors() map[string]Lang {
	m := readFile()
	langMap := make(map[string]Lang)
	fmt.Printf("Find %v languages in Github\n", len(m))
	for name, attrs := range m {
		//fmt.Printf("%s: %v \n", name, attrs)
		attrsMap, ok := attrs.(map[interface{}]interface{})
		color, okk := attrsMap["color"]
		stringColor := fmt.Sprintf("%s", color)
		//remove space from name
		newName := strings.Replace(name, " ", "-", -1)
		if okk && ok {
			langMap[newName] = Lang{stringColor}
		} else {
			langMap[newName] = Lang{""}
		}
	}
	return langMap
}

func readFile() map[string]interface{} {
	// ファイルが存在すればダウンロードしない
	if _, err := os.Stat(langSrcFile); os.IsNotExist(err) {
		fmt.Println("start downloading...")
		downloadFile(langSrcFile, langSrcURL)
	}

	m := make(map[string]interface{})
	ymlBytes, err := ioutil.ReadFile(langSrcFile)
	checkErr(err)

	err = yaml.Unmarshal(ymlBytes, m)
	checkErr(err)
	return m
}

func downloadFile(filePath string, url string) (err error) {
	// Create the file
	out, err := os.Create(filePath)
	checkErr(err)
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	checkErr(err)
	defer resp.Body.Close()

	// Writer the body to file
	_, err = io.Copy(out, resp.Body)
	checkErr(err)

	return nil
}

func checkErr(err error) (hasErr bool) {
	hasErr = err == nil
	if err != nil {
		fmt.Println(err)
	}
	return hasErr
}

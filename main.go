package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	jsonFile, err := os.Open("data.json")

	check(err)

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var data [][]string

	json.Unmarshal(byteValue, &data)

	var sb strings.Builder

	for i := 0; i < len(data); i++ {
		if len(data[i][1]) == 0 {
			continue
		}
		sb.WriteString("<button aria-label=\"" + data[i][0] + "\" data-audio-src=\"" + data[i][2] + "\">" + data[i][1] + "</button>")
	}

	htmlFile, err := ioutil.ReadFile(".template.html")

	check(err)

	var nextHtmlFile string = strings.Replace(string(htmlFile), "%%HTML%%", sb.String(), 1)

	e := ioutil.WriteFile("index.html", []byte(nextHtmlFile), 0644)

	check(e)

	fmt.Println("Successfully wrote `index.html`")
}

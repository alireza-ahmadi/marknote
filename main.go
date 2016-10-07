package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/everdev/mack"
	"github.com/russross/blackfriday"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		log.Println("At least one argument is required")
		return
	}

	fmt.Println(args)

	title := ""
	if len(args) > 1 {
		title = args[1]
	} else {
		now := time.Now()
		title = now.Format("2006-01-02 15:04:05")
	}

	inputFile := args[0]

	content, err := ioutil.ReadFile(inputFile)
	if err != nil {
		log.Println(err)
		return
	}

	formattedContent := string(blackfriday.MarkdownCommon([]byte(content)))

	command := fmt.Sprintf(`create note title "%s" with html "%s" notebook {"Buffer"}`, title, formattedContent)

	if err = mack.Tell("Evernote", command); err != nil {
		log.Println(err)
	} else {
		log.Println("Note saved successfully")
		return
	}
}

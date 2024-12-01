package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"seanThakur.com/app/usernote"
)

func note() {
	title, content := getNoteData()
	notedata, err := usernote.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}
	notedata.Display()
	err = notedata.Save()
	if err != nil {
		fmt.Println("saving the file has failed")
		return
	}
	fmt.Println("Your file is saved")
}

func getNoteData() (string, string) {
	noteHeader := getUserData("Note header: ")

	noteContent := getUserData("Note Content: ")

	return noteHeader, noteContent
}

func getUserData(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}
	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}

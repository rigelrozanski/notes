package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	wb "github.com/rigelrozanski/wb/lib"
)

// name of the wb used to hold command notes
const wbName = "notes"

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("please provide the command to retrieve notes for")
		return
	}
	search := strings.Join(args, " ")

	// get the notes
	notesBz, found := wb.GetWBRaw(wbName)
	if !found {
		fmt.Println("can't find notes wb")
		return
	}

	// parse the notes file
	var notes map[string][]string
	json.Unmarshal(notesBz, &notes)
	if len(notes) == 0 {
		fmt.Println("notes wb is not json formatted")
		return
	}

	// print the notes
	note := notes[search]
	if len(note) == 0 {
		fmt.Println("can't find note")
		return
	}
	for _, line := range note {
		fmt.Println(line)
	}
}

package main

// https://note.com/knowledgework/n/nc4c0a24a9569
// ナレッジワークの記事を参考に実装

import (
	"fmt"
	"os/exec"
	"strings"
)

func parse(output string) map[string]string {
	rows := strings.Split(output, "\n")
	mimeTypes := make(map[string]string)

	for _, row := range rows {
		if row == "" {
			continue
		}
		row = strings.ReplaceAll(row, " ", "")
		ss := strings.Split(row, ":")
		mimeTypes[ss[0]] = ss[1]
	}
	return mimeTypes
}

func main() {

	cmd := exec.Command("sh", "-c", "file --mime-type *")
	// cmd := exec.Command("ls", "-l")
	out, err := cmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(out))

	mimeTypes := parse(string(out))
	for fn, mime := range mimeTypes {
		fmt.Println(fn, "=>", mime)
	}
}

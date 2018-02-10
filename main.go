package main

import (
	"bufio"
	"bytes"
	"fmt"
	"mime"
	"net/http"
	"os"
	"strings"
	"time"
)

const fixXJavascript = false

func init() {
	if fixXJavascript {

		buf := bytes.NewBuffer([]byte(`application/javascript				js
`))

		scanner := bufio.NewScanner(buf)
		for scanner.Scan() {
			fields := strings.Fields(scanner.Text())
			if len(fields) <= 1 || fields[0][0] == '#' {
				continue
			}
			mimeType := fields[0]
			for _, ext := range fields[1:] {
				if ext[0] == '#' {
					break
				}
				err := mime.AddExtensionType("."+ext, mimeType)
				if err != nil {
					fmt.Println(err)
				}
			}
		}
		if err := scanner.Err(); err != nil {
			panic(err)
		}

	}
}

func main() {
	f, err := os.Open("./foo.js")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeContent(w, r, "foo.js", time.Time{}, f)
	})

	http.ListenAndServe(":8000", handler)
}

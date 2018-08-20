package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type LogEntry struct {
	MessageTemplate string
	Properties      map[string]interface{}
}

func main() {
	r := regexp.MustCompile(`{(.*?)}`)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		var entry LogEntry
		line := scanner.Text()
		err := json.Unmarshal([]byte(line), &entry)

		if err == nil {
			rendered := r.ReplaceAllStringFunc(entry.MessageTemplate, func(m string) string {
				key := strings.Trim(m, "{}")
				return fmt.Sprintf("%v", entry.Properties[key])
			})

			fmt.Println(rendered)
		}
	}
}

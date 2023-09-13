package exec

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

func SearchForTags(s, key, value string) string {

	spattern := fmt.Sprintf("%s = \"%s\"", key, value)

	pattern := regexp.QuoteMeta(spattern)

	// Compile the regular expression pattern.
	regex := regexp.MustCompile(pattern)

	// Search for the pattern in the input string.
	match := regex.FindString(s)

	// Check if a match was found.
	if match != "" {
		return match
	} else {
		log.Println("Match not found")
		return ""
	}

}

func RemoveKey(s string) string {
	pattern := `".*?"`

	// Compile the regular expression pattern.
	regex := regexp.MustCompile(pattern)

	// Find the first match in the input string.
	match := regex.FindString(s)

	// Check if a match was found.
	if match != "" {
		// Remove the surrounding double quotes from the match.
		result := match[1 : len(match)-1]
		return result
	} else {
		return ""
	}
}

func TagCheck(s string) (string, string) {
	parts := strings.Split(s, ":")
	var key, value string
	if len(parts) == 2 {
		key = parts[0]
		value = parts[1]
	} else {
		log.Fatalln("Tagging should be in format key:value, eg. env:dev")
	}
	return key, value
}

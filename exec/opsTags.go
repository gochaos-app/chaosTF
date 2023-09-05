package exec

import (
	"log"
	"regexp"
)

func SearchForTags(s string, spattern string) string {
	escapedTag := regexp.QuoteMeta(spattern)
	pattern := escapedTag + `\s*=\s*"([^"]+)"`

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

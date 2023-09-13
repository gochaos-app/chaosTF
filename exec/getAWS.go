package exec

import (
	"log"
	"os"

	"github.com/hashicorp/hcl/v2/hclwrite"
)

func GetAWSResources(filepath string, tagKey string, tagValue string) []string {
	file, err := Open(filepath)

	if err != nil {
		log.Println("Error: ", err)
		os.Exit(1)
	}

	mapServices := CloudServices()

	var tokens hclwrite.Tokens

	var resourcesArray []string

	for _, block := range file.Body().Blocks() {
		if block.Type() == "resource" {
			if value, found := mapServices[block.Labels()[0]]; found {
				blockTags := block.Body().GetAttribute("tags")
				if blockTags != nil {
					tokenTags := blockTags.BuildTokens(tokens)
					unTag := string(tokenTags.Bytes()[:])
					Y := SearchForTags(unTag, tagKey, tagValue)
					if Y != "" {
						service := value
						resourcesArray = append(resourcesArray, service)
					}
					log.Println("Tag", tagKey+":"+tagValue, "found in block", block.Labels()[0], block.Labels()[1])
				} else {
					log.Println("Tags not found in block", block.Labels()[0], block.Labels()[1])
				}

			}
		}
	}
	return resourcesArray
}

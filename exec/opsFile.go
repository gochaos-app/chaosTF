package exec

import (
	"errors"
	"io/ioutil"
	"log"
	"os"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclwrite"
)

// Open the hcl file
func Open(filepath string) (*hclwrite.File, error) {
	content, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	file, diags := hclwrite.ParseConfig(content, filepath, hcl.Pos{Line: 1, Column: 1})
	if diags.HasErrors() {
		err := errors.New("an error occurred")
		if err != nil {
			return nil, err
		}
	}
	return file, nil
}

func ExecConfig(filepath string, pattern string) {
	resourcesFile := GetResources(filepath, pattern)
	providersFile := GetProviders(filepath)
	CreateConfig(resourcesFile, providersFile)
	CreateFile(resourcesFile, providersFile)
}

func GetRegionVariable(filepath string) string {
	file, err := Open(filepath)

	if err != nil {
		log.Println("Error: ", err)
		os.Exit(1)
	}

	var tokens hclwrite.Tokens
	var region string

	for _, block := range file.Body().Blocks() {
		if block.Type() == "variable" {
			blockTokens := block.Body().GetAttribute("default").Expr().BuildTokens(tokens)
			region = string(blockTokens.Bytes()[:])
		}
	}
	return region
}

func GetProviders(filepath string) map[string]string {
	file, err := Open(filepath)

	if err != nil {
		log.Println("Error: ", err)
		os.Exit(1)
	}

	mapProviders := CloudType()
	var region string

	var tokens hclwrite.Tokens
	providersFile := make(map[string]string)

	for _, block := range file.Body().Blocks() {
		if block.Type() == "provider" {
			if value, found := mapProviders[block.Labels()[0]]; found {
				blockAttribute := block.Body().GetAttribute("region")
				if blockAttribute != nil {
					blockTokens := blockAttribute.Expr().BuildTokens(tokens)
					unRegion := string(blockTokens.Bytes()[:])

					//Checking if region points to variable
					if unRegion[0:3] == "var" {
						region = RemoveKey(GetRegionVariable(filepath))
						providersFile[value] = region

					} else {
						region = RemoveKey(unRegion)
						providersFile[value] = region
					}

				} else {
					log.Println("Region not found in provider")

				}
			}
		}
	}
	return providersFile
}

func GetResources(filepath string, pattern string) map[string]string {
	file, err := Open(filepath)

	if err != nil {
		log.Println("Error: ", err)
		os.Exit(1)
	}
	mapServices := CloudServices()

	var tokens hclwrite.Tokens

	resourcesFile := make(map[string]string)

	for _, block := range file.Body().Blocks() {
		if block.Type() == "resource" {
			if value, found := mapServices[block.Labels()[0]]; found {
				blockTags := block.Body().GetAttribute("tags")
				if blockTags != nil {
					tokenTags := blockTags.BuildTokens(tokens)
					unTag := string(tokenTags.Bytes()[:])
					Y := SearchForTags(unTag, pattern)
					if Y != "" {
						tag := Y
						service := value
						resourcesFile[service] = tag
					}
				} else {
					log.Println("Tags not found in block")
				}

			}
		}
	}
	return resourcesFile
}

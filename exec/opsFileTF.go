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

func ExecConfig(filepath, key, value, chaosFile string) {

	provider, region := GetProviders(filepath)

	var resourcesFile []string
	switch provider {
	case "aws":
		resourcesFile = GetAWSResources(filepath, key, value)
	case "do":
		log.Fatalln("Provider not supported...")
	case "kubernetes":
		log.Fatalln("Provider not supported...")
	case "gcp":
		log.Fatalln("Provider not supported...")
	}

	CreateFile(resourcesFile, provider, region, key, value, chaosFile)
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

func GetProviders(filepath string) (string, string) {
	file, err := Open(filepath)

	if err != nil {
		log.Println("Error: ", err)
		os.Exit(1)
	}

	mapProviders := CloudType()
	var provider string
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
					provider = value
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

	if len(providersFile) > 1 {
		log.Println("Multiple providers found")
		log.Fatalln("Multiple cloud providers not supported")
	}

	return provider, region
}

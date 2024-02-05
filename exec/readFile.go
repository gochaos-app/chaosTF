package exec

import (
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclwrite"
)

var tokens hclwrite.Tokens

func getResources(file *hclwrite.File) ([]*hclwrite.Block, []string) {
	resourceBlocks := file.Body().Blocks()
	var resource []string
	var blocks []*hclwrite.Block
	for _, block := range resourceBlocks {
		if block.Type() == "resource" {
			resource = append(resource, block.Labels()[0])
			blocks = append(blocks, block)
		}
	}
	return blocks, resource
}

func openFile(path string) (*hclwrite.File, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	file, diags := hclwrite.ParseConfig(content, path, hcl.Pos{Line: 1, Column: 1})
	if diags.HasErrors() {
		return nil, errors.New("error parsing file")
	}
	return file, nil
}

func getLabels(block *hclwrite.Block, tag string) string {
	var unTag string
	blockTags := block.Body().GetAttribute("tags")

	parts := strings.Split(tag, ":")
	var key, value string
	if len(parts) == 2 {
		key = parts[0]
		value = parts[1]
	} else {
		log.Fatalln("Tagging should be in format key:value, eg. env:dev")
	}

	if blockTags != nil {
		tokenTags := blockTags.BuildTokens(tokens)
		pattern := fmt.Sprintf("%s = \"%s\"", key, value)
		regex := regexp.MustCompile(pattern)

		//match := regex.FindString(s)
		unTag = string(tokenTags.Bytes()[:])
		match := regex.FindString(unTag)

		// Check if a match was found.
		if match != "" {
			return match
		} else {
			return ""
		}
	}
	return ""
}

func LogicSingleFileRead(path, tags string) {
	hclFile, err := openFile(path)
	if err != nil {
		log.Fatal(err)
	}
	blocks, chaosResources := getResources(hclFile)

	if len(chaosResources) == 0 {
		fmt.Println("No resources detected")
		fmt.Println("ChaosTF doesnt support no resources. Please use at least one resource")
		return
	}
	var chaosResourceArray []string
	var providersArray []string
	var tagsResource []string
	var provider string
	if len(blocks) != len(chaosResources) {
		fmt.Println("The number of resources in the file does not match the number of chaos resources")
		return
	}

	for i := range chaosResources {
		service := chaosResources[i]
		block := blocks[i]

		provider = strings.Split(service, "_")[0]
		chaosResource := cloudServices()(service)

		if chaosResource == "" {
			continue
		} else {
			chaosResourceArray = append(chaosResourceArray, chaosResource)
			provider = cloudType()(provider)
			tagsResource = append(tagsResource, getLabels(block, tags))
			providersArray = append(providersArray, provider)

		}
		fmt.Println(provider)
		fmt.Println(chaosResourceArray)
		fmt.Println(tagsResource)
		fmt.Println(providersArray)
		fmt.Println("------")
	}

}

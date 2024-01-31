package exec

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclwrite"
)

func GetProviders(file *hclwrite.File) []string {
	providers := file.Body().Blocks()
	var provider []string
	for _, block := range providers {
		if block.Type() == "provider" {
			provider = append(provider, block.Labels()[0])
		}
	}
	fmt.Println(provider)
	return provider
}

func OpenFile(path string) (*hclwrite.File, error) {
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

func LogicRead(path string) {
	hclFile, err := OpenFile(path)
	if err != nil {
		log.Fatal(err)
	}
	GetProviders(hclFile)
}

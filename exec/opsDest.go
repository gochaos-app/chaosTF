package exec

import (
	"fmt"
	"log"
	"os"

	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/zclconf/go-cty/cty"
)

func CreateConfig(resources map[string]string, providers map[string]string) {
	for key_reso,  := range resources {
		for cloud, region := range providers {
			case "aws":

			case "digitalocean":
				fmt.Println("do")
			case "kubernetes":
				fmt.Println("k8s")
			case "gcp":
				fmt.Println("google")
			}
		}
	}

}

func CreateFile(resources map[string]string, providers map[string]string) {

	fmt.Println(resources)
	fmt.Println(providers)

	chaosFile := hclwrite.NewEmptyFile()
	hclFile, err := os.OpenFile("chaos-config.hcl", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println("Error: ", err)
	}
	rootBody := chaosFile.Body()
	// Setting app and description
	rootBody.SetAttributeValue("App", cty.StringVal("ChaosTF"))
	rootBody.SetAttributeValue("Description", cty.StringVal("Automatically generated by ChaosTF"))
	rootBody.AppendNewline()

	job := rootBody.AppendNewBlock("job", []string{"aws", "ec2"})
	jobBody := job.Body()
	jobBody.SetAttributeValue("region", cty.StringVal("us-east-1"))

	hclFile.Write(chaosFile.Bytes())
}

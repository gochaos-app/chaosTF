package config

type GenConfig struct {
	App         string      `hcl:"app"`
	Description string      `hcl:"description"`
	Job         []JobConfig `hcl:"job,block"`
}

type JobConfig struct {
	Region    string      `hcl:"region,optional"`
	Namespace string      `hcl:"namespace,optional"`
	Project   string      `hcl:"project,optional"`
	Cloud     string      `hcl:"cloud,label"`
	Service   string      `hcl:"service,label"`
	Chaos     ChaosConfig `hcl:"config,block"`
}

type ChaosConfig struct {
	Tag   string `hcl:"tag"`
	Chaos string `hcl:"chaos"`
	Count int    `hcl:"count"`
}

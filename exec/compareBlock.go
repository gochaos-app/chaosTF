package exec

func cloudType() func(string) string {
	innerMap := map[string]string{
		"aws":          "aws",
		"google":       "gcp",
		"kubernetes":   "kubernetes",
		"digitalocean": "do",
	}
	return func(key string) string {
		return innerMap[key]
	}
}

func cloudServices() func(string) string {
	innerMap := map[string]string{
		"aws_instance":              "ec2",
		"aws_s3_bucket":             "s3",
		"aws_lambda_function":       "lambda",
		"aws_autoscaling_group":     "ec2_autoscaling",
		"digitalocean_droplet":      "droplet",
		"digitalocean_loadbalancer": "load_balancer",
		"google_compute_instance":   "vm",
		"kubernetes_pod":            "pod",
		"kubernetes_deployment":     "deployment",
		"kubernetes_daemonset":      "daemonSet",
	}
	return func(key string) string {
		return innerMap[key]
	}
}

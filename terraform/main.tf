variable "name" {
  type = string
  default = "World"
}

terraform {
  required_providers {
    local = {
      source = "hashicorp/local"
      version = "2.4.0"
    }
  }
}

resource "local_file" "foo" {
  content  = "Hello, ${var.name}"
  filename = "${path.module}/hello.txt"
}
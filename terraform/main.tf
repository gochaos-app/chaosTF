terraform {
  required_providers {
    local = {
      source = "hashicorp/local"
      version = "2.4.0"
    }
  }
}

variable "region" {
  default = "us-east-1"
}

provider "aws" {
 region ="us-east-1"

}

resource "aws_instance" "foo" {
  ami           = "ami-005e54dee72cc1d00" # us-west-2
  instance_type = "t2.micro"

}

resource "aws_instance" "boo" {
  ami           = "ami-005e54dee72cc1d00" # us-west-2
  instance_type = "t2.micro"

  tags = {
    Name = "Hola"
    App = "example"
    env = "dev"
  }
}

resource "aws_instance" "test" {
  ami           = "ami-005e54dee72cc1d00" # us-west-2
  instance_type = "t2.micro"

  tags = {
    Name = "Hola"
    App = "example"
    env = "dev"
  }
}


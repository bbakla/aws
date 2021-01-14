variable "aws_region" {
  default = "eu-west-1"
}

variable "tags" {
  type    = map(string)
  default = null
}

variable "prefix" {
  default = "lambdaRestApi"
}

variable "pipeline_environment" {
  default = "dev"
}

variable "partition_key" {
  default = "email"
}

variable "dynamodb_table_name" {
  type = string
  default = "lambdaingouser"
}

locals {
  tags = {
    Name = "lambdaRestAPI"
    Prefix = var.prefix
  }
}
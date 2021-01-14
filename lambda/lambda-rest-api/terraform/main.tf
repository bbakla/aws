provider "aws" {
  region = var.aws_region
}

module "lambda_gateway_function" {
  source = "git::https://gitlab.com/aws46/terraform/modules/lambda-function.git?ref=tags/0.1.2"

  lambda_handler_file_path            = "../lambda/cmd/main.go "
  lambda_handler_zip_file_destination = "../lambda/bin"
  lambda_handler_zip_file_name        = "main.zip"

  tags                 = var.tags != null ? merge(var.tags, local.tags) : local.tags
  pipeline_environment = var.pipeline_environment
  prefix               = join("", [var.prefix, "apigateway"])


   statements = [
    {
      sid       = "${var.prefix}AllowDynamo"
      effect    = "Allow"
      resources = [module.dynamodb_table.table_arn]
      action = [
        "dynamodb:PutItem",
        "dynamodb:GetItem",
        "dynamodb:Scan",
        "dynamodb:DeleteItem"
      ]
    }
  ]
}
module "dynamodb_table" {
  source = "git::https://github.com/cloudposse/terraform-aws-dynamodb.git?ref=tags/0.22.0"

  name =  var.dynamodb_table_name

  hash_key = var.partition_key
  //range_key = var.sort_key
  enable_autoscaler            = false
  autoscale_min_read_capacity  = 1
  autoscale_min_write_capacity = 1
}

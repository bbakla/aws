
resource "aws_api_gateway_rest_api" "gateway-api-attach-lambda" {
  name        = "attach-lambda"
  description = "This API attaches to a lambda to access the records in dynamodb"

  endpoint_configuration {
    types = [
      "REGIONAL"]
  }
}

module "api_using_remote_module" {
  source = "git::https://gitlab.com/aws46/terraform/modules/api-gateway.git?ref=tags/0.1.3"


  pipeline_environment = var.pipeline_environment
  prefix = "lambdarest"

  region = var.aws_region

  is_with_vpc_link_connected = false
  is_resource_attached_to_lambda = true
  is_cors_enabled = false


  lambda_invoke_arn = module.lambda_gateway_function.lambda_function_invoke_arn
  lambda_function_name = module.lambda_gateway_function.lambda_function_name
  lambda_arn = module.lambda_gateway_function.lambda_function_arn


  //authorizer_lambda_function_name = module.authorizer.lambda_function_name


  rest_api_id = aws_api_gateway_rest_api.gateway-api-attach-lambda.id
  parent_resource_id = aws_api_gateway_rest_api.gateway-api-attach-lambda.root_resource_id
  resource_http_method = "ANY"
  path_part = "api"
  allow_origin = "*"
  allow_credentials = false

  integration_input_type = "AWS_PROXY"

  tags = local.tags
}
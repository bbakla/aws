output "api_gateway_url" {
  value = module.api_using_remote_module.aws_api_gateway_deployment_url
}

output "http_resource_path" {
  value = module.api_using_remote_module.http_resource_path
}

output "http_resource_path_part" {
  value = module.api_using_remote_module.http_resource_path_part
}

output "lambda_function_name" {
  value = module.lambda_gateway_function.lambda_function_name
}

/*
EXAMPLE OUTPUT
-----------------------------

*/
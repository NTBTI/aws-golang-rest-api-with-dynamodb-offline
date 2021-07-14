# aws-golang-rest-api-with-dynamodb-offline

This is from the repo of Serverless examples, specifically the GO with DyanmoDB https://github.com/serverless/examples/tree/master/aws-golang-rest-api-with-dynamodb

I changed it to make seperate main.go files for each function, and to make it work with serverless-offline and DynamoDB-local in a Docker container


## Remember you had to do these things to make it work...

* You use Node 14.15.5 to install Serverless and all plugins https://github.com/dherault/serverless-offline/issues/1151
  * If you don't everything except GET calls will NOT work. They will not even get to the Lambda or the logs, it will simply hang

* Remove CGO from build params https://github.com/serverless/serverless/issues/8539#issuecomment-741752887

* GO only works with serverless-offline when using Docker ( --useDocker) https://github.com/dherault/serverless-offline/issues/381#issuecomment-696563404

* DynamoDB local in a container with admin GUI https://github.com/instructure/dynamo-local-admin-docker

* Remember to add --dockerNetwork to hit the DynamoDB container
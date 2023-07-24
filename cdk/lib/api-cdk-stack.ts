import {Construct} from 'constructs';
import * as cdk from 'aws-cdk-lib';
import { ApiStageProps } from './api-stage-props';

/**
 * The stack for running the api
 */
export class ApiCdkStack extends cdk.Stack {
  /**
   * The URL of the API Gateway endpoint
   */
  public readonly urlOutput: cdk.CfnOutput;
 
  constructor(scope: Construct, id: string, apiProps: ApiStageProps, props?: cdk.StackProps) {
    super(scope, id, props);

    // The Lambda function that contains the functionality
    const handler = new cdk.aws_lambda.Function(this, 'Lambda', {
      runtime: cdk.aws_lambda.Runtime.FROM_IMAGE,
      handler: cdk.aws_lambda.Handler.FROM_IMAGE,
      //grab the imgage from the infra account
      code: cdk.aws_lambda.Code.fromEcrImage(
        cdk.aws_ecr.Repository.fromRepositoryArn(this, 'ecr-repo', 'arn:aws:ecr:us-east-2:409783107836:repository/go-api'), {
        tagOrDigest: apiProps.version
      }),
      timeout: cdk.Duration.seconds(30),
      functionName: `${apiProps.stage}-api-handler`,
    });

    // An API Gateway to make the Lambda web-accessible
    const gw = new cdk.aws_apigateway.LambdaRestApi(this, 'Gateway', {
      restApiName: 'go-api',
      deployOptions: {
        stageName: apiProps.stage
      },
      description: 'Endpoint for the api',
      handler,
    });

    this.urlOutput = new cdk.CfnOutput(this, 'Url', {
      value: gw.url,
    });
  }
}
# cdk

This directory contains the CDK deployment pipeline code for the api

## Files and Directories

* **/bin** - the directory for the entry point of the program
    * [cdk.ts](./bin/cdk.ts) - the entrypoint for the CDK app
* **/lib** - the directory containing the stack files and stages
    * [api-app-stage](./lib/api-app-stage.ts) - the files that defines each application stage for the deploy pipeline. Essentially just a wrapper for the API stack
    * [api-cdk-stack](./lib/api-cdk-stack.ts) - the stack that defines every resource the API needs to run
    * [api-pipeline-stack](./lib/api-pipeline-stack.ts) - the file that defines the CI/CD pipeline built in CDK Pipelines for the api
    * [api-props](./lib/api-props.ts) - the parameters the api stack needs to run

## Testing Your Changes

To test your changes are effective and perform the desired outcome a couple approaches should be considered
* If making changes to the API stack itself (i.e. to [api-cdk-stack.ts](./lib/api-cdk-stack.ts) or files it uses); then you simply should run `cdk diff` or `cdk synth` and examine the results for your desired change.
* If modifying the pipeline, then it's important you make your changes to a **parallel pipeline first that is based off of your branch**. Failure to do this could disrupt the production deployment pipeline of the main branch, resulting in disruption of development or the API itself. 
    * To achieve this, simply update the ID of the stack in [cdk.ts](./bin/cdk.ts) (to something like `pipelinestack-bheil`) and the branch that the following line points to in [api-pipeline-stack](./lib/api-pipeline-stack.ts)
    * Once you save everything, run ```cdk deploy``` to deploy your changes to CodePipeline in the Infrastructure account. 

## Useful Commands

* `npm run build`   compile typescript to js
* `npm run watch`   watch for changes and compile
* `npm run test`    perform the jest unit tests
* `cdk deploy`      deploy this stack to your default AWS account/region
* `cdk diff`        compare deployed stack with current state
* `cdk synth`       emits the synthesized CloudFormation template


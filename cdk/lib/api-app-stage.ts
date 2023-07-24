import * as cdk from 'aws-cdk-lib'
import { Construct } from 'constructs';
import { ApiCdkStack } from './api-cdk-stack';
import { ApiStageProps } from './api-stage-props';

export class ApiAppStage extends cdk.Stage{

    constructor(parent: Construct, id: string, apiProps: ApiStageProps, props?: cdk.StageProps){
        super(parent, id, props);

        const stack = new ApiCdkStack(this, 'api-app-stack', apiProps, props);
    }

}
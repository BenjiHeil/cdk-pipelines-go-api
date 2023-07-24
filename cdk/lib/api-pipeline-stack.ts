import * as cdk from 'aws-cdk-lib';
import * as pipelines from 'aws-cdk-lib/pipelines'
import * as iam from 'aws-cdk-lib/aws-iam'
import * as codebuild from 'aws-cdk-lib/aws-codebuild'
import { Construct } from 'constructs';
import { ApiStageProps } from './api-stage-props';
import { ApiAppStage } from './api-app-stage';

export class ApiPipelineStack extends cdk.Stack {

    constructor(scope: Construct, id: string, version: string, props?: cdk.StackProps) {
        super(scope, id, props);

        const pipeline = new pipelines.CodePipeline(this, 'Api-Pipeline', {
            crossAccountKeys: true,
            pipelineName: 'ApiPipeline',
            synth: new pipelines.ShellStep('Synth', {
                input: pipelines.CodePipelineSource.gitHub('BenjiHeil/cdk-pipelines-go-api', 'main'),
                commands: [
                    'cd cdk',
                    'npm ci',
                    'npm run build',
                    'npx cdk synth',
                    'mv cdk.out ../cdk.out'
                ],
            }),
            dockerEnabledForSelfMutation: true,
            dockerEnabledForSynth: true,
        });

        // sanbox account
        this.addAppStage(pipeline, 'sandbox', '682134063971', 'us-east-2', {
            stage: 'sandbox',
            version,
            commands: ['./build.sh']
        });

        // stage account
        this.addAppStage(pipeline, 'stage', '036162860521', 'us-east-2', {
            stage: 'stage',
            version,
            commands: ['./stage.sh']
            
        });

        // prod account
        this.addAppStage(pipeline, 'prod', '465364827210', 'us-east-2', {
            stage: 'prod',
            version,
            commands: []
        });
    }

    addAppStage(pipeline: pipelines.PipelineBase, stageId: string, accountId: string, region: string, apiProps: ApiStageProps) {
        const stage = pipeline.addStage(new ApiAppStage(this, stageId, apiProps, {
            env: {
                account: accountId,
                region: region
            },
        }),
            {
                pre: [
                    new pipelines.CodeBuildStep('build-image', {
                        commands: apiProps.commands,
                        role: new iam.Role(this, `build-${apiProps.stage}-image-role`, {
                            inlinePolicies: {
                                ['ecr-allow-all']: this.ecrPolicy,
                                ['sts-assume-role']: this.stsPolicy
                            },
                            assumedBy: new iam.ServicePrincipal('codebuild.amazonaws.com')
                        }),
                        buildEnvironment: {
                            buildImage: codebuild.LinuxBuildImage.STANDARD_7_0,
                            privileged: true
                        }
                    }),
                ],
            });
    }

    private ecrPolicy = new cdk.aws_iam.PolicyDocument({
        statements: [
            new cdk.aws_iam.PolicyStatement({
                sid: 'AllowAllECR',
                effect: cdk.aws_iam.Effect.ALLOW,
                resources: ['*'],
                actions: ["ecr:*"]
            })
        ]
    })

    private stsPolicy = new cdk.aws_iam.PolicyDocument({
        statements: [
            new cdk.aws_iam.PolicyStatement({
                sid: 'AllowStsAssumeRole',
                effect: cdk.aws_iam.Effect.ALLOW,
                resources: ['*'],
                actions: ["sts:AssumeRole"]
            })
        ]
    })
}

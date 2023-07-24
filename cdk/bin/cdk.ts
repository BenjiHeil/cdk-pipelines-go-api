#!/usr/bin/env node
import 'source-map-support/register';
import * as cdk from 'aws-cdk-lib';
import { ApiPipelineStack } from '../lib/api-pipeline-stack';
import * as fs from 'fs';

//read semver file



const app = new cdk.App();
new ApiPipelineStack(app, 'ApiCdkPipelineStack', fs.readFileSync('../semver').toString(), {
  env: { account: '409783107836', region: 'us-east-2' }
});
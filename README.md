# Serverless using AWS CDK in Go

Serverless project for Go development with CDK using CDK Pipelines for continuous delivery.

**NOTICE**: Go support for AWS CDK is still in Developer Preview. This implies that APIs may change.

# Getting started
## Tools to install

* [Go >= 1.16](https://golang.org/doc/install)
* [Node.js ≥ 10.13.0](https://nodejs.org)
* [AWS CDK Toolkit](https://docs.aws.amazon.com/cdk/latest/guide/cli.html)

## How to deploy

`cdk bootstrap` deploys the CDK Toolkit stack. This is required for
using [CDK Assets](https://docs.aws.amazon.com/cdk/latest/guide/assets.html) to bundle Go lambda binary into CDK app.
For more info see [Bootstrapping](https://docs.aws.amazon.com/cdk/latest/guide/bootstrapping.html)

`cdk deploy` deploys the CI/CD stack. The pipeline is self-mutating, which means that if you add new application stages
in the source code, or new stacks to application, the pipeline will automatically reconfigure itself to deploy those new
stages and stacks.

## How it works

Pipeline gets the code from repository specified in `cdk/pipeline.go` file.

```
Owner:      jsii.String("czubocha"),
Repo:       jsii.String("aws-cdk-pipelines-go-serverless"),
```

It uses GitHub OAuth token which need to be provided in AWS Systems Manager Parameter Store in `github-token` String
parameter.   
For instructions to generate an GitHub OAuth token see
[Creating a personal access token.](https://docs.github.com/en/github/authenticating-to-github/keeping-your-account-and-data-secure/creating-a-personal-access-token)
The token should have the scopes `repo` and `admin:repo_hook`.

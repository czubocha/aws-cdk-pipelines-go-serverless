package main

import (
	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awscodepipeline"
	"github.com/aws/aws-cdk-go/awscdk/awscodepipelineactions"
	"github.com/aws/aws-cdk-go/awscdk/awsssm"
	"github.com/aws/aws-cdk-go/awscdk/pipelines"
	"github.com/aws/constructs-go/constructs/v3"
	"github.com/aws/jsii-runtime-go"
)

func NewPipelineStack(scope constructs.Construct, id string, props *awscdk.StackProps) awscdk.Stack {

	stack := awscdk.NewStack(scope, &id, props)

	cloudAssemblyArtifact := awscodepipeline.NewArtifact(nil)
	sourceOutputArtifact := awscodepipeline.NewArtifact(nil)
	gitHubToken := awsssm.StringParameter_ValueForStringParameter(stack, jsii.String("github-token"), nil)
	pipeline := pipelines.NewCdkPipeline(stack, jsii.String("pipeline"), &pipelines.CdkPipelineProps{
		CrossAccountKeys:      jsii.Bool(false), // only needed to perform cross-account deployments
		CloudAssemblyArtifact: cloudAssemblyArtifact,
		SourceAction: awscodepipelineactions.NewGitHubSourceAction(&awscodepipelineactions.GitHubSourceActionProps{
			ActionName: jsii.String("GitHub"),
			Output:     sourceOutputArtifact,
			OauthToken: awscdk.NewSecretValue(gitHubToken, nil),
			Owner:      jsii.String("czubocha"),
			Repo:       jsii.String("aws-cdk-pipelines-go-serverless"),
		}),
		SynthAction: pipelines.NewSimpleSynthAction(&pipelines.SimpleSynthActionProps{
			CloudAssemblyArtifact: cloudAssemblyArtifact,
			SourceArtifact:        sourceOutputArtifact,
			InstallCommands: jsii.Strings(
				// installing Go 1.16 as this version is required by aws-cdk-go (because of files embedding feature)
				// and AWS CodeBuild don't provide Docker images with Go 1.16 out of the box at the time of writing
				// https://docs.aws.amazon.com/codebuild/latest/userguide/runtime-versions.html
				"cd /root/.goenv && git pull && cd -",
				"goenv install 1.16.5",
				"goenv global 1.16.5",
				"npm i -g aws-cdk"),
			SynthCommand: jsii.String("cdk synth"),
		}),
	})

	staging := awscdk.NewStage(stack, jsii.String("staging"), nil)
	NewAppStack(staging, "app", nil)
	pipeline.AddApplicationStage(staging, nil)

	return stack
}

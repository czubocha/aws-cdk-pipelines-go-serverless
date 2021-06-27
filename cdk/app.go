package main

import (
	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awslambdago"
	"github.com/aws/constructs-go/constructs/v3"
	"github.com/aws/jsii-runtime-go"
)

func NewAppStack(scope constructs.Construct, id string, props *awscdk.StackProps) awscdk.Stack {

	stack := awscdk.NewStack(scope, &id, props)

	awslambdago.NewGoFunction(stack, jsii.String("hello"), &awslambdago.GoFunctionProps{
		Entry: jsii.String("functions/hello"),
	})

	return stack
}

package main

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

func main() {
	app := awscdk.NewApp(nil)

	NewPipelineStack(app, "cicd", nil)

	app.Synth(nil)
}

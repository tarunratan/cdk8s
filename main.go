package main

import (
	"getting-started-with-cdk8s-go/imports/k8s"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

type yarnChartProps struct {
	cdk8s.ChartProps
}

func NewyarnChart(scope constructs.Construct, id string, props *yarnChartProps) cdk8s.Chart {
	var cprops cdk8s.ChartProps
	if props != nil {
		cprops = props.ChartProps
	}
	chart := cdk8s.NewChart(scope, jsii.String(id), &cprops)

	selector := &k8s.LabelSelector{MatchLabels: &map[string]*string{"app": jsii.String("hello-yarn")}}

	labels := &k8s.ObjectMeta{Labels: &map[string]*string{"app": jsii.String("")}}

	yarnContainer := &k8s.Container{Name: jsii.String("yarn-app-container"), Image: jsii.String("yarn-demo-app"), Ports: &[]*k8s.ContainerPort{{ContainerPort: jsii.Number(3000)}}}

	deployment := k8s.NewKubeDeployment(chart, jsii.String("deployment"),
		&k8s.KubeDeploymentProps{
			Spec: &k8s.DeploymentSpec{
				Replicas: jsii.Number(1),
				Selector: selector,
				Template: &k8s.PodTemplateSpec{
					Metadata: labels,
					Spec: &k8s.PodSpec{
						Containers: &[]*k8s.Container{yarnContainer}}}}})

	service := k8s.NewKubeService(chart, jsii.String("service"), &k8s.KubeServiceProps{
		Spec: &k8s.ServiceSpec{
			Type:     jsii.String("LoadBalancer"),
			Ports:    &[]*k8s.ServicePort{{Port: jsii.Number(3000), TargetPort: k8s.IntOrString_FromNumber(jsii.Number(3000))}},
			Selector: &map[string]*string{"app": jsii.String("hello-yarn")}}})

	deployment.AddDependency(service)

	/*cdk8s.NewInclude(chart, jsii.String("existing service"), &cdk8s.IncludeProps{Url: jsii.String("service.yaml")})

	cdk8s.NewHelm(chart, jsii.String("bitnami yarn helm chart"), &cdk8s.HelmProps{
	Chart:  jsii.String("bitnami/yarn"),
	Values: &map[string]interface{}{"service.type": "ClusterIP"}})*/

	return chart
}

func main() {
	app := cdk8s.NewApp(nil)
	NewyarnChart(app, "yarn", nil)
	app.Synth()
}

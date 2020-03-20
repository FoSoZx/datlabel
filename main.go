package datlabel

import (
	"context"
	ce "github.com/FoSoZx/datlabel/error"
	"github.com/FoSoZx/datlabel/result"
	"github.com/FoSoZx/datlabel/utils"
	"github.com/docker/docker/api/types"
)

// Given a container id, the function returns the current labels only, without
// any field description.
func GetLabelsFromContainer(containerId string) (result.LabelResult, error) {
	cli := utils.NewDockerClient()
	containerDetails, err := cli.ContainerInspect(context.Background(),
		containerId)

	if err != nil {
		return nil, ce.NewNoSuchElement(containerId)
	}

	return result.NewLabelResult(containerDetails.Config.Labels), nil
}

// Given a service id, the function returns the service labels without any filed
// description
func GetLabelsFromService(serviceId string) (result.LabelResult, error) {
	cli := utils.NewDockerClient()
	serviceDetails, _, err := cli.ServiceInspectWithRaw(context.Background(),
		serviceId,
		types.ServiceInspectOptions{})

	if err != nil {
		return nil, ce.NewNoSuchElement(serviceId)
	}

	return result.NewLabelResult(serviceDetails.Spec.Labels), nil
}

// TODO finish ContainerResult implementation and perform container search
func ContainersFromLabels(label *result.Label) (result.ContainerResult, error) {
	return nil, nil
}

// TODO create ServiceResult implementation and perform service search
func ServicesFromLabels(label *result.Label) ([]string, error) {
	return []string{}, nil
}

// The idea here is to return all the labels a stack has, in order to collect
// them in a list
func GetLabelsFromStack(stackName string) ([]string, error) {
	// Steps to get the services in a Stack deployment:
	// 1 - Get all the services with the label "com.docker.stack.namespace"
	// 2 - Select all the services that have the stackName desired
	// 3 - From here, perform filtering and return the union of the labels of
	//     all the services in the stack
	return []string{}, nil
}

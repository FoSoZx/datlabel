package datlabel

import (
	"context"
	ce "github.com/FoSoZx/datlabel/error"
	"github.com/FoSoZx/datlabel/result"
	"github.com/FoSoZx/datlabel/utils"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"fmt"
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
		serviceId)
//		types.ServiceInspectOptions{})

	if err != nil {
		return nil, ce.NewNoSuchElement(serviceId)
	}

	return result.NewLabelResult(serviceDetails.Spec.Labels), nil
}

func ContainersFromLabel(label *result.Label) (result.ContainerResult, error) {
	cli := utils.NewDockerClient()

	filters := filters.NewArgs()
	filters.Add("label", fmt.Sprintf("%s=%s", label.Name(), label.Value()))

	containers, err := cli.ContainerList(context.Background(),
		types.ContainerListOptions{Filters: filters})

	if err != nil {
		return nil, ce.NewNoSuchElement(label.Name())
	}

	return result.NewContainerResult(containers), nil
}

func ServicesFromLabel(label *result.Label) (result.ServiceResult, error) {
	cli := utils.NewDockerClient()

	filters := filters.NewArgs()
	filters.Add("label", fmt.Sprintf("%s=%s", label.Name(), label.Value()))

	services, err := cli.ServiceList(context.Background(),
		types.ServiceListOptions{Filters: filters})

	if err != nil {
		return nil, ce.NewNoSuchElement(label.Name())
	}

	return result.NewServiceResult(services), nil
}

// The idea here is to return all the labels a stack has, in order to collect
// them in a list
func GetLabelsFromStack(stackName string) (result.LabelResult, error) {
	// Steps to get the services in a Stack deployment:
	// 1 - Get all the services with the label "com.docker.stack.namespace"
	// 2 - Select all the services that have the stackName desired
	//
	// 1 and 2 such as: docker service ls --filter "label=com.docker.stack.namespace=${stackName}" --format "{{.ID}}"
	//
	// 3 - From here, perform filtering and return the union of the labels of
	//     all the services in the stack

	cli := utils.NewDockerClient()

	filters := filters.NewArgs()
	filters.Add("label", fmt.Sprintf("%s=%s", "com.docker.stack.namespace", stackName))

	services, err := cli.ServiceList(context.Background(),
		types.ServiceListOptions{Filters: filters})

	if err != nil {
		return nil, ce.NewNoSuchElement(stackName)
	}

	labels := map[string]string{}

	for _, service := range services {
		if service.Spec.Labels != nil {
			for k, v := range service.Spec.Labels {
				labels[k] = v
			}
		}
	}

	return result.NewLabelResult(labels), nil
}

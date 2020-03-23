package result

import (
	"github.com/docker/docker/api/types/swarm"
)

// Struct that represents a container. It contains a pointer for the
// docker-defined structure and a list of labels.
// Generating []Label dynamically after the container started doesn't
// seem to be a reliable method, since it causes some tests to fail (
// see container_result_test.go/TestItShouldReturnRightContainers).
// The labels are then cached without relying on the Docker struct.
type Service struct {
	rawServiceDefinition *swarm.Service
	labels                 []Label
}

// Getter method to return the original docker container structure
func (c *Service) RawServiceDefinition() *swarm.Service {
	return c.rawServiceDefinition
}

// Getter method to return a list of labels
func (c *Service) Labels() []Label {
	return c.labels
}

// Getter method to return the container id
func (c *Service) Id() string {
	return c.rawServiceDefinition.ID
}

// Represents the result for a container search.
// It allows to get the list of containers found and to filter them
type ServiceResult interface {
	Result
	Services() []Service
	Filter(
		filter func(service *Service) *Service) (ServiceResult, error)
}

// Real ContainerResult interface implementation
type serviceResultImpl struct {
	ServiceResult
	services []Service
}

// Getter method to obtain the list of containers
func (c *serviceResultImpl) Services() []Service {
	return c.services
}

// Performs filtering operation on all the containers.
// A new ContainerResult is returned at the end of the operation,
// enabling the possibility to perform additional filtering.
func (c *serviceResultImpl) Filter(
	filter func(service *Service) *Service) (ServiceResult, error) {
	var result []Service
	for _, value := range c.services {
		filterResult := filter(&value)
		if filterResult != nil {
			result = append(result, *filterResult)
		}
	}

	return &serviceResultImpl{
		services: result,
	}, nil
}

// Returns a new ServiceResult object from a list of Docker Container types
func NewServiceResult(toEncapsulate []swarm.Service) ServiceResult {
	var services []Service
	for _, value := range toEncapsulate {

		var labels []Label
		for key, value := range value.Spec.Labels {
			labels = append(labels, Label{
				name:  key,
				value: value,
			})
		}

		services = append(services, Service{
			rawServiceDefinition: &value,
			labels:               labels,
		})
	}

	return &serviceResultImpl{
		services: services,
	}
}
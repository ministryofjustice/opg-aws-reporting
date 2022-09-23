package service

import (
	"fmt"
	"untagged/pkg/aws/endpoint"
)

// ServicesInPartition returns a slice of service names for the partition name passed along
func ServicesInPartition(partitionName string) ([]string, error) {

	resolver := endpoint.Resolver()
	partitions := endpoint.Partitions(resolver)
	services := []string{}
	foundService := false

	for _, partition := range partitions {
		if partition.ID() == partitionName {
			foundService = true
			for serviceID, _ := range partition.Services() {
				services = append(services, serviceID)
			}
		}

	}

	if !foundService {
		return services, fmt.Errorf("failed to find aws partition [%s]", partitionName)
	}
	return services, nil

}

// AllServices returns a slice of service names
func AllServices(partitionName string) ([]string, error) {

	resolver := endpoint.Resolver()
	partitions := endpoint.Partitions(resolver)
	services := []string{}

	for _, partition := range partitions {
		for serviceID, _ := range partition.Services() {
			services = append(services, serviceID)
		}
	}

	return services, nil
}

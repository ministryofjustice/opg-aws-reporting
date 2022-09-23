package region

import (
	"fmt"
	"untagged/pkg/aws/endpoint"
)

// RegionsInPartition returns a slice of region names for the partition name passed along
func RegionsInPartition(partitionName string) ([]string, error) {

	resolver := endpoint.Resolver()
	partitions := endpoint.Partitions(resolver)
	regions := []string{}
	foundPartition := false

	for _, partition := range partitions {
		if partition.ID() == partitionName {
			foundPartition = true
			for regionID, _ := range partition.Regions() {
				regions = append(regions, regionID)
			}
		}

	}

	if !foundPartition {
		return regions, fmt.Errorf("failed to find aws partition [%s]", partitionName)
	}
	return regions, nil

}

// AllRegions returns a slice of region names
func AllRegions() ([]string, error) {

	resolver := endpoint.Resolver()
	partitions := endpoint.Partitions(resolver)
	regions := []string{}

	for _, partition := range partitions {
		for regionID, _ := range partition.Regions() {
			regions = append(regions, regionID)
		}

	}

	return regions, nil

}

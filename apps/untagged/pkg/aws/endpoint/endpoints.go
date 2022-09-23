package endpoint

import "github.com/aws/aws-sdk-go/aws/endpoints"

func Resolver() endpoints.Resolver {
	return endpoints.DefaultResolver()
}

func Partitions(resolver endpoints.Resolver) []endpoints.Partition {
	return resolver.(endpoints.EnumPartitions).Partitions()
}

package main

import (
	"untagged/pkg/aws/availableservices/Iam"
	"untagged/pkg/aws/serviceregistry"
)

func init() {
	serviceregistry.RegisterTypes([]interface{}{
		(*Iam.Iam)(nil),
	})
}

func main() {

}

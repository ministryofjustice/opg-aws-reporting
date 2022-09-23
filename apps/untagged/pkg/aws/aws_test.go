package aws

import (
	"fmt"
	"sync"
	"testing"
	"untagged/pkg/aws/availableservices"
	"untagged/pkg/aws/availableservices/Iam"
	"untagged/pkg/aws/service"
	"untagged/pkg/aws/serviceregistry"
	"untagged/pkg/helpers"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/gammazero/workerpool"
)

// setup some registered types
func init() {
	serviceregistry.RegisterTypes([]interface{}{
		(*Iam.Iam)(nil),
	})
}

func TestAWS(t *testing.T) {
	services, _ := service.ServicesInPartition("aws")
	wp := workerpool.New(10)
	mu := &sync.Mutex{}

	for _, serviceName := range services {
		structName, _ := helpers.ConvertToStructName(serviceName)

		registered := serviceregistry.IsRegistered(structName)

		if registered {
			fmt.Printf("[%v]\n- structName: [%v]\n- structRegistered: [%v]\n", serviceName, structName, registered)

			newSess := session.New(&aws.Config{Region: aws.String("eu-west-1")})
			sess := session.Must(newSess, nil)

			instance := serviceregistry.Instance(structName).(availableservices.Service)
			resources, _ := instance.Resources(wp, mu, sess)
			instance.UntaggedResources(resources)

			//fmt.Printf("%v\n", untagged)

		}

	}
}

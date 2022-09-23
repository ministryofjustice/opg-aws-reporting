package availableservices

import (
	"sync"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/gammazero/workerpool"
)

// Service is the interface for all resources to process
type Service interface {
	// Resources should return a map keyed by ID with a list of tags
	Resources(wp *workerpool.WorkerPool, mu *sync.Mutex, s *session.Session) (map[string][]string, error)
	// UntaggedResources should return a slice of resource IDs and any error
	UntaggedResources(resources map[string][]string) ([]string, error)
}

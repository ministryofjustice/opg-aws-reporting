package service

import (
	"testing"
)

// TestServicesInPartition checks some data is returned for a correct partition
// and an error for a fake partition
func TestServicesInPartition(t *testing.T) {

	services, err := ServicesInPartition("aws")
	serviceLength := len(services)

	if err != nil {
		t.Errorf("[ServicesInPartition] failed to find services for [aws] partition:\n%v", err)
	}
	if serviceLength < 3 {
		t.Errorf("[ServicesInPartition] failed, expected at least 3 services, found [%v]", serviceLength)
	}

	services, err = ServicesInPartition("aws-that-does-not-exist")
	serviceLength = len(services)

	if err == nil {
		t.Errorf("[ServicesInPartition] should have generated an error for a fake region")
	}
	if serviceLength != 0 {
		t.Errorf("[ServicesInPartition] failed, expected 0 services, found [%v]", serviceLength)
	}
}

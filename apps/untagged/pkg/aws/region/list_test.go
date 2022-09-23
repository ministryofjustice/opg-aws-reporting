package region

import (
	"testing"
)

// TestRegionsInPartition checks some data is returned for a correct partition
// and an error for a fake partition
func TestRegionsInPartition(t *testing.T) {

	regions, err := RegionsInPartition("aws")
	regionLength := len(regions)

	if err != nil {
		t.Errorf("[RegionsInPartition] failed to find regions for [aws] partition:\n%v", err)
	}
	if regionLength < 3 {
		t.Errorf("[RegionsInPartition] failed, expected at least 3 regions, found [%v]", regionLength)
	}

	regions, err = RegionsInPartition("aws-that-does-not-exist")
	regionLength = len(regions)

	if err == nil {
		t.Errorf("[RegionsInPartition] should have generated an error for a fake region")
	}
	if regionLength != 0 {
		t.Errorf("[RegionsInPartition] failed, expected 0 regions, found [%v]", regionLength)
	}
}

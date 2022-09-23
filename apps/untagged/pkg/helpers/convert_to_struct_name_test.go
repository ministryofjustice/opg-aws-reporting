package helpers

import (
	"testing"
)

// TestConvertToStructName trry out a few service names and make sure they
// come back as expected
func TestConvertToStructName(t *testing.T) {

	testData := map[string]string{
		"data.jobs.iot": "DataJobsIot",
		"s3":            "S3",
		"sms-voice":     "SmsVoice",
		"ec2":           "Ec2",
		"R 53":          "R53",
		"R|53":          "R53",
	}

	for original, expected := range testData {
		converted, err := ConvertToStructName(original)

		if err != nil {
			t.Errorf("[ConvertToStructName] should not return an error:\n %v", err)
		}

		if converted != expected {
			t.Errorf("[ConvertToStructName] expected [%v], but recieved [%v]", expected, converted)
		}

	}

}

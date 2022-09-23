package helpers

import (
	"regexp"
	"strings"
)

// ConvertToStructName takes the aws service name (such as s3 / network-firewall) to
// a suitable Struct formatted name (S3 / NetworkFirewall)
func ConvertToStructName(str string) (string, error) {

	// lower case all
	str = strings.ToLower(str)
	// replace non alphanumerics with a space
	nonAlphanumericRegex := regexp.MustCompile(`[^a-zA-Z0-9 ]+`)
	str = string(
		nonAlphanumericRegex.ReplaceAll([]byte(str), []byte(" ")))

	// make title case
	str = strings.Title(str)
	// remove spaces
	str = strings.ReplaceAll(str, " ", "")
	return str, nil
}

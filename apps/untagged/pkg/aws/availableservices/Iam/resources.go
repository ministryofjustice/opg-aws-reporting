package Iam

import (
	"sync"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iam"
	"github.com/gammazero/workerpool"

	"github.com/rs/zerolog/log"
)

// Resource will find all resources create, use the arn or ID
// as the map key with a slice containing the tags
func (i Iam) Resources(
	wp *workerpool.WorkerPool,
	mu *sync.Mutex,
	s *session.Session,
) (map[string][]string, error) {

	log.Debug().Str("stack", "Iam::Resources").Msg("Start")

	resources := make(map[string][]string)
	svc := iam.New(s)

	// get base roles list
	roles, listError := i.ListRoles(svc)

	if listError != nil {
		return nil, listError
	}

	// get role details
	rolesAndTags, detailError := i.GetRoleDetail(wp, mu, svc, roles)
	if detailError != nil {
		return nil, detailError
	}

	// merge into the resources map
	for k, v := range rolesAndTags {
		resources[k] = v
	}

	log.Info().
		Str("stack", "Iam::Resources").
		Int("rolesAndTagsCount", len(rolesAndTags)).
		Int("resourcesCount", len(resources)).
		Msg("Result")

	return resources, nil
}

// GetRoleDetail uses the passed worker pool to fetch the tags for
// each role in the roles map
// Does this as there are likely dozens of roles, would be slow
// Returns map of arn => tags
// -----
// TODO - nested error handling
func (i *Iam) GetRoleDetail(
	wp *workerpool.WorkerPool,
	mu *sync.Mutex,
	svc *iam.IAM,
	roles map[string]string,
) (map[string][]string, error) {

	log.Debug().Str("stack", "Iam::GetRoleDetail").Msg("Start")

	rolesAndTags := make(map[string][]string)

	for arn, name := range roles {
		roleName := name
		roleArn := arn
		wp.Submit(func() {

			input := &iam.GetRoleInput{RoleName: aws.String(roleName)}
			detail, _ := svc.GetRole(input)

			log.Debug().
				Str("stack", "Iam::GetRoleDetail::wp").
				Str("RoleName", roleName).
				Str("ResultRoleName", *detail.Role.RoleName).
				Int("Tags", len(detail.Role.Tags)).
				Msg("GetRole result")

			mu.Lock()
			rolesAndTags[roleArn] = []string{}
			for _, tag := range detail.Role.Tags {
				rolesAndTags[arn] = append(rolesAndTags[arn], tag.GoString())
			}
			mu.Unlock()

		})
	}
	wp.StopWait()

	return rolesAndTags, nil
}

// ListRoles will go get all roles & return just their arns & names
func (i *Iam) ListRoles(svc *iam.IAM) (map[string]string, error) {

	log.Debug().Str("stack", "Iam::ListRoles").Msg("Start")

	arnAndRoleName := make(map[string]string)

	input := &iam.ListRolesInput{}
	result, err := svc.ListRoles(input)

	log.Debug().
		Str("stack", "Iam::ListRoles").
		Int("count", len(result.Roles)).
		Msg("Fetched roles from API")

	if err != nil {
		return arnAndRoleName, err
	}

	for _, role := range result.Roles {
		arn := string(*role.Arn)
		name := string(*role.RoleName)
		arnAndRoleName[arn] = name
	}

	log.Debug().
		Str("stack", "Iam::ListRoles").
		Int("count", len(arnAndRoleName)).
		Msg("arnAndRoleName")

	return arnAndRoleName, nil
}

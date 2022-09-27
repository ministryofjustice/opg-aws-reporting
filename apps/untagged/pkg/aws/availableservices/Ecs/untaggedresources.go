package Ecs

import (
	"untagged/pkg/helpers"

	"github.com/rs/zerolog/log"
)

func (e Ecs) UntaggedResources(resources map[string][]string) ([]string, error) {
	untagged, err := helpers.UntaggedResources(resources)

	log.Info().
		Str("stack", "Ecs::UntaggedResources").
		Int("untaggedResourcesCount", len(untagged)).
		Int("resourcesCount", len(resources)).
		Msg("Result")

	return untagged, err
}

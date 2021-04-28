// Code generated by gen.py. DO NOT EDIT.

package expconf

import (
	"github.com/santhosh-tekuri/jsonschema/v2"

	"github.com/determined-ai/determined/master/pkg/schemas"
)

func (r ReproducibilityConfigV0) ExperimentSeed() uint32 {
	if r.RawExperimentSeed == nil {
		panic("You must call WithDefaults on ReproducibilityConfigV0 before .RawExperimentSeed")
	}
	return *r.RawExperimentSeed
}

func (r *ReproducibilityConfigV0) SetExperimentSeed(val uint32) {
	r.RawExperimentSeed = &val
}

func (r ReproducibilityConfigV0) ParsedSchema() interface{} {
	return schemas.ParsedReproducibilityConfigV0()
}

func (r ReproducibilityConfigV0) SanityValidator() *jsonschema.Schema {
	return schemas.GetSanityValidator("http://determined.ai/schemas/expconf/v0/reproducibility.json")
}

func (r ReproducibilityConfigV0) CompletenessValidator() *jsonschema.Schema {
	return schemas.GetCompletenessValidator("http://determined.ai/schemas/expconf/v0/reproducibility.json")
}

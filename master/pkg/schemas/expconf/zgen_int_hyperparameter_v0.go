// Code generated by gen.py. DO NOT EDIT.

package expconf

import (
	"github.com/santhosh-tekuri/jsonschema/v2"

	"github.com/determined-ai/determined/master/pkg/schemas"
)

func (i IntHyperparameterV0) Minval() int {
	return i.RawMinval
}

func (i *IntHyperparameterV0) SetMinval(val int) {
	i.RawMinval = val
}

func (i IntHyperparameterV0) Maxval() int {
	return i.RawMaxval
}

func (i *IntHyperparameterV0) SetMaxval(val int) {
	i.RawMaxval = val
}

func (i IntHyperparameterV0) Count() *int {
	return i.RawCount
}

func (i *IntHyperparameterV0) SetCount(val *int) {
	i.RawCount = val
}

func (i IntHyperparameterV0) ParsedSchema() interface{} {
	return schemas.ParsedIntHyperparameterV0()
}

func (i IntHyperparameterV0) SanityValidator() *jsonschema.Schema {
	return schemas.GetSanityValidator("http://determined.ai/schemas/expconf/v0/hyperparameter-int.json")
}

func (i IntHyperparameterV0) CompletenessValidator() *jsonschema.Schema {
	return schemas.GetCompletenessValidator("http://determined.ai/schemas/expconf/v0/hyperparameter-int.json")
}

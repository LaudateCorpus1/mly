package domain

import (
	"context"
	"github.com/viant/mly/common"
	"github.com/viant/mly/common/storable"
	"reflect"
)

//Transformer represents output transformer
type Transformer func(ctx context.Context, signature *Signature, output interface{}) (common.Storable, error)

//Transform transform default model output
func Transform(ctx context.Context, signature *Signature, output interface{}) (common.Storable, error) {
	result := storable.New([]reflect.Type{reflect.TypeOf(output)})
	name := signature.Output.Name
	var outputValue interface{}
	switch val := output.(type) {
	case [][]float32:
		outputValue = val[0][0]
	case [][]float64:
		outputValue = val[0][0]
	case [][]string:
		outputValue = val[0][0]
	case [][]int64:
		outputValue = val[0][0]

	}
	err := result.Set(func(pair common.Pair) error {
		return pair(name, outputValue)
	})
	return result, err
}

package predict

import (
	"fmt"

	tf "github.com/tensorflow/tensorflow/tensorflow/go"
)

//  Predict predict default
// Valid values are scalars, slices, and arrays. Every element of a slice must have the same length so
// that the resulting Tensor has a valid shape.
func Predict(feedsName, fetchesName string, data interface{}) ([]*tf.Tensor, error) {
	m := GetModel(DefaultModelName)

	inputData, err := tf.NewTensor(data)
	if err != nil {
		return nil, err
	}

	input := map[tf.Output]*tf.Tensor{
		m.Graph.Operation(fmt.Sprint(feedsName)).Output(0): inputData,
	}
	output := []tf.Output{
		m.Graph.Operation(fmt.Sprint(fetchesName)).Output(0),
	}
	return m.Session.Run(input, output, nil)

}

// PredictWithModel predict with model
func PredictWithModel(tfModel *tf.SavedModel, feeds map[tf.Output]*tf.Tensor, fetches []tf.Output, targets []*tf.Operation) ([]*tf.Tensor, error) {
	return tfModel.Session.Run(
		feeds, fetches, targets,
	)
}

func PredictWithOperations(tfModel *tf.SavedModel, inputTensor *tf.Tensor, inputOperation, outputOperation string) ([]*tf.Tensor, error) {
	return tfModel.Session.Run(
		map[tf.Output]*tf.Tensor{
			tfModel.Graph.Operation(inputOperation).Output(0): inputTensor,
		},
		[]tf.Output{
			tfModel.Graph.Operation(outputOperation).Output(0),
		},
		nil,
	)
}

// PredictMini predict mini
func PredictMini(tfModel *tf.SavedModel, inputTensor *tf.Tensor) ([]*tf.Tensor, error) {
	return tfModel.Session.Run(
		map[tf.Output]*tf.Tensor{
			tfModel.Graph.Operation("input_tensor").Output(0): inputTensor,
		},
		[]tf.Output{
			tfModel.Graph.Operation("output_tensor").Output(0),
		},
		nil,
	)
}

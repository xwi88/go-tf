package predict

import (
	tf "github.com/tensorflow/tensorflow/tensorflow/go"
)

func predict(tfModel *tf.SavedModel, feeds map[tf.Output]*tf.Tensor, fetches []tf.Output, targets []*tf.Operation) ([]*tf.Tensor, error) {
	return tfModel.Session.Run(
		feeds, fetches, targets,
	)
}

func predictWithOperations(tfModel *tf.SavedModel, inputTensor *tf.Tensor, inputOperation, outputOperation string) ([]*tf.Tensor, error) {
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

func predictMini(tfModel *tf.SavedModel, inputTensor *tf.Tensor) ([]*tf.Tensor, error) {
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

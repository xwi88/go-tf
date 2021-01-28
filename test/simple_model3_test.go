package main

import (
	"fmt"
	"testing"

	tf "github.com/tensorflow/tensorflow/tensorflow/go"
)

func TestSimpleModel3(t *testing.T) {
	model, err := tf.LoadSavedModel("nasnet",
		[]string{"atag"}, nil)
	if err != nil {
		t.Fatal(err)
	}
	output, err := model.Session.Run(
		map[tf.Output]*tf.Tensor{
			model.Graph.Operation("input_1").Output(0): input,
		},
		[]tf.Output{
			model.Graph.Operation("predictions/Softmax").Output(0),
		},
		nil,
	)
	if err != nil {
		t.Fatal(err)
	}

	predictions, ok := output[0].Value().([][]float32)
	if !ok {
		t.Fatal(fmt.Sprintf("output has unexpected type %T", output[0].Value()))
	}
	fmt.Println(predictions)
}

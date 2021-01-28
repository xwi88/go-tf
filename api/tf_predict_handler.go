package api

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	tf "github.com/tensorflow/tensorflow/tensorflow/go"
)

func TFPredictHandler(ctx *fiber.Ctx) (err error) {
	tfModel, err := tf.LoadSavedModel("model", nil, nil)
	if err != nil {
		log.Fatal(err)
	}

	output, err := predict(ctx, tfModel)

	if err != nil {
		log.Fatal(err)
	}
	ret := fmt.Sprint(output[0].Value())
	fmt.Println(ret)

	ret2 := fmt.Sprint(output[0].Value().([][]float32))
	err = ctx.SendString(ret)
	fmt.Println(ret2)
	return
}

func predict(ctx *fiber.Ctx, tfModel *tf.SavedModel) ([]*tf.Tensor, error) {
	var predReqData struct{ Inputs [][][][]float32 }
	err := ctx.BodyParser(&predReqData)
	if err != nil {
		log.Fatal(err)
	}

	inputTensor, err := tf.NewTensor(predReqData)
	if err != nil {
		log.Fatal(err)
	}

	return tfModel.Session.Run(
		map[tf.Output]*tf.Tensor{
			tfModel.Graph.Operation("input_tensor").Output(0): inputTensor,
		},
		[]tf.Output{
			tfModel.Graph.Operation("softmax_tensor").Output(0),
		},
		nil,
	)
}

package api

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	tf "github.com/tensorflow/tensorflow/tensorflow/go"

	"github.com/xwi88/go-tf/predict"
)

func TFPredictHandler(ctx *fiber.Ctx) (err error) {
	output, err := predictDeal(ctx)

	if err != nil {
		log.Println(err)
	}
	ret := fmt.Sprint(output[0].Value())

	// ret2 := fmt.Sprint(output[0].Value().([]float32))
	err = ctx.SendString(ret)
	return
}

func predictDeal(ctx *fiber.Ctx) ([]*tf.Tensor, error) {
	var predReqData struct {
		Inputs []float32 `json:"data"`
	}
	err := ctx.BodyParser(&predReqData)
	if err != nil {
		log.Println(err)
	}
	log.Printf("body:%+v", predReqData)

	return predict.Predict("x", "y", predReqData.Inputs)
}

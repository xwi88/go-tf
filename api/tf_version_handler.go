package api

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	tf "github.com/tensorflow/tensorflow/tensorflow/go"
	"github.com/tensorflow/tensorflow/tensorflow/go/op"

	"github.com/xwi88/go-tf/tfm"
)

// VersionHandler TensorFlow 版本信息
func TFVersionHandler(ctx *fiber.Ctx) (err error) {
	output, err := tfm.EmptySession.Run(nil, []tf.Output{tfm.EmptyConst}, nil)
	if err != nil {
		panic(err)
	}
	ret := fmt.Sprint(output[0].Value())
	err = ctx.SendString(ret)
	fmt.Println(ret)
	return
}

// VersionHandler3 TensorFlow 版本信息
func TFVersionHandler3(ctx *fiber.Ctx) (err error) {
	err = ctx.SendString(tf.Version())
	return
}

// VersionHandler2 TensorFlow 版本信息
func TFVersionHandler2(ctx *fiber.Ctx) (err error) {
	// Construct a graph with an operation that produces a string constant.
	s := op.NewScope()
	c := op.Const(s, "Hello from TensorFlow version "+tf.Version())
	graph, err := s.Finalize()
	if err != nil {
		panic(err)
	}

	// Execute the graph in a session.
	sess, err := tf.NewSession(graph, nil)
	if err != nil {
		panic(err)
	}
	output, err := sess.Run(nil, []tf.Output{c}, nil)
	if err != nil {
		panic(err)
	}
	err = ctx.SendString(fmt.Sprint(output[0].Value()))
	return
}

package tfm

import (
	tf "github.com/tensorflow/tensorflow/tensorflow/go"
	"github.com/tensorflow/tensorflow/tensorflow/go/op"
)

var (
	EmptyScope   *op.Scope
	EmptyGraph   *tf.Graph
	EmptySession *tf.Session
	EmptyConst   tf.Output
)

func init() {
	var err error
	EmptyScope = op.NewScope()
	EmptyConst = op.Const(EmptyScope, "Hello from TensorFlow version "+tf.Version())
	EmptyGraph, err = EmptyScope.Finalize()
	if err != nil {
		panic(err)
	}
	EmptySession, err = tf.NewSession(EmptyGraph, nil)
	if err != nil {
		panic(err)
	}
}

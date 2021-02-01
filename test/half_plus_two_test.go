package main

import (
	"math/rand"
	"testing"
	"time"

	tf "github.com/tensorflow/tensorflow/tensorflow/go"
	"github.com/tensorflow/tensorflow/tensorflow/go/op"
)

func runHalfPlusTwo() {
	// Let's describe what we want: create the graph

	// We want to define two placeholder to fill at runtime
	// the first placeholder A will be a [2, 2] tensor of integers
	// the second placeholder x will be a [2, 1] tensor of intergers

	// Then we want to compute Y = Ax

	// Create the first node of the graph: an empty node, the root of our graph
	root := op.NewScope()

	// x := op.Placeholder(root.SubScope("input"), tf.Double, op.PlaceholderShape(tf.MakeShape(2, 1)))
	x := op.Placeholder(root.SubScope("input"), tf.Double)

	// Define the operation node that accepts  x as inputs
	y := op.Const(root.SubScope("input"), 2.0)
	z := op.Const(root.SubScope("input"), 2.0)
	// product := op.Div(root, x, y)
	// product = op.Add(root, product, z)

	product := op.Add(root, op.Div(root, x, y), z)

	// Every time we passed a `Scope` to an operation, we placed that
	// operation **under** that scope.
	// As you can see, we have an empty scope (created with NewScope): the empty scope
	// is the root of our graph and thus we denote it with "/".

	// Now we ask tensorflow to build the graph from our definition.
	// The concrete graph is created from the "abstract" graph we defined
	// using the combination of scope and op.

	graph, err := root.Finalize()
	if err != nil {
		// It's useless trying to handle this error in any way:
		// if we defined the graph wrongly we have to manually fix the definition.

		// It's like a SQL query: if the query is not syntactically valid
		// we have to rewrite it
		panic(err.Error())
	}

	// If here: our graph is syntatically valid.
	// We can now place it within a Session and execute it.

	var sess *tf.Session
	sess, err = tf.NewSession(graph, &tf.SessionOptions{})
	if err != nil {
		panic(err.Error())
	}

	// In order to use placeholders, we have to create the Tensors
	// containing the values to feed into the network
	var inputX *tf.Tensor

	// inputX, err = tf.NewTensor([]float64{1.0, 2.0, 3.0})

	rand.Seed(time.Now().UnixNano())
	inputS := generateSliceFloat(int(rand.Int31n(20)) + 1)
	inputX, err = tf.NewTensor(inputS)

	var results []*tf.Tensor
	if results, err = sess.Run(map[tf.Output]*tf.Tensor{
		x: inputX,
	}, []tf.Output{product}, nil); err != nil {
		panic(err.Error())
	}
	_ = results
	// for _, result := range results {
	// 	log.Printf(" input: %+v\n", inputS)
	// 	log.Printf("output: %+v\n", result.Value().([]float64))
	// }
}

func TestHalfPlusTwo(t *testing.T) {
	runHalfPlusTwo()
}

func Benchmark_HalfPlusTwo(b *testing.B) {
	n := b.N
	for i := 0; i < n; i++ {
		runHalfPlusTwo()
	}
}

func generateSliceFloat(size int) (s []float64) {
	rand.Seed(time.Now().UnixNano())
	s = make([]float64, size)
	for i := 0; i < size; i++ {
		s[i] = float64(rand.Int63n(200))
	}
	return s
}

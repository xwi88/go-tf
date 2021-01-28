package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	tf "github.com/tensorflow/tensorflow/tensorflow/go"
)

func main() {
	modelFile := "../models/saved_model_half_plus_two_cpu.pb"
	// modelFile := "../models/tensorflow_inception_graph.pb"
	if filesExist(modelFile) != nil {
		log.Fatalf("modelFile not exist")
	}

	model, err := ioutil.ReadFile(modelFile)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("read modelFile")

	// Construct an in-memory graph from the serialized form.
	graph := tf.NewGraph()
	if err := graph.Import(model, ""); err != nil {
		log.Fatal(err)
	}
	log.Printf("import graph")
	// Create a session for inference over graph.
	session, err := tf.NewSession(graph, nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("session create")
	defer session.Close()

	inputSlice := []float32{1.0, 2.0, 5.0}
	tensor, err := tf.NewTensor(inputSlice)
	if err != nil {
		log.Fatal(err)
	}

	output, err := session.Run(
		map[tf.Output]*tf.Tensor{
			graph.Operation("input").Output(0): tensor,
		},
		[]tf.Output{
			graph.Operation("output").Output(0),
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	probabilities := output[0].Value().([][]float32)[0]
	log.Printf("output:%+v", output)
	log.Printf("probabilities:%+v", probabilities)
}

func filesExist(files ...string) error {
	for _, f := range files {
		if _, err := os.Stat(f); err != nil {
			return fmt.Errorf("unable to stat %s: %v", f, err)
		}
	}
	return nil
}

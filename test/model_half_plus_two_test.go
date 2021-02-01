package main

import (
	"log"
	"math/rand"
	"testing"
	"time"

	tf "github.com/tensorflow/tensorflow/tensorflow/go"
)

var sm *tf.SavedModel

func runLoadModelHalfPlusTwo() {

	// log.Printf("tf.LoadSavedModel :%+v", sm)
	// log.Printf("tf.LoadSavedModel graph :%+v", sm.Graph)
	// log.Printf("tf.LoadSavedModel session :%+v", sm.Session)

	rand.Seed(time.Now().UnixNano())
	inputS := generateSliceFloat2(int(rand.Int31n(20)) + 1)
	inputData, err := tf.NewTensor(inputS)

	input := map[tf.Output]*tf.Tensor{
		sm.Graph.Operation("x").Output(0): inputData,
	}
	output := []tf.Output{
		sm.Graph.Operation("y").Output(0),
	}

	// devices, err := sm.Session.ListDevices()
	// log.Printf("devices:%+v, err:%v", devices, err)
	//
	// operations := sm.Graph.Operations()
	// for i, operation := range operations {
	// 	log.Printf("[%v] operation, name=%+v, inputs=%v, outputs=%v",
	// 		i, operation.Name(), operation.NumInputs(), operation.NumOutputs())
	//
	// }

	results, err := sm.Session.Run(input, output, nil)
	if err != nil {
		log.Fatalf("tf.LoadSavedModel err: %+v", err)
	}

	_ = results
	// for _, result := range results {
	// 	fmt.Printf(" input: %+v\n", inputS)
	// 	fmt.Printf("output: %+v\n", result.Value().([]float32))
	// }

}

func TestRunLoadModelHalfPlusTwo(t *testing.T) {
	loadSavedModelHalfPlusTwoCPU()
	runLoadModelHalfPlusTwo()
}

func Benchmark_RunLoadModelHalfPlusTwo(b *testing.B) {
	loadSavedModelHalfPlusTwoCPU()
	n := b.N
	for i := 0; i < n; i++ {
		runLoadModelHalfPlusTwo()
	}
}

func loadSavedModelHalfPlusTwoCPU() {
	var err error
	if sm == nil || sm.Graph == nil || sm.Session == nil {
		sm, err = tf.LoadSavedModel("../testdata/saved_model_half_plus_two_cpu/000001", []string{"serve"}, nil)
		if err != nil {
			log.Fatalf("tf.LoadSavedModel err: %v", err)
		}
	}
}

func generateSliceFloat2(size int) (s []float32) {
	rand.Seed(time.Now().UnixNano())
	s = make([]float32, size)
	for i := 0; i < size; i++ {
		s[i] = float32(rand.Int63n(200))
	}
	return s
}

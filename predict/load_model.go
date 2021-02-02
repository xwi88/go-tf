package predict

import (
	"log"
	"sync"

	tf "github.com/tensorflow/tensorflow/tensorflow/go"
)

const DefaultModelName = "latest"

var savedModels map[string]*tf.SavedModel
var loadLock sync.RWMutex

func LoadModel(exportDir string, tags []string, options *tf.SessionOptions, name string) (err error) {
	// TODO check exportDir
	loadLock.Lock()
	defer loadLock.Unlock()

	tfModel, err := tf.LoadSavedModel(exportDir, tags, options)
	if err != nil {
		return
	}

	if savedModels == nil {
		savedModels = make(map[string]*tf.SavedModel)
	}

	if len(name) == 0 {
		name = DefaultModelName
	}

	if _, ok := savedModels[name]; ok {
		savedModels[name] = tfModel
		log.Printf("repeated load model %v, with tags: %+v, in: %v", name, tags, exportDir)
	} else {
		log.Printf("first load model %v, with tags: %+v, in: %v", name, tags, exportDir)
	}
	savedModels[name] = tfModel
	return
}

func GetModel(name string) (model *tf.SavedModel) {
	loadLock.RLock()
	defer loadLock.RUnlock()
	return savedModels[name]
}

func ListModelNames() (names []string) {
	loadLock.RLock()
	defer loadLock.RUnlock()
	for name, _ := range savedModels {
		names = append(names, name)
	}
	return
}

func init() {
	// 	TODO load default model
	savedModels = make(map[string]*tf.SavedModel)
}

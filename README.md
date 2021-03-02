# QuickStart

## Sites

- [tensorflow](https://github.com/tensorflow/tensorflow)
- [tensorflow-serving](https://github.com/tensorflow/serving)

## Go API

>https://www.tensorflow.org/install/lang_go

## C lib install

- [download](https://www.tensorflow.org/install/lang_c#download)
- [current support:libtensorflow-cpu-darwin-x86_64-2.3.0](https://storage.googleapis.com/tensorflow/libtensorflow/libtensorflow-cpu-darwin-x86_64-2.3.0.tar.gz)

```bash
# Standard Install
sudo tar -C /usr/local -xzf (downloaded file)

# Non sys dir install
tar -C ~/mydir -xzf libtensorflow-cpu-darwin-x86_64-2.3.0.tar.gz

#ldconfig
export LIBRARY_PATH=$LIBRARY_PATH:~/mydir/lib
export DYLD_LIBRARY_PATH=$DYLD_LIBRARY_PATH:~/mydir/lib
```

## Run

```bash
# test tensorflow version
make test
make run
```

## Tips

- [version incompatible ref](https://github.com/tensorflow/tensorflow/issues/41808)
    - https://github.com/photoprism/photoprism/pull/775
- [hack fixed, only ref](https://github.com/tensorflow/tensorflow/blob/master/tensorflow/go/README.md)
- [saved_model_half_plus_two](https://github.com/tensorflow/serving/blob/master/tensorflow_serving/servables/tensorflow/testdata/saved_model_half_plus_two.py)
    - input, output, tags

## Module

```bash
go 1.15

require (
	github.com/golang/protobuf v1.4.3 // indirect
	github.com/tensorflow/tensorflow v2.1.3+incompatible
)
```

## Test

- HTTP: `https://github.com/giltene/wrk2`
- Bench stats: `go get golang.org/x/perf/cmd/benchstat`

## Tool

```bash
# CenotOS Install
Install GetPageSpeed repository:
# yum install https://extras.getpagespeed.com/release-el7-latest.rpm
Install wrk rpm package:
# yum install wrk

# MAC Install
brew install wrk
```

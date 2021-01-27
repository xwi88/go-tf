# QuickStart

## Go API

>https://www.tensorflow.org/install/lang_go

## C lib install

>https://www.tensorflow.org/install/lang_c#download

```bash
# Standard Install
sudo tar -C /usr/local -xzf (downloaded file)

# Non sys dir install
tar -C ~/c_lib -xzf libtensorflow-cpu-darwin-x86_64-2.3.0.tar.gz

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

## Module

```bash
go 1.15

require (
	github.com/golang/protobuf v1.4.3 // indirect
	github.com/tensorflow/tensorflow v2.1.3+incompatible
)
```

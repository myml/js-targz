package main

import (
	"archive/tar"
	"errors"
	"fmt"
	"io"
	"js-targz/internal/targz"

	gopherjs "github.com/gopherjs/gopherjs/js"
)

type blob struct {
	*gopherjs.Object
}

func (b *blob) Size() int64 {
	return b.Get("size").Int64()
}

func (b *blob) ReadAt(p []byte, offset int64) (n int, err error) {
	c := make(chan struct{})
	size := cap(p)
	b.
		Call("slice", offset, offset+int64(size)).
		Call("arrayBuffer").
		Call("then", func(buff *gopherjs.Object) {
			data := gopherjs.Global.Get("Uint8Array").New(buff).Interface().([]byte)
			n = copy(p, data)
		}).
		Call("catch", func(e *gopherjs.Object) {
			gopherjs.Global.Get("console").Call("log", e)
			err = errors.New("array buffer catch")
		}).
		Call("finally", func() {
			close(c)
		})
	<-c
	return
}

type promiseExecutor func(resolve *gopherjs.Object, reject *gopherjs.Object)

func newTarGzPromise(r io.Reader, callback *gopherjs.Object) promiseExecutor {
	return func(resolve *gopherjs.Object, reject *gopherjs.Object) {
		go func() {
			errBreak := fmt.Errorf("break")
			call := func(h *tar.Header) error {
				if callback == gopherjs.Undefined {
					return nil
				}
				err := callback.Invoke(h)
				if err.Bool() {
					return nil
				}
				return errBreak
			}
			err := targz.TarGz(r, call)
			if err != nil && !errors.Is(err, errBreak) {
				reject.Invoke(err)
				return
			}
			resolve.Invoke()
		}()
	}
}

func main() {
	targz := func(jsBlob *gopherjs.Object, callback *gopherjs.Object) *gopherjs.Object {
		b := blob{Object: jsBlob}
		r := io.NewSectionReader(&b, 0, b.Size())
		return gopherjs.Global.Get("Promise").New(newTarGzPromise(r, callback))
	}
	if gopherjs.Module != gopherjs.Undefined {
		gopherjs.Module.Get("exports").Set("targz", targz)
	} else {
		gopherjs.Global.Set("targz", targz)
	}
}

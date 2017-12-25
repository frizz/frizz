package wcache

import (
	"errors"

	"github.com/gopherjs/gopherjs/js"
)

type Cache struct {
	*js.Object
}

func NewCache(o *js.Object) *Cache {
	var c Cache
	c.Object = o
	return &c
}

func New(name string) (*Cache, error) {
	caches := js.Global.Get("caches")

	if caches == js.Undefined {
		// Caches not supported on some browsers (e.g. Safari)
		return NewCache(nil), nil
	}

	openReq := caches.Call("open", name)

	type response struct {
		cache *Cache
		err   error
	}
	ch := make(chan response)

	openReq.Call("then", func(o *js.Object) {
		go func() {
			ch <- response{cache: NewCache(o)}
		}()
	}, func(err *js.Object) {
		go func() {
			ch <- response{err: errors.New(err.String())}
		}()
	})

	val := <-ch
	return val.cache, val.err
}

func (c *Cache) Get(name string) (bool, []byte, error) {

	if c.Object == nil {
		// Caches not supported on some browsers (e.g. Safari)
		return false, nil, nil
	}

	req := js.Global.Get("Request").New(name)

	type response struct {
		found bool
		body  []byte
		err   error
	}
	ch := make(chan response)

	c.Call("match", req).Call("then", func(val *js.Object) {

		if val == js.Undefined {
			go func() {
				ch <- response{found: false}
			}()
			return
		}

		bch := make(chan []byte)
		val.Call("arrayBuffer").Call("then", func(ob *js.Object) {
			go func() {
				bch <- js.Global.Get("Uint8Array").New(ob).Interface().([]byte)
			}()
		})

		go func() {
			body := <-bch
			ch <- response{found: true, body: body}
		}()

	}).Call("catch", func(err *js.Object) {
		go func() {
			ch <- response{found: false, err: errors.New(err.String())}
		}()
	})

	val := <-ch

	if val.err != nil {
		return false, nil, val.err
	}

	return val.found, val.body, nil
}

func (c *Cache) Put(name string, data []byte) error {

	if c.Object == nil {
		// Caches not supported on some browsers (e.g. Safari)
		return nil
	}

	ch := make(chan error)

	req := js.Global.Get("Request").New(name)

	rsp := js.Global.Get("Response").New(
		js.NewArrayBuffer(data),
		map[string]interface{}{
			"status":  200,
			"headers": js.Global.Get("Headers").New(),
		},
	)

	c.Call("put", req, rsp).Call("then", func() {
		go func() {
			close(ch)
		}()
	}).Call("catch", func(err *js.Object) {
		go func() {
			ch <- errors.New(err.String())
		}()
	})

	if err := <-ch; err != nil {
		return err
	}

	return nil
}

package bootstrap

import (
	"sync"
)

func NewParallel(workers int) *Parallel {
	p := &Parallel{
		In: make(chan func() error),
		wg: &sync.WaitGroup{},
	}
	for i := 0; i < workers; i++ {
		go func() {
			for payload := range p.In {
				p.wg.Add(1)
				defer p.wg.Done()
				if err := payload(); err != nil {
					p.err = err
				}
			}
		}()
	}
	return p
}

type Parallel struct {
	In  chan func() error
	err error
	wg  *sync.WaitGroup
}

func (p *Parallel) Finish() error {
	close(p.In)  // signal to all the workers to exit
	p.wg.Wait()  // wait for payloads to finish
	return p.err // return the stored error
}

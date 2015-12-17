package tree

// isAsyncAndNotLoaded is true if the branch loads it's contents asynchronously, and the contents
// are not yet loaded
func (b *Branch) isAsyncAndNotLoaded() bool {
	if async, ok := b.self.(AsyncInterface); ok && !async.Loaded() {
		return true
	}
	return false
}

// ensureContentLoaded ensures the content is loaded. For asynchronous branches, it initializes the
// load event and returns a done chanel. For synchronous branches (or asynchronous branches that are
// already loaded), it returns nil so that the calling code can react synchronously.
func (b *Branch) ensureContentLoaded() (done chan bool, success bool) {

	if async, ok := b.self.(AsyncInterface); ok && !async.Loaded() {

		done = make(chan bool, 1)

		// load content asynchronously
		responseChannel, loading := async.LoadContent(b.tree.Conn, b.tree.Fail)

		if loading {
			// if we're already in the middle of a async request, we should cancel the operation
			return nil, false
		}

		go func() {
			<-responseChannel
			b.updateOpenerIcon()
			done <- true
		}()

		return done, true

	} else {
		// if item is not async or content is already loaded, just open the node.
		return nil, true
	}
}
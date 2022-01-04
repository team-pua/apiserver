package s3

import (
	"context"

	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/apiserver/pkg/storage"
)

type watcher struct {
}

// watchChan implements watch.Interface.
type watchChan struct {
}

// Stops watching. Will close the channel returned by ResultChan(). Releases
// any resources used by the watch.
func (wc *watchChan) Stop() {

}

// Returns a chan which will receive all the events. If an error occurs
// or Stop() is called, the implementation will close this channel and
// release any resources used by the watch.
func (wc *watchChan) ResultChan() <-chan watch.Event {
	return make(chan watch.Event)
}

func (w *watcher) Watch(ctx context.Context, key string, rev int64, recursive, progressNotify bool, pred storage.SelectionPredicate) (watch.Interface, error) {
	return &watchChan{}, nil
}

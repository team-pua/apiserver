/*
Copyright 2016 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package s3

import (
	"context"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/apiserver/pkg/storage"
	"k8s.io/apiserver/pkg/storage/storagebackend"
)

type store struct {
	versioner storage.Versioner
	watcher   *watcher
}

// it should implement storage interface
// New returns an etcd3 implementation of storage.Interface.
func New(c storagebackend.TransportConfig, newFunc func() runtime.Object) storage.Interface {
	return &store{}
}

// Watch implements storage.Interface.Watch.
func (s *store) Watch(ctx context.Context, key string, opts storage.ListOptions) (watch.Interface, error) {
	rev := 0
	return s.watcher.Watch(ctx, key, int64(rev), opts.Recursive, opts.ProgressNotify, opts.Predicate)
}

// Versioner implements storage.Interface.Versioner.
func (s *store) Versioner() storage.Versioner {
	return s.versioner
}

// Get implements storage.Interface.Get.
func (s *store) Get(ctx context.Context, key string, opts storage.GetOptions, out runtime.Object) error {
	return nil
}

// Create implements storage.Interface.Create.
func (s *store) Create(ctx context.Context, key string, obj, out runtime.Object, ttl uint64) error {
	return nil
}

// Delete implements storage.Interface.Delete.
func (s *store) Delete(
	ctx context.Context, key string, out runtime.Object, preconditions *storage.Preconditions,
	validateDeletion storage.ValidateObjectFunc, cachedExistingObject runtime.Object) error {
	return nil
}

// GuaranteedUpdate implements storage.Interface.GuaranteedUpdate.
func (s *store) GuaranteedUpdate(
	ctx context.Context, key string, out runtime.Object, ignoreNotFound bool,
	preconditions *storage.Preconditions, tryUpdate storage.UpdateFunc, cachedExistingObject runtime.Object) error {
	return nil
}

// GetToList implements storage.Interface.GetToList.
func (s *store) GetToList(ctx context.Context, key string, listOpts storage.ListOptions, listObj runtime.Object) error {
	return nil
}

func (s *store) Count(key string) (int64, error) {
	return 0, nil
}

// List implements storage.Interface.List.
func (s *store) List(ctx context.Context, key string, opts storage.ListOptions, listObj runtime.Object) error {
	return nil
}

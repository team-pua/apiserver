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

package factory

import (
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apiserver/pkg/storage"
	"k8s.io/apiserver/pkg/storage/s3"
	"k8s.io/apiserver/pkg/storage/storagebackend"
	"k8s.io/klog/v2"
)

func newS3Storage(c storagebackend.ConfigForResource, newFunc func() runtime.Object) (storage.Interface, DestroyFunc, error) {
	destroyFunc := func() {}

	return s3.New(c.Transport, newFunc), destroyFunc, nil
}

func newS3HealthCheck(c storagebackend.Config) (func() error, error) {
	klog.Info("calling newS3HealthCheck")
	return func() error {
		return nil
	}, nil
}

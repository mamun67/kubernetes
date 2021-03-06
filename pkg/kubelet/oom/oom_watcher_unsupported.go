// +build !linux

/*
Copyright 2019 The Kubernetes Authors.

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

package oom

import (
	"k8s.io/api/core/v1"
	"k8s.io/client-go/tools/record"
)

type oomWatcherUnsupported struct{}

var _ OOMWatcher = new(oomWatcherUnsupported)

// NewOOMWatcher creates a fake one here
func NewOOMWatcher(_ record.EventRecorder) OOMWatcher {
	return &oomWatcherUnsupported{}
}

func (ow *oomWatcherUnsupported) Start(_ *v1.ObjectReference) error {
	return nil
}

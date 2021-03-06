/*
Copyright © 2018-2019 The OpenEBS Authors

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

package payload

import (
	"github.com/container-storage-interface/spec/lib/go/csi"
)

// CreateVolumeResponseBuilder helps building an
// instance of csi CreateVolumeResponse
type CreateVolumeResponseBuilder struct {
	response *csi.CreateVolumeResponse
}

// NewCreateVolumeResponseBuilder returns a new
// instance of CreateVolumeResponseBuilder
func NewCreateVolumeResponseBuilder() *CreateVolumeResponseBuilder {
	return &CreateVolumeResponseBuilder{
		response: &csi.CreateVolumeResponse{
			Volume: &csi.Volume{},
		},
	}
}

// WithName sets the name against the
// CreateVolumeResponse instance
func (b *CreateVolumeResponseBuilder) WithName(name string) *CreateVolumeResponseBuilder {
	b.response.Volume.VolumeId = name
	return b
}

// WithName sets the capacity against the
// CreateVolumeResponse instance
func (b *CreateVolumeResponseBuilder) WithCapacity(capacity int64) *CreateVolumeResponseBuilder {
	b.response.Volume.CapacityBytes = capacity
	return b
}

// WithContext sets the context against the
// CreateVolumeResponse instance
func (b *CreateVolumeResponseBuilder) WithContext(ctx map[string]string) *CreateVolumeResponseBuilder {
	b.response.Volume.VolumeContext = ctx
	return b
}

// Build returns the constructed instance
// of csi CreateVolumeResponse
func (b *CreateVolumeResponseBuilder) Build() *csi.CreateVolumeResponse {
	return b.response
}

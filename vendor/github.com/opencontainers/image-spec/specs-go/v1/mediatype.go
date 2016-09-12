// Copyright 2016 The Linux Foundation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1

const (
	// MediaTypeDescriptor specifies the mediaType for a content descriptor.
	MediaTypeDescriptor = "application/vnd.oci.descriptor.v1+json"

	// MediaTypeImageManifest specifies the mediaType for an image manifest.
	MediaTypeImageManifest = "application/vnd.oci.image.manifest.v1+json"

	// MediaTypeImageManifestList specifies the mediaType for an image manifest list.
	MediaTypeImageManifestList = "application/vnd.oci.image.manifest.list.v1+json"

	// MediaTypeImageLayer is the mediaType used for layers referenced by the manifest.
	MediaTypeImageLayer = "application/vnd.oci.image.layer.tar+gzip"

	// MediaTypeImageConfig specifies the mediaType for the image configuration.
	MediaTypeImageConfig = "application/vnd.oci.image.config.v1+json"
)
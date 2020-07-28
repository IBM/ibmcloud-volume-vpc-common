/**
 * Copyright 2020 IBM Corp.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package instances

const (
	instanceIDParam                = "instance-id"
	clusterIDParam                 = "cluster-id"
	volumeIDParam                  = "volume-id"
	attachmentIDParam              = "id"
	instanceIDPath                 = "/{" + instanceIDParam + "}"
	volumeAttachmentPath           = "volume_attachments"
	instanceIDvolumeAttachmentPath = instanceIDPath + "/" + volumeAttachmentPath
	instanceIDattachmentIDPath     = instanceIDvolumeAttachmentPath + "/{" + attachmentIDParam + "}"

	// VpcPathPrefix  VPC URL path prefix
	VpcPathPrefix = "v1/instances"

	// IksPathPrefix  IKS URL path prefix
	IksPathPrefix = "v2/storage/vpc/"

	// IksClusterQueryKey ...
	IksClusterQueryKey = "cluster"

	// IksWorkerQueryKey ...
	IksWorkerQueryKey = "worker"

	// IksVolumeQueryKey ...
	IksVolumeQueryKey = "volumeID"

	// IksVolumeAttachmentIDQueryKey ...
	IksVolumeAttachmentIDQueryKey = "volumeAttachmentID"
)

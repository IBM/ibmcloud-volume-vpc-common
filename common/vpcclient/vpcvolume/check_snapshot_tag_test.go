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

package vpcvolume_test

import (
	//"github.com/IBM/ibmcloud-volume-vpc/common/vpcclient/models"
	"net/http"
	"testing"

	"github.com/IBM/ibmcloud-volume-vpc/common/vpcclient/riaas/test"
	"github.com/IBM/ibmcloud-volume-vpc/common/vpcclient/vpcvolume"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestCheckSnapshotTag(t *testing.T) {
	// Setup new style zap logger
	logger, _ := GetTestContextLogger()
	defer logger.Sync()

	testCases := []struct {
		name string

		// backend url
		url string

		// Response
		status  int
		content string

		// Expected return
		expectErr string
		verify    func(*testing.T, error)
	}{
		{
			name:   "Verify that the correct endpoint is invoked",
			status: http.StatusNoContent,
			url:    vpcvolume.Version + "/volumes/volume1/snapshots/snapshot1/tags/test-tag",
		}, {
			name:      "Verify that a 404 is returned to the caller",
			status:    http.StatusNotFound,
			url:       vpcvolume.Version + "/volumes/volume1/snapshots/snapshot1/tags/test-tag",
			content:   "{\"errors\":[{\"message\":\"testerr\"}]}",
			expectErr: "Trace Code:, testerr Please check ",
		}, {
			name:    "Verify that the snapshot is parsed correctly",
			status:  http.StatusNotFound,
			url:     vpcvolume.Version + "/volumes/volume1/snapshots/snapshot1/tags/test-tag",
			content: "{\"id\":\"snapshot1\",\"name\":\"snapshot1\",\"status\":\"pending\"}",
			verify: func(t *testing.T, err error) {
				assert.NotNil(t, err)
			},
		}, {
			name:   "Verify that the snapshot tag exists",
			status: http.StatusOK,
			url:    vpcvolume.Version + "/volumes/volume1/snapshots/snapshot1/tags/test-tag",
			verify: func(t *testing.T, err error) {
				assert.Nil(t, err)
			},
		},
	}

	for _, testcase := range testCases {
		t.Run(testcase.name, func(t *testing.T) {
			mux, client, teardown := test.SetupServer(t)
			emptyString := ""
			test.SetupMuxResponse(t, mux, testcase.url, http.MethodGet, &emptyString, testcase.status, testcase.content, nil)

			defer teardown()

			logger.Info("Test case being executed", zap.Reflect("testcase", testcase.name))

			snapshotService := vpcvolume.NewSnapshotManager(client)

			err := snapshotService.CheckSnapshotTag("volume1", "snapshot1", "test-tag", logger)

			// vpc snapshot functionality is not yet ready. It would return error for now
			if testcase.verify != nil {
				testcase.verify(t, err)
			}
		})
	}
}

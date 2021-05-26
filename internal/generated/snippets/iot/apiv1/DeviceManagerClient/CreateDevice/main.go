// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// [START cloudiot_v1_generated_DeviceManager_CreateDevice_sync]

package main

import (
	"context"

	iot "cloud.google.com/go/iot/apiv1"
	iotpb "google.golang.org/genproto/googleapis/cloud/iot/v1"
)

func main() {
	ctx := context.Background()
	c, err := iot.NewDeviceManagerClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &iotpb.CreateDeviceRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.CreateDevice(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

// [END cloudiot_v1_generated_DeviceManager_CreateDevice_sync]
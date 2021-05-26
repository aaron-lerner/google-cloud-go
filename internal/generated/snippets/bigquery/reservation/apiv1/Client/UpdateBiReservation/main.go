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

// [START bigqueryreservation_v1_generated_ReservationService_UpdateBiReservation_sync]

package main

import (
	"context"

	reservation "cloud.google.com/go/bigquery/reservation/apiv1"
	reservationpb "google.golang.org/genproto/googleapis/cloud/bigquery/reservation/v1"
)

func main() {
	ctx := context.Background()
	c, err := reservation.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &reservationpb.UpdateBiReservationRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.UpdateBiReservation(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

// [END bigqueryreservation_v1_generated_ReservationService_UpdateBiReservation_sync]
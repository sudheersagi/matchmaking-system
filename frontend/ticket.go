// Copyright 2019 Google LLC
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

package main

import (
	"math/rand"

	"open-match.dev/open-match/pkg/pb"
)

// Ticket generates a Ticket with a mode search field that has one of the
// randomly selected modes.
func makeTicket() *pb.Ticket {
	ticket := &pb.Ticket{
		SearchFields: &pb.SearchFields{
			Tags: []string{
				gameMode(),
			},
			StringArgs: map[string]string{
				"attributes.role": role(),
			},
		},
	}

	return ticket
}

func gameMode() string {
	modes := []string{"mode.demo", "mode.1v1", "mode.battleroyale", "mode.2v2"}
	return modes[rand.Intn(len(modes))]
}

func role() string {
	roles := []string{"role.dps", "role.support", "role.tank", "role.demo"}
	return roles[rand.Intn(len(roles))]
}

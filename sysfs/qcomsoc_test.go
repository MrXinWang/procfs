// Copyright 2019 The Prometheus Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//go:build linux
// +build linux

package sysfs

import (
	"reflect"
	"testing"
)

func TestQcomSoC(t *testing.T) {
	fs, err := NewFS(sysTestFixtures)
	if err != nil {
		t.Fatal(err)
	}

	s, err := fs.QcomSoC()
	if err != nil {
		t.Fatal(err)
	}

	soc := socInfo{
		ChipID:       "SA_LEMANSAU_IVI_ADAS",
		SocID:        "SA8775P",
		SocRevision:  "1.0",
		Machine:      "SA8775P",
		Family:       "Snapdragon",
		SerialNumber: "3929179670",
	}

	if !reflect.DeepEqual(soc, s) {
		t.Errorf("Result not correct: want %v, have %v", soc, s)
	}
}

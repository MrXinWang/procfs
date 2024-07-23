// Copyright 2024 Qualcomm Technologies, Inc.
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
	"strconv"

	"github.com/prometheus/procfs/internal/util"
)

const (
	chipIDPath  = "kernel/debug/qcom_socinfo/chip_id"
	socIDPath   = "devices/soc0/soc_id"
	machinePath = "devices/soc0/machine"
	familyPath  = "devices/soc0/family"
)

type socInfo struct {
	ChipID  string // /sys/kernel/debug/qcom_socinfo/chip_id
	SocID   string // /sys/devices/soc0/soc_id
	Machine string // /sys/devices/soc0/machine
	Family  string // /devices/soc0/family
}

// qcomSoC parses Qualcomm SoC related info from
// /sys/{kernel/debug/qcom_socinfo, devices/soc0}/...
func (fs FS) QcomSoC() (socInfo, error) {
	var soc socInfo

	ChipID, err := util.SysReadFile(fs.sys.Path(chipIDPath))
	if err != nil {
		return socInfo{}, err
	}

	soc.ChipID = string(ChipID)

	SocID, err := util.ReadUintFromFile(fs.sys.Path(socIDPath))
	if err != nil {
		return socInfo{}, err
	}

	switch SocID {
	case 532:
		soc.SocID = "SA8255P"
	case 533:
		soc.SocID = "SA8650P"
	case 534:
		soc.SocID = "SA8775P"
	default:
		soc.SocID = strconv.FormatUint(SocID, 10)
	}

	machine, err := util.SysReadFile(fs.sys.Path(machinePath))
	if err != nil {
		return socInfo{}, err
	}

	soc.Machine = string(machine)

	Family, err := util.SysReadFile(fs.sys.Path(familyPath))
	if err != nil {
		return socInfo{}, err
	}

	soc.Family = string(Family)

	return soc, nil
}

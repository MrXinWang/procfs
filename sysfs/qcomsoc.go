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
	"os"
	"strconv"

	"github.com/MrXinWang/procfs/internal/util"
)

const (
	chipIDPath       = "kernel/debug/qcom_socinfo/chip_id"
	socIDPath        = "devices/soc0/soc_id"
	socRevisionPath  = "devices/soc0/revision"
	machinePath      = "devices/soc0/machine"
	familyPath       = "devices/soc0/family"
	serialNumberPath = "devices/soc0/serial_number"
)

type socInfo struct {
	ChipID       string // /sys/kernel/debug/qcom_socinfo/chip_id
	SocID        string // /sys/devices/soc0/soc_id
	SocRevision  string // /sys/devices/soc0/revision
	Machine      string // /sys/devices/soc0/machine
	Family       string // /sys/devices/soc0/family
	SerialNumber string // /sys/devices/soc0/serial_number
}

// qcomSoC parses Qualcomm SoC related info from
// /sys/{kernel/debug/qcom_socinfo, devices/soc0}/...
func (fs FS) QcomSoC() (socInfo, error) {
	var soc socInfo

	ChipID, err := util.SysReadFile(fs.sys.Path(chipIDPath))
	if err != nil && !os.IsNotExist(err) {
		return socInfo{}, err
	}

	soc.ChipID = string(ChipID)

	SocID, err := util.ReadUintFromFile(fs.sys.Path(socIDPath))
	if err != nil && !os.IsNotExist(err) {
		return socInfo{}, err
	}

	switch SocID {
	case 460:
		soc.SocID = "SA8295P"
	case 532:
		soc.SocID = "SA8255P"
	case 533:
		soc.SocID = "SA8650P"
	case 534:
		soc.SocID = "SA8775P"
	case 535:
		soc.SocID = "SA8630P"
	case 605:
		soc.SocID = "SA8620P"
	case 606:
		soc.SocID = "SA7255P"
	case 607:
		soc.SocID = "SA8610P"
	case 648:
		soc.SocID = "SA8797P"
	default:
		soc.SocID = strconv.FormatUint(SocID, 10)
	}

	SocRevision, err := util.SysReadFile(fs.sys.Path(socRevisionPath))
	if err != nil && !os.IsNotExist(err) {
		return socInfo{}, err
	}

	soc.SocRevision = string(SocRevision)

	machine, err := util.SysReadFile(fs.sys.Path(machinePath))
	if err != nil && !os.IsNotExist(err) {
		return socInfo{}, err
	}

	soc.Machine = string(machine)

	Family, err := util.SysReadFile(fs.sys.Path(familyPath))
	if err != nil && !os.IsNotExist(err) {
		return socInfo{}, err
	}

	soc.Family = string(Family)

	SerialNumber, err := util.SysReadFile(fs.sys.Path(serialNumberPath))
	if err != nil && !os.IsNotExist(err) {
		return socInfo{}, err
	}

	soc.SerialNumber = string(SerialNumber)

	return soc, nil
}

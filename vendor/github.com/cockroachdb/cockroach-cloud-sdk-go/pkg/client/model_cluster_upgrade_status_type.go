// Copyright 2023 The Cockroach Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// CockroachDB Cloud API
// API version: 2023-04-10

package client

import (
	"encoding/json"
	"fmt"
)

// ClusterUpgradeStatusType  - FINALIZED: The cluster is running the latest available CockroachDB version, and all upgrades have been finalized.  - MAJOR_UPGRADE_RUNNING: An major version upgrade is currently in progress.  - UPGRADE_AVAILABLE: An upgrade is available. If preview builds are enabled for the parent organization, this could indicate that a preview upgrade is available.  - PENDING_FINALIZATION: An upgrade is complete, but pending finalization. Upgrades are automatically finalized after 72 hours. For more information, see https://www.cockroachlabs.com/docs/stable/upgrade-cockroach-version.html  - ROLLBACK_RUNNING: A rollback operation is currently in progress.
type ClusterUpgradeStatusType string

// List of ClusterUpgradeStatus.Type.
const (
	CLUSTERUPGRADESTATUSTYPE_FINALIZED             ClusterUpgradeStatusType = "FINALIZED"
	CLUSTERUPGRADESTATUSTYPE_MAJOR_UPGRADE_RUNNING ClusterUpgradeStatusType = "MAJOR_UPGRADE_RUNNING"
	CLUSTERUPGRADESTATUSTYPE_UPGRADE_AVAILABLE     ClusterUpgradeStatusType = "UPGRADE_AVAILABLE"
	CLUSTERUPGRADESTATUSTYPE_PENDING_FINALIZATION  ClusterUpgradeStatusType = "PENDING_FINALIZATION"
	CLUSTERUPGRADESTATUSTYPE_ROLLBACK_RUNNING      ClusterUpgradeStatusType = "ROLLBACK_RUNNING"
)

// All allowed values of ClusterUpgradeStatusType enum.
var AllowedClusterUpgradeStatusTypeEnumValues = []ClusterUpgradeStatusType{
	"FINALIZED",
	"MAJOR_UPGRADE_RUNNING",
	"UPGRADE_AVAILABLE",
	"PENDING_FINALIZATION",
	"ROLLBACK_RUNNING",
}

func (v *ClusterUpgradeStatusType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ClusterUpgradeStatusType(value)
	for _, existing := range AllowedClusterUpgradeStatusTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ClusterUpgradeStatusType", value)
}

// NewClusterUpgradeStatusTypeFromValue returns a pointer to a valid ClusterUpgradeStatusType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewClusterUpgradeStatusTypeFromValue(v string) (*ClusterUpgradeStatusType, error) {
	ev := ClusterUpgradeStatusType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ClusterUpgradeStatusType: valid values are %v", v, AllowedClusterUpgradeStatusTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ClusterUpgradeStatusType) IsValid() bool {
	for _, existing := range AllowedClusterUpgradeStatusTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ClusterUpgradeStatus.Type value.
func (v ClusterUpgradeStatusType) Ptr() *ClusterUpgradeStatusType {
	return &v
}

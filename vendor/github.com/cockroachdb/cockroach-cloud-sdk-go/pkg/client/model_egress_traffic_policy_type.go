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

// EgressTrafficPolicyType  - UNSPECIFIED: UNSPECIFIED signifies the egress traffic policy is unspecified.  - ERROR: ERROR signifies there has been an internal server error during an update to the egress traffic policy.  - ALLOW_ALL: ALLOW_ALL signifies all outbound connections from CockroachDB are allowed.  - DEFAULT_DENY: DEFAULT_DENY signifies that CockroachDB can only initiate network connections to destinations explicitly allowed by the user or CockroachDB Cloud operators.  - UPDATING: UPDATING signifies the egress traffic policy is updating.
type EgressTrafficPolicyType string

// List of EgressTrafficPolicy.Type.
const (
	EGRESSTRAFFICPOLICYTYPE_UNSPECIFIED  EgressTrafficPolicyType = "UNSPECIFIED"
	EGRESSTRAFFICPOLICYTYPE_ERROR        EgressTrafficPolicyType = "ERROR"
	EGRESSTRAFFICPOLICYTYPE_ALLOW_ALL    EgressTrafficPolicyType = "ALLOW_ALL"
	EGRESSTRAFFICPOLICYTYPE_DEFAULT_DENY EgressTrafficPolicyType = "DEFAULT_DENY"
	EGRESSTRAFFICPOLICYTYPE_UPDATING     EgressTrafficPolicyType = "UPDATING"
)

// All allowed values of EgressTrafficPolicyType enum.
var AllowedEgressTrafficPolicyTypeEnumValues = []EgressTrafficPolicyType{
	"UNSPECIFIED",
	"ERROR",
	"ALLOW_ALL",
	"DEFAULT_DENY",
	"UPDATING",
}

func (v *EgressTrafficPolicyType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EgressTrafficPolicyType(value)
	for _, existing := range AllowedEgressTrafficPolicyTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EgressTrafficPolicyType", value)
}

// NewEgressTrafficPolicyTypeFromValue returns a pointer to a valid EgressTrafficPolicyType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewEgressTrafficPolicyTypeFromValue(v string) (*EgressTrafficPolicyType, error) {
	ev := EgressTrafficPolicyType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EgressTrafficPolicyType: valid values are %v", v, AllowedEgressTrafficPolicyTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v EgressTrafficPolicyType) IsValid() bool {
	for _, existing := range AllowedEgressTrafficPolicyTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EgressTrafficPolicy.Type value.
func (v EgressTrafficPolicyType) Ptr() *EgressTrafficPolicyType {
	return &v
}

package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = AccessPackageAssignment{}

type AccessPackageAssignment struct {
	// Read-only. Nullable. Supports $filter (eq) on the id property and $expand query parameters.
	AccessPackage *AccessPackage `json:"accessPackage,omitempty"`

	// Read-only. Supports $filter (eq) on the id property and $expand query parameters.
	AssignmentPolicy *AccessPackageAssignmentPolicy `json:"assignmentPolicy,omitempty"`

	// Information about all the custom extension calls that were made during the access package assignment workflow.
	CustomExtensionCalloutInstances *[]CustomExtensionCalloutInstance `json:"customExtensionCalloutInstances,omitempty"`

	// The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example,
	// midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
	ExpiredDateTime nullable.Type[string] `json:"expiredDateTime,omitempty"`

	// When the access assignment is to be in place. Read-only.
	Schedule *EntitlementManagementSchedule `json:"schedule,omitempty"`

	// The state of the access package assignment. The possible values are: delivering, partiallyDelivered, delivered,
	// expired, deliveryFailed, unknownFutureValue. Read-only. Supports $filter (eq).
	State *AccessPackageAssignmentState `json:"state,omitempty"`

	// More information about the assignment lifecycle. Possible values include Delivering, Delivered,
	// NearExpiry1DayNotificationTriggered, or ExpiredNotificationTriggered. Read-only.
	Status nullable.Type[string] `json:"status,omitempty"`

	// The subject of the access package assignment. Read-only. Nullable. Supports $expand. Supports $filter (eq) on
	// objectId.
	Target *AccessPackageSubject `json:"target,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s AccessPackageAssignment) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = AccessPackageAssignment{}

func (s AccessPackageAssignment) MarshalJSON() ([]byte, error) {
	type wrapper AccessPackageAssignment
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AccessPackageAssignment: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AccessPackageAssignment: %+v", err)
	}

	delete(decoded, "accessPackage")
	delete(decoded, "assignmentPolicy")
	delete(decoded, "expiredDateTime")
	delete(decoded, "schedule")
	delete(decoded, "state")
	delete(decoded, "status")
	delete(decoded, "target")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.accessPackageAssignment"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AccessPackageAssignment: %+v", err)
	}

	return encoded, nil
}
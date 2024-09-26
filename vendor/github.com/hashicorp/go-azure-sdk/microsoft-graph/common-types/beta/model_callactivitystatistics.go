package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ActivityStatistics = CallActivityStatistics{}

type CallActivityStatistics struct {
	// Time spent on calls outside of working hours, which is based on the user's Outlook calendar setting for work hours.
	// The value is represented in ISO 8601 format for durations.
	AfterHours *string `json:"afterHours,omitempty"`

	// Fields inherited from ActivityStatistics

	// The type of activity for which statistics are returned. The possible values are: call, chat, email, focus, and
	// meeting.
	Activity *AnalyticsActivityType `json:"activity,omitempty"`

	// Total hours spent on the activity. The value is represented in ISO 8601 format for durations.
	Duration *string `json:"duration,omitempty"`

	// Date when the activity ended, expressed in ISO 8601 format for calendar dates. For example, the property value could
	// be '2019-07-03' that follows the YYYY-MM-DD format.
	EndDate *string `json:"endDate,omitempty"`

	// Date when the activity started, expressed in ISO 8601 format for calendar dates. For example, the property value
	// could be '2019-07-04' that follows the YYYY-MM-DD format.
	StartDate *string `json:"startDate,omitempty"`

	// The time zone that the user sets in Microsoft Outlook is used for the computation. For example, the property value
	// could be 'Pacific Standard Time.'
	TimeZoneUsed nullable.Type[string] `json:"timeZoneUsed,omitempty"`

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

func (s CallActivityStatistics) ActivityStatistics() BaseActivityStatisticsImpl {
	return BaseActivityStatisticsImpl{
		Activity:     s.Activity,
		Duration:     s.Duration,
		EndDate:      s.EndDate,
		StartDate:    s.StartDate,
		TimeZoneUsed: s.TimeZoneUsed,
		Id:           s.Id,
		ODataId:      s.ODataId,
		ODataType:    s.ODataType,
	}
}

func (s CallActivityStatistics) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = CallActivityStatistics{}

func (s CallActivityStatistics) MarshalJSON() ([]byte, error) {
	type wrapper CallActivityStatistics
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling CallActivityStatistics: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling CallActivityStatistics: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.callActivityStatistics"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling CallActivityStatistics: %+v", err)
	}

	return encoded, nil
}
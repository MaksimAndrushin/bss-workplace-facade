package model

import bss_workplace_facade "github.com/ozonmp/bss-workplace-facade/pkg/bss-workplace-facade"

func CreateEventFromProtoEvent(protoEvent bss_workplace_facade.WorkplaceEvent) *WorkplaceEvent {
	workplaceEvent := WorkplaceEvent{
		Type:        EventType(protoEvent.GetEventType()),
		Status:      EventStatus(protoEvent.GetEventStatus()),
		Entity:      &Workplace{
			ID: protoEvent.GetWorkplace().GetId(),
			Name: protoEvent.GetWorkplace().GetName(),
			Size: protoEvent.GetWorkplace().GetSize(),
			Created: protoEvent.GetWorkplace().GetCreated().AsTime(),
		},
	}

	return &workplaceEvent
}

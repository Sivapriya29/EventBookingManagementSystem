package events

import (
	"event-booking/constants"
	"event-booking/daos"
	"event-booking/database/models"
	"event-booking/dtos"
	"event-booking/utils/context"
	"log"
	"time"

	"github.com/google/uuid"
)

type Event struct {
	event daos.EventDAO
}

func New() *Event {
	return &Event{
		event: daos.NewEvent(),
	}
}

func (e *Event) EventFromEventReq(req *dtos.EventReq) *models.Event {
	if req.ID == "" {
		req.ID = uuid.New().String()
	}
	return &models.Event{
		ID:                req.ID,
		Event_name:        req.Event_name,
		Event_description: req.Event_description,
		Event_date:        req.Event_date,
		Event_time:        req.Event_time,
		Event_type:        req.Event_type,
		Location:          req.Location,
		Speaker_name:      req.Speaker_name,
		Organizer_name:    req.Organizer_name,
		Capacity:          req.Capacity,
		Per_person_price:  req.Per_person_price,
	}
}

func (e *Event) eventModelToDTO(event *models.Event) *dtos.EventRes {
	return &dtos.EventRes{
		ID:                event.ID,
		Event_name:        event.Event_name,
		Event_description: event.Event_description,
		Event_date:        event.Event_date,
		Event_time:        event.Event_time,
		Event_type:        event.Event_type,
		Location:          event.Location,
		Speaker_name:      event.Speaker_name,
		Organizer_name:    event.Organizer_name,
		Capacity:          event.Capacity,
		Per_person_price:  event.Per_person_price,
	}
}

func (e *Event) updateEventFromReq(event *models.Event, req *dtos.EventReq) *models.Event {
	event.Event_name = req.Event_name
	event.Event_description = req.Event_description
	event.Event_date = req.Event_date
	event.Event_time = req.Event_time
	event.Event_type = req.Event_type
	event.Location = req.Location
	event.Speaker_name = req.Speaker_name
	event.Organizer_name = req.Organizer_name
	event.Capacity = req.Capacity
	event.Per_person_price = req.Per_person_price
	event.Updated_at = time.Now()

	return event
}

func (e *Event) CreateEvent(ctx *context.Context, req *dtos.EventReq) error {

	event := e.EventFromEventReq(req)

	if ctx.User.Role != constants.RoleAdmin {
		return constants.ErrUnauthorizedToCreateEvent
	}
	return e.event.Create(ctx, event)
}

func (e *Event) GetEvent(ctx *context.Context, id string) (*dtos.EventRes, error) {
	event, err := e.event.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return e.eventModelToDTO(event), nil
}

func (e *Event) UpdateEvent(ctx *context.Context, id string, req *dtos.EventReq) error {

	if ctx.User.Role != constants.RoleAdmin {
		return constants.ErrUnauthorizedToUpdateEvent
	}

	event, err := e.event.Get(ctx, id)
	if err != nil {
		log.Println("Unable to get event. Err:", err)
		return err
	}

	event = e.updateEventFromReq(event, req)
	return e.event.Update(ctx, id, event)
}

func (e *Event) DeleteEvent(ctx *context.Context, id string) error {

	if ctx.User.Role != constants.RoleAdmin {
		return constants.ErrUnauthorizedToDeleteEvent
	}

	return e.event.Delete(ctx, id)
}

func (e *Event) GetAllEvents(ctx *context.Context) ([]*dtos.EventRes, error) {
	events, err := e.event.GetAllEvents(ctx)
	if err != nil {
		log.Println("Unable to fetch all events. Err: ", err)
		return nil, err
	}

	var result []*dtos.EventRes
	for _, event := range events {
		result = append(result, e.eventModelToDTO(event))
	}
	return result, nil
}

package daos

import (
	"event-booking/database/models"
	"event-booking/utils/context"
	"log"
)

type EventDAO interface {
	Create(ctx *context.Context, event *models.Event) error
	Upsert(ctx *context.Context, event *models.Event) error
	Get(ctx *context.Context, id string) (*models.Event, error)
	Update(ctx *context.Context, id string, updatedEvent *models.Event) error
	Delete(ctx *context.Context, id string) error
	GetAllEvents(ctx *context.Context) ([]*models.Event, error)
	UpdateCapacity(ctx *context.Context, eventID string, capacityChange int) error
}

func NewEvent() EventDAO {
	return &Event{}
}

type Event struct {
}

func (e *Event) Create(ctx *context.Context, event *models.Event) error {
	err := ctx.DB.Table("events").Create(event).Error
	if err != nil {
		log.Println("Unable to create event. Err:", err)
		return err
	}
	return nil
}

func (e *Event) Upsert(ctx *context.Context, event *models.Event) error {
	err := ctx.DB.Table("events").Save(event).Error
	if err != nil {
		log.Println("Unable to upsert event. Err:", err)
		return err
	}
	return nil
}

func (e *Event) Get(ctx *context.Context, id string) (*models.Event, error) {
	event := &models.Event{}
	err := ctx.DB.Table("events").First(event, "id = ?", id).Error
	if err != nil {
		log.Println("Unable to get event. Err:", err)
		return nil, err
	}
	return event, nil
}

func (e *Event) Update(ctx *context.Context, id string, updatedEvent *models.Event) error {
	err := ctx.DB.Table("events").Where("id = ?", id).Updates(updatedEvent).Error
	if err != nil {
		log.Println("Unable to update event.Err:", err)
		return err
	}
	return nil
}

func (e *Event) UpdateCapacity(ctx *context.Context, eventID string, capacityChange int) error {
	event := &models.Event{}
	err := ctx.DB.First(event, "id = ?", eventID).Error
	if err != nil {
		log.Println("Unable to find event. Err:", err)
		return err
	}

	event.Capacity += capacityChange
	if event.Capacity < 0 {
		event.Capacity = 0
	}

	err = ctx.DB.Save(event).Error
	if err != nil {
		log.Println("Unable to update event capacity. Err:", err)
		return err
	}

	return nil
}

func (e *Event) Delete(ctx *context.Context, id string) error {
	err := ctx.DB.Table("events").Delete(&models.Event{
		ID: id,
	}).Error
	if err != nil {
		log.Println("Unable to delete event. Err:", err)
		return err
	}
	return nil
}

func (e *Event) GetAllEvents(ctx *context.Context) ([]*models.Event, error) {
	var events []*models.Event
	err := ctx.DB.Table("events").Find(&events).Error

	if err != nil {
		log.Println("Unable to get all events. Err: ", err)
		return nil, err
	}

	return events, nil
}

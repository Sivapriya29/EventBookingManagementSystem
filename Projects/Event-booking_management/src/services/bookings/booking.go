package bookings

import (
	"errors"
	"event-booking/constants"
	"event-booking/daos"
	"event-booking/database/models"
	"event-booking/dtos"
	"event-booking/utils/context"
	"log"
	"time"

	"github.com/google/uuid"
)

type Booking struct {
	booking daos.BookingDAO
	event   daos.EventDAO
}

func New() *Booking {
	return &Booking{
		booking: daos.NewBooking(),
		event:   daos.NewEvent(),
	}
}

func (b *Booking) CreateBooking(ctx *context.Context, req *dtos.BookingReq) (*models.Booking, error) {

	if ctx.User.Role != constants.RoleUser {
		return nil, constants.ErrUnauthorizedToCreateEvent
	}

	event, err := b.event.Get(ctx, req.Event_id)
	if err != nil {
		log.Println("Unable to get event. Err:", err)
		return nil, err
	}

	if event.Capacity < req.Number_of_tickets {
		return nil, errors.New("not enough tickets available")
	}

	// Create booking
	booking := &models.Booking{
		ID:                uuid.New().String(),
		Event_id:          req.Event_id,
		Event_name:        event.Event_name,
		User_id:           req.User_id,
		Number_of_tickets: req.Number_of_tickets,
		Total_amount:      req.Total_amount,
		Created_at:        time.Now(),
	}

	err = b.booking.Create(ctx, booking)
	if err != nil {
		return nil, err
	}

	// Decrease event capacity
	event.Capacity -= req.Number_of_tickets
	err = b.event.Update(ctx, req.Event_id, event)
	if err != nil {
		return nil, err
	}

	return booking, nil
}

func (b *Booking) GetBooking(ctx *context.Context, id string) (*dtos.BookingRes, error) {
	booking, err := b.booking.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return &dtos.BookingRes{
		ID:                booking.ID,
		Event_id:          booking.Event_id,
		User_id:           booking.User_id,
		Number_of_tickets: booking.Number_of_tickets,
		Total_amount:      booking.Total_amount,
		Created_at:        booking.Created_at,
	}, nil
}

func (b *Booking) DeleteBooking(ctx *context.Context, id string) error {
	booking, err := b.booking.Delete(ctx, id)
	if err != nil {
		return err
	}

	// Increase event capacity
	err = b.event.UpdateCapacity(ctx, booking.Event_id, booking.Number_of_tickets)
	if err != nil {
		return err
	}

	return nil
}

func (b *Booking) GetAllBookings(ctx *context.Context) ([]*dtos.BookingRes, error) {
	bookings, err := b.booking.GetAllBookings(ctx)
	if err != nil {
		log.Println("Unable to fetch all bookings. Err: ", err)
		return nil, err
	}

	var result []*dtos.BookingRes
	for _, booking := range bookings {
		event, err := b.event.Get(ctx, booking.Event_id)
		if err != nil {
			log.Println("Unable to fetch event for booking. Err: ", err)
			continue
		}
		result = append(result, &dtos.BookingRes{
			ID:                booking.ID,
			Event_id:          booking.Event_id,
			EventName:         event.Event_name,
			User_id:           booking.User_id,
			Number_of_tickets: booking.Number_of_tickets,
			Total_amount:      booking.Total_amount,
			Created_at:        booking.Created_at,
		})
	}
	return result, nil
}

func (b *Booking) GetBookingsByUserID(ctx *context.Context, userID string) ([]models.Booking, error) {
	return b.booking.GetByUserID(ctx, userID)
}

func (b *Booking) GetBookingsByEventID(ctx *context.Context, eventID string) ([]models.Booking, error) {
	return b.booking.GetByEventID(ctx, eventID)
}

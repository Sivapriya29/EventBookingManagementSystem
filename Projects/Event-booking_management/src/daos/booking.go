package daos

import (
	"errors"
	"event-booking/database/models"
	"event-booking/utils/context"
	"log"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BookingDAO interface {
	Create(ctx *context.Context, booking *models.Booking) error
	Upsert(ctx *context.Context, booking *models.Booking) error
	Get(ctx *context.Context, id string) (*models.Booking, error)
	Update(ctx *context.Context, id string, updatedBooking *models.Booking) error
	Delete(ctx *context.Context, id string) (*models.Booking, error)
	GetAllBookings(ctx *context.Context) ([]*models.Booking, error)
	GetByUserID(ctx *context.Context, userID string) ([]models.Booking, error)
	GetByEventID(ctx *context.Context, eventID string) ([]models.Booking, error)
}

func NewBooking() BookingDAO {
	return &Booking{}
}

type Booking struct {
}

func (b *Booking) Create(ctx *context.Context, booking *models.Booking) error {
	err := ctx.DB.Table("bookings").Create(booking).Error
	if err != nil {
		log.Println("Unable to create booking. Err:", err)
		return err
	}
	return nil
}

func (b *Booking) Upsert(ctx *context.Context, booking *models.Booking) error {
	err := ctx.DB.Table("bookings").Save(booking).Error
	if err != nil {
		log.Println("Unable to upsert booking. Err:", err)
		return err
	}
	return nil
}

func (b *Booking) Get(ctx *context.Context, id string) (*models.Booking, error) {
	booking := &models.Booking{}
	err := ctx.DB.Table("bookings").First(booking, "id = ?", id).Error
	if err != nil {
		log.Println("Unable to get booking. Err:", err)
		return nil, err
	}
	return booking, nil
}

func (b *Booking) Update(ctx *context.Context, id string, updatedBooking *models.Booking) error {
	err := ctx.DB.Table("bookings").Where("id = ?", id).Updates(updatedBooking).Error
	if err != nil {
		log.Println("Unable to update booking.Err:", err)
		return err
	}
	return nil
}

func (b *Booking) Delete(ctx *context.Context, id string) (*models.Booking, error) {
	booking := &models.Booking{}
	err := ctx.DB.First(booking, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		log.Println("Unable to find booking. Err:", err)
		return nil, err
	}

	err = ctx.DB.Delete(booking).Error
	if err != nil {
		log.Println("Unable to delete booking. Err:", err)
		return nil, err
	}
	return booking, nil
}

func (b *Booking) GetAllBookings(ctx *context.Context) ([]*models.Booking, error) {
	var bookings []*models.Booking
	err := ctx.DB.Table("bookings").Find(&bookings).Error

	if err != nil {
		log.Println("Unable to get all bookings. Err: ", err)
		return nil, err
	}

	return bookings, nil
}

func (b *Booking) GetByUserID(ctx *context.Context, userID string) ([]models.Booking, error) {

	if userID == "" {
		return nil, errors.New("userID cannot be empty")
	}

	if _, err := uuid.Parse(userID); err != nil {
		log.Println("invalid user_id format or user_id not found. Err:", err)
		return nil, err
	}

	var bookings []models.Booking
	err := ctx.DB.Where("user_id = ?", userID).Find(&bookings).Error
	if err != nil {
		log.Println("Unable to get bookings by user ID. Err:", err)
		return nil, err
	}
	return bookings, nil
}

func (b *Booking) GetByEventID(ctx *context.Context, eventID string) ([]models.Booking, error) {

	if eventID == "" {
		return nil, errors.New("eventID cannot be empty")
	}

	if _, err := uuid.Parse(eventID); err != nil {
		log.Println("invalid event_id format or event_id not found. Err:", err)
		return nil, err
	}

	var bookings []models.Booking
	err := ctx.DB.Where("event_id = ?", eventID).Find(&bookings).Error
	if err != nil {
		log.Println("Unable to get bookings by event ID. Err:", err)
		return nil, err
	}
	return bookings, nil
}

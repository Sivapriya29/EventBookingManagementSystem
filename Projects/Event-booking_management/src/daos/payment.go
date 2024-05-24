package daos

import (
	"event-booking/database/models"
	"event-booking/utils/context"
	"log"
)

type PaymentDAO interface {
	Create(ctx *context.Context, payment *models.Payment) error
	Upsert(ctx *context.Context, payment *models.Payment) error
	Get(ctx *context.Context, id string) (*models.Payment, error)
	Update(ctx *context.Context, id string, updatedBooking *models.Payment) error
	Delete(ctx *context.Context, id string) error
	GetAllPayments(ctx *context.Context) ([]*models.Payment, error)
	CheckBookingIDExists(ctx *context.Context, booking_id string) (bool, error)
}

func NewPayment() PaymentDAO {
	return &Payment{}
}

type Payment struct {
}

func (p *Payment) CheckBookingIDExists(ctx *context.Context, booking_id string) (bool, error) {
	var cnt int
	err := ctx.DB.Table("payments").Select("count(*)").Where("booking_id=?", booking_id).Scan(&cnt).Error
	if err != nil {
		log.Println("Unable to create payment. Err:", err)
		return false, err
	}
	return cnt > 0, nil
}

func (p *Payment) Create(ctx *context.Context, payment *models.Payment) error {
	err := ctx.DB.Table("payments").Create(payment).Error
	if err != nil {
		log.Println("Unable to create payment. Err:", err)
		return err
	}
	return nil
}

func (p *Payment) Upsert(ctx *context.Context, payment *models.Payment) error {
	err := ctx.DB.Table("payments").Save(payment).Error
	if err != nil {
		log.Println("Unable to upsert payment. Err:", err)
		return err
	}
	return nil
}

func (p *Payment) Get(ctx *context.Context, id string) (*models.Payment, error) {
	payment := &models.Payment{}
	err := ctx.DB.Table("payments").First(payment, "id = ?", id).Error
	if err != nil {
		log.Println("Unable to get payment. Err:", err)
		return nil, err
	}
	return payment, nil
}

func (p *Payment) Update(ctx *context.Context, id string, updatedPayment *models.Payment) error {
	err := ctx.DB.Table("payments").Where("id = ?", id).Updates(updatedPayment).Error
	if err != nil {
		log.Println("Unable to update payment.Err:", err)
		return err
	}
	return nil
}

func (p *Payment) Delete(ctx *context.Context, id string) error {

	err := ctx.DB.Table("payments").Delete(&models.Payment{
		ID: id,
	}).Error
	if err != nil {
		log.Println("Unable to delete payment. Err:", err)
		return err
	}
	return nil
}

func (p *Payment) GetAllPayments(ctx *context.Context) ([]*models.Payment, error) {
	var payments []*models.Payment
	err := ctx.DB.Table("payments").Find(&payments).Error

	if err != nil {
		log.Println("Unable to get all payments. Err: ", err)
		return nil, err
	}

	return payments, nil
}

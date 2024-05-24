package payments

import (
	"event-booking/daos"
	"event-booking/database/models"
	"event-booking/dtos"
	"event-booking/utils/context"
	"log"
	"time"

	"github.com/google/uuid"
)

type Payment struct {
	payment daos.PaymentDAO
}

func New() *Payment {
	return &Payment{
		payment: daos.NewPayment(),
	}
}

func (p *Payment) PaymentFromReq(req *dtos.PaymentReq) *models.Payment {
	if req.ID == "" {
		req.ID = uuid.New().String()
	}
	return &models.Payment{
		ID:           req.ID,
		Booking_id:   req.Booking_id,
		User_id:      req.User_id,
		Event_id:     req.Event_id,
		Amount:       req.Amount,
		Card_number:  req.Card_number,
		Expiry_month: req.Expiry_month,
		Expiry_year:  req.Expiry_year,
		Cvv:          req.Cvv,
		Card_holder:  req.Card_holder,
	}
}

func (p *Payment) paymentModelToDTO(payment *models.Payment) *dtos.PaymentRes {
	return &dtos.PaymentRes{
		ID:           payment.ID,
		Booking_id:   payment.Booking_id,
		User_id:      payment.User_id,
		Event_id:     payment.Event_id,
		Amount:       payment.Amount,
		Payment_date: payment.Payment_date,
	}
}

func (p *Payment) CreatePayment(ctx *context.Context, req *dtos.PaymentReq) error {

	payment := p.PaymentFromReq(req)

	// if ok, _ := p.payment.CheckBookingIDExists(ctx, payment.Booking_id); ok {
	// 	return constants.ErrBookingTaken
	// }

	return p.payment.Create(ctx, payment)
}

func (p *Payment) GetPayment(ctx *context.Context, id string) (*dtos.PaymentRes, error) {
	payment, err := p.payment.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return p.paymentModelToDTO(payment), nil
}

func (p *Payment) UpdatePayment(ctx *context.Context, id string, req *dtos.PaymentReq) error {
	payment, err := p.payment.Get(ctx, id)
	if err != nil {
		log.Println("Unable to get payment. Err:", err)
		return err
	}
	payment.User_id = req.User_id
	payment.Event_id = req.Event_id
	payment.Amount = req.Amount
	payment.Payment_date = time.Now()

	return p.payment.Update(ctx, id, payment)
}

func (p *Payment) DeletePayment(ctx *context.Context, id string) error {
	return p.payment.Delete(ctx, id)
}

func (p *Payment) GetAllPayments(ctx *context.Context) ([]*dtos.PaymentRes, error) {
	payments, err := p.payment.GetAllPayments(ctx)
	if err != nil {
		log.Println("Unable to fetch all payments. Err: ", err)
		return nil, err
	}

	var result []*dtos.PaymentRes
	for _, payment := range payments {
		result = append(result, p.paymentModelToDTO(payment))
	}
	return result, nil
}

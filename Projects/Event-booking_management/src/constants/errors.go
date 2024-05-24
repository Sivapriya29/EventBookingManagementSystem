package constants

import "errors"

var (
	ErrEmailTaken                   = errors.New("email taken")
	ErrMobileTaken                  = errors.New("Mobile taken")
	ErrInvalidCredentials           = errors.New("Invalid email or password")
	ErrInvalidEmailOrPassword       = errors.New("Invalid email or password")
	ErrInvalidEmailOrPasswordOrRole = errors.New("Invalid email or password or role")
	ErrEventNotFound                = errors.New("Event not found")
	ErrBookingNotFound              = errors.New("Booking not found")
	ErrFeedbackNotFound             = errors.New("Feedback not found")
	ErrPaymentNotFound              = errors.New("Payment not found")
	ErrRefreshTokenExpired          = errors.New("Refresh token expired")
	ErrAccessTokenExpired           = errors.New("Access token expired")
	ErrUnauthorizedToCreateEvent    = errors.New("Unathorized to create event")
	ErrUnauthorizedToUpdateEvent    = errors.New("Unathorized to update event")
	ErrUnauthorizedToDeleteEvent    = errors.New("Unathorized to delete event")
	ErrBookingTaken                 = errors.New("Payment already done for this booking")
)

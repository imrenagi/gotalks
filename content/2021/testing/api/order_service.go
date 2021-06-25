package api

type Invoice struct {
	URL string
}

// STARTPAYMENTINTERFACE, OMIT

// PaymentService is another gRPC server handling payment creation etc
type PaymentService interface {
	GenerateInvoice(orderID string) (*Invoice, error)
}

// STOPPAYMENTINTERFACE, OMIT

// STARTORDER, OMIT
type OrderService struct {
	PaymentService PaymentService // good candidate for mock // HL
}

// Finalize will generate the invoice URL used for payment
func (o *OrderService) Finalize(ID string) (string, error) {
	// some setup ...

	invoice, err := o.PaymentService.GenerateInvoice(ID) // HL
	if err != nil {
		return "", err
	}

	return invoice.URL, nil
}

// STOPORDER, OMIT

type Location struct {
	Name    string
	Address string
}

type PinPoint struct{}

// STARTLOCATION,OMIT
type LocationService interface {
	CreatePinPoint(l Location) (*PinPoint, error)
}

// STOPLOCATION,OMIT

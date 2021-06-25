package api_test

// STARTMOCK, OMIT
import (
	"fmt"
	"testing"

	"github.com/imrenagi/gotalks/content/2021/testing/api"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock" // HL
)

type PaymentServiceMock struct {
	mock.Mock // HL
}

func (m *PaymentServiceMock) GenerateInvoice(ID string) (string, error) {
	args := m.Called(ID) // verification // HL
	return args.String(0), args.Error(1)
}

// STOPMOCK, OMIT

// STARTFINALIZE, OMIT
func TestFinalize(t *testing.T) {
	m := new(PaymentServiceMock)

	m.On("GenerateInvoice", mock.Anything). // HL
						Return(&Invoice{URL: "example.com/invoices/1"}, nil) // HL

	orderService := api.OrderService{PaymentService: m}
	url, err := orderService.Finalize("1234")

	assert.NoError(err)
	assert.Equal(t, "example.com/invoices/1", url)
}

// STOPFINALIZE, OMIT

func TestFinalize_Error(t *testing.T) {
	m := new(PaymentServiceMock)

	// STARTFINALIZEERROR, OMIT
	m.On("GenerateInvoice", mock.Anything). // HL
						Return("", fmt.Errorf("random error")) // HL
	// STOPFINALIZEERROR, OMIT

	orderService := api.OrderService{PaymentService: m}
	url, err := orderService.Finalize("1234")

	assert.NoError(err)
	assert.Equal(t, "example.com/invoices/1", url)
}

// STARTVERIFICATION, OMIT
func TestFinalize_WithVerification(t *testing.T) {
	m := new(PaymentServiceMock) // OMIT

	m.On("GenerateInvoice", "1234"). // HL
						Return(&Invoice{URL: "example.com/invoices/1"}, nil)

	orderService := api.OrderService{PaymentService: m}
	url, err := orderService.Finalize("1234") // HL

	assert.NoError(err)
	assert.Equal(t, "example.com/invoices/1", url)

	m.AssertExpectations(t)                        // HL
	m.AssertNumberOfCalls(t, "GenerateInvoice", 1) // HL
}

// STOPVERIFICATION, OMIT

type LocationServiceeMock struct {
	mock.Mock
}

func (m *PaymentServiceMock) CreatePinPoint(l Location) (*PinPoint, error) {
	args := m.Called(l) // verification // HL
	return args.Get(0).(*PinPoint), args.Error(1)
}

// STARTCOMPLEXVER, OMIT
func TestVerifyCreatePoint(t *testing.T) {
	// some processing

	m := new(LocationServiceeMock)
	m.On("CreatePinPoint", 
		mock.MatchedBy(func (l Location) bool) { // HL
			assert.NotEmpty(t, l.Name)
			assert.Equal(t, "1600 Villa st", l.Address)
			return true
	})
}

// STOPCOMPLEXVER, OMIT

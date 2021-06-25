package api_test

import (
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/stretchr/testify/mock"
)

// STARTMOCKS3, OMIT
type MockS3 struct {
	mock.Mock
}

func (m MockS3) PutObject(input *s3.PutObjectInput) (*s3.PutObjectOutput, error) {
	args := m.Called(input)
	return args.Get(0).(*s3.PutObjectOutput), args.Error(1)
}

// STOPMOCKS3, OMIT

// STARTTEST, OMIT
func TestUploader_Upload(t *testing.T) {
  m := new(MockS3)
  u := UserService{Uploader: m}

  m.On("PutObject", 
		mock.MatchedBy(func(i *s3.PutObjectInput) bool { // HL
			// other assertion ...
			assert.Equal(t, "imre", *i.Key)
			return true
  }).
		Return(&s3.PutObjectOutput{}, nil) // HL

  err := u.UploadProfile(&User{Name: "imre"}, []byte(`hello`)) // HL
  assert.NoError(t, err)
  m.AssertExpectations(t)
}

// STOPTEST, OMIT

package api

// STARTUSERSERVICE, OMIT

import (
	"bytes" //OMIT
	//OMIT
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

type UserService struct {
	Uploader *aws.S3 // watch out for the concrete here // HL
}

// uploadFile uploads an object.
func (u UserService) UploadProfile(user *User, picture []byte) error { // HL
	input := &s3.PutObjectInput{
		Body:   aws.ReadSeekCloser(bytes.NewBuffer(picture)),
		Bucket: aws.String("examplebucket"),
		Key:    aws.String(user.Name),
	}
	_, err := u.Uploader.PutObject(input) // hard to test // HL
	// ....
	if err != nil { // OMIT
		return err // OMIT
	} // OMIT
	// OMIT
	return nil // OMIT
}

// STOPUSERSERVICE, OMIT

type User struct {
	Name string
}

// STARTUPLOADER, OMIT
// Uploader is interface for file upload to aws s3
type Uploader interface {
	PutObject(input *s3.PutObjectInput) (*s3.PutObjectOutput, error)
}

// STOPUPLOADER, OMIT

// STARTNEWSERVICE,OMIT
type UserService struct {
	Uploader Uploader
}

// STOPNEWSERVICE,OMIT

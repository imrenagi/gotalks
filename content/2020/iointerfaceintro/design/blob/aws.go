package blob

import (
	"bytes"
	"context"
	"io"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func NewAWS(bucket, file string) *awsObj {
	sess := session.Must(session.NewSession())
	svc := s3.New(sess)
	return &awsObj{
		client: svc,
		bucket: bucket,
		object: file,
	}
}

//STARTAWSOBJ, OMIT
type awsObj struct {
	client         *s3.S3
	bucket, object string
}

func (a *awsObj) Writer() io.WriteCloser {
	return a
}

//STOPAWSOBJ, OMIT

//STARTIOIMPLEMENTATION, OMIT
func (a awsObj) Write(p []byte) (int, error) {
	//STARTAWSS3, OMIT
	out, err := a.client.PutObjectWithContext(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(a.bucket),
		Key:    aws.String(a.object),
		Body:   bytes.NewReader(p),
	})
	//STOPAWSS3, OMIT
	if err != nil {
		return 0, err
	}

	return len(p), nil
}

func (a awsObj) Close() error {
	return nil
}

//STOPIOIMPLEMENTATION, OMIT

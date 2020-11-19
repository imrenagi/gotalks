package blob

// STARTIMPORT, OMIT
import (
	"context" // OMIT
	"io"      // OMIT
	"log"     // OMIT

	//import other packages
	"cloud.google.com/go/storage"
)

// STOPIMPORT, OMIT

// STARTGCS, OMIT
type gcs struct {
	client *storage.Client
	bucket string
	object string
}

func (g gcs) Writer() io.WriteCloser {
	// gcs sdk luckily return instance of io.Writer and io.Closer
	return g.client.Bucket(g.bucket).Object(g.object).NewWriter(context.Background())
}

// STOPGCS, OMIT

func NewGCS(bucket, file string) *gcs {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	return &gcs{
		client: client,
		bucket: bucket,
		object: file,
	}
}

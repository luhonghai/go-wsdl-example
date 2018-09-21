package aws

import (
	"testing"
	"time"

	"github.com/magiconair/properties/assert"
)

func TestListAllMyBuckets(t *testing.T) {
	auth := &BasicAuth{}
	s3 := NewAmazonS3("", false, auth)

	request := &ListAllMyBuckets{
		Timestamp: time.Now(),
	}
	resp, err := s3.ListAllMyBuckets(request)

	if err != nil {
		t.Error("Could not request", err)
	} else {
		assert.Equal(t, resp.ListAllMyBucketsResponse.Owner, "test")
	}
}

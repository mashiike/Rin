package rin_test

import (
	"testing"

	rin "github.com/fujiwara/Rin"
)

var eventStr = `{
  "Records": [
    {
      "eventVersion": "2.0",
      "eventSource": "aws:s3",
      "awsRegion": "ap-northeast-1",
      "eventTime": "2015-04-21T04:55:48.282Z",
      "eventName": "ObjectCreated:Put",
      "userIdentity": {
        "principalId": "AWS:AIDAITB24YMP65EXRRFHC"
      },
      "requestParameters": {
        "sourceIPAddress": "10.115.144.24"
      },
      "responseElements": {
        "x-amz-request-id": "C223B09A2672B58C",
        "x-amz-id-2": "lwnmR96s31UoVCw5ozvg+jV+heZKoheJ+KBoWinmnfl1RzxVUn48R+Baha1vUyW0"
      },
      "s3": {
        "s3SchemaVersion": "1.0",
        "configurationId": "test",
        "bucket": {
          "name": "test.bucket.test",
          "ownerIdentity": {
            "principalId": "A3RIPTMLB7ZZQI"
          },
          "arn": "arn:aws:s3:::test.bucket.test"
        },
        "object": {
          "key": "foo/bar.json",
          "size": 443,
          "eTag": "86fcdfb65af50a994cf63ddd280cea0d"
        }
      }
    }
  ]
}`

func TestParseEvent(t *testing.T) {
	var event rin.Event
	event, err := rin.ParseEvent([]byte(eventStr))
	if err != nil {
		t.Error("json decode error", err)
	}
	r := event.Records[0]
	if r.EventName != "ObjectCreated:Put" {
		t.Error("unexpected EventName", r.EventName)
	}
	if r.EventSource != "aws:s3" {
		t.Error("unexpected EventSource", r.EventSource)
	}
	if r.AWSRegion != "ap-northeast-1" {
		t.Error("unexpected AWSRegion", r.AWSRegion)
	}
	if r.S3.Bucket.Name != "test.bucket.test" {
		t.Error("unexpected bucket name", r.S3.Bucket.Name)
	}
	if r.S3.Object.Key != "foo/bar.json" {
		t.Error("unexpected key", r.S3.Object.Key)
	}
}

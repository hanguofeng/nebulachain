package main

import (
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"os"
)

type StorageUtil struct {
	Bucket string
	Region string
	svc    *s3.S3
	ssn    *session.Session
}

func (this *StorageUtil) init() error {
	config := &aws.Config{}
	if this.Region != "" {
		config.Region = &this.Region
	}

	this.ssn = session.New(config)
	svc := s3.New(this.ssn)
	this.svc = svc
	if !this.isBucketExists(this.Bucket) {
		return errors.New(fmt.Sprintf("bucket %s not exists", this.Bucket))
	}

	return nil
}

func (this *StorageUtil) isBucketExists(name string) bool {
	input := &s3.HeadBucketInput{Bucket: &name}
	_, err := this.svc.HeadBucket(input)
	if err == nil {
		return true
	} else {
		return false
	}
}

func (this *StorageUtil) getFileInfo(fileName string) (*s3.HeadObjectOutput, error) {
	input := &s3.HeadObjectInput{Bucket: &this.Bucket, Key: &fileName}
	info, err := this.svc.HeadObject(input)
	return info, err
}

func (this *StorageUtil) upload(fileName string, fileLocation string, fingerPrint string) (string, error) {
	f, err := os.Open(fileLocation)
	defer f.Close()

	if err != nil {
		return "", err
	}

	metadata := map[string]*string{}
	metadata["fingerprint"] = &fingerPrint

	uploader := s3manager.NewUploader(this.ssn)

	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket:   &this.Bucket,
		Key:      &fileName,
		Body:     f,
		Metadata: metadata,
	})

	if err != nil {
		return "", err
	} else {
		return result.Location, nil
	}
}

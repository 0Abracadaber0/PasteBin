package handler

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"strings"
)

const (
	vkCloudHotboxEndpoint = "https://hb.ru-msk.vkcs.cloud"
	defaultRegion         = "ru-msk"
)

func UploadText(hash string, body string) {
	sess, _ := session.NewSession()

	svc := s3.New(sess, aws.NewConfig().WithEndpoint(vkCloudHotboxEndpoint).WithRegion(defaultRegion))

	bucket := "pastebin-by-abracadaber"

	key := hash + ".txt"

	if _, err := svc.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
		Body:   strings.NewReader(body),
	}); err != nil {
		fmt.Printf("Unable to upload %q to %q, %v\n", body, bucket, err)
	} else {
		fmt.Printf("File %q uploaded to bucket %q\n", body, bucket)
	}

}

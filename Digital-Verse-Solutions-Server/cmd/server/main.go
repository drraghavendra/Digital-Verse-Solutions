package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/rs/zerolog/log"
	"gitlab.com/digitalverse/blockchain/pkg/registry"
)

func createDownloader(c *cfg) *s3manager.Downloader {
	awsCfg := aws.Config{
		S3ForcePathStyle: aws.Bool(true),
		Credentials:      credentials.NewStaticCredentials(c.AccessKey, c.SecretKey, ""),
		Region:           aws.String(c.Region),
	}
	sess, err := session.NewSession(&awsCfg)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to create aws session")
	}
	return s3manager.NewDownloader(sess)
}

func main() {
	downloader := createDownloader(c)
	downloaderFunc := func(src string) (*os.File, error) {
		tmpdir := os.TempDir()
		file, err := ioutil.TempFile(tmpdir, "s3temp")
		if err != nil {
			return nil, err
		}

		input := &s3.GetObjectInput{
			Bucket: aws.String(c.Bucket),
			Key:    aws.String(src),
		}
		if _, err := downloader.Download(file, input); err != nil {
			return nil, err
		}

		return file, nil
	}

	r := createRouter(downloaderFunc, registry.CreateRegistry(bc))
	port := c.ListenPort
	if port == 0 {
		port = 8080
	}
	if err := r.Run(fmt.Sprintf(":%d", port)); err != nil {
		log.Fatal().Err(err).Msg("failed to start server")
	}
}

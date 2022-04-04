package blockchain

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/rs/zerolog/log"
)

func downloadFromS3(accessKey, secretKey, region, bucket string, src string, file *os.File) error {
	awsCfg := aws.Config{
		S3ForcePathStyle: aws.Bool(true),
		Credentials:      credentials.NewStaticCredentials(accessKey, secretKey, ""),
		Region:           aws.String(region),
	}
	sess, err := session.NewSession(&awsCfg)
	if err != nil {
		log.Error().Err(err).Msg("session.NewSession")
		return err
	}

	downloader := s3manager.NewDownloader(sess)

	param := s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(src),
	}
	if _, err := downloader.Download(file, &param); err != nil {
		log.Error().Err(err).Msg("downloader.Download")
		return err
	}

	return nil
}

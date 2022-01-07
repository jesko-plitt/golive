package cloud

import (
	"context"

	"github.com/ao-concepts/logging"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func ProvideMinioClient(log logging.Logger) *minio.Client {
	cfg := ProvideMinioConfig()

	minioClient, err := minio.New(cfg.URL, &minio.Options{
		Creds: credentials.NewStaticV4(cfg.AccessKey, cfg.SecretKey, ""),
	})

	if err != nil {
		log.ErrFatal(err)
	}

	err = minioClient.MakeBucket(context.Background(), cfg.BucketName, minio.MakeBucketOptions{Region: cfg.BucketLocation})

	if err != nil {
		exists, errBucketExists := minioClient.BucketExists(context.Background(), cfg.BucketName)

		if errBucketExists != nil {
			log.ErrFatal(errBucketExists)
		}

		if !exists {
			log.Fatal("minio: bucket %s has not been created", cfg.BucketName)
		}
	}

	return minioClient
}

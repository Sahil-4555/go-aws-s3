package configs

import (
	"go-aws-s3/log"
	"os"

	"github.com/joho/godotenv"
)

func Port() string {
	err := godotenv.Load(".env")
	if err != nil {
		log.GetLog().Info("ERROR : ", "Error loading .env file.")
	}

	return os.Getenv("PORT")
}

func AccessKey() string {
	err := godotenv.Load(".env")
	if err != nil {
		log.GetLog().Info("ERROR : ", "Error loading .env file.")
	}

	return os.Getenv(("AWS_ACCESS_KEY"))
}

func SecretKey() string {
	err := godotenv.Load(".env")
	if err != nil {
		log.GetLog().Info("ERROR : ", "Error loading .env file.")
	}

	return os.Getenv(("AWS_SECRET_KEY"))
}

func Region() string {
	err := godotenv.Load(".env")
	if err != nil {
		log.GetLog().Info("ERROR : ", "Error loading .env file.")
	}

	return os.Getenv(("AWS_REGION"))
}

func Bucket() string {
	err := godotenv.Load(".env")
	if err != nil {
		log.GetLog().Info("ERROR : ", "Error loading .env file.")
	}

	return os.Getenv(("S3_BUCKET"))
}

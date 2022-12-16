package gostorage

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type Aws struct {
	client *s3.Client
}

func NewAws() *Aws {

	ac := &Aws{}

	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic("configuration error, " + err.Error())
	}

	ac.client = s3.NewFromConfig(cfg)

	return ac
}

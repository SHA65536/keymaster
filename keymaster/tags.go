package keymaster

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
)

var keymasterTags = []*ec2.Tag{{
	Key:   aws.String("CreatedByKeymaster"),
	Value: aws.String("True"),
}}

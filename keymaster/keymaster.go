package keymaster

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/schollz/progressbar/v3"
)

type Keymaster struct {
	sess *session.Session
}

type Key struct {
	Name      string
	Region    string
	PublicKey string
}

// MakeKeymaster creates a Keymaster object using shared config
func MakeKeymaster() (*Keymaster, error) {
	bar := progressbar.Default(1, "Connecting to AWS")
	sess, err := session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	})
	if err != nil {
		return nil, err
	}
	bar.Add(1)
	return &Keymaster{sess: sess}, nil
}

// ListKeys returns all key-pairs in the given regions
func (km *Keymaster) ListKeys(regions ...string) (map[string][]Key, error) {
	var res = map[string][]Key{}
	bar := progressbar.Default(int64(len(regions)), "Listing keys in given regions")
	for _, region := range regions {
		// Create ec2 service for each region
		svc := ec2.New(km.sess, &aws.Config{Region: aws.String(region)})

		// Get keys from AWS
		result, err := svc.DescribeKeyPairs(&ec2.DescribeKeyPairsInput{
			IncludePublicKey: aws.Bool(true),
		})
		if err != nil {
			return nil, err
		}
		bar.Add(1)
		if len(result.KeyPairs) == 0 {
			continue
		}
		temp := make([]Key, len(result.KeyPairs))
		for i, keyPair := range result.KeyPairs {
			temp[i] = Key{
				Name:      *keyPair.KeyName,
				Region:    region,
				PublicKey: *keyPair.PublicKey,
			}
		}
		res[region] = temp
	}
	return res, nil
}

// CreateKey creates a key-pair in the given region using a key
func (km *Keymaster) CreateKey(region string, key Key) error {
	bar := progressbar.Default(2, "Creating key in new region")
	svc := ec2.New(km.sess, &aws.Config{Region: aws.String(region)})

	// Create key in new region
	input := &ec2.ImportKeyPairInput{
		KeyName:           aws.String(key.Name),
		PublicKeyMaterial: []byte(key.PublicKey),
	}
	result, err := svc.ImportKeyPair(input)
	if err != nil {
		return err
	}
	bar.Add(1)

	// Add tag
	_, err = svc.CreateTags(&ec2.CreateTagsInput{
		Resources: []*string{result.KeyPairId},
		Tags:      keymasterTags,
	})
	if err != nil {
		return err
	}

	bar.Add(1)
	return err
}

// GetAllRegions returns all AWS regions
func (km *Keymaster) GetAllRegions() ([]string, error) {
	bar := progressbar.Default(1, "Getting AWS regions")
	svc := ec2.New(km.sess)

	// Get all regions
	result, err := svc.DescribeRegions(&ec2.DescribeRegionsInput{})
	if err != nil {
		return nil, err
	}

	// Return region names
	var regions []string = make([]string, len(result.Regions))
	for i, region := range result.Regions {
		regions[i] = *region.RegionName
	}

	bar.Add(1)

	return regions, nil
}

package testcontainerx

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/docker/go-connections/nat"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"log"
	"net/url"
)

const (
	localImage = "amazon/dynamodb-local"
	port       = nat.Port("8000/tcp")
)

type DynamoDBLocalResolver struct {
	hostAndPort string
}

func (r *DynamoDBLocalResolver) ResolveEndpoint(
	_ context.Context,
	_ dynamodb.EndpointParameters,
) (endpoint smithyendpoints.Endpoint, err error) {
	return smithyendpoints.Endpoint{
		URI: url.URL{Host: r.hostAndPort, Scheme: "http"},
	}, nil
}

func StartDynamoDBContainer() *dynamodb.Client {
	im := localImage

	ctx := context.Background()
	req := testcontainers.ContainerRequest{
		Image:        im,
		ExposedPorts: []string{string(port)},
		WaitingFor:   wait.ForListeningPort(port),
		Name:         "dynamodb_local_tc2",
	}

	genericContainerReq := testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
		Reuse:            true,
	}

	container, err := testcontainers.GenericContainer(ctx, genericContainerReq)
	if err != nil {
		log.Fatalf("Could not start container: %v", err)
	}

	cfg, err := config.LoadDefaultConfig(ctx,
		config.WithCredentialsProvider(credentials.StaticCredentialsProvider{
			Value: aws.Credentials{
				AccessKeyID:     "DUMMYIDEXAMPLE",
				SecretAccessKey: "DUMMYEXAMPLEKEY",
			},
		}))
	if err != nil {
		log.Fatalf("Could not load config: %v", err)
	}

	mappedPort, err := container.MappedPort(ctx, "8000/tcp")
	if err != nil {
		log.Fatalf("Could not get mapped port: %v", err)
	}

	hostIP, err := container.Host(ctx)
	if err != nil {
		log.Fatalf("Could not get host: %v", err)
	}

	uri := fmt.Sprintf("%s:%s", hostIP, mappedPort.Port())

	cli := dynamodb.NewFromConfig(
		cfg,
		dynamodb.WithEndpointResolverV2(&DynamoDBLocalResolver{hostAndPort: uri}),
	)
	return cli
}

func CreateTable(cli *dynamodb.Client) error {
	// TODO - define and create a sensible ddb table for explore service
	return nil
}

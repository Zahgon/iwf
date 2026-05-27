// Copyright (c) 2021 Cadence workflow OSS organization
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package iwf

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/indeedeng/iwf/config"

	uclient "github.com/indeedeng/iwf/service/client"
	"github.com/indeedeng/iwf/service/common/log"
	"github.com/uber-go/tally/v4"
	"github.com/uber-go/tally/v4/prometheus"
	"github.com/urfave/cli"
	"go.temporal.io/sdk/client"
	"go.uber.org/cadence/.gen/go/cadence/workflowserviceclient"
	cclient "go.uber.org/cadence/client"
)

const serviceAPI = "api"
const serviceInterpreter = "interpreter"

// BuildCLI is the main entry point for the iwf server
func BuildCLI() *cli.App { _ = "STUB: not implemented"; return nil }

const DefaultCadenceDomain = "default"
const DefaultCadenceHostPort = "127.0.0.1:7833"

func start(c *cli.Context) { _ = "STUB: not implemented"; return }

// The client is a heavyweight object that should be created once per process.

// NOTE: this connectionOptions can be removed when upgrading temporal SDK to latest
// see https://docs.temporal.io/cloud/api-keys#sdk

// TODO improve the waiting with process signal

func launchTemporalService(
	svcName string, config config.Config, unifiedClient uclient.UnifiedClient, temporalClient client.Client,
	logger log.Logger, metrics client.MetricsHandler,
) {
	_ = "STUB: not implemented"
	return
}

func launchCadenceService(
	svcName string,
	config config.Config,
	unifiedClient uclient.UnifiedClient,
	service workflowserviceclient.Interface,
	domain string,
	closeFunc func(),
	logger log.Logger,
) {
	_ = "STUB: not implemented"
	return
}

func getServices(c *cli.Context) []string { _ = "STUB: not implemented"; return nil }

const _cadenceFrontendService = "cadence-frontend"
const _cadenceClientName = "cadence-client"

func BuildCadenceClient(service workflowserviceclient.Interface, domain string) (cclient.Client, error) {
	_ = "STUB: not implemented"
	return *new(cclient.Client), nil
}

func BuildCadenceServiceClient(hostPort string) (workflowserviceclient.Interface, func(), error) {
	_ = "STUB: not implemented"
	return *new(workflowserviceclient.Interface), nil, nil
}

// tally sanitizer options that satisfy Prometheus restrictions.
// This will rename metrics at the tally emission level, so metrics name we
// use maybe different from what gets emitted. In the current implementation
// it will replace - and . with _
var (
	safeCharacters = []rune{'_'}

	sanitizeOptions = tally.SanitizeOptions{
		NameCharacters: tally.ValidCharacters{
			Ranges:     tally.AlphanumericRange,
			Characters: safeCharacters,
		},
		KeyCharacters: tally.ValidCharacters{
			Ranges:     tally.AlphanumericRange,
			Characters: safeCharacters,
		},
		ValueCharacters: tally.ValidCharacters{
			Ranges:     tally.AlphanumericRange,
			Characters: safeCharacters,
		},
		ReplacementCharacter: tally.DefaultReplacementCharacter,
	}
)

func newPrometheusScope(c prometheus.Configuration, logger log.Logger) tally.Scope {
	_ = "STUB: not implemented"
	return *new(tally.Scope)
}

func CreateS3Client(cfg config.Config, ctx context.Context) *s3.Client {
	_ = "STUB: not implemented"
	return nil
}

// get the first active storage

// Create custom resolver for MinIO endpoint

// Load AWS config with custom credentials and endpoint

// Create S3 client with path-style addressing (required for MinIO)

func createBucketIfNotExists(ctx context.Context, client *s3.Client, bucketName string) {
	_ = "STUB: not implemented"
	// Check if bucket exists
	return
}

// Bucket doesn't exist, create it

// Its posible creating a bucket failed because the bucket was created by another service (api or interpreter)
// check the bucket still does not exist

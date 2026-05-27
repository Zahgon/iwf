package blobstore

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/indeedeng/iwf/config"
	"github.com/indeedeng/iwf/service/common/log"
	"go.temporal.io/sdk/client"
)

type blobStoreImpl struct {
	s3Client                    *s3.Client
	pathPrefix                  string // the Temporal namespace or Cadence domain + "/"
	activeStorage               config.BlobStorageConfig
	supportedStore              map[string]config.BlobStorageConfig // storeId as key
	logger                      log.Logger
	writeObjectErrorCounter     client.MetricsCounter
	readObjectErrorCounter      client.MetricsCounter
	writeObjectSuccessHistogram client.MetricsTimer
	readObjectSuccessHistogram  client.MetricsTimer
}

func NewBlobStore(
	s3Client *s3.Client,
	temporalOrCadenceNamespace string,
	storeConfig config.ExternalStorageConfig,
	logger log.Logger,
	metrics client.MetricsHandler,
) BlobStore {
	_ = "STUB: not implemented"
	return *new(BlobStore)
}

func (b *blobStoreImpl) WriteObject(ctx context.Context, workflowId, data string) (storeId, path string, err error) {
	_ = "STUB: not implemented"
	return "", "", nil
}

// yyyymmdd$workflowId/uuid
// Note: using $ here so that the listing can be much easier to implement for pagination

func (b *blobStoreImpl) ReadObject(ctx context.Context, storeId, path string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func putObject(ctx context.Context, client *s3.Client, bucketName string, key, content string) error {
	_ = "STUB: not implemented"
	return nil
}

func getObject(ctx context.Context, client *s3.Client, bucketName, key string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (b *blobStoreImpl) CountWorkflowObjectsForTesting(ctx context.Context, workflowId string) (int64, error) {
	_ = "STUB: not implemented"
	// Create the prefix to match objects for this workflowId for today
	return 0, nil
}

// List objects with the prefix (limited to 1000 objects as documented)

func (b *blobStoreImpl) DeleteWorkflowObjects(ctx context.Context, storeId, workflowPath string) error {
	_ = "STUB: not implemented"
	return nil
}

// Construct the prefix for all objects of this workflow

// Paginate through all objects and delete them in batches

// If no objects found, we're done

// Prepare objects for batch deletion

// Delete objects in batch

// Don't return successful deletions

// Log S3-specific request identifiers for debugging with AWS Support

// Check for any delete errors

// Check if there are more objects to process

func (b *blobStoreImpl) ListWorkflowPaths(ctx context.Context, input ListObjectPathsInput) (*ListObjectPathsOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Set continuation token if provided

// Extract workflow paths from common prefixes

// Remove the pathPrefix to get the workflow path (yyyymmdd$workflowId)

// Remove trailing "/" if present

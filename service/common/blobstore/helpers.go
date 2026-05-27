package blobstore

import (
	"context"

	"github.com/indeedeng/iwf/gen/iwfidl"
)

func WriteDataObjectsToExternalStorage(ctx context.Context, dataObjects []iwfidl.KeyValue, workflowId string, threashold int, blobStore BlobStore, isExternalStorageEnabled bool) error {
	_ = "STUB: not implemented"
	return nil
}

// Save data to external storage

// Clear data since it's now in external storage

func LoadDataObjectsFromExternalStorage(ctx context.Context, dataObjects []iwfidl.KeyValue, blobStore BlobStore) error {
	_ = "STUB: not implemented"
	return nil
}

/*
 * MinIO Cloud Storage, (C) 2017, 2018 MinIO, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package grpc

import (
	"context"
	"io"
	"net/http"

	"github.com/minio/cli"
	"github.com/minio/minio/pkg/auth"

	minio "github.com/minio/minio/cmd"
)

const (
	grpcBackend = "grpc"
)

func init() {
	const grpcHelpTemplate = `NAME:
  {{.HelpName}} - {{.Usage}}

USAGE:
	TODO
`

	minio.RegisterGatewayCommand(cli.Command{
		Name:               grpcBackend,
		Usage:              "GRPC",
		Action:             grpcGatewayMain,
		CustomHelpTemplate: grpcHelpTemplate,
		HideHelpCommand:    true,
	})
}

// Handler for 'minio gateway gcs' command line.
func grpcGatewayMain(ctx *cli.Context) {
	// TODO: Pull out CLI arguments as needed here.

	minio.StartGateway(ctx, &GRPC{})
}

// GRPC implements a minio.ObjectLayer which is called using GRPC
type GRPC struct {
	// TODO(cmaloney): GRPC connection details go here.
}

// Name returns the name of ObjectLayer.
func (g *GRPC) Name() string {
	return grpcBackend
}

// NewGatewayLayer returns gcs ObjectLayer.
func (g *GRPC) NewGatewayLayer(creds auth.Credentials) (minio.ObjectLayer, error) {
	// TODO(cmaloney): Validate configuration here
	// TODO(cmaloney): Initialize GRPC client.
	grpc := &grpcGateway{
		// TODO(cmaloney): GRPC client goes here.
	}

	// TODO(cmaloney): Start any background processes needed here.
	return grpc, nil
}

// Production - GCS gateway is production ready.
func (g *GRPC) Production() bool {
	// TODO(cmaloney): Toggle this when ready.
	return false
}

// TODO(cmaloney): Mapping multipart uploads.
// TODO(cmaloney): Helpers to convert grpc errors to minio object layer errors.

// gcsGateway - Implements gateway for MinIO and GCS compatible object storage servers.
type grpcGateway struct {
	minio.GatewayUnsupported
}

// Shutdown - save any gateway metadata to disk
// if necessary and reload upon next restart.
func (l *grpcGateway) Shutdown(ctx context.Context) error {
	return nil
}

// StorageInfo - Not relevant to GCS backend.
func (l *grpcGateway) StorageInfo(ctx context.Context) minio.StorageInfo {
	return minio.StorageInfo{}
}

// MakeBucketWithLocation - Create a new container. For now not allowed.
func (l *grpcGateway) MakeBucketWithLocation(ctx context.Context, bucket, location string) error {
	return minio.NotImplemented{}
}

// GetBucketInfo - Get bucket metadata..
func (l *grpcGateway) GetBucketInfo(ctx context.Context, bucket string) (minio.BucketInfo, error) {
	return minio.BucketInfo{}, minio.NotImplemented{}
}

// ListBuckets lists all buckets under your project-id on GCS.
func (l *grpcGateway) ListBuckets(ctx context.Context) (buckets []minio.BucketInfo, err error) {
	return nil, minio.NotImplemented{}
}

// DeleteBucket delete a bucket on GCS.
func (l *grpcGateway) DeleteBucket(ctx context.Context, bucket string) error {
	return minio.NotImplemented{}
}

// ListObjects - lists all blobs in GCS bucket filtered by prefix
func (l *grpcGateway) ListObjects(ctx context.Context, bucket string, prefix string, marker string, delimiter string, maxKeys int) (minio.ListObjectsInfo, error) {
	if maxKeys == 0 {
		return minio.ListObjectsInfo{}, nil
	}

	return minio.ListObjectsInfo{}, minio.NotImplemented{}
}

// ListObjectsV2 - lists all blobs in GCS bucket filtered by prefix
func (l *grpcGateway) ListObjectsV2(ctx context.Context, bucket, prefix, continuationToken, delimiter string, maxKeys int, fetchOwner bool, startAfter string) (minio.ListObjectsV2Info, error) {
	if maxKeys == 0 {
		return minio.ListObjectsV2Info{ContinuationToken: continuationToken}, nil
	}
	return minio.ListObjectsV2Info{}, minio.NotImplemented{}
}

// GetObjectNInfo - returns object info and locked object ReadCloser
func (l *grpcGateway) GetObjectNInfo(ctx context.Context, bucket, object string, rs *minio.HTTPRangeSpec, h http.Header, lockType minio.LockType, opts minio.ObjectOptions) (gr *minio.GetObjectReader, err error) {
	return nil, minio.NotImplemented{}
}

// GetObject - reads an object from GCS. Supports additional
// parameters like offset and length which are synonymous with
// HTTP Range requests.
//
// startOffset indicates the starting read location of the object.
// length indicates the total length of the object.
func (l *grpcGateway) GetObject(ctx context.Context, bucket string, key string, startOffset int64, length int64, writer io.Writer, etag string, opts minio.ObjectOptions) error {
	return minio.NotImplemented{}
}

// GetObjectInfo - reads object info and replies back ObjectInfo
func (l *grpcGateway) GetObjectInfo(ctx context.Context, bucket string, object string, opts minio.ObjectOptions) (minio.ObjectInfo, error) {
	return minio.ObjectInfo{}, minio.NotImplemented{}
}

// PutObject - Create a new object with the incoming data,
func (l *grpcGateway) PutObject(ctx context.Context, bucket string, key string, r *minio.PutObjReader, opts minio.ObjectOptions) (minio.ObjectInfo, error) {
	return minio.ObjectInfo{}, minio.NotImplemented{}
}

// DeleteObject - Deletes a blob in bucket
func (l *grpcGateway) DeleteObject(ctx context.Context, bucket string, object string) error {
	return minio.NotImplemented{}
}

func (l *grpcGateway) DeleteObjects(ctx context.Context, bucket string, objects []string) ([]error, error) {
	errs := make([]error, len(objects))
	for idx := range objects {
		errs[idx] = minio.NotImplemented{}
	}
	return errs, nil
}

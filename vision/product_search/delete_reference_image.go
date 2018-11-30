// Copyright 2018 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package productsearch

// [START vision_product_search_delete_reference_image]

import (
	"context"
	"fmt"
	"io"

	vision "cloud.google.com/go/vision/apiv1"
	visionpb "google.golang.org/genproto/googleapis/cloud/vision/v1"
)

func deleteReferenceImage(w io.Writer, projectID string, location string, productID string, referenceImageID string) error {
	ctx := context.Background()
	c, err := vision.NewProductSearchClient(ctx)
	if err != nil {
		return err
	}

	req := &visionpb.DeleteReferenceImageRequest{
		Name: fmt.Sprintf("projects/%s/locations/%s/products/%s/referenceImages/%s", projectID, location, productID, referenceImageID),
	}

	err = c.DeleteReferenceImage(ctx, req)
	if err != nil {
		return err
	}

	fmt.Fprintln(w, "Reference image deleted from product.")

	return nil
}

// [END vision_product_search_delete_reference_image]
package resourcecreator_test

import (
	"testing"

	"github.com/nais/naiserator/pkg/apis/nais.io/v1alpha1"
	"github.com/nais/naiserator/pkg/resourcecreator"
	"github.com/nais/naiserator/pkg/test/fixtures"
	"github.com/stretchr/testify/assert"
)

func TestGetGoogleStorageBucket(t *testing.T) {
	t.Run("bucket creation", func(t *testing.T) {
		bucketname := "buckowens"
		retentionPeriod := 3600
		app := fixtures.MinimalApplication()
		app.Spec.GCP = &v1alpha1.GCP{
			Buckets: []v1alpha1.CloudStorageBucket{
				{
					Name: bucketname,
					RetentionPeriod: retentionPeriod,
				},
			},
		}
		bucket := resourcecreator.GoogleStorageBucket(app, app.Spec.GCP.Buckets[0])
		assert.Equal(t, "buckowens", bucket.Name)
		assert.Equal(t, 3600, bucket.Spec.RetentionPolicy.RetentionPeriod)
		assert.Equal(t, resourcecreator.GoogleRegion, bucket.Spec.Location)
	})
}

package dataloaders

import (
	// "context"
	"time"

	"github.com/google/uuid"

	genloaders "github.com/cmsgov/easi-app/pkg/graph/generated/dataloaders"
	"github.com/cmsgov/easi-app/pkg/models"
	"github.com/cmsgov/easi-app/pkg/storage"
)

// NewAccessibilityRequestStatusRecordLoader creates a data loader for AccessibilityRequestStatusRecords
func NewAccessibilityRequestStatusRecordLoader( /*ctx context.Context, */ store storage.Store) *genloaders.AccessibilityRequestStatusRecordLoader {
	return genloaders.NewAccessibilityRequestStatusRecordLoader(genloaders.AccessibilityRequestStatusRecordLoaderConfig{
		Fetch: func(keys []uuid.UUID) ([]*models.AccessibilityRequestStatusRecord, []error) {
			var refs []*models.AccessibilityRequestStatusRecord
			statusRecords, err := store.FetchLatestAccessibilityRequestStatusRecordsForMultipleRequests( /*ctx,*/ keys)

			if err != nil {
				return nil, []error{err}
			}

			for i := range statusRecords {
				refs = append(refs, &statusRecords[i])
			}

			return refs, nil
		},
		Wait:     20 * time.Millisecond,
		MaxBatch: 100,
	})
}

package veriff_test

import (
	"testing"

	"github.com/brokeyourbike/veriff-api-client-go"
	"github.com/stretchr/testify/assert"
)

func TestTime(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		wantErr bool
	}{
		{"dob", "2023-07-21", false},
	}

	for i := range tests {
		test := tests[i]
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			var d veriff.Time

			err := d.UnmarshalJSON([]byte(test.value))
			if test.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}

}

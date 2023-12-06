package httperror

import (
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestErrorFromStatus(t *testing.T) {
	rows := []struct {
		description string
		input       int
		err         error
	}{
		{
			description: "Zero value",
			input:       0,
			err:         nil,
		},
		{
			description: "Default value",
			input:       http.StatusOK,
			err:         nil,
		},
		{
			description: "ErrStatusNotFound",
			input:       http.StatusNotFound,
			err:         ErrStatusNotFound,
		},
		{
			description: "Non-existent 599 error",
			input:       599,
			err:         newError(599),
		},
	}
	for rowIndex, row := range rows {
		t.Run(fmt.Sprintf("%d/%s", rowIndex, row.description), func(t *testing.T) {
			err := ErrorFromStatus(row.input)
			if row.err == nil {
				require.Nil(t, err)

				status := StatusFromError(err)
				assert.Equal(t, 0, status)
				return
			}
			require.NotNil(t, err)

			// Make sure that we got the right error back.
			assert.Equal(t, row.err, err)
			// Make sure that the error wraps ErrStatus.
			assert.True(t, errors.Is(err, ErrStatus))

			// Make sure that the internal status is what we expect (via helper).
			status := StatusFromError(err)
			assert.Equal(t, row.input, status)

			// Make sure that the internal status is what we expect (via Error).
			var e *Error
			if assert.True(t, errors.As(err, &e)) {
				assert.Equal(t, row.input, e.Code())
			}
		})
	}
}

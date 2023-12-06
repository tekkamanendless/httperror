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
		output      error
	}{
		{
			description: "Zero value",
			input:       0,
			output:      nil,
		},
		{
			description: "Default value",
			input:       http.StatusOK,
			output:      nil,
		},
		{
			description: "ErrStatusNotFound",
			input:       http.StatusNotFound,
			output:      ErrStatusNotFound,
		},
		{
			description: "Non-existent 500 error",
			input:       599,
			output:      ErrStatus,
		},
	}
	for rowIndex, row := range rows {
		t.Run(fmt.Sprintf("%d/%s", rowIndex, row.description), func(t *testing.T) {
			output := ErrorFromStatus(row.input)
			if row.output == nil {
				require.Nil(t, output)
				return
			}
			require.NotNil(t, output)
			assert.Equal(t, row.output, output)
			assert.True(t, errors.Is(output, ErrStatus))
		})
	}
}

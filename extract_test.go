package ziputil_test

import (
	"testing"

	"github.com/Ollinar/ziputil"
	"github.com/stretchr/testify/require"
)

func TestExtract(t *testing.T) {
	err := ziputil.ExtractFromPath("./testSample/testz.zip", "./testResult/TestExtract")
	require.NoError(t, err)
}

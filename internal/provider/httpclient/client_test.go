package httpclient_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/julien-noblet/download-geofabrik/internal/provider/httpclient"
)

func TestNew(t *testing.T) {
	t.Parallel()

	client := httpclient.New()
	require.NotNil(t, client)
	assert.NotNil(t, client.Transport)
}

func TestNewWithTimeout(t *testing.T) {
	t.Parallel()

	client := httpclient.NewWithTimeout(15 * time.Second)
	require.NotNil(t, client)
	assert.NotNil(t, client.Transport)
}

func TestNewCrawlerClient(t *testing.T) {
	t.Parallel()

	client := httpclient.NewCrawlerClient(25)
	require.NotNil(t, client)
	assert.NotNil(t, client.Transport)
}

func TestNewCrawlerClientWithTimeout(t *testing.T) {
	t.Parallel()

	client := httpclient.NewCrawlerClientWithTimeout(15*time.Second, 25)
	require.NotNil(t, client)
	assert.NotNil(t, client.Transport)
}

package api

import (
	"os"
	"testing"
	"time"

	db "github.com/21toffy/relinc/db/sqlc"
	"github.com/gin-gonic/gin"

	"github.com/21toffy/relinc/util"
	"github.com/stretchr/testify/require"
)

func newTestServer(t *testing.T, store db.Store) *Server {
	config := util.Config{
		TokenSymmetricKey:   util.SecondRandomString(32),
		AccessTokenDuration: time.Minute,
	}

	server, err := NewServer(config, &store)
	require.NoError(t, err)

	return server
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)

	os.Exit(m.Run())
}

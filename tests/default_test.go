package test

import (
	"testing"

	"github.com/go-phd/ssf"
	"github.com/stretchr/testify/require"
)

type User struct {
	ID       int64 //主键
	Username string
	Password string
}

// TestGet is a sample to run an endpoint test
func TestGet(t *testing.T) {

	contents, err := ssf.RestClient.Get("DBManager", "v1/users")
	require.Nil(t, err, "err must be nil")

	t.Logf(contents)
}

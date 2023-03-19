package inmemory_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"solaiman.me/cardgameapi/src/repository/inmemory"
)

func TestCanInstantiateInMemoryRepo(t *testing.T) {
	repo := inmemory.NewInMemoryRepository()

	assert.NotEmpty(t, repo, "Repo instance is an instance of Game Repository interface")
}

package context

import (
	"context"
	"testing"
)

func TestClient(t *testing.T) {
	err := Client(context.TODO())
	t.Log(err)
}

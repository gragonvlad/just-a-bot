package client

import (
	"context"
	"fmt"
	"os"
	"testing"
)

func TestClient(t *testing.T) {
	c := NewClient(os.Getenv("TOKEN"))

	users := c.RelatedUsers(context.Background(), "wrfly")
	for _, user := range users {
		fmt.Println(user)
	}
	fmt.Println(len(users))
}
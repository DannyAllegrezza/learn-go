package github

import (
	"context"
	"fmt"
	"testing"
)

func TestPayload_NoPanic(t *testing.T) {
	ctx := context.TODO()
	actual, err := GetUsersStarredResults(ctx, "dannyallegrezza")

	if err != nil {
		t.Errorf("There was an error")
	}

	fmt.Printf("Found a total of %v stars", len(actual))

	for _, result := range actual {

		fmt.Println(result.Name)
	}
}

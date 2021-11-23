package tests

import (
	"example-ci/helpers"
	"example-ci/models"
	"reflect"
	"testing"
)

var expectedPosts = &[]models.Post{
	{
		Id:          3,
		Content:     "Testing post with id 3",
		PublishedAt: "2021-08-19",
	},
	{
		Id:          4,
		Content:     "Testing post with id 4",
		PublishedAt: "2021-08-09",
	},
}

func TestDecodePosts(t *testing.T) {
	gotPosts, err := helpers.DecodePosts("samples/posts.json")
	if err != nil {
		t.Errorf("getting posts: %v", err)
	}

	if !reflect.DeepEqual(gotPosts, expectedPosts) {
		t.Fail()
	}
}

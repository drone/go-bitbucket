package bitbucket

import (
	"testing"
)

func Test_Pullrequests(t *testing.T) {

	// Create Pullrequest comment
	// NOTE: Create Pullrequest in testRepo before Testing.
	//       And the PR need `README.md` updated.
	comment, err := client.Pullrequests.CreateComment(
		testUser,
		testRepo,
		1,                      // PR id
		"README.md",            // Comment file
		2,                      // Comment line
		"This is test comment", // Comment body
	)
	if err != nil {
		t.Error(err)
	}
	if comment == nil {
		t.Error("Pullrequest Comment create failed")
	}


	// Delete Pullrequest comment
	comment, err = client.Pullrequests.DeleteComment(
		testUser,
		testRepo,
		1,                 // PR id
		comment.CommentId, // Comment id
	)
	if err != nil {
		t.Error(err)
	}
	if comment == nil {
		t.Error("Pullrequest Comment delete failed")
	}
}

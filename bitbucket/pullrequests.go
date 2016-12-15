package bitbucket

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
)

type PullrequestComment struct {
	Username        string      `json:"username"`
	PullRequestId   int         `json:"pull_request_id"`
	CommentId       int         `json:"comment_id"`
	DisplayMame     string      `json:"display_name"`
	Deleted         bool        `json:"deleted"`
	UtcLastUpdated  string      `json:"utc_last_updated"`
	FilenameHash    string      `json:"filename_hash"`
	BaseRev         string      `json:"base_rev"`
	Filename        string      `json:"filename"`
	Content         string      `json:"content"`
	ContentRendered string      `json:"content_rendered"`
	UserAvatarUrl   string      `json:"user_avatar_url"`
	LineFrom        json.Number `json:"line_from"`
	LineTo          json.Number `json:"line_to"`
	DestRev         string      `json:"dest_rev"`
	UtcCreatedOn    string      `json:"utc_created_on"`
	Anchor          string      `json:"anchor"`
	IsSpam          bool        `json:"is_spam"`
}

type PullrequestResource struct {
	client *Client
}

// Create Pullrequest Comment
func (r *PullrequestResource) CreateComment(owner, slug string, prId int, filename string, lineFrom int, content string) (*PullrequestComment, error) {
	values := url.Values{}
	values.Add("content", content)
	values.Add("filename", filename)
	values.Add("line_from", strconv.Itoa(lineFrom))

	c := PullrequestComment{}
	path := fmt.Sprintf("/repositories/%s/%s/pullrequests/%s/comments", owner, slug, strconv.Itoa(prId))
	if err := r.client.do("POST", path, nil, values, &c); err != nil {
		return nil, err
	}

	return &c, nil
}

// Delete Pullrequest Comment
func (r *PullrequestResource) DeleteComment(owner, slug string, prId int, commentId int) (*PullrequestComment, error) {
	c := PullrequestComment{}
	path := fmt.Sprintf("/repositories/%s/%s/pullrequests/%s/comments/%s", owner, slug, strconv.Itoa(prId), strconv.Itoa(commentId))
	if err := r.client.do("DELETE", path, nil, nil, &c); err != nil {
		return nil, err
	}

	return &c, nil
}
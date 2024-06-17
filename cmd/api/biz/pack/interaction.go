package pack

import (
	"tiktok/cmd/api/biz/model/api"
	"tiktok/kitex_gen/interaction"
)

func BuildComment(c *interaction.Comment) *api.Comment {
	return &api.Comment{
		UID:        c.Uid,
		Vid:        c.Vid,
		Cid:        c.Cid,
		ParentID:   c.ParentId,
		LikeCount:  c.LikeCount,
		ChildCount: c.ChildCount,
		Content:    c.Content,
		CreatedAt:  c.CreatedAt,
		UpdatedAt:  c.UpdatedAt,
		DeletedAt:  c.DeletedAt,
	}
}

func BuildComments(comments []*interaction.Comment) []*api.Comment {
	if len(comments) == 0 {
		return nil
	}

	resp := make([]*api.Comment, len(comments))
	for _, comment := range comments {
		resp = append(resp, BuildComment(comment))
	}
	return resp
}

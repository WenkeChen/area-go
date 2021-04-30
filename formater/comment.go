package formater

import (
	"AreaGo/model"
	"time"
)

type CommentListFormat struct {
	Id   uint                `json:"id"`
	List []CommentPairFormat `json:"list"`
}

type CommentPairFormat struct {
	Parent  CommentFormat `json:"parent"`
	Comment CommentFormat `json:"comment"`
}

type CommentFormat struct {
	Id        uint      `json:"id"`
	Content   string    `json:"content"`
	Useragent string    `json:"useragent"`
	CreatedAt time.Time `json:"created_at"`
}

func BuildCommentList(comments, parents []model.Comment) []CommentPairFormat {
	var commentPairs = make([]CommentPairFormat, len(comments))
	parentsMap := BuildParentsMap(parents)
	for i, comment := range comments {
		var pair CommentPairFormat
		pair.Comment = BuildCommentFormat(comment)
		pair.Parent = parentsMap[comment.Pid]
		commentPairs[i] = pair
	}

	return commentPairs
}

func BuildParentsMap(parents []model.Comment) map[uint]CommentFormat {
	var parentsMap = make(map[uint]CommentFormat, len(parents))
	for _, comment := range parents {
		parentsMap[comment.Id] = BuildCommentFormat(comment)
	}

	return parentsMap
}

func BuildCommentFormat(comment model.Comment) CommentFormat {
	return CommentFormat{
		Id:        comment.Id,
		Content:   comment.Content,
		Useragent: comment.Useragent,
		CreatedAt: comment.CreatedAt,
	}
}

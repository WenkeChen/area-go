package service

import (
	"AreaGo/model"
	"AreaGo/utils"
)

func GetPostComments(postId string, pageStr string) ([]model.Comment, error) {
	var comments []model.Comment
	size, offset, err := utils.GetOffset(pageStr)
	if err != nil {
		return nil, err
	}
	if err := model.Db.Where("post_id = ?", postId).
		Where("approved = ?", 1).
		Preload("User").
		Offset(offset).Limit(size).Find(&comments).Error; err != nil {
		return nil, err
	}

	return comments, nil
}

type CommentItem struct {
	Comment  model.Comment
	Children []CommentItem
}

func BuildCommentsTree(comments []model.Comment, args ...uint) []CommentItem {
	var tree []CommentItem
	if args == nil {
		args = append(args, 0)
	}
	for _, comment := range comments {
		if comment.Pid == args[0] {
			var item CommentItem
			item.Comment = comment
			item.Children = BuildCommentsTree(comments, comment.Id)
			tree = append(tree, item)
		}
	}
	return tree
}

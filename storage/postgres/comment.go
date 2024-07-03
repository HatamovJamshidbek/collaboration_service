package postgres

import (
	pb "collaboration_service/genproto"
	"collaboration_service/help"
	"database/sql"
	"time"
)

type CommentRepository struct {
	Db *sql.DB
}

func NewCommentRepositoryRepository(db *sql.DB) *CommentRepository {
	return &CommentRepository{Db: db}
}

func (repo CommentRepository) CreateComment(comment *pb.CreateCommitRequest) (*pb.Void, error) {
	_, err := repo.Db.Exec("insert into comments(composition_id,user_id,content,created_at) values ($1,$2,$3,$4)", comment.CompositionId, comment.UserId, comment.Content, time.Now())
	if err != nil {
		return nil, err
	}
	return &pb.Void{}, nil
}

//	func (repo CommentRepository) UpdateComment(comment *models.Comment, id string) (interface{}, error) {
//		return repo.Db.Exec("update comments set composition_id=$1,user_id=$2,content=$3,updated_at=$4 where id=$5 and deleted_at=0)", comment.Composition_id, comment.UpdatedAt, comment.Content, time.Now(), id)
//	}
//
//	func (repo CommentRepository) DeleteComment(id string) (interface{}, error) {
//		return repo.Db.Exec("update comments set deleted_at=$1 where id=$2 and deleted_at is null)", time.Now(), id)
//	}
//
//	func (repo CommentRepository) GetCommentById(id string) (*models.Comment, error) {
//		rows, err := repo.Db.Query("select composition_id,user_id ,conetnt from comments  where id=$1 and deleted_at is null)", id)
//		if err != nil {
//			return nil, err
//		}
//		var comment models.Comment
//		for rows.Next() {
//			err := rows.Scan(&comment.Composition_id, &comment.User_id, &comment.Content)
//			if err != nil {
//				return nil, err
//			}
//			return &comment, nil
//		}
//		return nil, err
//	}
func (repo CommentRepository) GetComment(comment *pb.GetCommitRequest) (*pb.CommitsResponse, error) {
	var (
		params = make(map[string]interface{})
		arr    []interface{}
		limit  string
		offset string
	)
	filter := ""
	if len(comment.CompositionId) > 0 {
		params["composition_id"] = comment.CompositionId
		filter += " and composition_id = :composition_id "

	}
	if len(comment.UserId) > 0 {
		params["user_id"] = comment.UserId
		filter += " and user_id = :user_id "

	}
	if len(comment.Content) > 0 {
		params["content"] = comment.Content
		filter += " and content = :content "

	}

	if comment.LimitOffset.Limit > 0 {
		params["limit"] = comment.LimitOffset.Limit
		limit = ` LIMIT :limit`

	}
	if comment.LimitOffset.Offset > 0 {
		params["offset"] = comment.LimitOffset.Offset
		limit = ` OFFSET :offset`

	}
	query := "select composition_id,user_id ,content from comments  where  deleted_at is null and composition_id=$1"

	query = query + filter + limit + offset
	query, arr = strorage.ReplaceQueryParams(query, params)
	rows, err := repo.Db.Query(query, arr...)
	arr = append(arr, comment.CompositionId)
	var comments []*pb.CommitResponse
	for rows.Next() {
		var comment pb.CommitResponse
		err := rows.Scan(&comment.CompositionId, &comment.UserId, &comment.Content)
		if err != nil {
			return nil, err
		}
		comments = append(comments, &comment)
	}
	return &pb.CommitsResponse{CommitsResponse: comments}, err
}

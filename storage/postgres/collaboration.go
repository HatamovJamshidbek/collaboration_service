package postgres

import (
	pb "collaboration_service/genproto"
	strorage "collaboration_service/help"
	"database/sql"
	"time"
)

type CollaborationRepository struct {
	Db *sql.DB
}

func NewCollaborationRepositoryRepository(db *sql.DB) *CollaborationRepository {
	return &CollaborationRepository{Db: db}
}

//	func (repo CollaborationRepository) CreateCollaboration(compositionId string, collaboration *models.Collaboration) (interface{}, error) {
//		return repo.Db.Exec("insert into collaborations(composition_id,user_id,role,created_at)", compositionId, collaboration.User_Id, collaboration.Role, time.Now())
//	}
func (repo CollaborationRepository) UpdateCollaboration(collaboration *pb.UpdateCollaborationRequest) (*pb.Void, error) {
	_, err := repo.Db.Exec("update collaborations set composition_id=$1,user_id=$2,role=$3,updated_at=$4 where user_id=$5 and deleted_at=0 and composition_id)", collaboration.CompositionId, collaboration.Userid, collaboration.Role, time.Now(), collaboration.Userid)
	if err != nil {
		return nil, err
	}
	return &pb.Void{}, nil
}

func (repo CollaborationRepository) DeleteCollaboration(collaboration *pb.DeleteCollaborationRequest) (*pb.Void, error) {
	_, err := repo.Db.Exec("update collaborations set deleted_at=$1 where user_id=$2 and deleted_at is null and collobartion_id=$3)", time.Now(), collaboration.Userid, collaboration.CompositionId)
	if err != nil {
		return nil, err
	}
	return &pb.Void{}, nil
}

//	func (repo CollaborationRepository) GetCollaborationById(compositionId string,userId string) (*models.Collaboration, error) {
//		rows, err := repo.Db.Query("select composition_id,user_id ,role from collaborations  where id=$1 and deleted_at is null)", id)
//		if err != nil {
//			return nil, err
//		}
//		var collaboration models.Collaboration
//		for rows.Next() {
//			err := rows.Scan(&collaboration.Composition_id, &collaboration.User_Id, &collaboration.Role)
//			if err != nil {
//				return nil, err
//			}
//			return &collaboration, nil
//		}
//		return nil, err
//	}
func (repo CollaborationRepository) GetCollaboration(collaboration *pb.GetCollaboratorsRequest) (*pb.CollaborationsResponse, error) {
	var (
		params = make(map[string]interface{})
		arr    []interface{}
		limit  string
		offset string
	)
	filter := ""
	if len(collaboration.CompositionId) > 0 {
		params["composition_id"] = collaboration.CompositionId
		filter += " and composition_id = :composition_id "

	}
	if len(collaboration.UserId) > 0 {
		params["user_id"] = collaboration.UserId
		filter += " and user_id = :user_id "

	}
	if len(collaboration.Role) > 0 {
		params["role"] = collaboration.Role
		filter += " and role = :role "

	}

	if collaboration.LimitOffset.Limit > 0 {
		params["limit"] = collaboration.LimitOffset.Limit
		limit = ` LIMIT :limit`

	}
	if collaboration.LimitOffset.Offset > 0 {
		params["offset"] = collaboration.LimitOffset.Offset
		limit = ` OFFSET :offset`

	}
	query := "select composition_id,user_id ,role from collaborations  where  deleted_at is null and composition_id=$1"

	query = query + filter + limit + offset
	query, arr = strorage.ReplaceQueryParams(query, params)
	arr = append(arr, collaboration.CompositionId)
	rows, err := repo.Db.Query(query, arr...)
	var collaborations []*pb.CollaborationResponse
	for rows.Next() {
		var collaboration pb.CollaborationResponse
		err := rows.Scan(&collaboration.CompositionId, &collaboration.Userid, &collaboration.Role)
		if err != nil {
			return nil, err
		}
		collaborations = append(collaborations, &collaboration)
	}
	return &pb.CollaborationsResponse{Collaborations: collaborations}, err
}

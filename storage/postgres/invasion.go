package postgres

import (
	pb "collaboration_service/genproto"
	"database/sql"
	"time"
)

type InvasionRepository struct {
	Db *sql.DB
}

func NewInvasionRepository(db *sql.DB) *InvasionRepository {
	return &InvasionRepository{Db: db}
}

func (repo InvasionRepository) CreateInvasion(invasion *pb.CreateInviteRequest) (*pb.Void, error) {
	_, err := repo.Db.Exec("insert into invasions(composition_id,invert_id,invitee_id,status,created_at)", invasion.CompositionId, invasion.InvertId, invasion.InviteeId, invasion.Status, time.Now())
	if err != nil {
		return nil, err
	}
	return &pb.Void{}, nil
}
func (repo InvasionRepository) UpdateInvasion(invasion *pb.UpdateInviteRequest) (*pb.Void, error) {
	_, err := repo.Db.Exec("update invasions set composition_id=$1,user_id=$2,title=$3,file_url=$4,updated_at=$5 where id=$6 and deleted_at=0)", invasion.CompositionId, invasion.InvertId, invasion.InviteeId, invasion.Status, time.Now(), invasion.Id)
	if err != nil {
		return nil, err
	}
	return &pb.Void{}, nil
}

//func (repo InvasionRepository) DeleteInvasion(id string) (interface{}, error) {
//	return repo.Db.Exec("update invasions set deleted_at=$1 where id=$2 and deleted_at is null)", time.Now(), id)
//}
//func (repo InvasionRepository) GetInvasionById(id string) (*models.Invasion, error) {
//	rows, err := repo.Db.Query("select composition_id,user_id ,title, file_url from invasions  where id=$1 and deleted_at is null)", id)
//	if err != nil {
//		return nil, err
//	}
//	var invasion models.Invasion
//	for rows.Next() {
//		err := rows.Scan(&invasion.Composition_id, &invasion.Invert_id, &invasion.Invete_id, &invasion.Status)
//		if err != nil {
//			return nil, err
//		}
//		return &invasion, nil
//	}
//	return nil, err
//}
//func (repo InvasionRepository) GetInvasion(invasion *models.Invasion) (*[]models.Invasion, error) {
//	var (
//		params = make(map[string]interface{})
//		arr    []interface{}
//		limit  string
//		offset string
//	)
//	filter := ""
//	if len(invasion.Composition_id) > 0 {
//		params["composition_id"] = invasion.Composition_id
//		filter += " and composition_id = :composition_id "
//
//	}
//	if len(invasion.Invert_id) > 0 {
//		params["invert_id"] = invasion.Invert_id
//		filter += " and invert_id = :invert_id "
//
//	}
//	if len(invasion.Status) > 0 {
//		params["status"] = invasion.Status
//		filter += " and status = :status "
//
//	}
//	if len(invasion.Invete_id) > 0 {
//		params["invete_id"] = invasion.Invete_id
//		filter += " and invete_id = :invete_id "
//
//	}
//
//	if invasion.Limit > 0 {
//		params["limit"] = invasion.Limit
//		limit = ` LIMIT :limit`
//
//	}
//	if invasion.Offset > 0 {
//		params["offset"] = invasion.Offset
//		limit = ` OFFSET :offset `
//
//	}
//
//	query := "select id,composition_id,invert_id,invete_id,status from invasions  where  deleted_at is null"
//
//	query = query + filter + limit + offset
//	query, arr = strorage.ReplaceQueryParams(query, params)
//	rows, err := repo.Db.Query(query, arr...)
//	var invasions []models.Invasion
//	for rows.Next() {
//		var invasion models.Invasion
//		err := rows.Scan(&invasion.Id, &invasion.Composition_id, &invasion.Invert_id, &invasion.Invete_id, &invasion.Status)
//		if err != nil {
//			return nil, err
//		}
//		invasions = append(invasions, invasion)
//	}
//	return &invasions, err
//}

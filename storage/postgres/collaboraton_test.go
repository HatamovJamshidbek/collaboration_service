package postgres

import (
	"collaboration_service/config"
	pb "collaboration_service/genproto"
	"fmt"
	"reflect"
	"testing"
)

func TestCollaborationRepository_DeleteCollaboration(t *testing.T) {
	cnf := config.Config{}
	db, err := ConnectionDb(&cnf)
	if err != nil {
		//fmt.Println("salom=================")
		//panic(err)
	}
	col := NewCollaborationRepositoryRepository(db)
	collaboration := pb.DeleteCollaborationRequest{
		CompositionId: "",
		Userid:        "",
	}
	response, err := col.DeleteCollaboration(&collaboration)
	if err != nil {
		fmt.Println("+++++++++++")
	}
	if !reflect.DeepEqual(response, &pb.Void{}) {
		t.Errorf("Response does not match expected value.\nGot: %+v\nExpected: %+v", response, &pb.Void{})
	}

}

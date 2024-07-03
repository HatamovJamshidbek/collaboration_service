package postgres

import (
	"collaboration_service/config"
	pb "collaboration_service/genproto"
	"fmt"
	"reflect"
	"testing"
)

func TestCommentRepository_CreateComment(t *testing.T) {
	cnf := config.Config{}
	db, err := ConnectionDb(&cnf)
	if err != nil {
		//fmt.Println("salom=================")
		//panic(err)
	}
	comm := NewCommentRepositoryRepository(db)
	comment := pb.CreateCommitRequest{
		CompositionId: "ad6b59c5-8d12-4b9e-ab1e-4d3ba7620c01",
		UserId:        "d3dcbdff-de1c-452d-94da-2bb783f1016a",
		Content:       "salom",
	}
	response, err := comm.CreateComment(&comment)
	if err != nil {
		fmt.Println("++++++", err)
	}
	if !reflect.DeepEqual(response, &pb.Void{}) {
		t.Errorf("Response does not match expected value.\nGot: %+v\nExpected: %+v", response, &pb.Void{})

	}
}
func TestCommentRepository_GetComment(t *testing.T) {
	cnf := config.Config{}
	db, err := ConnectionDb(&cnf)
	if err != nil {
		fmt.Println("salom=================")
		panic(err)
	}

	comm := NewCommentRepositoryRepository(db)
	comment := pb.GetCommitRequest{
		CompositionId: "ad6b59c5-8d12-4b9e-ab1e-4d3ba7620c01",
		UserId:        "d3dcbdff-de1c-452d-94da-2bb783f1016a",
	}
	fmt.Println(&comment, comm)
	//response, err := comm.GetComment(&comment)
	//if err != nil {
	//	fmt.Println("++++++++", err)
	//	panic(err)
	//}
	//if !reflect.DeepEqual(response, &pb.CommitsResponse{}) {
	//	t.Errorf("Response does not match expected value.\nGot: %+v\nExpected: %+v", response, &pb.Void{})
	//
	//}

}

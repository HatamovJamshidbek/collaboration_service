package postgres

import (
	"collaboration_service/config"
	pb "collaboration_service/genproto"
	"fmt"
	"reflect"
	"testing"
)

func TestInvasionRepository_CreateInvasion(t *testing.T) {
	cnf := config.Config{}
	db, err := ConnectionDb(&cnf)
	if err != nil {
		//fmt.Println("salom=================")
		//panic(err)
	}
	inv := NewInvasionRepository(db)
	invite := pb.CreateInviteRequest{
		CompositionId: "ad6b59c5-8d12-4b9e-ab1e-4d3ba7620c01",
		InviteeId:     "d3dcbdff-de1c-452d-94da-2bb783f1016a",
		InvertId:      "d3dcbdff-de1c-452d-94da-2bb783f1016a",
		Status:        "pending",
	}
	response, err := inv.CreateInvite(&invite)
	if err != nil {
		fmt.Println("++++++", err)
	}
	if !reflect.DeepEqual(response, &pb.Void{}) {
		t.Errorf("Response does not match expected value.\nGot: %+v\nExpected: %+v", response, &pb.Void{})

	}

}

func TestInvasionRepository_UpdateInvasion(t *testing.T) {
	cnf := config.Config{}
	db, err := ConnectionDb(&cnf)
	if err != nil {
		//fmt.Println("salom=================")
		//panic(err)
	}
	inv := NewInvasionRepository(db)
	invite := pb.UpdateInviteRequest{
		Id:            "e5f6e346-27d3-465d-a738-e934d7791f7c",
		CompositionId: "ad6b59c5-8d12-4b9e-ab1e-4d3ba7620c01",
		InviteeId:     "d3dcbdff-de1c-452d-94da-2bb783f1016a",
		InvertId:      "d3dcbdff-de1c-452d-94da-2bb783f1016a",
		Status:        "pending",
	}
	response, err := inv.UpdateInvite(&invite)
	if err != nil {
		fmt.Println("++++++", err)
	}
	if !reflect.DeepEqual(response, &pb.Void{}) {
		t.Errorf("Response does not match expected value.\nGot: %+v\nExpected: %+v", response, &pb.Void{})

	}

}

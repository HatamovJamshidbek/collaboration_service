package service

import (
	pb "collaboration_service/genproto"
	"collaboration_service/storage/postgres"
	"context"
)

type CollaborationService struct {
	Inv *postgres.InvasionRepository
	Col *postgres.CollaborationRepository
	Com *postgres.CommentRepository
	pb.UnimplementedCollaborationServiceServer
}

func NewCollaborationService(inv *postgres.InvasionRepository, col *postgres.CollaborationRepository, com *postgres.CommentRepository) *CollaborationService {
	return &CollaborationService{Inv: inv, Col: col, Com: com}
}

func (service *CollaborationService) CreateInvite(ctx context.Context, in *pb.CreateInviteRequest) (*pb.Void, error) {
	response, err := service.Inv.CreateInvite(in)
	if err != nil {
		return nil, err
	}
	return response, nil
}
func (service *CollaborationService) UpdateInvite(ctx context.Context, in *pb.UpdateInviteRequest) (*pb.Void, error) {
	response, err := service.Inv.UpdateInvite(in)
	if err != nil {
		return nil, err
	}
	return response, nil
}
func (service *CollaborationService) GetCollaborators(ctx context.Context, in *pb.GetCollaboratorsRequest) (*pb.CollaborationsResponse, error) {
	response, err := service.Col.GetCollaborators(in)
	if err != nil {
		return nil, err
	}
	return response, nil
}
func (service *CollaborationService) UpdateCollaborators(ctx context.Context, in *pb.UpdateCollaborationRequest) (*pb.Void, error) {
	response, err := service.Col.UpdateCollaboration(in)
	if err != nil {
		return nil, err
	}
	return response, err
}
func (service *CollaborationService) DeleteCollaborators(ctx context.Context, in *pb.DeleteCollaborationRequest) (*pb.Void, error) {
	response, err := service.Col.DeleteCollaboration(in)
	if err != nil {
		return nil, err
	}
	return response, err
}
func (service *CollaborationService) CreateComment(ctx context.Context, in *pb.CreateCommitRequest) (*pb.Void, error) {
	response, err := service.Com.CreateComment(in)
	if err != nil {
		return nil, err
	}
	return response, err
}
func (service *CollaborationService) GetComment(ctx context.Context, in *pb.GetCommitRequest) (*pb.CommitsResponse, error) {
	response, err := service.Com.GetComment(in)
	if err != nil {
		return nil, err
	}
	return response, err
}

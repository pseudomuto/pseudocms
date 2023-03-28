package server

import (
	"context"

	"github.com/gofrs/uuid"
	v1 "github.com/pseudomuto/pseudocms/pkg/api/v1"
	"github.com/pseudomuto/pseudocms/pkg/models"
)

// AdminService creates a new v1.AdminServiceServer.
func AdminService(rf RepoFactory) v1.AdminServiceServer {
	return &adminService{repoFactory: rf}
}

type adminService struct {
	repoFactory RepoFactory
}

func (s *adminService) CreateDefinition(
	ctx context.Context,
	r *v1.CreateDefinitionRequest,
) (*v1.CreateDefinitionResponse, error) {
	def := models.DefinitionFromProto(&v1.Definition{
		Name:        r.Name,
		Description: r.Description,
		Fields:      r.Fields,
	})

	if err := s.repoFactory.Definitions().Create(def, models.CreateOptions{Eager: true}); err != nil {
		return nil, err
	}

	return &v1.CreateDefinitionResponse{
		Definition: def.ToProto(),
	}, nil
}

func (s *adminService) GetDefinition(
	ctx context.Context,
	r *v1.GetDefinitionRequest,
) (*v1.GetDefinitionResponse, error) {
	id, err := uuid.FromString(r.Id)
	if err != nil {
		return nil, err
	}

	def, err := s.repoFactory.Definitions().Find(id, models.FindOptions{Eager: true})
	if err != nil {
		return nil, err
	}

	return &v1.GetDefinitionResponse{Definition: def.ToProto()}, nil
}

func (s *adminService) CreateField(
	ctx context.Context,
	r *v1.CreateFieldRequest,
) (*v1.CreateFieldResponse, error) {
	id, err := uuid.FromString(r.DefinitionId)
	if err != nil {
		return nil, err
	}

	field := models.FieldFromProto(&v1.Field{
		Name:        r.Name,
		Description: r.Description,
		FieldType:   r.FieldType,
		Constraints: r.Constraints,
	})
	field.DefinitionID = id

	if err := s.repoFactory.Fields().Create(field, models.CreateOptions{}); err != nil {
		return nil, err
	}

	return &v1.CreateFieldResponse{
		Field: field.ToProto(),
	}, nil
}

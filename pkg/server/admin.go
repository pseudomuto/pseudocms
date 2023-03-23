package server

import (
	"context"

	"github.com/gofrs/uuid"
	v1 "github.com/pseudomuto/pseudocms/pkg/api/v1"
	"github.com/pseudomuto/pseudocms/pkg/models"
)

// AdminService creates a new v1.AdminServiceServer.
func AdminService(dr DefinitionsRepo, fr FieldsRepo) v1.AdminServiceServer {
	return &adminService{defs: dr, fields: fr}
}

type adminService struct {
	defs   DefinitionsRepo
	fields FieldsRepo
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

	if err := s.defs.Create(def, models.CreateOptions{Eager: true}); err != nil {
		return nil, err
	}

	return &v1.CreateDefinitionResponse{
		Definition: def.ToProto(),
	}, nil
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

	if err := s.fields.Create(field, models.CreateOptions{}); err != nil {
		return nil, err
	}

	return &v1.CreateFieldResponse{
		Field: field.ToProto(),
	}, nil
}

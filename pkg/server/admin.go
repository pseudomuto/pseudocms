package server

import (
	"context"

	v1 "github.com/pseudomuto/pseudocms/pkg/api/v1"
	"github.com/pseudomuto/pseudocms/pkg/models"
)

// AdminService creates a new v1.AdminServiceServer.
func AdminService(r DefinitionsRepo) v1.AdminServiceServer {
	return &adminService{db: r}
}

type adminService struct {
	db DefinitionsRepo
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

	if err := s.db.Create(def, models.CreateOptions{Eager: true}); err != nil {
		return nil, err
	}

	return &v1.CreateDefinitionResponse{
		Definition: def.ToProto(),
	}, nil
}

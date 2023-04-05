package server

import (
	"context"

	"github.com/gofrs/uuid"
	v1 "github.com/pseudomuto/pseudocms/pkg/api/v1"
	"github.com/pseudomuto/pseudocms/pkg/models"
	"github.com/pseudomuto/pseudocms/pkg/repo"
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

	if err := s.repoFactory.Definitions().Create(def, repo.CreateOptions{Eager: true}); err != nil {
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

	def, err := s.repoFactory.Definitions().Find(id, repo.FindOptions{Eager: true})
	if err != nil {
		return nil, err
	}

	return &v1.GetDefinitionResponse{Definition: def.ToProto()}, nil
}

func (s *adminService) ListDefinitions(
	r *v1.ListDefinitionsRequest,
	stream v1.AdminService_ListDefinitionsServer,
) error {
	opts := []repo.ListOption{
		repo.Eager(r.Eager),
		repo.OrderBy(r.OrderBy, r.SortDirection.ToSQL()),
	}

	if r.AfterKey != "" {
		opts = append(opts, repo.AfterKey(r.AfterKey))
	}

	// If the caller specified a max results value, just get the one page.
	if r.MaxResults.GetValue() != 0 {
		opts = append(opts, repo.PageSize(int(r.MaxResults.GetValue())))
		_, err := s.listDefinitions(stream, opts)
		return err
	}

	// Otherwise, auto-paginate all results.
	var err error
	for {
		opts, err = s.listDefinitions(stream, opts)
		if err != nil {
			return err
		}

		if opts == nil {
			return nil
		}
	}
}

func (s *adminService) listDefinitions(
	stream v1.AdminService_ListDefinitionsServer,
	opts []repo.ListOption,
) ([]repo.ListOption, error) {
	defs, err := s.repoFactory.Definitions().List(opts...)
	if err != nil {
		return nil, err
	}

	for _, def := range defs.Results {
		if err := stream.Send(def.ToProto()); err != nil {
			return nil, err
		}
	}

	if defs.LastPage {
		return nil, nil
	}

	return append(opts, repo.AfterKey(defs.LastKey)), nil
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

	if err := s.repoFactory.Fields().Create(field, repo.CreateOptions{}); err != nil {
		return nil, err
	}

	return &v1.CreateFieldResponse{
		Field: field.ToProto(),
	}, nil
}

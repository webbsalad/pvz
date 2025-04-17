package v1

import (
	"context"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/webbsalad/pvz/internal/config"
	"github.com/webbsalad/pvz/internal/model"
	mock_item_repository "github.com/webbsalad/pvz/internal/repository/item/mock"
)

var (
	testPVZID, _       = model.NewPVZID("2b98ee88-7970-4e6f-b325-ccf3ce10909f")
	testProductID, _   = model.NewProductID("3b98ee88-7970-4e6f-b325-ccf3ce10909f")
	testReceptionID, _ = model.NewReceptionID("4b98ee88-7970-4e6f-b325-ccf3ce10909f")

	testProduct             = model.Product{ID: testProductID, ReceptionID: testReceptionID, Type: testType}
	testInProgressReception = model.Reception{ID: testReceptionID, PVZID: testPVZID, Status: model.IN_PROGRESS}
	testCloseReception      = model.Reception{ID: testReceptionID, PVZID: testPVZID, Status: model.CLOSE}

	testReceptionFilter = model.ReceptionFilter{PVZID: &testPVZID, Status: &testStatus}
	testProductFilter   = model.ProductFilter{ReceptionID: &testReceptionID}

	testType   = "одежда"
	testStatus = model.IN_PROGRESS
)

type serviceTestDeps struct {
	Service *Service

	ctx            context.Context
	itemRepository *mock_item_repository.MockRepository
}

func getTestDeps(t *testing.T) *serviceTestDeps {
	ctrl := gomock.NewController(t)
	itemRepository := mock_item_repository.NewMockRepository(ctrl)

	return &serviceTestDeps{
		Service: &Service{
			itemRepository: itemRepository,
			config:         config.Config{},
		},

		ctx:            context.Background(),
		itemRepository: itemRepository,
	}
}

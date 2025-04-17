package v1

import (
	"context"
	"testing"
	"time"

	gomock "github.com/golang/mock/gomock"
	"github.com/webbsalad/pvz/internal/config"
	"github.com/webbsalad/pvz/internal/model"
	mock_item_repository "github.com/webbsalad/pvz/internal/repository/item/mock"
	mock_pvz_repository "github.com/webbsalad/pvz/internal/repository/pvz/mock"
)

var (
	testPVZID, _       = model.NewPVZID("2b98ee88-7970-4e6f-b325-ccf3ce10909f")
	testProductID, _   = model.NewProductID("3b98ee88-7970-4e6f-b325-ccf3ce10909f")
	testReceptionID, _ = model.NewReceptionID("4b98ee88-7970-4e6f-b325-ccf3ce10909f")

	testProduct                = model.Product{ID: testProductID, ReceptionID: testReceptionID, Type: testType}
	testReception              = model.Reception{ID: testReceptionID, PVZID: testPVZID, Status: testStatus}
	testPVZ                    = model.PVZ{ID: testPVZID, City: testCity}
	testProducts               = []model.Product{testProduct}
	testReceptions             = []model.Reception{testReception}
	testPVZs                   = []model.PVZ{testPVZ}
	testReceptionWithProducts  = model.ReceptionWithProducts{Reception: testReception, Products: testProducts}
	testReceptionsWithProducts = []model.ReceptionWithProducts{testReceptionWithProducts}
	testPVZWithReceptions      = model.PVZWithReceptions{PVZ: testPVZ, Receptions: testReceptionsWithProducts}
	testPVZsWithReceptions     = []model.PVZWithReceptions{testPVZWithReceptions}

	testReceptionWithProductsWithoutProducts  = model.ReceptionWithProducts{Reception: testReception} // :D
	testReceptionsWithProductsWithoutProducts = []model.ReceptionWithProducts{testReceptionWithProductsWithoutProducts}
	testPVZWithReceptionsWithoutProducts      = model.PVZWithReceptions{PVZ: testPVZ, Receptions: testReceptionsWithProductsWithoutProducts}
	testPVZsWithReceptionsWithoutProducts     = []model.PVZWithReceptions{testPVZWithReceptionsWithoutProducts}

	testReceptionFilter = model.ReceptionFilter{From: &testFrom, To: &testTo}
	testPVZFilter       = model.PVZFilter{IDs: []model.PVZID{testPVZID}, Page: &testPage, Limit: &testLimit}
	testProductFilter   = model.ProductFilter{ReceptionID: &testReceptionID}

	testType   = "одежда"
	testCity   = "Санкт-Петербург"
	testStatus = model.CLOSE

	testPage    int32 = 2
	testLimit   int32 = 5
	testFrom, _       = time.Parse("2006-01-02T15:04:05.999999Z", "2020-04-17T11:07:39.780502Z")
	testTo, _         = time.Parse("2006-01-02T15:04:05.999999Z", "2029-04-17T11:07:39.780502Z")
)

type serviceTestDeps struct {
	Service *Service

	ctx            context.Context
	pvzRepository  *mock_pvz_repository.MockRepository
	itemRepository *mock_item_repository.MockRepository
}

func getTestDeps(t *testing.T) *serviceTestDeps {
	ctrl := gomock.NewController(t)
	pvzRepository := mock_pvz_repository.NewMockRepository(ctrl)
	itemRepository := mock_item_repository.NewMockRepository(ctrl)

	return &serviceTestDeps{
		Service: &Service{
			pvzRepository:  pvzRepository,
			itemRepository: itemRepository,
			config:         config.Config{},
		},

		ctx:            context.Background(),
		pvzRepository:  pvzRepository,
		itemRepository: itemRepository,
	}
}

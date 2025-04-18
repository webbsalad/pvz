package v1

import (
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/webbsalad/pvz/internal/model"
)

func TestService_AddProduct(t *testing.T) {
	t.Parallel()

	type args struct {
		userRole        model.Role
		receptionFilter model.ReceptionFilter
		reception       model.Reception
		pvzID           model.PVZID
		productType     string
		product         model.Product
		productFields   model.Product
	}

	type result struct {
		product model.Product
		err     error
	}

	type testCase struct {
		name   string
		args   args
		mocks  func(tc testCase, deps *serviceTestDeps)
		result result
	}

	testCases := []testCase{
		{
			name: "success",
			mocks: func(tc testCase, deps *serviceTestDeps) {
				deps.itemRepository.EXPECT().
					GetReceptionsByParams(gomock.Any(), tc.args.receptionFilter).
					Return([]model.Reception{tc.args.reception}, nil)

				deps.itemRepository.EXPECT().
					AddProduct(gomock.Any(), tc.args.productFields).
					Return(tc.args.product, nil)
			},
			args: args{
				userRole:        model.EMPLOYEE,
				receptionFilter: testReceptionFilter,
				reception:       testInProgressReception,
				pvzID:           testPVZID,
				productType:     testType,
				product:         testProduct,
				productFields: model.Product{
					ReceptionID: testReceptionID,
					Type:        testType,
				},
			},
			result: result{
				product: testProduct,
				err:     nil,
			},
		},
		{
			name:  "wrong role moderator",
			mocks: func(tc testCase, deps *serviceTestDeps) {},
			args: args{
				userRole:    model.MODERATOR,
				pvzID:       testPVZID,
				productType: testType,
			},
			result: result{
				err: model.ErrWrongRole,
			},
		},
		{
			name:  "wrong role client",
			mocks: func(tc testCase, deps *serviceTestDeps) {},
			args: args{
				userRole:    model.CLIENT,
				pvzID:       testPVZID,
				productType: testType,
			},
			result: result{
				err: model.ErrWrongRole,
			},
		},
		{
			name: "receptions not found",
			mocks: func(tc testCase, deps *serviceTestDeps) {
				deps.itemRepository.EXPECT().
					GetReceptionsByParams(gomock.Any(), tc.args.receptionFilter).
					Return([]model.Reception{}, model.ErrReceptionNotFound)

			},
			args: args{
				userRole:        model.EMPLOYEE,
				receptionFilter: testReceptionFilter,
				pvzID:           testPVZID,
				productType:     testType,
			},
			result: result{
				err: model.ErrReceptionNotFound,
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			deps := getTestDeps(t)

			tc.mocks(tc, deps)

			product, err := deps.Service.AddProduct(deps.ctx, tc.args.userRole, tc.args.pvzID, model.Product{Type: tc.args.productType})

			require.ErrorIs(t, err, tc.result.err)
			if err == nil {
				require.Equal(t, tc.result.product, product)
			}
		})
	}
}

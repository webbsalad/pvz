package v1

import (
	"fmt"
	"testing"
	"time"

	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/webbsalad/pvz/internal/model"
)

func TestService_GetPVZIntervalList(t *testing.T) {
	t.Parallel()

	type args struct {
		userRole        model.Role
		page            *int32
		limit           *int32
		from            *time.Time
		to              *time.Time
		receptionFilter model.ReceptionFilter
		receptions      []model.Reception
		pvzFilter       model.PVZFilter
		pvzs            []model.PVZ
		productFilter   model.ProductFilter
		products        []model.Product
	}

	type result struct {
		pvzsWithReceptions []model.PVZWithReceptions
		err                error
	}

	type testCase struct {
		name   string
		args   args
		mocks  func(tc testCase, deps *serviceTestDeps)
		result result
	}

	testCases := []testCase{
		{
			name: "success employee",
			mocks: func(tc testCase, deps *serviceTestDeps) {
				fmt.Println(tc.args.receptionFilter)
				deps.itemRepository.EXPECT().
					GetReceptionsByParams(gomock.Any(), tc.args.receptionFilter).
					Return(tc.args.receptions, nil)

				deps.pvzRepository.EXPECT().
					GetPVZsByParams(gomock.Any(), tc.args.pvzFilter).
					Return(tc.args.pvzs, nil)

				deps.itemRepository.EXPECT().
					GetProductssByParams(gomock.Any(), tc.args.productFilter).
					Return(tc.args.products, nil)
			},
			args: args{
				userRole:        model.EMPLOYEE,
				page:            &testPage,
				limit:           &testLimit,
				from:            &testFrom,
				to:              &testTo,
				receptionFilter: testReceptionFilter,
				pvzFilter:       testPVZFilter,
				productFilter:   testProductFilter,
				receptions:      testReceptions,
				products:        testProducts,
				pvzs:            testPVZs,
			},
			result: result{
				pvzsWithReceptions: testPVZsWithReceptions,
				err:                nil,
			},
		},
		{
			name: "success moderator",
			mocks: func(tc testCase, deps *serviceTestDeps) {
				fmt.Println(tc.args.receptionFilter)
				deps.itemRepository.EXPECT().
					GetReceptionsByParams(gomock.Any(), tc.args.receptionFilter).
					Return(tc.args.receptions, nil)

				deps.pvzRepository.EXPECT().
					GetPVZsByParams(gomock.Any(), tc.args.pvzFilter).
					Return(tc.args.pvzs, nil)

				deps.itemRepository.EXPECT().
					GetProductssByParams(gomock.Any(), tc.args.productFilter).
					Return(tc.args.products, nil)
			},
			args: args{
				userRole:        model.MODERATOR,
				page:            &testPage,
				limit:           &testLimit,
				from:            &testFrom,
				to:              &testTo,
				receptionFilter: testReceptionFilter,
				pvzFilter:       testPVZFilter,
				productFilter:   testProductFilter,
				receptions:      testReceptions,
				products:        testProducts,
				pvzs:            testPVZs,
			},
			result: result{
				pvzsWithReceptions: testPVZsWithReceptions,
				err:                nil,
			},
		},
		{
			name: "success (product not found)",
			mocks: func(tc testCase, deps *serviceTestDeps) {
				fmt.Println(tc.args.receptionFilter)
				deps.itemRepository.EXPECT().
					GetReceptionsByParams(gomock.Any(), tc.args.receptionFilter).
					Return(tc.args.receptions, nil)

				deps.pvzRepository.EXPECT().
					GetPVZsByParams(gomock.Any(), tc.args.pvzFilter).
					Return(tc.args.pvzs, nil)

				deps.itemRepository.EXPECT().
					GetProductssByParams(gomock.Any(), tc.args.productFilter).
					Return(tc.args.products, model.ErrProductNotFound)
			},
			args: args{
				userRole:        model.MODERATOR,
				page:            &testPage,
				limit:           &testLimit,
				from:            &testFrom,
				to:              &testTo,
				receptionFilter: testReceptionFilter,
				pvzFilter:       testPVZFilter,
				productFilter:   testProductFilter,
				receptions:      testReceptions,
				products:        nil,
				pvzs:            testPVZs,
			},
			result: result{
				pvzsWithReceptions: testPVZsWithReceptionsWithoutProducts,
				err:                nil,
			},
		},
		{
			name: "reception not found",
			mocks: func(tc testCase, deps *serviceTestDeps) {
				fmt.Println(tc.args.receptionFilter)
				deps.itemRepository.EXPECT().
					GetReceptionsByParams(gomock.Any(), tc.args.receptionFilter).
					Return(tc.args.receptions, model.ErrReceptionNotFound)

			},
			args: args{
				userRole:        model.EMPLOYEE,
				page:            &testPage,
				limit:           &testLimit,
				from:            &testFrom,
				to:              &testTo,
				receptionFilter: testReceptionFilter,
				receptions:      nil,
			},
			result: result{
				pvzsWithReceptions: nil,
				err:                model.ErrReceptionNotFound,
			},
		},
		{
			name: "pvz not found",
			mocks: func(tc testCase, deps *serviceTestDeps) {
				fmt.Println(tc.args.receptionFilter)
				deps.itemRepository.EXPECT().
					GetReceptionsByParams(gomock.Any(), tc.args.receptionFilter).
					Return(tc.args.receptions, nil)

				deps.pvzRepository.EXPECT().
					GetPVZsByParams(gomock.Any(), tc.args.pvzFilter).
					Return(tc.args.pvzs, model.ErrPVZNotFound)
			},
			args: args{
				userRole:        model.MODERATOR,
				page:            &testPage,
				limit:           &testLimit,
				from:            &testFrom,
				to:              &testTo,
				receptionFilter: testReceptionFilter,
				pvzFilter:       testPVZFilter,
				receptions:      testReceptions,
				pvzs:            nil,
			},
			result: result{
				pvzsWithReceptions: testPVZsWithReceptions,
				err:                model.ErrPVZNotFound,
			},
		},
		{
			name:  "wrong role",
			mocks: func(tc testCase, deps *serviceTestDeps) {},
			args: args{
				userRole: model.CLIENT,

				page:  &testPage,
				limit: &testLimit,
				from:  &testFrom,
				to:    &testTo,
			},
			result: result{
				pvzsWithReceptions: nil,
				err:                model.ErrWrongRole,
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			deps := getTestDeps(t)

			tc.mocks(tc, deps)

			pvzsWithReceptions, err := deps.Service.GetPVZIntervalList(deps.ctx, tc.args.userRole, tc.args.page, tc.args.limit, tc.args.from, tc.args.to)

			require.ErrorIs(t, err, tc.result.err)
			if err == nil {
				require.Equal(t, tc.result.pvzsWithReceptions, pvzsWithReceptions)

			}
		})
	}
}

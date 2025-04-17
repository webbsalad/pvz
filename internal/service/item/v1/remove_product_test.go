package v1

import (
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/webbsalad/pvz/internal/model"
)

func TestService_RemoveProduct(t *testing.T) {
	t.Parallel()

	type args struct {
		userRole        model.Role
		receptionFilter model.ReceptionFilter
		reception       model.Reception
		pvzID           model.PVZID
	}

	type result struct {
		err error
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
					RemoveProduct(gomock.Any(), tc.args.reception.ID).
					Return(nil)
			},
			args: args{
				userRole:        model.EMPLOYEE,
				receptionFilter: testReceptionFilter,
				reception:       testInProgressReception,
				pvzID:           testPVZID,
			},
			result: result{
				err: nil,
			},
		},
		{
			name:  "wrong role moderator",
			mocks: func(tc testCase, deps *serviceTestDeps) {},
			args: args{
				userRole: model.MODERATOR,
				pvzID:    testPVZID,
			},
			result: result{
				err: model.ErrWrongRole,
			},
		},
		{
			name:  "wrong role client",
			mocks: func(tc testCase, deps *serviceTestDeps) {},
			args: args{
				userRole: model.CLIENT,
				pvzID:    testPVZID,
			},
			result: result{
				err: model.ErrWrongRole,
			},
		},
		{
			name: "reception not found",
			mocks: func(tc testCase, deps *serviceTestDeps) {
				deps.itemRepository.EXPECT().
					GetReceptionsByParams(gomock.Any(), tc.args.receptionFilter).
					Return([]model.Reception{}, model.ErrReceptionNotFound)
			},
			args: args{
				userRole:        model.EMPLOYEE,
				receptionFilter: testReceptionFilter,
				pvzID:           testPVZID,
			},
			result: result{
				err: model.ErrReceptionNotFound,
			},
		},
		{
			name: "product not found",
			mocks: func(tc testCase, deps *serviceTestDeps) {
				deps.itemRepository.EXPECT().
					GetReceptionsByParams(gomock.Any(), tc.args.receptionFilter).
					Return([]model.Reception{tc.args.reception}, nil)

				deps.itemRepository.EXPECT().
					RemoveProduct(gomock.Any(), tc.args.reception.ID).
					Return(model.ErrProductNotFound)
			},
			args: args{
				userRole:        model.EMPLOYEE,
				receptionFilter: testReceptionFilter,
				reception:       testInProgressReception,
				pvzID:           testPVZID,
			},
			result: result{
				err: model.ErrProductNotFound,
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			deps := getTestDeps(t)

			tc.mocks(tc, deps)

			err := deps.Service.RemoveProduct(deps.ctx, tc.args.userRole, tc.args.pvzID)

			require.ErrorIs(t, err, tc.result.err)

		})
	}
}

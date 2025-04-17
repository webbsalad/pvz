package v1

import (
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/webbsalad/pvz/internal/model"
)

func TestService_CloseReception(t *testing.T) {
	t.Parallel()

	type args struct {
		userRole            model.Role
		receptionFilter     model.ReceptionFilter
		inProgressReception model.Reception
		closeReception      model.Reception
		pvzID               model.PVZID
	}

	type result struct {
		reception model.Reception
		err       error
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
					Return([]model.Reception{tc.args.inProgressReception}, nil)

				deps.itemRepository.EXPECT().
					UpdateReception(gomock.Any(), tc.args.closeReception).
					Return(tc.args.closeReception, nil)
			},
			args: args{
				userRole:            model.EMPLOYEE,
				receptionFilter:     testReceptionFilter,
				inProgressReception: testInProgressReception,
				closeReception:      testCloseReception,
				pvzID:               testPVZID,
			},
			result: result{
				reception: testCloseReception,
				err:       nil,
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
			name: "receptions not found on get",
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
			name: "receptions not found on update",
			mocks: func(tc testCase, deps *serviceTestDeps) {
				deps.itemRepository.EXPECT().
					GetReceptionsByParams(gomock.Any(), tc.args.receptionFilter).
					Return([]model.Reception{tc.args.inProgressReception}, nil)

				deps.itemRepository.EXPECT().
					UpdateReception(gomock.Any(), tc.args.closeReception).
					Return(model.Reception{}, model.ErrReceptionNotFound)
			},
			args: args{
				userRole:            model.EMPLOYEE,
				receptionFilter:     testReceptionFilter,
				inProgressReception: testInProgressReception,
				closeReception:      testCloseReception,
				pvzID:               testPVZID,
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

			reception, err := deps.Service.CloseReception(deps.ctx, tc.args.userRole, tc.args.pvzID)

			require.ErrorIs(t, err, tc.result.err)
			if err == nil {
				require.Equal(t, tc.result.reception, reception)
			}
		})
	}
}

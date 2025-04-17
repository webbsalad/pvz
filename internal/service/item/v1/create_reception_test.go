package v1

import (
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/webbsalad/pvz/internal/model"
)

func TestService_CreateReception(t *testing.T) {
	t.Parallel()

	type args struct {
		userRole        model.Role
		receptionFilter model.ReceptionFilter
		reception       model.Reception
		pvzID           model.PVZID
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
					Return([]model.Reception{}, model.ErrReceptionNotFound)

				deps.itemRepository.EXPECT().
					CreateReception(gomock.Any(), tc.args.pvzID).
					Return(tc.args.reception, nil)
			},
			args: args{
				userRole:        model.EMPLOYEE,
				receptionFilter: testReceptionFilter,
				pvzID:           testPVZID,
				reception:       testInProgressReception,
			},
			result: result{
				reception: testInProgressReception,
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
			name: "exist in progress reception",
			mocks: func(tc testCase, deps *serviceTestDeps) {
				deps.itemRepository.EXPECT().
					GetReceptionsByParams(gomock.Any(), tc.args.receptionFilter).
					Return([]model.Reception{tc.args.reception}, nil)

			},
			args: args{
				userRole:        model.EMPLOYEE,
				receptionFilter: testReceptionFilter,
				reception:       testInProgressReception,
				pvzID:           testPVZID,
			},
			result: result{
				err: model.ErrReceptionAlreadyExist,
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			deps := getTestDeps(t)

			tc.mocks(tc, deps)

			reception, err := deps.Service.CreateReception(deps.ctx, tc.args.userRole, tc.args.pvzID)

			require.ErrorIs(t, err, tc.result.err)
			if err == nil {
				require.Equal(t, tc.result.reception, reception)
			}
		})
	}
}

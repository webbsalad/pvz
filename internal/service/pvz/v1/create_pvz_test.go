package v1

import (
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/webbsalad/pvz/internal/model"
)

func TestService_CreatePVZ(t *testing.T) {
	t.Parallel()

	type args struct {
		userRole   model.Role
		pvz        model.PVZ
		createdPVZ model.PVZ
	}

	type result struct {
		pvz model.PVZ
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
				deps.pvzRepository.EXPECT().
					CreatePVZ(gomock.Any(), tc.args.pvz).
					Return(tc.args.createdPVZ, nil)
			},
			args: args{
				userRole: model.MODERATOR,
				pvz: model.PVZ{
					ID:   model.PVZID(testPVZID),
					City: testCity,
				},
				createdPVZ: model.PVZ{
					ID:   model.PVZID(testPVZID),
					City: testCity,
				},
			},
			result: result{
				pvz: model.PVZ{
					ID:   model.PVZID(testPVZID),
					City: testCity,
				},
				err: nil,
			},
		},
		{
			name:  "wrong client role",
			mocks: func(tc testCase, deps *serviceTestDeps) {},
			args: args{
				userRole: model.CLIENT,
				pvz: model.PVZ{
					ID:   model.PVZID(testPVZID),
					City: testCity,
				},
			},
			result: result{
				pvz: model.PVZ{},
				err: model.ErrWrongRole,
			},
		},

		{
			name:  "wrong employee role",
			mocks: func(tc testCase, deps *serviceTestDeps) {},
			args: args{
				userRole: model.EMPLOYEE,
				pvz: model.PVZ{
					ID:   model.PVZID(testPVZID),
					City: testCity,
				},
			},
			result: result{
				pvz: model.PVZ{},
				err: model.ErrWrongRole,
			},
		},
		{
			name: "pvz already exist",
			mocks: func(tc testCase, deps *serviceTestDeps) {
				deps.pvzRepository.EXPECT().
					CreatePVZ(gomock.Any(), tc.args.pvz).
					Return(tc.args.createdPVZ, model.ErrPVZAlreadyExist)
			},
			args: args{
				userRole: model.MODERATOR,
				pvz: model.PVZ{
					ID:   model.PVZID(testPVZID),
					City: testCity,
				},
				createdPVZ: model.PVZ{},
			},
			result: result{
				pvz: model.PVZ{},
				err: model.ErrPVZAlreadyExist,
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			deps := getTestDeps(t)

			tc.mocks(tc, deps)

			pvz, err := deps.Service.CreatePVZ(deps.ctx, tc.args.userRole, tc.args.pvz)

			require.ErrorIs(t, err, tc.result.err)
			if err == nil {

				require.Equal(t, tc.result.pvz, pvz)
			}
		})
	}
}

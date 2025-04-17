package v1

import (
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/webbsalad/pvz/internal/model"
)

func TestService_GetPVZList(t *testing.T) {
	t.Parallel()

	type args struct {
		pvzs []model.PVZ
	}

	type result struct {
		pvzs []model.PVZ
		err  error
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
					GetPVZsByParams(gomock.Any(), model.PVZFilter{}).
					Return(tc.args.pvzs, nil)
			},
			args: args{
				pvzs: testPVZs,
			},
			result: result{
				pvzs: testPVZs,
				err:  nil,
			},
		},
		{
			name: "pvzs not found",
			mocks: func(tc testCase, deps *serviceTestDeps) {
				deps.pvzRepository.EXPECT().
					GetPVZsByParams(gomock.Any(), model.PVZFilter{}).
					Return(tc.args.pvzs, model.ErrPVZNotFound)
			},
			args: args{
				pvzs: nil,
			},
			result: result{
				pvzs: nil,
				err:  model.ErrNotFound,
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			deps := getTestDeps(t)

			tc.mocks(tc, deps)

			pvzs, err := deps.Service.GetPVZList(deps.ctx)

			require.ErrorIs(t, err, tc.result.err)
			if err == nil {

				require.Equal(t, tc.result.pvzs, pvzs)
			}
		})
	}
}

package v1

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/webbsalad/pvz/internal/model"
	"github.com/webbsalad/pvz/internal/utils/jwt"
)

func TestService_DummyLogin(t *testing.T) {
	t.Parallel()

	type args struct {
		role model.Role
	}

	type result struct {
		token string
		err   error
	}

	type testCase struct {
		name   string
		args   args
		mocks  func(tc testCase, deps *serviceTestDeps)
		result result
	}

	testCases := []testCase{
		{
			name: "success client role",
			args: args{
				role: model.CLIENT,
			},
			mocks: func(tc testCase, deps *serviceTestDeps) {},
			result: result{
				token: testClientToken,
				err:   nil,
			},
		},
		{
			name: "success moderator role",
			args: args{
				role: model.MODERATOR,
			},
			mocks: func(tc testCase, deps *serviceTestDeps) {},
			result: result{
				token: testModeratorToken,
				err:   nil,
			},
		},
		{
			name: "success employee role",
			args: args{
				role: model.EMPLOYEE,
			},
			mocks: func(tc testCase, deps *serviceTestDeps) {},
			result: result{
				token: testEmployeeToken,
				err:   nil,
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			deps := getTestDeps(t)
			tc.mocks(tc, deps)

			token, err := deps.Service.DummyLogin(deps.ctx, tc.args.role)

			require.ErrorIs(t, err, tc.result.err)
			if err == nil {
				role, _ := jwt.ExtractClaimsFromToken(token, deps.Service.config.JWTSecret)
				resRole, _ := jwt.ExtractClaimsFromToken(tc.result.token, deps.Service.config.JWTSecret)

				require.Equal(t, resRole, role)
			}
		})
	}
}

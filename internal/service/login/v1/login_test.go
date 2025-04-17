package v1

import (
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/webbsalad/pvz/internal/model"
	"github.com/webbsalad/pvz/internal/utils/jwt"
)

func TestService_Login(t *testing.T) {
	t.Parallel()

	type args struct {
		email    string
		password string
		passhash string
		userID   model.UserID
		user     model.User
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
			name: "success",
			mocks: func(tc testCase, deps *serviceTestDeps) {
				deps.userRepository.EXPECT().
					GetUserID(gomock.Any(), tc.args.email).
					Return(tc.args.userID, nil)

				deps.userRepository.EXPECT().
					GetPassHash(gomock.Any(), tc.args.userID).
					Return(tc.args.passhash, nil)

				deps.userRepository.EXPECT().
					GetUser(gomock.Any(), tc.args.userID).
					Return(tc.args.user, nil)
			},
			args: args{
				email:    testEmail,
				password: testPassword,
				passhash: testHash,
				userID:   testUserID,
				user: model.User{
					ID:    testUserID,
					Email: testEmail,
					Role:  model.CLIENT,
				},
			},
			result: result{
				token: testClientToken,
				err:   nil,
			},
		},
		{
			name: "wrong password",
			mocks: func(tc testCase, deps *serviceTestDeps) {
				deps.userRepository.EXPECT().
					GetUserID(gomock.Any(), tc.args.email).
					Return(tc.args.userID, nil)

				deps.userRepository.EXPECT().
					GetPassHash(gomock.Any(), tc.args.userID).
					Return(tc.args.passhash, nil)
			},
			args: args{
				email:    testEmail,
				password: "wrong password",
				passhash: testHash,
				userID:   testUserID,
			},
			result: result{
				err: model.ErrWrongPassword,
			},
		},
		{
			name: "user not found",
			mocks: func(tc testCase, deps *serviceTestDeps) {
				deps.userRepository.EXPECT().
					GetUserID(gomock.Any(), tc.args.email).
					Return(tc.args.userID, model.ErrUserNotFound)

			},
			args: args{
				email:    testEmail,
				password: testPassword,
				userID:   model.UserID{},
			},
			result: result{
				err: model.ErrUserNotFound,
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			deps := getTestDeps(t)

			tc.mocks(tc, deps)

			token, err := deps.Service.Login(deps.ctx, tc.args.email, tc.args.password)

			require.ErrorIs(t, err, tc.result.err)
			if err == nil {
				role, _ := jwt.ExtractClaimsFromToken(token, deps.Service.config.JWTSecret)
				resRole, _ := jwt.ExtractClaimsFromToken(tc.result.token, deps.Service.config.JWTSecret)

				require.Equal(t, resRole, role)
			}
		})
	}
}

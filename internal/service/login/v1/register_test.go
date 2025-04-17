package v1

import (
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/webbsalad/pvz/internal/model"
)

func TestService_Register(t *testing.T) {
	t.Parallel()

	type args struct {
		user       model.User
		storedUser model.User
		password   string
	}

	type result struct {
		user model.User
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
				deps.userRepository.EXPECT().
					CreateUser(gomock.Any(), tc.args.user, hashMatcher{password: tc.args.password}).
					Return(tc.args.storedUser, nil)
			},
			args: args{
				storedUser: model.User{
					ID:    testUserID,
					Email: testEmail,
					Role:  model.CLIENT,
				},
				user: model.User{
					Email: testEmail,
					Role:  model.CLIENT,
				},
				password: testPassword,
			},
			result: result{
				user: model.User{
					ID:    testUserID,
					Email: testEmail,
					Role:  model.CLIENT,
				},
				err: nil,
			},
		},
		{
			name: "already exist",
			mocks: func(tc testCase, deps *serviceTestDeps) {
				deps.userRepository.EXPECT().
					CreateUser(gomock.Any(), tc.args.user, hashMatcher{password: tc.args.password}).
					Return(tc.args.storedUser, model.ErrUserAlreadyExist)
			},
			args: args{
				storedUser: model.User{},
				user: model.User{
					Email: testEmail,
					Role:  model.CLIENT,
				},
				password: testPassword,
			},
			result: result{
				err: model.ErrUserAlreadyExist,
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			deps := getTestDeps(t)

			tc.mocks(tc, deps)

			user, err := deps.Service.Register(deps.ctx, tc.args.user, tc.args.password)

			require.ErrorIs(t, err, tc.result.err)
			if err == nil {
				require.Equal(t, tc.result.user, user)
			}
		})
	}
}

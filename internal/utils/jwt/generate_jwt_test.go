package jwt

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/webbsalad/pvz/internal/model"
)

func Test_GenerateToken(t *testing.T) {
	t.Parallel()

	type args struct {
		secret string
		role   model.Role
	}

	type result struct {
		token string
		err   error
	}

	type testCase struct {
		name   string
		args   args
		result result
	}

	testCases := []testCase{
		{
			name: "success moderator",
			args: args{
				role:   model.MODERATOR,
				secret: testJWT_SECRET,
			},
			result: result{
				token: testModeratorToken,
				err:   nil,
			},
		},
		{
			name: "success employee",
			args: args{
				role:   model.EMPLOYEE,
				secret: testJWT_SECRET,
			},
			result: result{
				token: testEmployeeToken,
				err:   nil,
			},
		},
		{
			name: "success client",
			args: args{
				role:   model.CLIENT,
				secret: testJWT_SECRET,
			},
			result: result{
				token: testClientToken,
				err:   nil,
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			token, err := GenerateTokens(tc.args.role, tc.args.secret)

			require.ErrorIs(t, err, tc.result.err)
			if err == nil {
				role, _ := ExtractClaimsFromToken(token, testJWT_SECRET)
				resRole, _ := ExtractClaimsFromToken(tc.result.token, testJWT_SECRET)

				require.Equal(t, resRole, role)
			}

		})
	}
}

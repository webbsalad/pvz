package jwt

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/webbsalad/pvz/internal/model"
)

var (
	testJWT_SECRET     = "test secret"
	testClientToken    = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjMzMjE2OTk1ODEsInJvbGUiOiJjbGllbnQifQ.87XmEZIK8gySrCGYIrN1-1Ub4naQp5VrHncjlPArMkc"
	testModeratorToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjMzMjE2OTk2NjYsInJvbGUiOiJtb2RlcmF0b3IifQ.SAE1-ybXc_3ueXDDKkcA2Vz6Cgpp8gLjoYZuEwg4vzs"
	testEmployeeToken  = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjMzMjE2OTk3MDAsInJvbGUiOiJlbXBsb3llZSJ9.0mJKAurls2Mwb81E5lMMO7eqnfPSlFE73XVnhl3Rfmk"
	testExpiredToken   = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NDQyMDgzNTMsInJvbGUiOiJtb2RlcmF0b3IifQ.Tc8jsb2L8MCyqfE8-M66p2ussC8vPuEB68R1T4VjArk"
)

func Test_EctractClaims(t *testing.T) {
	t.Parallel()

	type args struct {
		token  string
		secret string
	}

	type result struct {
		role model.Role
		err  error
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
				token:  testModeratorToken,
				secret: testJWT_SECRET,
			},
			result: result{
				role: model.MODERATOR,
				err:  nil,
			},
		},
		{
			name: "success client",
			args: args{
				token:  testClientToken,
				secret: testJWT_SECRET,
			},
			result: result{
				role: model.CLIENT,
				err:  nil,
			},
		},
		{
			name: "success employee",
			args: args{
				token:  testEmployeeToken,
				secret: testJWT_SECRET,
			},
			result: result{
				role: model.EMPLOYEE,
				err:  nil,
			},
		},
		{
			name: "expired jwt",
			args: args{
				token:  testExpiredToken,
				secret: testJWT_SECRET,
			},
			result: result{
				err: model.ErrJwtExpired,
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			role, err := ExtractClaimsFromToken(tc.args.token, tc.args.secret)

			require.ErrorIs(t, err, tc.result.err)
			if err == nil {
				require.Equal(t, tc.result.role, role)
			}

		})
	}
}

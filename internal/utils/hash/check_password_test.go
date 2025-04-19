package hash

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/webbsalad/pvz/internal/model"
)

var (
	testPassword = "qwe123"
	testPassHash = "$2a$10$Lwcl8...xg/4aBecP2MAXuFTRcCbbXrohr0GphbVrWpWuWOpynEWC"
)

func Test_CheckPassword(t *testing.T) {
	t.Parallel()

	type args struct {
		password string
		passHash string
	}

	type result struct {
		err error
	}

	type testCase struct {
		name   string
		args   args
		result result
	}

	testCases := []testCase{
		{
			name: "success",
			args: args{
				password: testPassword,
				passHash: testPassHash,
			},
			result: result{
				err: nil,
			},
		},
		{
			name: "wrong password",
			args: args{
				password: testPassword,
				passHash: "wrong",
			},
			result: result{
				err: model.ErrPasswordMismatch,
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			err := CheckPassword(tc.args.passHash, tc.args.password)

			require.ErrorIs(t, err, tc.result.err)

		})
	}
}

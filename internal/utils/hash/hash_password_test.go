package hash

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_HashPassword(t *testing.T) {
	t.Parallel()

	type args struct {
		password string
	}

	type result struct {
		passHash string
		err      error
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
			},
			result: result{
				passHash: testPassHash,
				err:      nil,
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			passHash, err := HashPassword(tc.args.password)

			require.ErrorIs(t, err, tc.result.err)
			if err == nil {
				require.Equal(t, CheckPassword(passHash, tc.args.password), CheckPassword(tc.result.passHash, tc.args.password)) // idk how do it better...
			}

		})
	}
}

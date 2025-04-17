package v1

import (
	"context"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/webbsalad/pvz/internal/config"
	"github.com/webbsalad/pvz/internal/model"
	mock_user_repository "github.com/webbsalad/pvz/internal/repository/user/mock"
	"github.com/webbsalad/pvz/internal/utils/hash"
)

var (
	testUserID, _ = model.NewUserID("2b98ee88-7970-4e6f-b325-ccf3ce10909f")

	testEmail = "example.txt"

	testPassword = "qwe123"
	testHash     = "$2a$10$5yBPKaEe1uqUqauwCiK8ROHrIBmshQT2Wif45wYwQ2rQ0QYXCC1gS"

	testClientToken    = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjMzMjE2OTk1ODEsInJvbGUiOiJjbGllbnQifQ.87XmEZIK8gySrCGYIrN1-1Ub4naQp5VrHncjlPArMkc"
	testModeratorToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjMzMjE2OTk2NjYsInJvbGUiOiJtb2RlcmF0b3IifQ.SAE1-ybXc_3ueXDDKkcA2Vz6Cgpp8gLjoYZuEwg4vzs"
	testEmployeeToken  = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjMzMjE2OTk3MDAsInJvbGUiOiJlbXBsb3llZSJ9.0mJKAurls2Mwb81E5lMMO7eqnfPSlFE73XVnhl3Rfmk"
)

type serviceTestDeps struct {
	Service *Service

	ctx            context.Context
	userRepository *mock_user_repository.MockRepository
}

func getTestDeps(t *testing.T) *serviceTestDeps {
	ctrl := gomock.NewController(t)
	userRepository := mock_user_repository.NewMockRepository(ctrl)

	return &serviceTestDeps{
		Service: &Service{
			userRepository: userRepository,
			config: config.Config{
				JWTSecret: "test secret",
			},
		},

		ctx:            context.Background(),
		userRepository: userRepository,
	}
}

type hashMatcher struct {
	password string
}

func (hm hashMatcher) Matches(x interface{}) bool {
	hashedPassword, ok := x.(string)
	if !ok {
		return false
	}

	err := hash.CheckPassword(hashedPassword, hm.password)
	return err == nil
}

func (hm hashMatcher) String() string {
	return "matches hashed password"
}

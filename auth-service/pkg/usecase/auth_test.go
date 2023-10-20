package usecase

import (
	"auth-service/pkg/domain"
	"auth-service/pkg/mock"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestSignUp(t *testing.T) {

	testCases := map[string]struct {
		input     domain.User
		buildStub func(authRepo *mock.MockAuthRepo,
			tokenAuth *mock.MockTokenAuth, input domain.User)
		expectedOutput domain.User
		checkSameError bool // to notify is need to check the actual error and expected error are same
		expectedError  error
	}{
		"failed_to_check_user_already_exist_on_db_should_return_error": {
			input: domain.User{
				Email:    "user@gmail.com",
				Password: "password",
			},
			buildStub: func(authRepo *mock.MockAuthRepo, tokenAuth *mock.MockTokenAuth,
				input domain.User) {

				dbErr := errors.New("error from db")
				// expecting a call to auth repo is user already exist
				authRepo.EXPECT().IsUserExist("user@gmail.com").
					Times(1).Return(false, dbErr)
			},
			expectedOutput: domain.User{},
			checkSameError: false,
			expectedError:  errors.New("expecting db error"),
		},
		"user_already_exist_should_return_error": {
			input: domain.User{
				Email:    "already_exist_user@gmail.com",
				Password: "password",
			},
			buildStub: func(authRepo *mock.MockAuthRepo, tokenAuth *mock.MockTokenAuth,
				input domain.User) {

				// expecting call to auth repo for checking user already exist
				authRepo.EXPECT().IsUserExist("already_exist_user@gmail.com").
					Times(1).Return(true, nil) // return true for auth repo call
			},
			expectedOutput: domain.User{},
			checkSameError: true,
			expectedError:  ErrAlreadyExist,
		},
		"failed_save_user_on_db_should_return_error": {
			input: domain.User{
				Email:    "new_user@gmail.com",
				Password: "password",
			},
			buildStub: func(authRepo *mock.MockAuthRepo, tokenAuth *mock.MockTokenAuth,
				input domain.User) {

				// expecting call to auth repo for checking user already exist
				authRepo.EXPECT().IsUserExist("new_user@gmail.com").
					Times(1).Return(false, nil) // return false for use not exist

				dbErr := errors.New("failed to save user on db")
				// expecting a call to auth repo for save user on db
				authRepo.EXPECT().SaveUser(gomock.Any()).Return(input, dbErr)
			},
			expectedOutput: domain.User{},
			checkSameError: false,
			expectedError:  errors.New("some error"),
		},
		"successful_user_sign_up": {
			input: domain.User{
				Email:    "new_user@gmail.com",
				Password: "password",
			},
			buildStub: func(authRepo *mock.MockAuthRepo, tokenAuth *mock.MockTokenAuth,
				input domain.User) {
				// expecting call to auth repo for checking user already exist
				authRepo.EXPECT().IsUserExist("new_user@gmail.com").
					Times(1).Return(false, nil) // return false for use not exist

				// expecting a call to auth repo for save user on db
				authRepo.EXPECT().SaveUser(gomock.Any()).Return(input, nil)
			},
			expectedOutput: domain.User{
				Email:    "new_user@gmail.com",
				Password: "password",
			},
		},
	}

	for name, test := range testCases {

		test := test
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			// create a gomock controller with testing object
			ctl := gomock.NewController(t)
			// create a mock repository for the usecase
			mockAuthRepo := mock.NewMockAuthRepo(ctl)
			// create a mock token auth
			mockTokenAuth := mock.NewMockTokenAuth(ctl)

			// pass the mock auth repo and token auth to setup repo before using in usecase
			test.buildStub(mockAuthRepo, mockTokenAuth, test.input)

			// create auth usecase with the mock deps
			authUseCase := NewAuthUseCase(mockAuthRepo, mockTokenAuth)

			// run the sign up with test input
			actualOutput, actualErr := authUseCase.SignUp(test.input)

			// check the actual output and expected output is equal
			assert.Equal(t, test.expectedOutput, actualOutput)

			// if the test case expect and output error should be same or no error.
			if test.checkSameError || test.expectedError == nil {
				// then error should be same even if it's nil on both side.
				assert.Equal(t, test.expectedError, actualErr)
			} else {
				// otherwise just confirm it's an error
				assert.Error(t, actualErr)
			}
		})
	}
}

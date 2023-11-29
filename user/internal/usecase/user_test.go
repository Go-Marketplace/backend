package usecase_test

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/Go-Marketplace/backend/pkg/logger"
	mocks "github.com/Go-Marketplace/backend/user/internal/mocks/repo"
	"github.com/Go-Marketplace/backend/user/internal/model"
	"github.com/Go-Marketplace/backend/user/internal/usecase"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func userHelper(t *testing.T) (*usecase.UserUsecase, *mocks.MockUserRepo) {
	t.Helper()

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	logger := logger.New("debug")

	repo := mocks.NewMockUserRepo(mockCtrl)
	userUsecase := usecase.NewUserUsecase(repo, logger)

	return userUsecase, repo
}

func TestGetUser(t *testing.T) {
	t.Parallel()

	type args struct {
		ctx    context.Context
		userID uuid.UUID
	}

	ctx := context.Background()

	userID, _ := uuid.Parse("3e7cbbcf-ebe3-4808-b895-fe5eaf41be22")

	expectedUserFromRepo := &model.User{
		ID:        userID,
		FirstName: "test",
		LastName:  "test",
		Address:   "test",
		Role:      1,
	}
	expectedErrFromRepo := errors.New("test error")

	testcases := []struct {
		name         string
		args         args
		mock         func(repo *mocks.MockUserRepo)
		expectedUser *model.User
		expectedErr  error
	}{
		{
			name: "Successfully get user",
			args: args{
				ctx:    ctx,
				userID: userID,
			},
			mock: func(repo *mocks.MockUserRepo) {
				repo.EXPECT().GetUser(ctx, userID).Return(expectedUserFromRepo, nil).Times(1)
			},
			expectedUser: expectedUserFromRepo,
			expectedErr:  nil,
		},
		{
			name: "Got error when get user",
			args: args{
				ctx:    ctx,
				userID: userID,
			},
			mock: func(repo *mocks.MockUserRepo) {
				repo.EXPECT().GetUser(ctx, userID).Return(nil, expectedErrFromRepo).Times(1)
			},
			expectedUser: nil,
			expectedErr:  expectedErrFromRepo,
		},
	}

	for _, testcase := range testcases {
		testcase := testcase

		t.Run(testcase.name, func(t *testing.T) {
			t.Parallel()

			userUsecase, userRepo := userHelper(t)
			testcase.mock(userRepo)

			actualUser, actualErr := userUsecase.GetUser(
				testcase.args.ctx,
				testcase.args.userID,
			)

			assert.Equal(t, testcase.expectedUser, actualUser)
			assert.Equal(t, testcase.expectedErr, actualErr)
		})
	}
}

func TestGetAllUsers(t *testing.T) {
	t.Parallel()

	type args struct {
		ctx context.Context
	}

	ctx := context.Background()

	expectedUsersFromRepo := []*model.User{
		{
			ID:        uuid.New(),
			FirstName: "test",
			LastName:  "test",
			Address:   "test",
			Role:      1,
		},
	}
	expectedErrFromRepo := errors.New("test error")

	testcases := []struct {
		name          string
		args          args
		mock          func(repo *mocks.MockUserRepo)
		expectedUsers []*model.User
		expectedErr   error
	}{
		{
			name: "Successfully get all users",
			args: args{
				ctx: ctx,
			},
			mock: func(repo *mocks.MockUserRepo) {
				repo.EXPECT().GetAllUsers(ctx).Return(expectedUsersFromRepo, nil).Times(1)
			},
			expectedUsers: expectedUsersFromRepo,
			expectedErr:   nil,
		},
		{
			name: "Got error when get all users",
			args: args{
				ctx: ctx,
			},
			mock: func(repo *mocks.MockUserRepo) {
				repo.EXPECT().GetAllUsers(ctx).Return(nil, expectedErrFromRepo).Times(1)
			},
			expectedUsers: nil,
			expectedErr:   expectedErrFromRepo,
		},
	}

	for _, testcase := range testcases {
		testcase := testcase

		t.Run(testcase.name, func(t *testing.T) {
			t.Parallel()

			userUsecase, userRepo := userHelper(t)
			testcase.mock(userRepo)

			actualUsers, actualErr := userUsecase.GetAllUsers(
				testcase.args.ctx,
			)

			assert.Equal(t, testcase.expectedUsers, actualUsers)
			assert.Equal(t, testcase.expectedErr, actualErr)
		})
	}
}

func TestGetUserByEmail(t *testing.T) {
	t.Parallel()

	type args struct {
		ctx   context.Context
		email string
	}

	ctx := context.Background()
	mail := "test@test.ru"

	expectedUserFromRepo := &model.User{
		ID:        uuid.New(),
		FirstName: "test",
		LastName:  "test",
		Address:   "test",
		Role:      1,
	}
	expectedErrFromRepo := errors.New("test error")

	testcases := []struct {
		name         string
		args         args
		mock         func(repo *mocks.MockUserRepo)
		expectedUser *model.User
		expectedErr  error
	}{
		{
			name: "Successfully get user by email",
			args: args{
				ctx:   ctx,
				email: mail,
			},
			mock: func(repo *mocks.MockUserRepo) {
				repo.EXPECT().GetUserByEmail(ctx, mail).Return(expectedUserFromRepo, nil).Times(1)
			},
			expectedUser: expectedUserFromRepo,
			expectedErr:  nil,
		},
		{
			name: "Got error when get user by email",
			args: args{
				ctx:   ctx,
				email: mail,
			},
			mock: func(repo *mocks.MockUserRepo) {
				repo.EXPECT().GetUserByEmail(ctx, mail).Return(nil, expectedErrFromRepo).Times(1)
			},
			expectedUser: nil,
			expectedErr:  expectedErrFromRepo,
		},
	}

	for _, testcase := range testcases {
		testcase := testcase

		t.Run(testcase.name, func(t *testing.T) {
			t.Parallel()

			userUsecase, userRepo := userHelper(t)
			testcase.mock(userRepo)

			actualUser, actualErr := userUsecase.GetUserByEmail(
				testcase.args.ctx,
				testcase.args.email,
			)

			assert.Equal(t, testcase.expectedUser, actualUser)
			assert.Equal(t, testcase.expectedErr, actualErr)
		})
	}
}

func TestCreateUser(t *testing.T) {
	t.Parallel()

	type args struct {
		ctx  context.Context
		user model.User
	}

	ctx := context.Background()
	userID := uuid.New()
	testUser := model.User{
		ID:    userID,
		Email: "test@test.ru",
		Role:  1,
	}

	expectedErrFromRepo := errors.New("test error")

	testcases := []struct {
		name        string
		args        args
		mock        func(repo *mocks.MockUserRepo)
		expectedErr error
	}{
		{
			name: "Successfully create user",
			args: args{
				ctx:  ctx,
				user: testUser,
			},
			mock: func(repo *mocks.MockUserRepo) {
				repo.EXPECT().CreateUser(ctx, testUser).Return(nil).Times(1)
			},
			expectedErr: nil,
		},
		{
			name: "Got error when create user",
			args: args{
				ctx:  ctx,
				user: testUser,
			},
			mock: func(repo *mocks.MockUserRepo) {
				repo.EXPECT().CreateUser(ctx, testUser).Return(expectedErrFromRepo).Times(1)
			},
			expectedErr: expectedErrFromRepo,
		},
	}

	for _, testcase := range testcases {
		testcase := testcase

		t.Run(testcase.name, func(t *testing.T) {
			t.Parallel()

			userUsecase, userRepo := userHelper(t)
			testcase.mock(userRepo)

			actualErr := userUsecase.CreateUser(
				testcase.args.ctx,
				testcase.args.user,
			)

			assert.Equal(t, testcase.expectedErr, actualErr)
		})
	}
}

func TestDeleteUser(t *testing.T) {
	t.Parallel()

	type args struct {
		ctx    context.Context
		userID uuid.UUID
	}

	ctx := context.Background()
	userID := uuid.New()

	expectedErrFromRepo := errors.New("test error")

	testcases := []struct {
		name        string
		args        args
		mock        func(repo *mocks.MockUserRepo)
		expectedErr error
	}{
		{
			name: "Successfully delete user",
			args: args{
				ctx:    ctx,
				userID: userID,
			},
			mock: func(repo *mocks.MockUserRepo) {
				repo.EXPECT().DeleteUser(ctx, userID).Return(nil).Times(1)
			},
			expectedErr: nil,
		},
		{
			name: "Got error when delete user",
			args: args{
				ctx:    ctx,
				userID: userID,
			},
			mock: func(repo *mocks.MockUserRepo) {
				repo.EXPECT().DeleteUser(ctx, userID).Return(expectedErrFromRepo).Times(1)
			},
			expectedErr: expectedErrFromRepo,
		},
	}

	for _, testcase := range testcases {
		testcase := testcase

		t.Run(testcase.name, func(t *testing.T) {
			t.Parallel()

			userUsecase, userRepo := userHelper(t)
			testcase.mock(userRepo)

			actualErr := userUsecase.DeleteUser(
				testcase.args.ctx,
				testcase.args.userID,
			)

			assert.Equal(t, testcase.expectedErr, actualErr)
		})
	}
}

func TestUpdateUser(t *testing.T) {
	t.Parallel()

	type args struct {
		ctx  context.Context
		user model.User
	}

	ctx := context.Background()

	userID, _ := uuid.Parse("3e7cbbcf-ebe3-4808-b895-fe5eaf41be22")
	testUser := model.User{
		ID:      userID,
		Address: "test",
		Role:    1,
	}

	expectedUserFromRepo := &model.User{
		ID:        userID,
		FirstName: "test",
		LastName:  "test",
		Address:   "test",
		Role:      1,
	}
	expectedErrFromRepo := errors.New("test error")

	testcases := []struct {
		name         string
		args         args
		mock         func(repo *mocks.MockUserRepo)
		expectedUser *model.User
		expectedErr  error
	}{
		{
			name: "Successfully update user",
			args: args{
				ctx:  ctx,
				user: testUser,
			},
			mock: func(repo *mocks.MockUserRepo) {
				repo.EXPECT().UpdateUser(ctx, testUser).Return(nil).Times(1)
				repo.EXPECT().GetUser(ctx, userID).Return(expectedUserFromRepo, nil).Times(1)
			},
			expectedUser: expectedUserFromRepo,
			expectedErr:  nil,
		},
		{
			name: "Got error when update user",
			args: args{
				ctx:  ctx,
				user: testUser,
			},
			mock: func(repo *mocks.MockUserRepo) {
				repo.EXPECT().UpdateUser(ctx, testUser).Return(expectedErrFromRepo).Times(1)
			},
			expectedUser: nil,
			expectedErr:  fmt.Errorf("failed to update user: %w", expectedErrFromRepo),
		},
		{
			name: "Got error when get user",
			args: args{
				ctx:  ctx,
				user: testUser,
			},
			mock: func(repo *mocks.MockUserRepo) {
				repo.EXPECT().UpdateUser(ctx, testUser).Return(nil).Times(1)
				repo.EXPECT().GetUser(ctx, userID).Return(nil, expectedErrFromRepo).Times(1)
			},
			expectedUser: nil,
			expectedErr:  expectedErrFromRepo,
		},
	}

	for _, testcase := range testcases {
		testcase := testcase

		t.Run(testcase.name, func(t *testing.T) {
			t.Parallel()

			userUsecase, userRepo := userHelper(t)
			testcase.mock(userRepo)

			actualUser, actualErr := userUsecase.UpdateUser(
				testcase.args.ctx,
				testcase.args.user,
			)

			assert.Equal(t, testcase.expectedUser, actualUser)
			assert.Equal(t, testcase.expectedErr, actualErr)
		})
	}
}

func TestChangeUserRole(t *testing.T) {
	t.Parallel()

	type args struct {
		ctx  context.Context
		user model.User
	}

	ctx := context.Background()

	userID, _ := uuid.Parse("3e7cbbcf-ebe3-4808-b895-fe5eaf41be22")
	testUser := model.User{
		ID:      userID,
		Address: "test",
		Role:    1,
	}

	expectedUserFromRepo := &model.User{
		ID:        userID,
		FirstName: "test",
		LastName:  "test",
		Address:   "test",
		Role:      1,
	}
	expectedErrFromRepo := errors.New("test error")

	testcases := []struct {
		name         string
		args         args
		mock         func(repo *mocks.MockUserRepo)
		expectedUser *model.User
		expectedErr  error
	}{
		{
			name: "Successfully change user role",
			args: args{
				ctx:  ctx,
				user: testUser,
			},
			mock: func(repo *mocks.MockUserRepo) {
				repo.EXPECT().ChangeUserRole(ctx, testUser).Return(nil).Times(1)
				repo.EXPECT().GetUser(ctx, userID).Return(expectedUserFromRepo, nil).Times(1)
			},
			expectedUser: expectedUserFromRepo,
			expectedErr:  nil,
		},
		{
			name: "Got error when change user role",
			args: args{
				ctx:  ctx,
				user: testUser,
			},
			mock: func(repo *mocks.MockUserRepo) {
				repo.EXPECT().ChangeUserRole(ctx, testUser).Return(expectedErrFromRepo).Times(1)
			},
			expectedUser: nil,
			expectedErr:  fmt.Errorf("failed to change user role: %w", expectedErrFromRepo),
		},
		{
			name: "Got error when get user",
			args: args{
				ctx:  ctx,
				user: testUser,
			},
			mock: func(repo *mocks.MockUserRepo) {
				repo.EXPECT().ChangeUserRole(ctx, testUser).Return(nil).Times(1)
				repo.EXPECT().GetUser(ctx, userID).Return(nil, expectedErrFromRepo).Times(1)
			},
			expectedUser: nil,
			expectedErr:  expectedErrFromRepo,
		},
	}

	for _, testcase := range testcases {
		testcase := testcase

		t.Run(testcase.name, func(t *testing.T) {
			t.Parallel()

			userUsecase, userRepo := userHelper(t)
			testcase.mock(userRepo)

			actualUser, actualErr := userUsecase.ChangeUserRole(
				testcase.args.ctx,
				testcase.args.user,
			)

			assert.Equal(t, testcase.expectedUser, actualUser)
			assert.Equal(t, testcase.expectedErr, actualErr)
		})
	}
}

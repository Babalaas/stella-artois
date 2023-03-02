package repository_test

import (
	"babalaas/stella-artois/model"
	"babalaas/stella-artois/repository"
	"context"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type args struct {
	ctx         context.Context
	userProfile *model.UserProfile
}

type test_case struct {
	name    string
	args    args
	wantErr bool
}

func TestUserProfileRepository_Create_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{NamingStrategy: schema.NamingStrategy{SingularTable: true}})
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a gorm database connection", err)
	}

	tc_args := args{
		ctx: context.Background(),
		userProfile: &model.UserProfile{
			DisplayName: "JohnDoe",
			FirstName:   "John",
			LastName:    "Doe",
			Email:       "johndoe@example.com",
			Phone:       "1234567890",
			Gender:      "male",
			Birthdate:   time.Date(1990, 10, 10, 0, 0, 0, 0, time.UTC),
			Password:    "password",
		},
	}

	tc := test_case{
		name:    "User_Profile_Create_Success",
		args:    tc_args,
		wantErr: false,
	}

	repo := repository.NewTestUserProfileRepository(gormDB)

	t.Run(tc.name, func(t *testing.T) {
		ctx, cancel := context.WithTimeout(tc_args.ctx, time.Second)
		defer cancel()

		mock.ExpectBegin()
		mock.ExpectExec(`INSERT INTO "user_profile"`).
			WithArgs(
				sqlmock.AnyArg(),
				tc_args.userProfile.DisplayName,
				tc_args.userProfile.FirstName,
				tc_args.userProfile.LastName,
				tc_args.userProfile.Email,
				tc_args.userProfile.Phone,
				tc_args.userProfile.Gender,
				tc_args.userProfile.Birthdate,
				tc_args.userProfile.Password,
			).
			WillReturnResult(sqlmock.NewResult(1, 1))

		mock.ExpectCommit()

		id, err := repo.Create(ctx, tc.args.userProfile)
		if (err != nil) != tc.wantErr {
			t.Errorf("UserProfileRepository.Create() error = %v, wantErr %v", err, tc.wantErr)
			return
		}

		assert.NotEqualValues(t, uuid.Nil, id)
	})
}

/*
func TestUserProfileRepository_Create(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{NamingStrategy: schema.NamingStrategy{SingularTable: true}})
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a gorm database connection", err)
	}

	type args struct {
		ctx         context.Context
		userProfile *model.UserProfile
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				ctx: context.Background(),
				userProfile: &model.UserProfile{
					DisplayName: "JohnDoe",
					FirstName:   "John",
					LastName:    "Doe",
					Email:       "johndoe@example.com",
					Phone:       "1234567890",
					Gender:      "male",
					Birthdate:   time.Date(1990, 10, 10, 0, 0, 0, 0, time.UTC),
					Password:    "password",
				},
			},
			wantErr: false,
		},
		{
			name: "failure",
			args: args{
				ctx: context.Background(),
				userProfile: &model.UserProfile{
					DisplayName: "JaneDoe",
					FirstName:   "Jane",
					LastName:    "Doe",
					Email:       "janedoe@example.com",
					Phone:       "1234567890",
					Gender:      "female",
					Birthdate:   time.Date(1995, 5, 5, 0, 0, 0, 0, time.UTC),
					Password:    "password",
				},
			},
			wantErr: true,
		},
	}

	repo := repository.NewTestUserProfileRepository(gormDB)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(tt.args.ctx, time.Second)
			defer cancel()

			mock.ExpectBegin()
			mock.ExpectExec(`INSERT INTO "user_profile"`).
				WithArgs(
					sqlmock.AnyArg(),
					tt.args.userProfile.DisplayName,
					tt.args.userProfile.FirstName,
					tt.args.userProfile.LastName,
					tt.args.userProfile.Email,
					tt.args.userProfile.Phone,
					tt.args.userProfile.Gender,
					tt.args.userProfile.Birthdate,
					tt.args.userProfile.Password,
				).
				WillReturnResult(sqlmock.NewResult(1, 1))
			if tt.wantErr {
				mock.ExpectRollback()
			} else {
				mock.ExpectCommit()
			}

			err := repo.Create(ctx, tt.args.userProfile)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserProfileRepository.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
*/

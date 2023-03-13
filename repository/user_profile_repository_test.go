package repository_test

import (
	"babalaas/stella-artois/model"
	"babalaas/stella-artois/repository"
	"context"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func Test_Create(t *testing.T) {
	// Set up test dependencies
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("error setting up test dependencies: %v", err)
	}
	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		DriverName:           "postgres",
		DSN:                  "",
		PreferSimpleProtocol: true,
		WithoutReturning:     false,
		Conn:                 db,
	}), &gorm.Config{
		SkipDefaultTransaction: false,
		NamingStrategy:         &schema.NamingStrategy{TablePrefix: "", SingularTable: true, NameReplacer: nil, NoLowerCase: false},
	})

	if err != nil {
		t.Fatalf("error creating GORM database object: %v", err)
	}

	// Create test repository
	newRepo := repository.NewUserProfileRepository(gormDB)
	// repo := &userProfileRepository{DB: gorm.OpenDB(db)}

	// Create test context and user profile
	ctx := context.Background()
	userProfile := &model.UserProfile{
		DisplayName: "Big John Doe",
		FirstName:   "John",
		LastName:    "Doe",
		Email:       "john.doe@gmail.com",
		Phone:       "8148675309",
		Gender:      "Male",
		Birthdate:   time.Date(2000, 3, 13, 10, 23, 0, 0, time.UTC),
		Password:    "taco",
	}

	expectedID := uuid.New()

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO user_profile ("id","display_name","first_name", "last_name", "email", "phone", "gender", "birthdate", "password") VALUES`)).
		WithArgs(expectedID, userProfile.DisplayName, userProfile.FirstName, userProfile.LastName, userProfile.Email, userProfile.Phone, userProfile.Gender, userProfile.Birthdate, userProfile.Password).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	// Call Create method
	id, err := newRepo.Create(ctx, userProfile)

	// Check that the result is as expected
	if err != nil {
		t.Errorf("unexpected error from Create method: %v", err)
	}
	if id != expectedID {
		t.Errorf("unexpected ID returned from Create method: got %v, want %v", id, expectedID)
	}

	// Check that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unfulfilled expectations: %v", err)
	}
}

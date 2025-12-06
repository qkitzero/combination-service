package category

import (
	"errors"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"go.uber.org/mock/gomock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/qkitzero/combination-service/internal/domain/category"
	mockscategory "github.com/qkitzero/combination-service/mocks/domain/category"
	"github.com/qkitzero/combination-service/testutil"
)

func TestCreate(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		success bool
		setup   func(mock sqlmock.Sqlmock, category category.Category)
	}{
		{
			name:    "success create category",
			success: true,
			setup: func(mock sqlmock.Sqlmock, category category.Category) {
				mock.ExpectBegin()

				mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO "categories" ("id","name","created_at") VALUES ($1,$2,$3)`)).
					WithArgs(category.ID(), category.Name(), testutil.AnyTime{}).
					WillReturnResult(sqlmock.NewResult(1, 1))

				mock.ExpectCommit()
			},
		},
		{
			name:    "failure create category error",
			success: false,
			setup: func(mock sqlmock.Sqlmock, category category.Category) {
				mock.ExpectBegin()

				mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO "categories" ("id","name","created_at") VALUES ($1,$2,$3)`)).
					WithArgs(category.ID(), category.Name(), testutil.AnyTime{}).
					WillReturnError(errors.New("create category error"))

				mock.ExpectRollback()
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			sqlDB, mock, err := sqlmock.New()
			if err != nil {
				t.Errorf("failed to new sqlmock: %s", err)
			}

			gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: sqlDB}), &gorm.Config{})
			if err != nil {
				t.Errorf("failed to open gorm: %s", err)
			}

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockCategory := mockscategory.NewMockCategory(ctrl)
			mockCategory.EXPECT().ID().Return(category.CategoryID{UUID: uuid.New()}).AnyTimes()
			mockCategory.EXPECT().Name().Return(category.Name("test category")).AnyTimes()
			mockCategory.EXPECT().CreatedAt().Return(time.Now()).AnyTimes()

			tt.setup(mock, mockCategory)

			repo := NewCategoryRepository(gormDB)

			err = repo.Create(mockCategory)
			if tt.success && err != nil {
				t.Errorf("expected no error, but got %v", err)
			}
			if !tt.success && err == nil {
				t.Errorf("expected error, but got nil")
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

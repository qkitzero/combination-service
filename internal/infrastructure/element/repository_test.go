package element

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
	"github.com/qkitzero/combination-service/internal/domain/element"
	mockscategory "github.com/qkitzero/combination-service/mocks/domain/category"
	mockselement "github.com/qkitzero/combination-service/mocks/domain/element"
	"github.com/qkitzero/combination-service/testutil"
)

func TestCreate(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name          string
		success       bool
		numCategories int
		setup         func(mock sqlmock.Sqlmock, element element.Element)
	}{
		{
			name:          "success create element",
			success:       true,
			numCategories: 0,
			setup: func(mock sqlmock.Sqlmock, element element.Element) {
				mock.ExpectBegin()

				mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO "elements" ("id","name","created_at") VALUES ($1,$2,$3)`)).
					WithArgs(element.ID(), element.Name(), testutil.AnyTime{}).
					WillReturnResult(sqlmock.NewResult(1, 1))

				mock.ExpectCommit()
			},
		},
		{
			name:          "success create element",
			success:       true,
			numCategories: 1,
			setup: func(mock sqlmock.Sqlmock, element element.Element) {
				mock.ExpectBegin()

				mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO "elements" ("id","name","created_at") VALUES ($1,$2,$3)`)).
					WithArgs(element.ID(), element.Name(), testutil.AnyTime{}).
					WillReturnResult(sqlmock.NewResult(1, 1))

				mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO "element_category" ("element_id","category_id") VALUES ($1,$2)`)).
					WithArgs(element.ID(), element.Categories()[0].ID()).
					WillReturnResult(sqlmock.NewResult(1, 1))

				mock.ExpectCommit()
			},
		},
		{
			name:          "success create element",
			success:       true,
			numCategories: 2,
			setup: func(mock sqlmock.Sqlmock, element element.Element) {
				mock.ExpectBegin()

				mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO "elements" ("id","name","created_at") VALUES ($1,$2,$3)`)).
					WithArgs(element.ID(), element.Name(), testutil.AnyTime{}).
					WillReturnResult(sqlmock.NewResult(1, 1))

				mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO "element_category" ("element_id","category_id") VALUES ($1,$2),($3,$4)`)).
					WithArgs(element.ID(), element.Categories()[0].ID(), element.ID(), element.Categories()[1].ID()).
					WillReturnResult(sqlmock.NewResult(1, 1))

				mock.ExpectCommit()
			},
		},
		{
			name:          "failure create element error",
			success:       false,
			numCategories: 1,
			setup: func(mock sqlmock.Sqlmock, element element.Element) {
				mock.ExpectBegin()

				mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO "elements" ("id","name","created_at") VALUES ($1,$2,$3)`)).
					WithArgs(element.ID(), element.Name(), testutil.AnyTime{}).
					WillReturnError(errors.New("create element error"))

				mock.ExpectRollback()
			},
		},
		{
			name:          "failure create element category error",
			success:       false,
			numCategories: 1,
			setup: func(mock sqlmock.Sqlmock, element element.Element) {
				mock.ExpectBegin()

				mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO "elements" ("id","name","created_at") VALUES ($1,$2,$3)`)).
					WithArgs(element.ID(), element.Name(), testutil.AnyTime{}).
					WillReturnResult(sqlmock.NewResult(1, 1))

				mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO "element_category" ("element_id","category_id") VALUES ($1,$2)`)).
					WithArgs(element.ID(), element.Categories()[0].ID()).
					WillReturnError(errors.New("create element category error"))

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

			categories := make([]category.Category, tt.numCategories)
			for i := range tt.numCategories {
				mockCategory := mockscategory.NewMockCategory(ctrl)
				mockCategory.EXPECT().ID().Return(category.CategoryID{UUID: uuid.New()}).AnyTimes()
				categories[i] = mockCategory
			}
			mockElement := mockselement.NewMockElement(ctrl)
			mockElement.EXPECT().ID().Return(element.ElementID{UUID: uuid.New()}).AnyTimes()
			mockElement.EXPECT().Name().Return(element.Name("test element")).AnyTimes()
			mockElement.EXPECT().Categories().Return(categories).AnyTimes()
			mockElement.EXPECT().CreatedAt().Return(time.Now()).AnyTimes()

			tt.setup(mock, mockElement)

			repo := NewElementRepository(gormDB)

			err = repo.Create(mockElement)
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

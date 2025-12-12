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
	"github.com/qkitzero/combination-service/internal/domain/language"
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

				mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO "elements" ("id","name","language_code","created_at") VALUES ($1,$2,$3,$4)`)).
					WithArgs(element.ID(), element.Name(), element.Language(), testutil.AnyTime{}).
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

				mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO "elements" ("id","name","language_code","created_at") VALUES ($1,$2,$3,$4)`)).
					WithArgs(element.ID(), element.Name(), element.Language(), testutil.AnyTime{}).
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

				mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO "elements" ("id","name","language_code","created_at") VALUES ($1,$2,$3,$4)`)).
					WithArgs(element.ID(), element.Name(), element.Language(), testutil.AnyTime{}).
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

				mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO "elements" ("id","name","language_code","created_at") VALUES ($1,$2,$3,$4)`)).
					WithArgs(element.ID(), element.Name(), element.Language(), testutil.AnyTime{}).
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

				mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO "elements" ("id","name","language_code","created_at") VALUES ($1,$2,$3,$4)`)).
					WithArgs(element.ID(), element.Name(), element.Language(), testutil.AnyTime{}).
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

			mockCategories := make([]category.Category, tt.numCategories)
			for i := range mockCategories {
				mockCategory := mockscategory.NewMockCategory(ctrl)
				mockCategory.EXPECT().ID().Return(category.CategoryID{UUID: uuid.New()}).AnyTimes()
				mockCategories[i] = mockCategory
			}
			mockElement := mockselement.NewMockElement(ctrl)
			mockElement.EXPECT().ID().Return(element.ElementID{UUID: uuid.New()}).AnyTimes()
			mockElement.EXPECT().Name().Return(element.Name("test element")).AnyTimes()
			mockElement.EXPECT().Language().Return(language.Language("en")).AnyTimes()
			mockElement.EXPECT().Categories().Return(mockCategories).AnyTimes()
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
func TestFindAll(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		success bool
		setup   func(mock sqlmock.Sqlmock)
	}{
		{
			name:    "success find all",
			success: true,
			setup: func(mock sqlmock.Sqlmock) {
				elementID := uuid.New()
				categoryID := uuid.New()

				elementRows := sqlmock.NewRows([]string{"id", "name", "language_code", "created_at"}).
					AddRow(elementID, "element name", "en", time.Now())
				mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "elements"`)).
					WillReturnRows(elementRows)

				elementCategoryRows := sqlmock.NewRows([]string{"element_id", "category_id"}).
					AddRow(elementID, categoryID)
				mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "element_category" WHERE element_id IN ($1)`)).
					WithArgs(elementID).
					WillReturnRows(elementCategoryRows)

				categoryRows := sqlmock.NewRows([]string{"id", "name", "language_code", "created_at"}).
					AddRow(categoryID, "category name", "en", time.Now())
				mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "categories" WHERE id IN ($1)`)).
					WithArgs(categoryID).
					WillReturnRows(categoryRows)
			},
		},
		{
			name:    "success find all",
			success: true,
			setup: func(mock sqlmock.Sqlmock) {
				elementID := uuid.New()
				categoryID := uuid.New()

				elementRows := sqlmock.NewRows([]string{"id", "name", "language_code", "created_at"}).
					AddRow(elementID, "element name", "en", time.Now())
				mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "elements"`)).
					WillReturnRows(elementRows)

				elementCategoryRows := sqlmock.NewRows([]string{"element_id", "category_id"}).
					AddRow(elementID, categoryID)
				mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "element_category" WHERE element_id IN ($1)`)).
					WithArgs(elementID).
					WillReturnRows(elementCategoryRows)

				categoryRows := sqlmock.NewRows([]string{"id", "name", "language_code", "created_at"})
				mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "categories" WHERE id IN ($1)`)).
					WithArgs(categoryID).
					WillReturnRows(categoryRows)
			},
		},
		{
			name:    "failure find element error",
			success: false,
			setup: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "elements"`)).
					WillReturnError(errors.New("find element error"))
			},
		},
		{
			name:    "failure find element category error",
			success: false,
			setup: func(mock sqlmock.Sqlmock) {
				elementID := uuid.New()

				elementRows := sqlmock.NewRows([]string{"id", "name", "language_code", "created_at"}).
					AddRow(elementID, "element name", "en", time.Now())
				mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "elements"`)).
					WillReturnRows(elementRows)

				mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "element_category" WHERE element_id IN ($1)`)).
					WithArgs(elementID).
					WillReturnError(errors.New("find element category error"))
			},
		},
		{
			name:    "failure find category error",
			success: false,
			setup: func(mock sqlmock.Sqlmock) {
				elementID := uuid.New()
				categoryID := uuid.New()

				elementRows := sqlmock.NewRows([]string{"id", "name", "language_code", "created_at"}).
					AddRow(elementID, "element name", "en", time.Now())
				mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "elements"`)).
					WillReturnRows(elementRows)

				elementCategoryRows := sqlmock.NewRows([]string{"element_id", "category_id"}).
					AddRow(elementID, categoryID)
				mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "element_category" WHERE element_id IN ($1)`)).
					WithArgs(elementID).
					WillReturnRows(elementCategoryRows)

				mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "categories" WHERE id IN ($1)`)).
					WithArgs(categoryID).
					WillReturnError(errors.New("find category error"))
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

			tt.setup(mock)

			repo := NewElementRepository(gormDB)

			_, err = repo.FindAll()
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

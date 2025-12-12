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
	"github.com/qkitzero/combination-service/internal/domain/language"
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

				mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO "categories" ("id","name","language_code","created_at") VALUES ($1,$2,$3,$4)`)).
					WithArgs(category.ID(), category.Name(), category.Language(), testutil.AnyTime{}).
					WillReturnResult(sqlmock.NewResult(1, 1))

				mock.ExpectCommit()
			},
		},
		{
			name:    "failure create category error",
			success: false,
			setup: func(mock sqlmock.Sqlmock, category category.Category) {
				mock.ExpectBegin()

				mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO "categories" ("id","name","language_code","created_at") VALUES ($1,$2,$3,$4)`)).
					WithArgs(category.ID(), category.Name(), category.Language(), testutil.AnyTime{}).
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
			mockCategory.EXPECT().Language().Return(language.Language("en")).AnyTimes()
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

func TestFindByID(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		success bool
		id      category.CategoryID
		setup   func(mock sqlmock.Sqlmock, id category.CategoryID)
	}{
		{
			name:    "success find category by id",
			success: true,
			id:      category.CategoryID{UUID: uuid.New()},
			setup: func(mock sqlmock.Sqlmock, id category.CategoryID) {
				categoryRows := sqlmock.NewRows([]string{"id", "name", "language_code", "created_at"}).
					AddRow(id, "test category", "en", time.Now())
				mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "categories" WHERE id = $1 ORDER BY "categories"."id" LIMIT $2`)).
					WithArgs(id, 1).
					WillReturnRows(categoryRows)
			},
		},
		{
			name:    "failure category not found",
			success: false,
			id:      category.CategoryID{UUID: uuid.New()},
			setup: func(mock sqlmock.Sqlmock, id category.CategoryID) {
				mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "categories" WHERE id = $1 ORDER BY "categories"."id" LIMIT $2`)).
					WithArgs(id, 1).
					WillReturnError(gorm.ErrRecordNotFound)
			},
		},
		{
			name:    "failure find category error",
			success: false,
			id:      category.CategoryID{UUID: uuid.New()},
			setup: func(mock sqlmock.Sqlmock, id category.CategoryID) {
				mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "categories" WHERE id = $1 ORDER BY "categories"."id" LIMIT $2`)).
					WithArgs(id, 1).
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

			tt.setup(mock, tt.id)

			repo := NewCategoryRepository(gormDB)

			_, err = repo.FindByID(tt.id)
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
			name:    "success find all categories",
			success: true,
			setup: func(mock sqlmock.Sqlmock) {
				categoryRows := sqlmock.NewRows([]string{"id", "name", "languages_code", "created_at"}).
					AddRow(uuid.New(), "test category", "en", time.Now())
				mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "categories"`)).
					WillReturnRows(categoryRows)
			},
		},
		{
			name:    "failure find all category error",
			success: false,
			setup: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "categories"`)).
					WillReturnError(errors.New("find all category error"))
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

			tt.setup(mock)

			repo := NewCategoryRepository(gormDB)

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

func TestFindAllByIDs(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		success bool
		ids     []category.CategoryID
		setup   func(mock sqlmock.Sqlmock, ids []category.CategoryID)
	}{
		{
			name:    "success find all categories by ids",
			success: true,
			ids:     []category.CategoryID{},
			setup: func(mock sqlmock.Sqlmock, ids []category.CategoryID) {
			},
		},
		{
			name:    "success find all categories by ids",
			success: true,
			ids:     []category.CategoryID{{UUID: uuid.New()}},
			setup: func(mock sqlmock.Sqlmock, ids []category.CategoryID) {
				categoryRows := sqlmock.NewRows([]string{"id", "name", "language_code", "created_at"}).
					AddRow(ids[0], "test category", "en", time.Now())
				mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "categories" WHERE id IN ($1)`)).
					WithArgs(ids[0]).
					WillReturnRows(categoryRows)
			},
		},
		{
			name:    "success find all categories by ids",
			success: true,
			ids:     []category.CategoryID{{UUID: uuid.New()}, {UUID: uuid.New()}},
			setup: func(mock sqlmock.Sqlmock, ids []category.CategoryID) {
				categoryRows := sqlmock.NewRows([]string{"id", "name", "language_code", "created_at"}).
					AddRow(ids[0], "test category", "en", time.Now()).
					AddRow(ids[1], "test category", "en", time.Now())
				mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "categories" WHERE id IN ($1,$2)`)).
					WithArgs(ids[0], ids[1]).
					WillReturnRows(categoryRows)
			},
		},
		{
			name:    "failure find category error",
			success: false,
			ids:     []category.CategoryID{{UUID: uuid.New()}},
			setup: func(mock sqlmock.Sqlmock, ids []category.CategoryID) {
				mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "categories" WHERE id IN ($1)`)).
					WithArgs(ids[0]).
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

			tt.setup(mock, tt.ids)

			repo := NewCategoryRepository(gormDB)

			_, err = repo.FindAllByIDs(tt.ids)
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

package mysql

import (
	"context"
	"database/sql"

	"github.com/bxcodec/go-clean-arch/domain"
)

type AuthorRepository struct {
	DB *sql.DB
}

// NewMysqlAuthorRepository will create an implementation of author.Repository
func NewAuthorRepository() *AuthorRepository {
	return &AuthorRepository{}
}

func (m *AuthorRepository) getOne(ctx context.Context, query string, args ...interface{}) (res domain.Author, err error) {
	stmt, err := m.DB.PrepareContext(ctx, query)
	if err != nil {
		return domain.Author{}, err
	}
	row := stmt.QueryRowContext(ctx, args...)
	res = domain.Author{}

	err = row.Scan(
		&res.ID,
		&res.Name,
		&res.CreatedAt,
		&res.UpdatedAt,
	)
	return
}

func (m *AuthorRepository) GetByID(ctx context.Context, id int64) (domain.Author, error) {
	query := `SELECT id, name, created_at, updated_at FROM author WHERE id=?`
	return m.getOne(ctx, query, id)
}

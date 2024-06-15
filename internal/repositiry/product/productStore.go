package product

import (
	"errors"
	sq "github.com/Masterminds/squirrel"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	logger "go_admin_back/internal/loger"
	"go_admin_back/internal/models/product"
	"go_admin_back/internal/repositiry"
	"net/http"
)

type PGProductStore struct {
	store *repositiry.Store
}

func NewDBProductStore(store *repositiry.Store) *PGProductStore {
	return &PGProductStore{
		store: store,
	}
}

func (pg *PGProductStore) Create(cf *fiber.Ctx, p *product.Product) (*product.Product, error) {
	sqlBuilder := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	ctx := cf.Context()
	tx, err := pg.store.Db().Begin(ctx)
	if err != nil {
		logger.Log.Debug("Error While Create. Error Occurs in Begin Method", zap.Error(err))
		return nil, err
	}
	defer tx.Rollback(ctx)

	sqlSelect := sqlBuilder.Insert("products").
		Columns("catalog_id", "uuid", "alias", "name", "created_at", "updated_at", "is_deleted", "is_enabled").
		Values(p.CatalogId, p.Uuid, p.Alias, p.Name, p.CreatedAt, p.UpdatedAt, p.IsDeleted, p.IsEnabled).
		Suffix("RETURNING id")
	query, args, err := sqlSelect.ToSql()
	if err != nil {
		logger.Log.Debug("Error While Create. Error Building SQL", zap.Error(err))
		return nil, err
	}
	err = tx.QueryRow(ctx, query, args...).Scan(&p.Id)
	if err != nil {
		logger.Log.Debug("Error While Create. Rrror in Method QueryRow", zap.Error(err))
		msg := errors.Wrap(err, "Bad Request")
		err = errorDomain.NewCustomError(msg, http.StatusBadRequest)
		return nil, err
	}
}

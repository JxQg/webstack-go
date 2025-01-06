// Code generated by gen-table. DO NOT EDIT.
// Code generated by gen-table. DO NOT EDIT.
// Code generated by gen-table. DO NOT EDIT.

package repository

import (
	"time"

	"github.com/ch3nnn/webstack-go/internal/dal/model"
	"github.com/ch3nnn/webstack-go/internal/dal/query"

	"gorm.io/gen"
	"gorm.io/gen/field"
)

var _ iStSiteDao = (*stSiteDao)(nil)

// ------------------------------------
// StSite  ColumnName
// ------------------------------------
type iWhereStSiteFunc interface {
	WhereByID(id int) func(dao gen.Dao) gen.Dao
	WhereByCategoryID(categoryId int) func(dao gen.Dao) gen.Dao
	WhereByTitle(title string) func(dao gen.Dao) gen.Dao
	WhereByIcon(icon string) func(dao gen.Dao) gen.Dao
	WhereByDescription(description string) func(dao gen.Dao) gen.Dao
	WhereByURL(url string) func(dao gen.Dao) gen.Dao
	WhereByIsUsed(isUsed bool) func(dao gen.Dao) gen.Dao
	WhereByCreatedAt(createdAt time.Time) func(dao gen.Dao) gen.Dao
	WhereByUpdatedAt(updatedAt time.Time) func(dao gen.Dao) gen.Dao
	WhereByDeletedAt(deletedAt time.Time) func(dao gen.Dao) gen.Dao
}

// ------------------------------------
// StSite  Generate Function
// ------------------------------------
type iGenStSiteFunc interface {
	Create(m *model.StSite) (*model.StSite, error)
	Delete(whereFunc ...func(dao gen.Dao) gen.Dao) error
	DeletePhysical(whereFunc ...func(dao gen.Dao) gen.Dao) error
	Update(v interface{}, whereFunc ...func(dao gen.Dao) gen.Dao) (rowsAffected int64, err error)
	FindCount(whereFunc ...func(dao gen.Dao) gen.Dao) (int64, error)
	FindOne(whereFunc ...func(dao gen.Dao) gen.Dao) (*model.StSite, error)
	FindAll(whereFunc ...func(dao gen.Dao) gen.Dao) ([]*model.StSite, error)
	FindPage(page int, pageSize int, orderColumns []field.Expr, whereFunc ...func(dao gen.Dao) gen.Dao) ([]*model.StSite, int64, error)
	Scan(result interface{}, whereFunc ...func(dao gen.Dao) gen.Dao) (err error)
	ScanPage(page int, pageSize int, orderColumns []field.Expr, result interface{}, whereFunc ...func(dao gen.Dao) gen.Dao) (count int64, err error)
}

type iStSiteDao interface {
	iWhereStSiteFunc
	iGenStSiteFunc
}

type stSiteDao struct {
	stSiteDo query.IStSiteDo
}

func (s *stSiteDao) WhereByID(id int) func(dao gen.Dao) gen.Dao {
	return func(dao gen.Dao) gen.Dao {
		return dao.Where(query.StSite.ID.Eq(id))
	}
}

func (s *stSiteDao) WhereByCategoryID(categoryId int) func(dao gen.Dao) gen.Dao {
	return func(dao gen.Dao) gen.Dao {
		return dao.Where(query.StSite.CategoryID.Eq(categoryId))
	}
}

func (s *stSiteDao) WhereByTitle(title string) func(dao gen.Dao) gen.Dao {
	return func(dao gen.Dao) gen.Dao {
		return dao.Where(query.StSite.Title.Eq(title))
	}
}

func (s *stSiteDao) WhereByIcon(icon string) func(dao gen.Dao) gen.Dao {
	return func(dao gen.Dao) gen.Dao {
		return dao.Where(query.StSite.Icon.Eq(icon))
	}
}

func (s *stSiteDao) WhereByDescription(description string) func(dao gen.Dao) gen.Dao {
	return func(dao gen.Dao) gen.Dao {
		return dao.Where(query.StSite.Description.Eq(description))
	}
}

func (s *stSiteDao) WhereByURL(url string) func(dao gen.Dao) gen.Dao {
	return func(dao gen.Dao) gen.Dao {
		return dao.Where(query.StSite.URL.Eq(url))
	}
}

func (s *stSiteDao) WhereByIsUsed(isUsed bool) func(dao gen.Dao) gen.Dao {
	return func(dao gen.Dao) gen.Dao {
		return dao.Where(query.StSite.IsUsed.Is(isUsed))
	}
}

func (s *stSiteDao) WhereByCreatedAt(createdAt time.Time) func(dao gen.Dao) gen.Dao {
	return func(dao gen.Dao) gen.Dao {
		return dao.Where(query.StSite.CreatedAt.Eq(createdAt))
	}
}

func (s *stSiteDao) WhereByUpdatedAt(updatedAt time.Time) func(dao gen.Dao) gen.Dao {
	return func(dao gen.Dao) gen.Dao {
		return dao.Where(query.StSite.UpdatedAt.Eq(updatedAt))
	}
}

func (s *stSiteDao) WhereByDeletedAt(deletedAt time.Time) func(dao gen.Dao) gen.Dao {
	return func(dao gen.Dao) gen.Dao {
		return dao.Where(query.StSite.DeletedAt.Eq(deletedAt))
	}
}

func (s *stSiteDao) Create(m *model.StSite) (*model.StSite, error) {
	if err := s.stSiteDo.Create(m); err != nil {
		return nil, err
	}
	return s.FindOne(s.WhereByID(m.ID))
}

func (s *stSiteDao) FindCount(whereFunc ...func(dao gen.Dao) gen.Dao) (int64, error) {
	return s.stSiteDo.Scopes(whereFunc...).Count()
}

func (s *stSiteDao) FindOne(whereFunc ...func(dao gen.Dao) gen.Dao) (*model.StSite, error) {
	return s.stSiteDo.Scopes(whereFunc...).First()
}

func (s *stSiteDao) FindAll(whereFunc ...func(dao gen.Dao) gen.Dao) ([]*model.StSite, error) {
	return s.stSiteDo.Scopes(whereFunc...).Find()
}

func (s *stSiteDao) FindPage(page int, pageSize int, orderColumns []field.Expr, whereFunc ...func(dao gen.Dao) gen.Dao) ([]*model.StSite, int64, error) {
	return s.stSiteDo.Scopes(whereFunc...).Order(orderColumns...).FindByPage((page-1)*pageSize, pageSize)
}

// 注意 当通过 struct 更新时，GORM 只会更新非零字段
// 如果想确保指定字段被更新，使用 map 来完成更新操作
func (s *stSiteDao) Update(v interface{}, whereFunc ...func(dao gen.Dao) gen.Dao) (rowsAffected int64, err error) {
	info, err := s.stSiteDo.Scopes(whereFunc...).Updates(v)
	if err != nil {
		return rowsAffected, err
	}

	return info.RowsAffected, nil
}

func (s *stSiteDao) Delete(whereFunc ...func(dao gen.Dao) gen.Dao) error {
	if _, err := s.stSiteDo.Scopes(whereFunc...).Delete(); err != nil {
		return err
	}
	return nil
}

func (s *stSiteDao) DeletePhysical(whereFunc ...func(dao gen.Dao) gen.Dao) error {
	if _, err := s.stSiteDo.Unscoped().Scopes(whereFunc...).Delete(); err != nil {
		return err
	}
	return nil
}

func (s *stSiteDao) Scan(result interface{}, whereFunc ...func(dao gen.Dao) gen.Dao) (err error) {
	return s.stSiteDo.Scopes(whereFunc...).Scan(result)
}

func (s *stSiteDao) ScanPage(page int, pageSize int, orderColumns []field.Expr, result interface{}, whereFunc ...func(dao gen.Dao) gen.Dao) (count int64, err error) {
	return s.stSiteDo.Scopes(whereFunc...).Order(orderColumns...).ScanByPage(result, (page-1)*pageSize, pageSize)
}
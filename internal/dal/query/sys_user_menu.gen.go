// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/ch3nnn/webstack-go/internal/dal/model"
)

func newSysUserMenu(db *gorm.DB, opts ...gen.DOOption) sysUserMenu {
	_sysUserMenu := sysUserMenu{}

	_sysUserMenu.sysUserMenuDo.UseDB(db, opts...)
	_sysUserMenu.sysUserMenuDo.UseModel(&model.SysUserMenu{})

	tableName := _sysUserMenu.sysUserMenuDo.TableName()
	_sysUserMenu.ALL = field.NewAsterisk(tableName)
	_sysUserMenu.ID = field.NewInt(tableName, "id")
	_sysUserMenu.UserID = field.NewInt(tableName, "user_id")
	_sysUserMenu.MenuID = field.NewInt(tableName, "menu_id")
	_sysUserMenu.CreatedAt = field.NewTime(tableName, "created_at")
	_sysUserMenu.UpdatedAt = field.NewTime(tableName, "updated_at")
	_sysUserMenu.DeletedAt = field.NewTime(tableName, "deleted_at")

	_sysUserMenu.fillFieldMap()

	return _sysUserMenu
}

type sysUserMenu struct {
	sysUserMenuDo sysUserMenuDo

	ALL       field.Asterisk
	ID        field.Int
	UserID    field.Int
	MenuID    field.Int
	CreatedAt field.Time
	UpdatedAt field.Time
	DeletedAt field.Time

	fieldMap map[string]field.Expr
}

func (s sysUserMenu) Table(newTableName string) *sysUserMenu {
	s.sysUserMenuDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s sysUserMenu) As(alias string) *sysUserMenu {
	s.sysUserMenuDo.DO = *(s.sysUserMenuDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *sysUserMenu) updateTableName(table string) *sysUserMenu {
	s.ALL = field.NewAsterisk(table)
	s.ID = field.NewInt(table, "id")
	s.UserID = field.NewInt(table, "user_id")
	s.MenuID = field.NewInt(table, "menu_id")
	s.CreatedAt = field.NewTime(table, "created_at")
	s.UpdatedAt = field.NewTime(table, "updated_at")
	s.DeletedAt = field.NewTime(table, "deleted_at")

	s.fillFieldMap()

	return s
}

func (s *sysUserMenu) WithContext(ctx context.Context) ISysUserMenuDo {
	return s.sysUserMenuDo.WithContext(ctx)
}

func (s sysUserMenu) TableName() string { return s.sysUserMenuDo.TableName() }

func (s sysUserMenu) Alias() string { return s.sysUserMenuDo.Alias() }

func (s sysUserMenu) Columns(cols ...field.Expr) gen.Columns { return s.sysUserMenuDo.Columns(cols...) }

func (s *sysUserMenu) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *sysUserMenu) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 6)
	s.fieldMap["id"] = s.ID
	s.fieldMap["user_id"] = s.UserID
	s.fieldMap["menu_id"] = s.MenuID
	s.fieldMap["created_at"] = s.CreatedAt
	s.fieldMap["updated_at"] = s.UpdatedAt
	s.fieldMap["deleted_at"] = s.DeletedAt
}

func (s sysUserMenu) clone(db *gorm.DB) sysUserMenu {
	s.sysUserMenuDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s sysUserMenu) replaceDB(db *gorm.DB) sysUserMenu {
	s.sysUserMenuDo.ReplaceDB(db)
	return s
}

type sysUserMenuDo struct{ gen.DO }

type ISysUserMenuDo interface {
	gen.SubQuery
	Debug() ISysUserMenuDo
	WithContext(ctx context.Context) ISysUserMenuDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ISysUserMenuDo
	WriteDB() ISysUserMenuDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ISysUserMenuDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ISysUserMenuDo
	Not(conds ...gen.Condition) ISysUserMenuDo
	Or(conds ...gen.Condition) ISysUserMenuDo
	Select(conds ...field.Expr) ISysUserMenuDo
	Where(conds ...gen.Condition) ISysUserMenuDo
	Order(conds ...field.Expr) ISysUserMenuDo
	Distinct(cols ...field.Expr) ISysUserMenuDo
	Omit(cols ...field.Expr) ISysUserMenuDo
	Join(table schema.Tabler, on ...field.Expr) ISysUserMenuDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ISysUserMenuDo
	RightJoin(table schema.Tabler, on ...field.Expr) ISysUserMenuDo
	Group(cols ...field.Expr) ISysUserMenuDo
	Having(conds ...gen.Condition) ISysUserMenuDo
	Limit(limit int) ISysUserMenuDo
	Offset(offset int) ISysUserMenuDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ISysUserMenuDo
	Unscoped() ISysUserMenuDo
	Create(values ...*model.SysUserMenu) error
	CreateInBatches(values []*model.SysUserMenu, batchSize int) error
	Save(values ...*model.SysUserMenu) error
	First() (*model.SysUserMenu, error)
	Take() (*model.SysUserMenu, error)
	Last() (*model.SysUserMenu, error)
	Find() ([]*model.SysUserMenu, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SysUserMenu, err error)
	FindInBatches(result *[]*model.SysUserMenu, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.SysUserMenu) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ISysUserMenuDo
	Assign(attrs ...field.AssignExpr) ISysUserMenuDo
	Joins(fields ...field.RelationField) ISysUserMenuDo
	Preload(fields ...field.RelationField) ISysUserMenuDo
	FirstOrInit() (*model.SysUserMenu, error)
	FirstOrCreate() (*model.SysUserMenu, error)
	FindByPage(offset int, limit int) (result []*model.SysUserMenu, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ISysUserMenuDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s sysUserMenuDo) Debug() ISysUserMenuDo {
	return s.withDO(s.DO.Debug())
}

func (s sysUserMenuDo) WithContext(ctx context.Context) ISysUserMenuDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s sysUserMenuDo) ReadDB() ISysUserMenuDo {
	return s.Clauses(dbresolver.Read)
}

func (s sysUserMenuDo) WriteDB() ISysUserMenuDo {
	return s.Clauses(dbresolver.Write)
}

func (s sysUserMenuDo) Session(config *gorm.Session) ISysUserMenuDo {
	return s.withDO(s.DO.Session(config))
}

func (s sysUserMenuDo) Clauses(conds ...clause.Expression) ISysUserMenuDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s sysUserMenuDo) Returning(value interface{}, columns ...string) ISysUserMenuDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s sysUserMenuDo) Not(conds ...gen.Condition) ISysUserMenuDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s sysUserMenuDo) Or(conds ...gen.Condition) ISysUserMenuDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s sysUserMenuDo) Select(conds ...field.Expr) ISysUserMenuDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s sysUserMenuDo) Where(conds ...gen.Condition) ISysUserMenuDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s sysUserMenuDo) Order(conds ...field.Expr) ISysUserMenuDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s sysUserMenuDo) Distinct(cols ...field.Expr) ISysUserMenuDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s sysUserMenuDo) Omit(cols ...field.Expr) ISysUserMenuDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s sysUserMenuDo) Join(table schema.Tabler, on ...field.Expr) ISysUserMenuDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s sysUserMenuDo) LeftJoin(table schema.Tabler, on ...field.Expr) ISysUserMenuDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s sysUserMenuDo) RightJoin(table schema.Tabler, on ...field.Expr) ISysUserMenuDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s sysUserMenuDo) Group(cols ...field.Expr) ISysUserMenuDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s sysUserMenuDo) Having(conds ...gen.Condition) ISysUserMenuDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s sysUserMenuDo) Limit(limit int) ISysUserMenuDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s sysUserMenuDo) Offset(offset int) ISysUserMenuDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s sysUserMenuDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ISysUserMenuDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s sysUserMenuDo) Unscoped() ISysUserMenuDo {
	return s.withDO(s.DO.Unscoped())
}

func (s sysUserMenuDo) Create(values ...*model.SysUserMenu) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s sysUserMenuDo) CreateInBatches(values []*model.SysUserMenu, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s sysUserMenuDo) Save(values ...*model.SysUserMenu) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s sysUserMenuDo) First() (*model.SysUserMenu, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysUserMenu), nil
	}
}

func (s sysUserMenuDo) Take() (*model.SysUserMenu, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysUserMenu), nil
	}
}

func (s sysUserMenuDo) Last() (*model.SysUserMenu, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysUserMenu), nil
	}
}

func (s sysUserMenuDo) Find() ([]*model.SysUserMenu, error) {
	result, err := s.DO.Find()
	return result.([]*model.SysUserMenu), err
}

func (s sysUserMenuDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SysUserMenu, err error) {
	buf := make([]*model.SysUserMenu, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s sysUserMenuDo) FindInBatches(result *[]*model.SysUserMenu, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s sysUserMenuDo) Attrs(attrs ...field.AssignExpr) ISysUserMenuDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s sysUserMenuDo) Assign(attrs ...field.AssignExpr) ISysUserMenuDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s sysUserMenuDo) Joins(fields ...field.RelationField) ISysUserMenuDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s sysUserMenuDo) Preload(fields ...field.RelationField) ISysUserMenuDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s sysUserMenuDo) FirstOrInit() (*model.SysUserMenu, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysUserMenu), nil
	}
}

func (s sysUserMenuDo) FirstOrCreate() (*model.SysUserMenu, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysUserMenu), nil
	}
}

func (s sysUserMenuDo) FindByPage(offset int, limit int) (result []*model.SysUserMenu, count int64, err error) {
	result, err = s.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = s.Offset(-1).Limit(-1).Count()
	return
}

func (s sysUserMenuDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s sysUserMenuDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s sysUserMenuDo) Delete(models ...*model.SysUserMenu) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *sysUserMenuDo) withDO(do gen.Dao) *sysUserMenuDo {
	s.DO = *do.(*gen.DO)
	return s
}
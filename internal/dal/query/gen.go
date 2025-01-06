// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"
	"database/sql"

	"gorm.io/gorm"

	"gorm.io/gen"

	"gorm.io/plugin/dbresolver"
)

var (
	Q           = new(Query)
	StCategory  *stCategory
	StSite      *stSite
	SysMenu     *sysMenu
	SysUser     *sysUser
	SysUserMenu *sysUserMenu
)

func SetDefault(db *gorm.DB, opts ...gen.DOOption) {
	*Q = *Use(db, opts...)
	StCategory = &Q.StCategory
	StSite = &Q.StSite
	SysMenu = &Q.SysMenu
	SysUser = &Q.SysUser
	SysUserMenu = &Q.SysUserMenu
}

func Use(db *gorm.DB, opts ...gen.DOOption) *Query {
	return &Query{
		db:          db,
		StCategory:  newStCategory(db, opts...),
		StSite:      newStSite(db, opts...),
		SysMenu:     newSysMenu(db, opts...),
		SysUser:     newSysUser(db, opts...),
		SysUserMenu: newSysUserMenu(db, opts...),
	}
}

type Query struct {
	db *gorm.DB

	StCategory  stCategory
	StSite      stSite
	SysMenu     sysMenu
	SysUser     sysUser
	SysUserMenu sysUserMenu
}

func (q *Query) Available() bool { return q.db != nil }

func (q *Query) clone(db *gorm.DB) *Query {
	return &Query{
		db:          db,
		StCategory:  q.StCategory.clone(db),
		StSite:      q.StSite.clone(db),
		SysMenu:     q.SysMenu.clone(db),
		SysUser:     q.SysUser.clone(db),
		SysUserMenu: q.SysUserMenu.clone(db),
	}
}

func (q *Query) ReadDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Read))
}

func (q *Query) WriteDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Write))
}

func (q *Query) ReplaceDB(db *gorm.DB) *Query {
	return &Query{
		db:          db,
		StCategory:  q.StCategory.replaceDB(db),
		StSite:      q.StSite.replaceDB(db),
		SysMenu:     q.SysMenu.replaceDB(db),
		SysUser:     q.SysUser.replaceDB(db),
		SysUserMenu: q.SysUserMenu.replaceDB(db),
	}
}

type queryCtx struct {
	StCategory  IStCategoryDo
	StSite      IStSiteDo
	SysMenu     ISysMenuDo
	SysUser     ISysUserDo
	SysUserMenu ISysUserMenuDo
}

func (q *Query) WithContext(ctx context.Context) *queryCtx {
	return &queryCtx{
		StCategory:  q.StCategory.WithContext(ctx),
		StSite:      q.StSite.WithContext(ctx),
		SysMenu:     q.SysMenu.WithContext(ctx),
		SysUser:     q.SysUser.WithContext(ctx),
		SysUserMenu: q.SysUserMenu.WithContext(ctx),
	}
}

func (q *Query) Transaction(fc func(tx *Query) error, opts ...*sql.TxOptions) error {
	return q.db.Transaction(func(tx *gorm.DB) error { return fc(q.clone(tx)) }, opts...)
}

func (q *Query) Begin(opts ...*sql.TxOptions) *QueryTx {
	tx := q.db.Begin(opts...)
	return &QueryTx{Query: q.clone(tx), Error: tx.Error}
}

type QueryTx struct {
	*Query
	Error error
}

func (q *QueryTx) Commit() error {
	return q.db.Commit().Error
}

func (q *QueryTx) Rollback() error {
	return q.db.Rollback().Error
}

func (q *QueryTx) SavePoint(name string) error {
	return q.db.SavePoint(name).Error
}

func (q *QueryTx) RollbackTo(name string) error {
	return q.db.RollbackTo(name).Error
}
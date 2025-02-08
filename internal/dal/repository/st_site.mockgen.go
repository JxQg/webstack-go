// Code generated by MockGen. DO NOT EDIT.
// Source: internal/dal/repository/st_site.go
//
// Generated by this command:
//
//	mockgen -source=internal/dal/repository/st_site.go -destination internal/dal/repository/st_site.mockgen.go -package=repository
//

// Package repository is a generated GoMock package.
package repository

import (
	context "context"
	reflect "reflect"
	time "time"

	model "github.com/ch3nnn/webstack-go/internal/dal/model"
	gomock "go.uber.org/mock/gomock"
	gen "gorm.io/gen"
	field "gorm.io/gen/field"
)

// MockIStSiteDao is a mock of IStSiteDao interface.
type MockIStSiteDao struct {
	ctrl     *gomock.Controller
	recorder *MockIStSiteDaoMockRecorder
	isgomock struct{}
}

// MockIStSiteDaoMockRecorder is the mock recorder for MockIStSiteDao.
type MockIStSiteDaoMockRecorder struct {
	mock *MockIStSiteDao
}

// NewMockIStSiteDao creates a new mock instance.
func NewMockIStSiteDao(ctrl *gomock.Controller) *MockIStSiteDao {
	mock := &MockIStSiteDao{ctrl: ctrl}
	mock.recorder = &MockIStSiteDaoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIStSiteDao) EXPECT() *MockIStSiteDaoMockRecorder {
	return m.recorder
}

// LikeInByTitleOrDescOrURL mocks base method.
func (m *MockIStSiteDao) LikeInByTitleOrDescOrURL(search string) func(gen.Dao) gen.Dao {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LikeInByTitleOrDescOrURL", search)
	ret0, _ := ret[0].(func(gen.Dao) gen.Dao)
	return ret0
}

// LikeInByTitleOrDescOrURL indicates an expected call of LikeInByTitleOrDescOrURL.
func (mr *MockIStSiteDaoMockRecorder) LikeInByTitleOrDescOrURL(search any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LikeInByTitleOrDescOrURL", reflect.TypeOf((*MockIStSiteDao)(nil).LikeInByTitleOrDescOrURL), search)
}

// WhereByCategoryID mocks base method.
func (m *MockIStSiteDao) WhereByCategoryID(categoryId int) func(gen.Dao) gen.Dao {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WhereByCategoryID", categoryId)
	ret0, _ := ret[0].(func(gen.Dao) gen.Dao)
	return ret0
}

// WhereByCategoryID indicates an expected call of WhereByCategoryID.
func (mr *MockIStSiteDaoMockRecorder) WhereByCategoryID(categoryId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WhereByCategoryID", reflect.TypeOf((*MockIStSiteDao)(nil).WhereByCategoryID), categoryId)
}

// WhereByCreatedAt mocks base method.
func (m *MockIStSiteDao) WhereByCreatedAt(createdAt time.Time) func(gen.Dao) gen.Dao {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WhereByCreatedAt", createdAt)
	ret0, _ := ret[0].(func(gen.Dao) gen.Dao)
	return ret0
}

// WhereByCreatedAt indicates an expected call of WhereByCreatedAt.
func (mr *MockIStSiteDaoMockRecorder) WhereByCreatedAt(createdAt any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WhereByCreatedAt", reflect.TypeOf((*MockIStSiteDao)(nil).WhereByCreatedAt), createdAt)
}

// WhereByDeletedAt mocks base method.
func (m *MockIStSiteDao) WhereByDeletedAt(deletedAt time.Time) func(gen.Dao) gen.Dao {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WhereByDeletedAt", deletedAt)
	ret0, _ := ret[0].(func(gen.Dao) gen.Dao)
	return ret0
}

// WhereByDeletedAt indicates an expected call of WhereByDeletedAt.
func (mr *MockIStSiteDaoMockRecorder) WhereByDeletedAt(deletedAt any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WhereByDeletedAt", reflect.TypeOf((*MockIStSiteDao)(nil).WhereByDeletedAt), deletedAt)
}

// WhereByDescription mocks base method.
func (m *MockIStSiteDao) WhereByDescription(description string) func(gen.Dao) gen.Dao {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WhereByDescription", description)
	ret0, _ := ret[0].(func(gen.Dao) gen.Dao)
	return ret0
}

// WhereByDescription indicates an expected call of WhereByDescription.
func (mr *MockIStSiteDaoMockRecorder) WhereByDescription(description any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WhereByDescription", reflect.TypeOf((*MockIStSiteDao)(nil).WhereByDescription), description)
}

// WhereByID mocks base method.
func (m *MockIStSiteDao) WhereByID(id int) func(gen.Dao) gen.Dao {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WhereByID", id)
	ret0, _ := ret[0].(func(gen.Dao) gen.Dao)
	return ret0
}

// WhereByID indicates an expected call of WhereByID.
func (mr *MockIStSiteDaoMockRecorder) WhereByID(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WhereByID", reflect.TypeOf((*MockIStSiteDao)(nil).WhereByID), id)
}

// WhereByIcon mocks base method.
func (m *MockIStSiteDao) WhereByIcon(icon string) func(gen.Dao) gen.Dao {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WhereByIcon", icon)
	ret0, _ := ret[0].(func(gen.Dao) gen.Dao)
	return ret0
}

// WhereByIcon indicates an expected call of WhereByIcon.
func (mr *MockIStSiteDaoMockRecorder) WhereByIcon(icon any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WhereByIcon", reflect.TypeOf((*MockIStSiteDao)(nil).WhereByIcon), icon)
}

// WhereByIsUsed mocks base method.
func (m *MockIStSiteDao) WhereByIsUsed(isUsed bool) func(gen.Dao) gen.Dao {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WhereByIsUsed", isUsed)
	ret0, _ := ret[0].(func(gen.Dao) gen.Dao)
	return ret0
}

// WhereByIsUsed indicates an expected call of WhereByIsUsed.
func (mr *MockIStSiteDaoMockRecorder) WhereByIsUsed(isUsed any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WhereByIsUsed", reflect.TypeOf((*MockIStSiteDao)(nil).WhereByIsUsed), isUsed)
}

// WhereBySort mocks base method.
func (m *MockIStSiteDao) WhereBySort(sort int) func(gen.Dao) gen.Dao {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WhereBySort", sort)
	ret0, _ := ret[0].(func(gen.Dao) gen.Dao)
	return ret0
}

// WhereBySort indicates an expected call of WhereBySort.
func (mr *MockIStSiteDaoMockRecorder) WhereBySort(sort any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WhereBySort", reflect.TypeOf((*MockIStSiteDao)(nil).WhereBySort), sort)
}

// WhereByTitle mocks base method.
func (m *MockIStSiteDao) WhereByTitle(title string) func(gen.Dao) gen.Dao {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WhereByTitle", title)
	ret0, _ := ret[0].(func(gen.Dao) gen.Dao)
	return ret0
}

// WhereByTitle indicates an expected call of WhereByTitle.
func (mr *MockIStSiteDaoMockRecorder) WhereByTitle(title any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WhereByTitle", reflect.TypeOf((*MockIStSiteDao)(nil).WhereByTitle), title)
}

// WhereByURL mocks base method.
func (m *MockIStSiteDao) WhereByURL(url string) func(gen.Dao) gen.Dao {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WhereByURL", url)
	ret0, _ := ret[0].(func(gen.Dao) gen.Dao)
	return ret0
}

// WhereByURL indicates an expected call of WhereByURL.
func (mr *MockIStSiteDaoMockRecorder) WhereByURL(url any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WhereByURL", reflect.TypeOf((*MockIStSiteDao)(nil).WhereByURL), url)
}

// WhereByUpdatedAt mocks base method.
func (m *MockIStSiteDao) WhereByUpdatedAt(updatedAt time.Time) func(gen.Dao) gen.Dao {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WhereByUpdatedAt", updatedAt)
	ret0, _ := ret[0].(func(gen.Dao) gen.Dao)
	return ret0
}

// WhereByUpdatedAt indicates an expected call of WhereByUpdatedAt.
func (mr *MockIStSiteDaoMockRecorder) WhereByUpdatedAt(updatedAt any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WhereByUpdatedAt", reflect.TypeOf((*MockIStSiteDao)(nil).WhereByUpdatedAt), updatedAt)
}

// WithContext mocks base method.
func (m *MockIStSiteDao) WithContext(ctx context.Context) iCustomGenStSiteFunc {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithContext", ctx)
	ret0, _ := ret[0].(iCustomGenStSiteFunc)
	return ret0
}

// WithContext indicates an expected call of WithContext.
func (mr *MockIStSiteDaoMockRecorder) WithContext(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithContext", reflect.TypeOf((*MockIStSiteDao)(nil).WithContext), ctx)
}

// MockiCustomGenStSiteFunc is a mock of iCustomGenStSiteFunc interface.
type MockiCustomGenStSiteFunc struct {
	ctrl     *gomock.Controller
	recorder *MockiCustomGenStSiteFuncMockRecorder
	isgomock struct{}
}

// MockiCustomGenStSiteFuncMockRecorder is the mock recorder for MockiCustomGenStSiteFunc.
type MockiCustomGenStSiteFuncMockRecorder struct {
	mock *MockiCustomGenStSiteFunc
}

// NewMockiCustomGenStSiteFunc creates a new mock instance.
func NewMockiCustomGenStSiteFunc(ctrl *gomock.Controller) *MockiCustomGenStSiteFunc {
	mock := &MockiCustomGenStSiteFunc{ctrl: ctrl}
	mock.recorder = &MockiCustomGenStSiteFuncMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockiCustomGenStSiteFunc) EXPECT() *MockiCustomGenStSiteFuncMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m_2 *MockiCustomGenStSiteFunc) Create(m *model.StSite) (*model.StSite, error) {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "Create", m)
	ret0, _ := ret[0].(*model.StSite)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockiCustomGenStSiteFuncMockRecorder) Create(m any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockiCustomGenStSiteFunc)(nil).Create), m)
}

// Delete mocks base method.
func (m *MockiCustomGenStSiteFunc) Delete(whereFunc ...func(gen.Dao) gen.Dao) error {
	m.ctrl.T.Helper()
	varargs := []any{}
	for _, a := range whereFunc {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Delete", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockiCustomGenStSiteFuncMockRecorder) Delete(whereFunc ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockiCustomGenStSiteFunc)(nil).Delete), whereFunc...)
}

// DeletePhysical mocks base method.
func (m *MockiCustomGenStSiteFunc) DeletePhysical(whereFunc ...func(gen.Dao) gen.Dao) error {
	m.ctrl.T.Helper()
	varargs := []any{}
	for _, a := range whereFunc {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeletePhysical", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeletePhysical indicates an expected call of DeletePhysical.
func (mr *MockiCustomGenStSiteFuncMockRecorder) DeletePhysical(whereFunc ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePhysical", reflect.TypeOf((*MockiCustomGenStSiteFunc)(nil).DeletePhysical), whereFunc...)
}

// FindAll mocks base method.
func (m *MockiCustomGenStSiteFunc) FindAll(whereFunc ...func(gen.Dao) gen.Dao) ([]*model.StSite, error) {
	m.ctrl.T.Helper()
	varargs := []any{}
	for _, a := range whereFunc {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FindAll", varargs...)
	ret0, _ := ret[0].([]*model.StSite)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAll indicates an expected call of FindAll.
func (mr *MockiCustomGenStSiteFuncMockRecorder) FindAll(whereFunc ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockiCustomGenStSiteFunc)(nil).FindAll), whereFunc...)
}

// FindCount mocks base method.
func (m *MockiCustomGenStSiteFunc) FindCount(whereFunc ...func(gen.Dao) gen.Dao) (int64, error) {
	m.ctrl.T.Helper()
	varargs := []any{}
	for _, a := range whereFunc {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FindCount", varargs...)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindCount indicates an expected call of FindCount.
func (mr *MockiCustomGenStSiteFuncMockRecorder) FindCount(whereFunc ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindCount", reflect.TypeOf((*MockiCustomGenStSiteFunc)(nil).FindCount), whereFunc...)
}

// FindOne mocks base method.
func (m *MockiCustomGenStSiteFunc) FindOne(whereFunc ...func(gen.Dao) gen.Dao) (*model.StSite, error) {
	m.ctrl.T.Helper()
	varargs := []any{}
	for _, a := range whereFunc {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FindOne", varargs...)
	ret0, _ := ret[0].(*model.StSite)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOne indicates an expected call of FindOne.
func (mr *MockiCustomGenStSiteFuncMockRecorder) FindOne(whereFunc ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOne", reflect.TypeOf((*MockiCustomGenStSiteFunc)(nil).FindOne), whereFunc...)
}

// FindPage mocks base method.
func (m *MockiCustomGenStSiteFunc) FindPage(page, pageSize int, orderColumns []field.Expr, whereFunc ...func(gen.Dao) gen.Dao) ([]*model.StSite, int64, error) {
	m.ctrl.T.Helper()
	varargs := []any{page, pageSize, orderColumns}
	for _, a := range whereFunc {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FindPage", varargs...)
	ret0, _ := ret[0].([]*model.StSite)
	ret1, _ := ret[1].(int64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// FindPage indicates an expected call of FindPage.
func (mr *MockiCustomGenStSiteFuncMockRecorder) FindPage(page, pageSize, orderColumns any, whereFunc ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{page, pageSize, orderColumns}, whereFunc...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindPage", reflect.TypeOf((*MockiCustomGenStSiteFunc)(nil).FindPage), varargs...)
}

// FindSiteCategoryWithPage mocks base method.
func (m *MockiCustomGenStSiteFunc) FindSiteCategoryWithPage(page, pageSize int, result any, orderColumns []field.Expr, whereFunc ...func(gen.Dao) gen.Dao) (int64, error) {
	m.ctrl.T.Helper()
	varargs := []any{page, pageSize, result, orderColumns}
	for _, a := range whereFunc {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FindSiteCategoryWithPage", varargs...)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindSiteCategoryWithPage indicates an expected call of FindSiteCategoryWithPage.
func (mr *MockiCustomGenStSiteFuncMockRecorder) FindSiteCategoryWithPage(page, pageSize, result, orderColumns any, whereFunc ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{page, pageSize, result, orderColumns}, whereFunc...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindSiteCategoryWithPage", reflect.TypeOf((*MockiCustomGenStSiteFunc)(nil).FindSiteCategoryWithPage), varargs...)
}

// Scan mocks base method.
func (m *MockiCustomGenStSiteFunc) Scan(result any, whereFunc ...func(gen.Dao) gen.Dao) error {
	m.ctrl.T.Helper()
	varargs := []any{result}
	for _, a := range whereFunc {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Scan", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Scan indicates an expected call of Scan.
func (mr *MockiCustomGenStSiteFuncMockRecorder) Scan(result any, whereFunc ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{result}, whereFunc...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Scan", reflect.TypeOf((*MockiCustomGenStSiteFunc)(nil).Scan), varargs...)
}

// ScanPage mocks base method.
func (m *MockiCustomGenStSiteFunc) ScanPage(page, pageSize int, orderColumns []field.Expr, result any, whereFunc ...func(gen.Dao) gen.Dao) (int64, error) {
	m.ctrl.T.Helper()
	varargs := []any{page, pageSize, orderColumns, result}
	for _, a := range whereFunc {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ScanPage", varargs...)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ScanPage indicates an expected call of ScanPage.
func (mr *MockiCustomGenStSiteFuncMockRecorder) ScanPage(page, pageSize, orderColumns, result any, whereFunc ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{page, pageSize, orderColumns, result}, whereFunc...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ScanPage", reflect.TypeOf((*MockiCustomGenStSiteFunc)(nil).ScanPage), varargs...)
}

// Update mocks base method.
func (m *MockiCustomGenStSiteFunc) Update(v any, whereFunc ...func(gen.Dao) gen.Dao) (int64, error) {
	m.ctrl.T.Helper()
	varargs := []any{v}
	for _, a := range whereFunc {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Update", varargs...)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockiCustomGenStSiteFuncMockRecorder) Update(v any, whereFunc ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{v}, whereFunc...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockiCustomGenStSiteFunc)(nil).Update), varargs...)
}

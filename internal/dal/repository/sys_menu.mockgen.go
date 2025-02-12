// Code generated by MockGen. DO NOT EDIT.
// Source: internal/dal/repository/sys_menu.go
//
// Generated by this command:
//
//	mockgen -source=internal/dal/repository/sys_menu.go -destination internal/dal/repository/sys_menu.mockgen.go -package=repository
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

// MockISysMenuDao is a mock of ISysMenuDao interface.
type MockISysMenuDao struct {
	ctrl     *gomock.Controller
	recorder *MockISysMenuDaoMockRecorder
	isgomock struct{}
}

// MockISysMenuDaoMockRecorder is the mock recorder for MockISysMenuDao.
type MockISysMenuDaoMockRecorder struct {
	mock *MockISysMenuDao
}

// NewMockISysMenuDao creates a new mock instance.
func NewMockISysMenuDao(ctrl *gomock.Controller) *MockISysMenuDao {
	mock := &MockISysMenuDao{ctrl: ctrl}
	mock.recorder = &MockISysMenuDaoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockISysMenuDao) EXPECT() *MockISysMenuDaoMockRecorder {
	return m.recorder
}

// WhereByCreatedAt mocks base method.
func (m *MockISysMenuDao) WhereByCreatedAt(createdAt time.Time) func(gen.Dao) gen.Dao {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WhereByCreatedAt", createdAt)
	ret0, _ := ret[0].(func(gen.Dao) gen.Dao)
	return ret0
}

// WhereByCreatedAt indicates an expected call of WhereByCreatedAt.
func (mr *MockISysMenuDaoMockRecorder) WhereByCreatedAt(createdAt any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WhereByCreatedAt", reflect.TypeOf((*MockISysMenuDao)(nil).WhereByCreatedAt), createdAt)
}

// WhereByDeletedAt mocks base method.
func (m *MockISysMenuDao) WhereByDeletedAt(deletedAt time.Time) func(gen.Dao) gen.Dao {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WhereByDeletedAt", deletedAt)
	ret0, _ := ret[0].(func(gen.Dao) gen.Dao)
	return ret0
}

// WhereByDeletedAt indicates an expected call of WhereByDeletedAt.
func (mr *MockISysMenuDaoMockRecorder) WhereByDeletedAt(deletedAt any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WhereByDeletedAt", reflect.TypeOf((*MockISysMenuDao)(nil).WhereByDeletedAt), deletedAt)
}

// WhereByID mocks base method.
func (m *MockISysMenuDao) WhereByID(id int) func(gen.Dao) gen.Dao {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WhereByID", id)
	ret0, _ := ret[0].(func(gen.Dao) gen.Dao)
	return ret0
}

// WhereByID indicates an expected call of WhereByID.
func (mr *MockISysMenuDaoMockRecorder) WhereByID(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WhereByID", reflect.TypeOf((*MockISysMenuDao)(nil).WhereByID), id)
}

// WhereByIcon mocks base method.
func (m *MockISysMenuDao) WhereByIcon(icon string) func(gen.Dao) gen.Dao {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WhereByIcon", icon)
	ret0, _ := ret[0].(func(gen.Dao) gen.Dao)
	return ret0
}

// WhereByIcon indicates an expected call of WhereByIcon.
func (mr *MockISysMenuDaoMockRecorder) WhereByIcon(icon any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WhereByIcon", reflect.TypeOf((*MockISysMenuDao)(nil).WhereByIcon), icon)
}

// WhereByIsUsed mocks base method.
func (m *MockISysMenuDao) WhereByIsUsed(isUsed bool) func(gen.Dao) gen.Dao {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WhereByIsUsed", isUsed)
	ret0, _ := ret[0].(func(gen.Dao) gen.Dao)
	return ret0
}

// WhereByIsUsed indicates an expected call of WhereByIsUsed.
func (mr *MockISysMenuDaoMockRecorder) WhereByIsUsed(isUsed any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WhereByIsUsed", reflect.TypeOf((*MockISysMenuDao)(nil).WhereByIsUsed), isUsed)
}

// WhereByLevel mocks base method.
func (m *MockISysMenuDao) WhereByLevel(level int) func(gen.Dao) gen.Dao {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WhereByLevel", level)
	ret0, _ := ret[0].(func(gen.Dao) gen.Dao)
	return ret0
}

// WhereByLevel indicates an expected call of WhereByLevel.
func (mr *MockISysMenuDaoMockRecorder) WhereByLevel(level any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WhereByLevel", reflect.TypeOf((*MockISysMenuDao)(nil).WhereByLevel), level)
}

// WhereByLink mocks base method.
func (m *MockISysMenuDao) WhereByLink(link string) func(gen.Dao) gen.Dao {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WhereByLink", link)
	ret0, _ := ret[0].(func(gen.Dao) gen.Dao)
	return ret0
}

// WhereByLink indicates an expected call of WhereByLink.
func (mr *MockISysMenuDaoMockRecorder) WhereByLink(link any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WhereByLink", reflect.TypeOf((*MockISysMenuDao)(nil).WhereByLink), link)
}

// WhereByName mocks base method.
func (m *MockISysMenuDao) WhereByName(name string) func(gen.Dao) gen.Dao {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WhereByName", name)
	ret0, _ := ret[0].(func(gen.Dao) gen.Dao)
	return ret0
}

// WhereByName indicates an expected call of WhereByName.
func (mr *MockISysMenuDaoMockRecorder) WhereByName(name any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WhereByName", reflect.TypeOf((*MockISysMenuDao)(nil).WhereByName), name)
}

// WhereByPid mocks base method.
func (m *MockISysMenuDao) WhereByPid(pid int) func(gen.Dao) gen.Dao {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WhereByPid", pid)
	ret0, _ := ret[0].(func(gen.Dao) gen.Dao)
	return ret0
}

// WhereByPid indicates an expected call of WhereByPid.
func (mr *MockISysMenuDaoMockRecorder) WhereByPid(pid any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WhereByPid", reflect.TypeOf((*MockISysMenuDao)(nil).WhereByPid), pid)
}

// WhereBySort mocks base method.
func (m *MockISysMenuDao) WhereBySort(sort int) func(gen.Dao) gen.Dao {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WhereBySort", sort)
	ret0, _ := ret[0].(func(gen.Dao) gen.Dao)
	return ret0
}

// WhereBySort indicates an expected call of WhereBySort.
func (mr *MockISysMenuDaoMockRecorder) WhereBySort(sort any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WhereBySort", reflect.TypeOf((*MockISysMenuDao)(nil).WhereBySort), sort)
}

// WhereByUpdatedAt mocks base method.
func (m *MockISysMenuDao) WhereByUpdatedAt(updatedAt time.Time) func(gen.Dao) gen.Dao {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WhereByUpdatedAt", updatedAt)
	ret0, _ := ret[0].(func(gen.Dao) gen.Dao)
	return ret0
}

// WhereByUpdatedAt indicates an expected call of WhereByUpdatedAt.
func (mr *MockISysMenuDaoMockRecorder) WhereByUpdatedAt(updatedAt any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WhereByUpdatedAt", reflect.TypeOf((*MockISysMenuDao)(nil).WhereByUpdatedAt), updatedAt)
}

// WithContext mocks base method.
func (m *MockISysMenuDao) WithContext(ctx context.Context) iCustomGenSysMenuFunc {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithContext", ctx)
	ret0, _ := ret[0].(iCustomGenSysMenuFunc)
	return ret0
}

// WithContext indicates an expected call of WithContext.
func (mr *MockISysMenuDaoMockRecorder) WithContext(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithContext", reflect.TypeOf((*MockISysMenuDao)(nil).WithContext), ctx)
}

// MockiCustomGenSysMenuFunc is a mock of iCustomGenSysMenuFunc interface.
type MockiCustomGenSysMenuFunc struct {
	ctrl     *gomock.Controller
	recorder *MockiCustomGenSysMenuFuncMockRecorder
	isgomock struct{}
}

// MockiCustomGenSysMenuFuncMockRecorder is the mock recorder for MockiCustomGenSysMenuFunc.
type MockiCustomGenSysMenuFuncMockRecorder struct {
	mock *MockiCustomGenSysMenuFunc
}

// NewMockiCustomGenSysMenuFunc creates a new mock instance.
func NewMockiCustomGenSysMenuFunc(ctrl *gomock.Controller) *MockiCustomGenSysMenuFunc {
	mock := &MockiCustomGenSysMenuFunc{ctrl: ctrl}
	mock.recorder = &MockiCustomGenSysMenuFuncMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockiCustomGenSysMenuFunc) EXPECT() *MockiCustomGenSysMenuFuncMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m_2 *MockiCustomGenSysMenuFunc) Create(m *model.SysMenu) (*model.SysMenu, error) {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "Create", m)
	ret0, _ := ret[0].(*model.SysMenu)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockiCustomGenSysMenuFuncMockRecorder) Create(m any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockiCustomGenSysMenuFunc)(nil).Create), m)
}

// Delete mocks base method.
func (m *MockiCustomGenSysMenuFunc) Delete(whereFunc ...func(gen.Dao) gen.Dao) error {
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
func (mr *MockiCustomGenSysMenuFuncMockRecorder) Delete(whereFunc ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockiCustomGenSysMenuFunc)(nil).Delete), whereFunc...)
}

// DeletePhysical mocks base method.
func (m *MockiCustomGenSysMenuFunc) DeletePhysical(whereFunc ...func(gen.Dao) gen.Dao) error {
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
func (mr *MockiCustomGenSysMenuFuncMockRecorder) DeletePhysical(whereFunc ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePhysical", reflect.TypeOf((*MockiCustomGenSysMenuFunc)(nil).DeletePhysical), whereFunc...)
}

// FindAll mocks base method.
func (m *MockiCustomGenSysMenuFunc) FindAll(whereFunc ...func(gen.Dao) gen.Dao) ([]*model.SysMenu, error) {
	m.ctrl.T.Helper()
	varargs := []any{}
	for _, a := range whereFunc {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FindAll", varargs...)
	ret0, _ := ret[0].([]*model.SysMenu)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAll indicates an expected call of FindAll.
func (mr *MockiCustomGenSysMenuFuncMockRecorder) FindAll(whereFunc ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockiCustomGenSysMenuFunc)(nil).FindAll), whereFunc...)
}

// FindCount mocks base method.
func (m *MockiCustomGenSysMenuFunc) FindCount(whereFunc ...func(gen.Dao) gen.Dao) (int64, error) {
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
func (mr *MockiCustomGenSysMenuFuncMockRecorder) FindCount(whereFunc ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindCount", reflect.TypeOf((*MockiCustomGenSysMenuFunc)(nil).FindCount), whereFunc...)
}

// FindOne mocks base method.
func (m *MockiCustomGenSysMenuFunc) FindOne(whereFunc ...func(gen.Dao) gen.Dao) (*model.SysMenu, error) {
	m.ctrl.T.Helper()
	varargs := []any{}
	for _, a := range whereFunc {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FindOne", varargs...)
	ret0, _ := ret[0].(*model.SysMenu)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOne indicates an expected call of FindOne.
func (mr *MockiCustomGenSysMenuFuncMockRecorder) FindOne(whereFunc ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOne", reflect.TypeOf((*MockiCustomGenSysMenuFunc)(nil).FindOne), whereFunc...)
}

// FindPage mocks base method.
func (m *MockiCustomGenSysMenuFunc) FindPage(page, pageSize int, orderColumns []field.Expr, whereFunc ...func(gen.Dao) gen.Dao) ([]*model.SysMenu, int64, error) {
	m.ctrl.T.Helper()
	varargs := []any{page, pageSize, orderColumns}
	for _, a := range whereFunc {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FindPage", varargs...)
	ret0, _ := ret[0].([]*model.SysMenu)
	ret1, _ := ret[1].(int64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// FindPage indicates an expected call of FindPage.
func (mr *MockiCustomGenSysMenuFuncMockRecorder) FindPage(page, pageSize, orderColumns any, whereFunc ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{page, pageSize, orderColumns}, whereFunc...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindPage", reflect.TypeOf((*MockiCustomGenSysMenuFunc)(nil).FindPage), varargs...)
}

// Scan mocks base method.
func (m *MockiCustomGenSysMenuFunc) Scan(result any, whereFunc ...func(gen.Dao) gen.Dao) error {
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
func (mr *MockiCustomGenSysMenuFuncMockRecorder) Scan(result any, whereFunc ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{result}, whereFunc...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Scan", reflect.TypeOf((*MockiCustomGenSysMenuFunc)(nil).Scan), varargs...)
}

// ScanPage mocks base method.
func (m *MockiCustomGenSysMenuFunc) ScanPage(page, pageSize int, orderColumns []field.Expr, result any, whereFunc ...func(gen.Dao) gen.Dao) (int64, error) {
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
func (mr *MockiCustomGenSysMenuFuncMockRecorder) ScanPage(page, pageSize, orderColumns, result any, whereFunc ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{page, pageSize, orderColumns, result}, whereFunc...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ScanPage", reflect.TypeOf((*MockiCustomGenSysMenuFunc)(nil).ScanPage), varargs...)
}

// Update mocks base method.
func (m *MockiCustomGenSysMenuFunc) Update(v any, whereFunc ...func(gen.Dao) gen.Dao) (int64, error) {
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
func (mr *MockiCustomGenSysMenuFuncMockRecorder) Update(v any, whereFunc ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{v}, whereFunc...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockiCustomGenSysMenuFunc)(nil).Update), varargs...)
}

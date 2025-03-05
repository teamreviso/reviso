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

	"github.com/teamreviso/code/pkg/models"
)

func newWaitlistUser(db *gorm.DB, opts ...gen.DOOption) waitlistUser {
	_waitlistUser := waitlistUser{}

	_waitlistUser.waitlistUserDo.UseDB(db, opts...)
	_waitlistUser.waitlistUserDo.UseModel(&models.WaitlistUser{})

	tableName := _waitlistUser.waitlistUserDo.TableName()
	_waitlistUser.ALL = field.NewAsterisk(tableName)
	_waitlistUser.Email = field.NewString(tableName, "email")
	_waitlistUser.CreatedAt = field.NewTime(tableName, "created_at")
	_waitlistUser.UpdatedAt = field.NewTime(tableName, "updated_at")
	_waitlistUser.AllowAccess = field.NewBool(tableName, "allow_access")

	_waitlistUser.fillFieldMap()

	return _waitlistUser
}

type waitlistUser struct {
	waitlistUserDo

	ALL         field.Asterisk
	Email       field.String
	CreatedAt   field.Time
	UpdatedAt   field.Time
	AllowAccess field.Bool

	fieldMap map[string]field.Expr
}

func (w waitlistUser) Table(newTableName string) *waitlistUser {
	w.waitlistUserDo.UseTable(newTableName)
	return w.updateTableName(newTableName)
}

func (w waitlistUser) As(alias string) *waitlistUser {
	w.waitlistUserDo.DO = *(w.waitlistUserDo.As(alias).(*gen.DO))
	return w.updateTableName(alias)
}

func (w *waitlistUser) updateTableName(table string) *waitlistUser {
	w.ALL = field.NewAsterisk(table)
	w.Email = field.NewString(table, "email")
	w.CreatedAt = field.NewTime(table, "created_at")
	w.UpdatedAt = field.NewTime(table, "updated_at")
	w.AllowAccess = field.NewBool(table, "allow_access")

	w.fillFieldMap()

	return w
}

func (w *waitlistUser) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := w.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (w *waitlistUser) fillFieldMap() {
	w.fieldMap = make(map[string]field.Expr, 4)
	w.fieldMap["email"] = w.Email
	w.fieldMap["created_at"] = w.CreatedAt
	w.fieldMap["updated_at"] = w.UpdatedAt
	w.fieldMap["allow_access"] = w.AllowAccess
}

func (w waitlistUser) clone(db *gorm.DB) waitlistUser {
	w.waitlistUserDo.ReplaceConnPool(db.Statement.ConnPool)
	return w
}

func (w waitlistUser) replaceDB(db *gorm.DB) waitlistUser {
	w.waitlistUserDo.ReplaceDB(db)
	return w
}

type waitlistUserDo struct{ gen.DO }

type IWaitlistUserDo interface {
	gen.SubQuery
	Debug() IWaitlistUserDo
	WithContext(ctx context.Context) IWaitlistUserDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IWaitlistUserDo
	WriteDB() IWaitlistUserDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IWaitlistUserDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IWaitlistUserDo
	Not(conds ...gen.Condition) IWaitlistUserDo
	Or(conds ...gen.Condition) IWaitlistUserDo
	Select(conds ...field.Expr) IWaitlistUserDo
	Where(conds ...gen.Condition) IWaitlistUserDo
	Order(conds ...field.Expr) IWaitlistUserDo
	Distinct(cols ...field.Expr) IWaitlistUserDo
	Omit(cols ...field.Expr) IWaitlistUserDo
	Join(table schema.Tabler, on ...field.Expr) IWaitlistUserDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IWaitlistUserDo
	RightJoin(table schema.Tabler, on ...field.Expr) IWaitlistUserDo
	Group(cols ...field.Expr) IWaitlistUserDo
	Having(conds ...gen.Condition) IWaitlistUserDo
	Limit(limit int) IWaitlistUserDo
	Offset(offset int) IWaitlistUserDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IWaitlistUserDo
	Unscoped() IWaitlistUserDo
	Create(values ...*models.WaitlistUser) error
	CreateInBatches(values []*models.WaitlistUser, batchSize int) error
	Save(values ...*models.WaitlistUser) error
	First() (*models.WaitlistUser, error)
	Take() (*models.WaitlistUser, error)
	Last() (*models.WaitlistUser, error)
	Find() ([]*models.WaitlistUser, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.WaitlistUser, err error)
	FindInBatches(result *[]*models.WaitlistUser, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*models.WaitlistUser) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IWaitlistUserDo
	Assign(attrs ...field.AssignExpr) IWaitlistUserDo
	Joins(fields ...field.RelationField) IWaitlistUserDo
	Preload(fields ...field.RelationField) IWaitlistUserDo
	FirstOrInit() (*models.WaitlistUser, error)
	FirstOrCreate() (*models.WaitlistUser, error)
	FindByPage(offset int, limit int) (result []*models.WaitlistUser, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IWaitlistUserDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (w waitlistUserDo) Debug() IWaitlistUserDo {
	return w.withDO(w.DO.Debug())
}

func (w waitlistUserDo) WithContext(ctx context.Context) IWaitlistUserDo {
	return w.withDO(w.DO.WithContext(ctx))
}

func (w waitlistUserDo) ReadDB() IWaitlistUserDo {
	return w.Clauses(dbresolver.Read)
}

func (w waitlistUserDo) WriteDB() IWaitlistUserDo {
	return w.Clauses(dbresolver.Write)
}

func (w waitlistUserDo) Session(config *gorm.Session) IWaitlistUserDo {
	return w.withDO(w.DO.Session(config))
}

func (w waitlistUserDo) Clauses(conds ...clause.Expression) IWaitlistUserDo {
	return w.withDO(w.DO.Clauses(conds...))
}

func (w waitlistUserDo) Returning(value interface{}, columns ...string) IWaitlistUserDo {
	return w.withDO(w.DO.Returning(value, columns...))
}

func (w waitlistUserDo) Not(conds ...gen.Condition) IWaitlistUserDo {
	return w.withDO(w.DO.Not(conds...))
}

func (w waitlistUserDo) Or(conds ...gen.Condition) IWaitlistUserDo {
	return w.withDO(w.DO.Or(conds...))
}

func (w waitlistUserDo) Select(conds ...field.Expr) IWaitlistUserDo {
	return w.withDO(w.DO.Select(conds...))
}

func (w waitlistUserDo) Where(conds ...gen.Condition) IWaitlistUserDo {
	return w.withDO(w.DO.Where(conds...))
}

func (w waitlistUserDo) Order(conds ...field.Expr) IWaitlistUserDo {
	return w.withDO(w.DO.Order(conds...))
}

func (w waitlistUserDo) Distinct(cols ...field.Expr) IWaitlistUserDo {
	return w.withDO(w.DO.Distinct(cols...))
}

func (w waitlistUserDo) Omit(cols ...field.Expr) IWaitlistUserDo {
	return w.withDO(w.DO.Omit(cols...))
}

func (w waitlistUserDo) Join(table schema.Tabler, on ...field.Expr) IWaitlistUserDo {
	return w.withDO(w.DO.Join(table, on...))
}

func (w waitlistUserDo) LeftJoin(table schema.Tabler, on ...field.Expr) IWaitlistUserDo {
	return w.withDO(w.DO.LeftJoin(table, on...))
}

func (w waitlistUserDo) RightJoin(table schema.Tabler, on ...field.Expr) IWaitlistUserDo {
	return w.withDO(w.DO.RightJoin(table, on...))
}

func (w waitlistUserDo) Group(cols ...field.Expr) IWaitlistUserDo {
	return w.withDO(w.DO.Group(cols...))
}

func (w waitlistUserDo) Having(conds ...gen.Condition) IWaitlistUserDo {
	return w.withDO(w.DO.Having(conds...))
}

func (w waitlistUserDo) Limit(limit int) IWaitlistUserDo {
	return w.withDO(w.DO.Limit(limit))
}

func (w waitlistUserDo) Offset(offset int) IWaitlistUserDo {
	return w.withDO(w.DO.Offset(offset))
}

func (w waitlistUserDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IWaitlistUserDo {
	return w.withDO(w.DO.Scopes(funcs...))
}

func (w waitlistUserDo) Unscoped() IWaitlistUserDo {
	return w.withDO(w.DO.Unscoped())
}

func (w waitlistUserDo) Create(values ...*models.WaitlistUser) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Create(values)
}

func (w waitlistUserDo) CreateInBatches(values []*models.WaitlistUser, batchSize int) error {
	return w.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (w waitlistUserDo) Save(values ...*models.WaitlistUser) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Save(values)
}

func (w waitlistUserDo) First() (*models.WaitlistUser, error) {
	if result, err := w.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*models.WaitlistUser), nil
	}
}

func (w waitlistUserDo) Take() (*models.WaitlistUser, error) {
	if result, err := w.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*models.WaitlistUser), nil
	}
}

func (w waitlistUserDo) Last() (*models.WaitlistUser, error) {
	if result, err := w.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*models.WaitlistUser), nil
	}
}

func (w waitlistUserDo) Find() ([]*models.WaitlistUser, error) {
	result, err := w.DO.Find()
	return result.([]*models.WaitlistUser), err
}

func (w waitlistUserDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.WaitlistUser, err error) {
	buf := make([]*models.WaitlistUser, 0, batchSize)
	err = w.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (w waitlistUserDo) FindInBatches(result *[]*models.WaitlistUser, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return w.DO.FindInBatches(result, batchSize, fc)
}

func (w waitlistUserDo) Attrs(attrs ...field.AssignExpr) IWaitlistUserDo {
	return w.withDO(w.DO.Attrs(attrs...))
}

func (w waitlistUserDo) Assign(attrs ...field.AssignExpr) IWaitlistUserDo {
	return w.withDO(w.DO.Assign(attrs...))
}

func (w waitlistUserDo) Joins(fields ...field.RelationField) IWaitlistUserDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Joins(_f))
	}
	return &w
}

func (w waitlistUserDo) Preload(fields ...field.RelationField) IWaitlistUserDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Preload(_f))
	}
	return &w
}

func (w waitlistUserDo) FirstOrInit() (*models.WaitlistUser, error) {
	if result, err := w.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*models.WaitlistUser), nil
	}
}

func (w waitlistUserDo) FirstOrCreate() (*models.WaitlistUser, error) {
	if result, err := w.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*models.WaitlistUser), nil
	}
}

func (w waitlistUserDo) FindByPage(offset int, limit int) (result []*models.WaitlistUser, count int64, err error) {
	result, err = w.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = w.Offset(-1).Limit(-1).Count()
	return
}

func (w waitlistUserDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = w.Count()
	if err != nil {
		return
	}

	err = w.Offset(offset).Limit(limit).Scan(result)
	return
}

func (w waitlistUserDo) Scan(result interface{}) (err error) {
	return w.DO.Scan(result)
}

func (w waitlistUserDo) Delete(models ...*models.WaitlistUser) (result gen.ResultInfo, err error) {
	return w.DO.Delete(models)
}

func (w *waitlistUserDo) withDO(do gen.Dao) *waitlistUserDo {
	w.DO = *do.(*gen.DO)
	return w
}

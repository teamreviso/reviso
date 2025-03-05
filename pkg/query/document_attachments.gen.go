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

func newDocumentAttachment(db *gorm.DB, opts ...gen.DOOption) documentAttachment {
	_documentAttachment := documentAttachment{}

	_documentAttachment.documentAttachmentDo.UseDB(db, opts...)
	_documentAttachment.documentAttachmentDo.UseModel(&models.DocumentAttachment{})

	tableName := _documentAttachment.documentAttachmentDo.TableName()
	_documentAttachment.ALL = field.NewAsterisk(tableName)
	_documentAttachment.ID = field.NewString(tableName, "id")
	_documentAttachment.UserID = field.NewString(tableName, "user_id")
	_documentAttachment.DocumentID = field.NewString(tableName, "document_id")
	_documentAttachment.S3ID = field.NewString(tableName, "s3_id")
	_documentAttachment.Filename = field.NewString(tableName, "filename")
	_documentAttachment.ContentType = field.NewString(tableName, "content_type")
	_documentAttachment.Size = field.NewInt64(tableName, "size")
	_documentAttachment.CreatedAt = field.NewTime(tableName, "created_at")
	_documentAttachment.UpdatedAt = field.NewTime(tableName, "updated_at")

	_documentAttachment.fillFieldMap()

	return _documentAttachment
}

type documentAttachment struct {
	documentAttachmentDo

	ALL         field.Asterisk
	ID          field.String
	UserID      field.String
	DocumentID  field.String
	S3ID        field.String
	Filename    field.String
	ContentType field.String
	Size        field.Int64
	CreatedAt   field.Time
	UpdatedAt   field.Time

	fieldMap map[string]field.Expr
}

func (d documentAttachment) Table(newTableName string) *documentAttachment {
	d.documentAttachmentDo.UseTable(newTableName)
	return d.updateTableName(newTableName)
}

func (d documentAttachment) As(alias string) *documentAttachment {
	d.documentAttachmentDo.DO = *(d.documentAttachmentDo.As(alias).(*gen.DO))
	return d.updateTableName(alias)
}

func (d *documentAttachment) updateTableName(table string) *documentAttachment {
	d.ALL = field.NewAsterisk(table)
	d.ID = field.NewString(table, "id")
	d.UserID = field.NewString(table, "user_id")
	d.DocumentID = field.NewString(table, "document_id")
	d.S3ID = field.NewString(table, "s3_id")
	d.Filename = field.NewString(table, "filename")
	d.ContentType = field.NewString(table, "content_type")
	d.Size = field.NewInt64(table, "size")
	d.CreatedAt = field.NewTime(table, "created_at")
	d.UpdatedAt = field.NewTime(table, "updated_at")

	d.fillFieldMap()

	return d
}

func (d *documentAttachment) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := d.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (d *documentAttachment) fillFieldMap() {
	d.fieldMap = make(map[string]field.Expr, 9)
	d.fieldMap["id"] = d.ID
	d.fieldMap["user_id"] = d.UserID
	d.fieldMap["document_id"] = d.DocumentID
	d.fieldMap["s3_id"] = d.S3ID
	d.fieldMap["filename"] = d.Filename
	d.fieldMap["content_type"] = d.ContentType
	d.fieldMap["size"] = d.Size
	d.fieldMap["created_at"] = d.CreatedAt
	d.fieldMap["updated_at"] = d.UpdatedAt
}

func (d documentAttachment) clone(db *gorm.DB) documentAttachment {
	d.documentAttachmentDo.ReplaceConnPool(db.Statement.ConnPool)
	return d
}

func (d documentAttachment) replaceDB(db *gorm.DB) documentAttachment {
	d.documentAttachmentDo.ReplaceDB(db)
	return d
}

type documentAttachmentDo struct{ gen.DO }

type IDocumentAttachmentDo interface {
	gen.SubQuery
	Debug() IDocumentAttachmentDo
	WithContext(ctx context.Context) IDocumentAttachmentDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IDocumentAttachmentDo
	WriteDB() IDocumentAttachmentDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IDocumentAttachmentDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IDocumentAttachmentDo
	Not(conds ...gen.Condition) IDocumentAttachmentDo
	Or(conds ...gen.Condition) IDocumentAttachmentDo
	Select(conds ...field.Expr) IDocumentAttachmentDo
	Where(conds ...gen.Condition) IDocumentAttachmentDo
	Order(conds ...field.Expr) IDocumentAttachmentDo
	Distinct(cols ...field.Expr) IDocumentAttachmentDo
	Omit(cols ...field.Expr) IDocumentAttachmentDo
	Join(table schema.Tabler, on ...field.Expr) IDocumentAttachmentDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IDocumentAttachmentDo
	RightJoin(table schema.Tabler, on ...field.Expr) IDocumentAttachmentDo
	Group(cols ...field.Expr) IDocumentAttachmentDo
	Having(conds ...gen.Condition) IDocumentAttachmentDo
	Limit(limit int) IDocumentAttachmentDo
	Offset(offset int) IDocumentAttachmentDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IDocumentAttachmentDo
	Unscoped() IDocumentAttachmentDo
	Create(values ...*models.DocumentAttachment) error
	CreateInBatches(values []*models.DocumentAttachment, batchSize int) error
	Save(values ...*models.DocumentAttachment) error
	First() (*models.DocumentAttachment, error)
	Take() (*models.DocumentAttachment, error)
	Last() (*models.DocumentAttachment, error)
	Find() ([]*models.DocumentAttachment, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.DocumentAttachment, err error)
	FindInBatches(result *[]*models.DocumentAttachment, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*models.DocumentAttachment) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IDocumentAttachmentDo
	Assign(attrs ...field.AssignExpr) IDocumentAttachmentDo
	Joins(fields ...field.RelationField) IDocumentAttachmentDo
	Preload(fields ...field.RelationField) IDocumentAttachmentDo
	FirstOrInit() (*models.DocumentAttachment, error)
	FirstOrCreate() (*models.DocumentAttachment, error)
	FindByPage(offset int, limit int) (result []*models.DocumentAttachment, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IDocumentAttachmentDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (d documentAttachmentDo) Debug() IDocumentAttachmentDo {
	return d.withDO(d.DO.Debug())
}

func (d documentAttachmentDo) WithContext(ctx context.Context) IDocumentAttachmentDo {
	return d.withDO(d.DO.WithContext(ctx))
}

func (d documentAttachmentDo) ReadDB() IDocumentAttachmentDo {
	return d.Clauses(dbresolver.Read)
}

func (d documentAttachmentDo) WriteDB() IDocumentAttachmentDo {
	return d.Clauses(dbresolver.Write)
}

func (d documentAttachmentDo) Session(config *gorm.Session) IDocumentAttachmentDo {
	return d.withDO(d.DO.Session(config))
}

func (d documentAttachmentDo) Clauses(conds ...clause.Expression) IDocumentAttachmentDo {
	return d.withDO(d.DO.Clauses(conds...))
}

func (d documentAttachmentDo) Returning(value interface{}, columns ...string) IDocumentAttachmentDo {
	return d.withDO(d.DO.Returning(value, columns...))
}

func (d documentAttachmentDo) Not(conds ...gen.Condition) IDocumentAttachmentDo {
	return d.withDO(d.DO.Not(conds...))
}

func (d documentAttachmentDo) Or(conds ...gen.Condition) IDocumentAttachmentDo {
	return d.withDO(d.DO.Or(conds...))
}

func (d documentAttachmentDo) Select(conds ...field.Expr) IDocumentAttachmentDo {
	return d.withDO(d.DO.Select(conds...))
}

func (d documentAttachmentDo) Where(conds ...gen.Condition) IDocumentAttachmentDo {
	return d.withDO(d.DO.Where(conds...))
}

func (d documentAttachmentDo) Order(conds ...field.Expr) IDocumentAttachmentDo {
	return d.withDO(d.DO.Order(conds...))
}

func (d documentAttachmentDo) Distinct(cols ...field.Expr) IDocumentAttachmentDo {
	return d.withDO(d.DO.Distinct(cols...))
}

func (d documentAttachmentDo) Omit(cols ...field.Expr) IDocumentAttachmentDo {
	return d.withDO(d.DO.Omit(cols...))
}

func (d documentAttachmentDo) Join(table schema.Tabler, on ...field.Expr) IDocumentAttachmentDo {
	return d.withDO(d.DO.Join(table, on...))
}

func (d documentAttachmentDo) LeftJoin(table schema.Tabler, on ...field.Expr) IDocumentAttachmentDo {
	return d.withDO(d.DO.LeftJoin(table, on...))
}

func (d documentAttachmentDo) RightJoin(table schema.Tabler, on ...field.Expr) IDocumentAttachmentDo {
	return d.withDO(d.DO.RightJoin(table, on...))
}

func (d documentAttachmentDo) Group(cols ...field.Expr) IDocumentAttachmentDo {
	return d.withDO(d.DO.Group(cols...))
}

func (d documentAttachmentDo) Having(conds ...gen.Condition) IDocumentAttachmentDo {
	return d.withDO(d.DO.Having(conds...))
}

func (d documentAttachmentDo) Limit(limit int) IDocumentAttachmentDo {
	return d.withDO(d.DO.Limit(limit))
}

func (d documentAttachmentDo) Offset(offset int) IDocumentAttachmentDo {
	return d.withDO(d.DO.Offset(offset))
}

func (d documentAttachmentDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IDocumentAttachmentDo {
	return d.withDO(d.DO.Scopes(funcs...))
}

func (d documentAttachmentDo) Unscoped() IDocumentAttachmentDo {
	return d.withDO(d.DO.Unscoped())
}

func (d documentAttachmentDo) Create(values ...*models.DocumentAttachment) error {
	if len(values) == 0 {
		return nil
	}
	return d.DO.Create(values)
}

func (d documentAttachmentDo) CreateInBatches(values []*models.DocumentAttachment, batchSize int) error {
	return d.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (d documentAttachmentDo) Save(values ...*models.DocumentAttachment) error {
	if len(values) == 0 {
		return nil
	}
	return d.DO.Save(values)
}

func (d documentAttachmentDo) First() (*models.DocumentAttachment, error) {
	if result, err := d.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*models.DocumentAttachment), nil
	}
}

func (d documentAttachmentDo) Take() (*models.DocumentAttachment, error) {
	if result, err := d.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*models.DocumentAttachment), nil
	}
}

func (d documentAttachmentDo) Last() (*models.DocumentAttachment, error) {
	if result, err := d.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*models.DocumentAttachment), nil
	}
}

func (d documentAttachmentDo) Find() ([]*models.DocumentAttachment, error) {
	result, err := d.DO.Find()
	return result.([]*models.DocumentAttachment), err
}

func (d documentAttachmentDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.DocumentAttachment, err error) {
	buf := make([]*models.DocumentAttachment, 0, batchSize)
	err = d.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (d documentAttachmentDo) FindInBatches(result *[]*models.DocumentAttachment, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return d.DO.FindInBatches(result, batchSize, fc)
}

func (d documentAttachmentDo) Attrs(attrs ...field.AssignExpr) IDocumentAttachmentDo {
	return d.withDO(d.DO.Attrs(attrs...))
}

func (d documentAttachmentDo) Assign(attrs ...field.AssignExpr) IDocumentAttachmentDo {
	return d.withDO(d.DO.Assign(attrs...))
}

func (d documentAttachmentDo) Joins(fields ...field.RelationField) IDocumentAttachmentDo {
	for _, _f := range fields {
		d = *d.withDO(d.DO.Joins(_f))
	}
	return &d
}

func (d documentAttachmentDo) Preload(fields ...field.RelationField) IDocumentAttachmentDo {
	for _, _f := range fields {
		d = *d.withDO(d.DO.Preload(_f))
	}
	return &d
}

func (d documentAttachmentDo) FirstOrInit() (*models.DocumentAttachment, error) {
	if result, err := d.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*models.DocumentAttachment), nil
	}
}

func (d documentAttachmentDo) FirstOrCreate() (*models.DocumentAttachment, error) {
	if result, err := d.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*models.DocumentAttachment), nil
	}
}

func (d documentAttachmentDo) FindByPage(offset int, limit int) (result []*models.DocumentAttachment, count int64, err error) {
	result, err = d.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = d.Offset(-1).Limit(-1).Count()
	return
}

func (d documentAttachmentDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = d.Count()
	if err != nil {
		return
	}

	err = d.Offset(offset).Limit(limit).Scan(result)
	return
}

func (d documentAttachmentDo) Scan(result interface{}) (err error) {
	return d.DO.Scan(result)
}

func (d documentAttachmentDo) Delete(models ...*models.DocumentAttachment) (result gen.ResultInfo, err error) {
	return d.DO.Delete(models)
}

func (d *documentAttachmentDo) withDO(do gen.Dao) *documentAttachmentDo {
	d.DO = *do.(*gen.DO)
	return d
}

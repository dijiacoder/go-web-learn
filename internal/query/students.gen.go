// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"
	"github.com/dijiacoder/go-web-learn/internal/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"
)

func newStudent(db *gorm.DB, opts ...gen.DOOption) student {
	_student := student{}

	_student.studentDo.UseDB(db, opts...)
	_student.studentDo.UseModel(&model.Student{})

	tableName := _student.studentDo.TableName()
	_student.ALL = field.NewAsterisk(tableName)
	_student.ID = field.NewInt32(tableName, "id")
	_student.Name = field.NewString(tableName, "name")
	_student.IDCard = field.NewString(tableName, "id_card")
	_student.AdmissionDate = field.NewTime(tableName, "admission_date")
	_student.ClassID = field.NewInt32(tableName, "class_id")
	_student.Birthdate = field.NewTime(tableName, "birthdate")
	_student.ParentName = field.NewString(tableName, "parent_name")
	_student.ParentPhone = field.NewString(tableName, "parent_phone")
	_student.CreatedAt = field.NewTime(tableName, "created_at")
	_student.UpdatedAt = field.NewTime(tableName, "updated_at")
	_student.IsDel = field.NewInt32(tableName, "is_del")

	_student.fillFieldMap()

	return _student
}

type student struct {
	studentDo

	ALL           field.Asterisk
	ID            field.Int32
	Name          field.String // 学生姓名
	IDCard        field.String // 身份证号
	AdmissionDate field.Time   // 入学日期
	ClassID       field.Int32  // 所属班级
	Birthdate     field.Time   // 生日
	ParentName    field.String // 家长姓名
	ParentPhone   field.String // 家长联系电话
	CreatedAt     field.Time   // 创建时间
	UpdatedAt     field.Time   // 更新时间
	IsDel         field.Int32  // 是否删除（0: 未删除, 1: 已删除）

	fieldMap map[string]field.Expr
}

func (s student) Table(newTableName string) *student {
	s.studentDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s student) As(alias string) *student {
	s.studentDo.DO = *(s.studentDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *student) updateTableName(table string) *student {
	s.ALL = field.NewAsterisk(table)
	s.ID = field.NewInt32(table, "id")
	s.Name = field.NewString(table, "name")
	s.IDCard = field.NewString(table, "id_card")
	s.AdmissionDate = field.NewTime(table, "admission_date")
	s.ClassID = field.NewInt32(table, "class_id")
	s.Birthdate = field.NewTime(table, "birthdate")
	s.ParentName = field.NewString(table, "parent_name")
	s.ParentPhone = field.NewString(table, "parent_phone")
	s.CreatedAt = field.NewTime(table, "created_at")
	s.UpdatedAt = field.NewTime(table, "updated_at")
	s.IsDel = field.NewInt32(table, "is_del")

	s.fillFieldMap()

	return s
}

func (s *student) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *student) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 11)
	s.fieldMap["id"] = s.ID
	s.fieldMap["name"] = s.Name
	s.fieldMap["id_card"] = s.IDCard
	s.fieldMap["admission_date"] = s.AdmissionDate
	s.fieldMap["class_id"] = s.ClassID
	s.fieldMap["birthdate"] = s.Birthdate
	s.fieldMap["parent_name"] = s.ParentName
	s.fieldMap["parent_phone"] = s.ParentPhone
	s.fieldMap["created_at"] = s.CreatedAt
	s.fieldMap["updated_at"] = s.UpdatedAt
	s.fieldMap["is_del"] = s.IsDel
}

func (s student) clone(db *gorm.DB) student {
	s.studentDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s student) replaceDB(db *gorm.DB) student {
	s.studentDo.ReplaceDB(db)
	return s
}

type studentDo struct{ gen.DO }

type IStudentDo interface {
	gen.SubQuery
	Debug() IStudentDo
	WithContext(ctx context.Context) IStudentDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IStudentDo
	WriteDB() IStudentDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IStudentDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IStudentDo
	Not(conds ...gen.Condition) IStudentDo
	Or(conds ...gen.Condition) IStudentDo
	Select(conds ...field.Expr) IStudentDo
	Where(conds ...gen.Condition) IStudentDo
	Order(conds ...field.Expr) IStudentDo
	Distinct(cols ...field.Expr) IStudentDo
	Omit(cols ...field.Expr) IStudentDo
	Join(table schema.Tabler, on ...field.Expr) IStudentDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IStudentDo
	RightJoin(table schema.Tabler, on ...field.Expr) IStudentDo
	Group(cols ...field.Expr) IStudentDo
	Having(conds ...gen.Condition) IStudentDo
	Limit(limit int) IStudentDo
	Offset(offset int) IStudentDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IStudentDo
	Unscoped() IStudentDo
	Create(values ...*model.Student) error
	CreateInBatches(values []*model.Student, batchSize int) error
	Save(values ...*model.Student) error
	First() (*model.Student, error)
	Take() (*model.Student, error)
	Last() (*model.Student, error)
	Find() ([]*model.Student, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Student, err error)
	FindInBatches(result *[]*model.Student, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Student) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IStudentDo
	Assign(attrs ...field.AssignExpr) IStudentDo
	Joins(fields ...field.RelationField) IStudentDo
	Preload(fields ...field.RelationField) IStudentDo
	FirstOrInit() (*model.Student, error)
	FirstOrCreate() (*model.Student, error)
	FindByPage(offset int, limit int) (result []*model.Student, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IStudentDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s studentDo) Debug() IStudentDo {
	return s.withDO(s.DO.Debug())
}

func (s studentDo) WithContext(ctx context.Context) IStudentDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s studentDo) ReadDB() IStudentDo {
	return s.Clauses(dbresolver.Read)
}

func (s studentDo) WriteDB() IStudentDo {
	return s.Clauses(dbresolver.Write)
}

func (s studentDo) Session(config *gorm.Session) IStudentDo {
	return s.withDO(s.DO.Session(config))
}

func (s studentDo) Clauses(conds ...clause.Expression) IStudentDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s studentDo) Returning(value interface{}, columns ...string) IStudentDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s studentDo) Not(conds ...gen.Condition) IStudentDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s studentDo) Or(conds ...gen.Condition) IStudentDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s studentDo) Select(conds ...field.Expr) IStudentDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s studentDo) Where(conds ...gen.Condition) IStudentDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s studentDo) Order(conds ...field.Expr) IStudentDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s studentDo) Distinct(cols ...field.Expr) IStudentDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s studentDo) Omit(cols ...field.Expr) IStudentDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s studentDo) Join(table schema.Tabler, on ...field.Expr) IStudentDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s studentDo) LeftJoin(table schema.Tabler, on ...field.Expr) IStudentDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s studentDo) RightJoin(table schema.Tabler, on ...field.Expr) IStudentDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s studentDo) Group(cols ...field.Expr) IStudentDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s studentDo) Having(conds ...gen.Condition) IStudentDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s studentDo) Limit(limit int) IStudentDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s studentDo) Offset(offset int) IStudentDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s studentDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IStudentDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s studentDo) Unscoped() IStudentDo {
	return s.withDO(s.DO.Unscoped())
}

func (s studentDo) Create(values ...*model.Student) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s studentDo) CreateInBatches(values []*model.Student, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s studentDo) Save(values ...*model.Student) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s studentDo) First() (*model.Student, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Student), nil
	}
}

func (s studentDo) Take() (*model.Student, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Student), nil
	}
}

func (s studentDo) Last() (*model.Student, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Student), nil
	}
}

func (s studentDo) Find() ([]*model.Student, error) {
	result, err := s.DO.Find()
	return result.([]*model.Student), err
}

func (s studentDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Student, err error) {
	buf := make([]*model.Student, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s studentDo) FindInBatches(result *[]*model.Student, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s studentDo) Attrs(attrs ...field.AssignExpr) IStudentDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s studentDo) Assign(attrs ...field.AssignExpr) IStudentDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s studentDo) Joins(fields ...field.RelationField) IStudentDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s studentDo) Preload(fields ...field.RelationField) IStudentDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s studentDo) FirstOrInit() (*model.Student, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Student), nil
	}
}

func (s studentDo) FirstOrCreate() (*model.Student, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Student), nil
	}
}

func (s studentDo) FindByPage(offset int, limit int) (result []*model.Student, count int64, err error) {
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

func (s studentDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s studentDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s studentDo) Delete(models ...*model.Student) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *studentDo) withDO(do gen.Dao) *studentDo {
	s.DO = *do.(*gen.DO)
	return s
}

package dao

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"reflect"
	"strings"
)

const (
	MainDBLabel = "Main"
)

type (
	SqlDBSetter interface {
		SetDB(db *gorm.DB)
	}

	DBLabelGetter interface {
		GetDBLabel() string
	}

	DBCollector interface {
		RegisterDB(label string, db *gorm.DB)
		Provide(daoModel SqlDBSetter, args ...SqlDBSetter) error
		// ProvideDB Deprecated
		ProvideDB(daoModel DaoModel, args ...DaoModel) error
		ContextManager(label string, handle func(db *gorm.DB)) error
	}
)

// DaoModel
// Deprecated: 使用InjectSqlDao注入dao
type DaoModel struct {
	Dao   SqlDBSetter
	Model DBLabelGetter
}

// DBCollection db实例管理
type DBCollection struct {
	Collection map[string]*gorm.DB
}

// InjectSqlDao dao注入
// @param container: struct指针， 有一个DBCollector成员，至少一个SqlDBSetter成员
func InjectSqlDao(container interface{}) error {
	var dbCollector DBCollector
	var daoList []SqlDBSetter

	mutReflector := reflect.ValueOf(container).Elem()
	fieldNum := mutReflector.NumField()

	for fieldIndex := 0; fieldIndex < fieldNum; fieldIndex++ {
		member := mutReflector.Field(fieldIndex)

		switch concrete := member.Interface().(type) {
		case DBCollector:
			dbCollector = concrete
		case SqlDBSetter:
			daoList = append(daoList, concrete)
		default:
			continue
		}
	}

	if dbCollector == nil {
		return fmt.Errorf("DBCollector not exist")
	} else if len(daoList) == 0 {
		return fmt.Errorf("SqlDBSetter not exist")
	} else if len(daoList) > 1 {
		return dbCollector.Provide(daoList[0], daoList[1:]...)
	}
	return dbCollector.Provide(daoList[0])
}

func (collector *DBCollection) RegisterDB(label string, db *gorm.DB) {
	collector.Collection[label] = db
}

// ProvideDB
// Deprecated: 使用Provide替代
func (collector *DBCollection) ProvideDB(daoModel DaoModel, args ...DaoModel) error {
	daoModels := append([]DaoModel{}, daoModel)
	for _, arg := range args {
		daoModels = append(daoModels, arg)
	}

	for _, daoModel := range daoModels {
		db, ok := collector.Collection[daoModel.Model.GetDBLabel()]
		if !ok {
			return fmt.Errorf("db not exists, label not ")
		}

		daoModel.Dao.SetDB(db)
	}
	return nil
}

// Provide
// dao和args会被合并一个slice中遍历
// 调用示例
// c.ProvideDB(dao1) // 向一个dao提供db
// c.ProvideDB(dao1, dao2, dao3) // 向多个dao提供db
func (collector *DBCollection) Provide(dao SqlDBSetter, args ...SqlDBSetter) error {
	daoList := append([]SqlDBSetter{}, dao)
	for _, arg := range args {
		daoList = append(daoList, arg)
	}

	for _, dao := range daoList {
		var label string

		// 通过反射注入db到dao
		mutReflector := reflect.ValueOf(dao).Elem()
		fieldNum := mutReflector.NumField()

		// 通过DBLabelGetter确定dao所需要的db实例
		for fieldIndex := 0; fieldIndex < fieldNum; fieldIndex++ {
			member := mutReflector.Field(fieldIndex)
			if model, ok := member.Interface().(DBLabelGetter); ok {
				label = model.GetDBLabel()
				break
			}
		}

		// 没有model成员或model没有GetDBLabel方法 先检查dao是否为DBLabelGetter
		// model 和 dao 都没有的情况下使用默认值
		if strings.TrimSpace(label) != "" {
		} else if dBLabelGetter, ok := dao.(DBLabelGetter); ok {
			label = dBLabelGetter.GetDBLabel()
		} else {
			label = MainDBLabel
		}

		db, ok := collector.Collection[label]
		if !ok {
			return fmt.Errorf("db not exists, label: %s", MainDBLabel)
		}
		dao.SetDB(db)
	}
	return nil
}

func (collector *DBCollection) ContextManager(label string, handle func(db *gorm.DB)) error {
	db, ok := collector.Collection[label]
	if !ok {
		return fmt.Errorf("db not exists, label not ")
	}

	handle(db)
	return nil
}

// 基础db dao操作封装

type SqlDao struct {
	db *gorm.DB
}

func (dao *SqlDao) SetDB(db *gorm.DB) {
	dao.db = db
}

func (dao *SqlDao) GetDB() *gorm.DB {
	return dao.db
}

func (dao SqlDao) Select(model interface{}, columns []string, query interface{}, args ...interface{}) *gorm.DB {
	return dao.db.Select(columns[0], columns[1:]).Where(query, args).Find(model)
}

func (dao SqlDao) First(model interface{}, conds ...interface{}) *gorm.DB {
	return dao.db.First(model, conds)
}

func (dao SqlDao) Last(model interface{}, conds ...interface{}) *gorm.DB {
	return dao.db.Last(model, conds)
}

func (dao SqlDao) Find(model interface{}, query interface{}, args ...interface{}) *gorm.DB {
	return dao.db.Where(query, args...).Find(model)
}

func (dao SqlDao) Create(model interface{}) *gorm.DB {
	return dao.db.Create(model)
}

func (dao SqlDao) Update(model interface{}, field string, value interface{}, args ...interface{}) *gorm.DB {
	if len(args) > 0 {
		return dao.db.Model(model).Where(args[0], args[1:]).Update(field, value)
	}
	return dao.db.Model(model).Update(field, value)
}

func (dao SqlDao) Updates(model interface{}, values interface{}, columns ...interface{}) *gorm.DB {
	if len(columns) > 0 {
		return dao.db.Model(model).Select(columns[0], columns[1:]...).Updates(values)
	}
	return dao.db.Model(model).Updates(values)
}

func (dao SqlDao) UpdatesWithOmit(model interface{}, values interface{}, columns ...string) *gorm.DB {
	return dao.db.Model(model).Omit(columns...).Updates(values)
}

func (dao SqlDao) BatchUpdate(model interface{}, values interface{}, query interface{}, args ...interface{}) *gorm.DB {
	return dao.db.Model(model).Where(query, args...).Updates(values)
}

// Upsert Create or Update
// @param clauseColumn: 只有在clauseColumn中的字段unique下才能执行upsert
func (dao SqlDao) Upsert(model interface{}, clauseColumn []clause.Column, assignmentColumns []string) *gorm.DB {
	return dao.db.Clauses(clause.OnConflict{
		Columns:   clauseColumn,
		DoUpdates: clause.AssignmentColumns(assignmentColumns),
	}).Create(model)
}

func (dao SqlDao) PhysicalDeleteByID(model interface{}) *gorm.DB {
	return dao.db.Delete(model)
}

func (dao SqlDao) PhysicalDelete(model interface{}, query interface{}, args ...interface{}) *gorm.DB {
	return dao.db.Where(query, args...).Delete(model)
}

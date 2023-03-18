package dao

import (
	"context"
	"fmt"

	"go.uber.org/zap"

	"github.com/guidomantilla/go-feather-sql/pkg/datasource"
	feather_sql "github.com/guidomantilla/go-feather-sql/pkg/sql"
)

type DefaultCrudDao struct {
	//datasourceContext datasource.RelationalDatasourceContext
	driverName        feather_sql.DriverName
	paramHolder       feather_sql.ParamHolder
	statementCreate   string
	statementUpdate   string
	statementDelete   string
	statementFindById string
	statementFindAll  string
}

func NewDefaultCrudDao(datasourceContext datasource.RelationalDatasourceContext, table string, model any) *DefaultCrudDao {

	statementCreate, err := feather_sql.CreateInsertSQL(datasourceContext.GetDriverName(), table, model, datasourceContext.GetParamHolder().EvalValueOnly())
	if err != nil {
		zap.L().Fatal(fmt.Sprintf("server starting up - error setting up %s dao: %s", table, err.Error()))
	}

	statementUpdate, err := feather_sql.CreateUpdateSQL(datasourceContext.GetDriverName(), table, model, feather_sql.PkColumnFilter, datasourceContext.GetParamHolder().EvalNameValue())
	if err != nil {
		zap.L().Fatal(fmt.Sprintf("server starting up - error setting up %s dao: %s", table, err.Error()))
	}

	statementDelete, err := feather_sql.CreateDeleteSQL(datasourceContext.GetDriverName(), table, model, feather_sql.PkColumnFilter, datasourceContext.GetParamHolder().EvalNameValue())
	if err != nil {
		zap.L().Fatal(fmt.Sprintf("server starting up - error setting up %s dao: %s", table, err.Error()))
	}

	statementFindById, err := feather_sql.CreateSelectSQL(datasourceContext.GetDriverName(), table, model, feather_sql.PkColumnFilter, datasourceContext.GetParamHolder().EvalNameValue())
	if err != nil {
		zap.L().Fatal(fmt.Sprintf("server starting up - error setting up %s dao: %s", table, err.Error()))
	}

	statementFindAll, err := feather_sql.CreateSelectSQL(datasourceContext.GetDriverName(), table, model, nil, nil)
	if err != nil {
		zap.L().Fatal(fmt.Sprintf("server starting up - error setting up %s dao: %s", table, err.Error()))
	}

	return &DefaultCrudDao{
		driverName:        datasourceContext.GetDriverName(),
		paramHolder:       datasourceContext.GetParamHolder(),
		statementCreate:   statementCreate,
		statementUpdate:   statementUpdate,
		statementDelete:   statementDelete,
		statementFindById: statementFindById,
		statementFindAll:  statementFindAll,
	}
}

func (dao *DefaultCrudDao) Save(ctx context.Context, args ...any) (*int64, error) {

	var err error
	var serial *int64
	ctx = context.WithValue(ctx, feather_sql.DriverNameContext{}, dao.driverName)
	if serial, err = WriteContext(ctx, dao.statementCreate, args...); err != nil {
		return nil, err
	}

	return serial, nil
}

func (dao *DefaultCrudDao) Update(ctx context.Context, args ...any) error {

	var err error
	ctx = context.WithValue(ctx, feather_sql.DriverNameContext{}, dao.driverName)
	if _, err = WriteContext(ctx, dao.statementUpdate, args...); err != nil {
		return err
	}

	return nil
}

func (dao *DefaultCrudDao) Delete(ctx context.Context, id any) error {

	var err error
	ctx = context.WithValue(ctx, feather_sql.DriverNameContext{}, dao.driverName)
	if _, err = WriteContext(ctx, dao.statementDelete, id); err != nil {
		return err
	}

	return nil
}

func (dao *DefaultCrudDao) FindById(ctx context.Context, id any, args ...any) error {

	var err error
	ctx = context.WithValue(ctx, feather_sql.DriverNameContext{}, dao.driverName)
	if err = ReadRowContext(ctx, dao.statementFindById, id, args...); err != nil {
		return err
	}

	return nil
}

func (dao *DefaultCrudDao) ExistsById(ctx context.Context, id any, args ...any) bool {

	var err error
	ctx = context.WithValue(ctx, feather_sql.DriverNameContext{}, dao.driverName)
	if err = dao.FindById(ctx, id, args...); err != nil {
		return false
	}

	return true

}

func (dao *DefaultCrudDao) FindAll(ctx context.Context, fn ReadFunction) error {

	var err error
	ctx = context.WithValue(ctx, feather_sql.DriverNameContext{}, dao.driverName)
	if err = ReadContext(ctx, dao.statementFindAll, fn); err != nil {
		return err
	}

	return nil
}

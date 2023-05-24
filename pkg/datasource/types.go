package datasource

import (
	"database/sql"

	feather_sql "github.com/guidomantilla/go-feather-sql/pkg/sql"
)

var (
	_ RelationalDatasourceContext = (*DefaultRelationalDatasourceContext)(nil)
	_ RelationalDatasource        = (*DefaultRelationalDatasource)(nil)
)

type RelationalDatasourceContext interface {
	GetDriverName() feather_sql.DriverName
	GetParamHolder() feather_sql.ParamHolder
	GetUrl() string
}

//

type OpenDatasourceFunc func(driverName, datasourceUrl string) (*sql.DB, error)

type RelationalDatasource interface {
	GetDatabase() (*sql.DB, error)
}

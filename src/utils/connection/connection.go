package connection

import (
	"fmt"

	"entgo.io/ent/dialect"
	"github.com/irdaislakhuafa/GoAttendEasy/src/schema/generated"
	"github.com/irdaislakhuafa/GoAttendEasy/src/utils/config"
	"github.com/irdaislakhuafa/GoAttendEasy/src/utils/errors"
	"github.com/irdaislakhuafa/GoAttendEasy/src/utils/errors/codes"
	"github.com/irdaislakhuafa/GoAttendEasy/src/utils/operator"
	_ "github.com/lib/pq"
)

func OpenConnection(cfg config.AppConfig) (*generated.Client, error) {
	datasource := ""
	db := cfg.DB
	switch cfg.DB.Driver {
	case dialect.Postgres:
		datasource = fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=%v", db.Host, db.Port, db.Username, db.DatabaseName, db.Password, operator.Ternary(db.SSL, "enable", "disable"))
	default:
		return nil, errors.NewWithCode(codes.CodeDBNotSupported, "database driver '%s' not supported", cfg.DB.Driver)
	}

	c, err := generated.Open(cfg.DB.Driver, datasource)
	if err != nil {
		return nil, errors.NewWithCode(codes.CodeDBConnectionError, err.Error())
	}

	return c, nil
}

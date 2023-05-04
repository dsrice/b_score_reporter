package database

import (
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
	"gorm.io/gorm"
	"strings"
)

type myError struct {
	errs []string
}

func (e *myError) add(stage string, err error) {
	if err == nil {
		return
	}

	e.errs = append(e.errs, "stage="+stage+":"+err.Error())
}

func (e myError) toError() error {
	if len(e.errs) == 0 {
		return nil
	}

	return e
}

func (e myError) Error() string {
	return strings.Join(e.errs, ";")
}

func SetCallBack(db *gorm.DB) error {
	e := myError{
		errs: make([]string, 0, 12),
	}

	err := db.Callback().Create().Before("gorm:create").Register("update_created_at", before)
	e.add("update_created_at", err)
	err = db.Callback().Create().After("gorm:create").Register("update_created_at", after)
	e.add("update_created_at", err)

	err = db.Callback().Update().Before("gorm:update").Register("before_update", before)
	e.add("before_update", err)
	err = db.Callback().Update().After("gorm:update").Register("after_update", after)
	e.add("after_update", err)

	err = db.Callback().Query().Before("gorm:query").Register("before_query", before)
	e.add("before_query", err)
	err = db.Callback().Query().After("gorm:query").Register("after_query", after)
	e.add("after_query", err)

	err = db.Callback().Delete().Before("gorm:delete").Register("before_delete", before)
	e.add("before_delete", err)
	err = db.Callback().Delete().After("gorm:delete").Register("after_delete", after)
	e.add("after_delete", err)

	err = db.Callback().Row().Before("gorm:row").Register("before_row", before)
	e.add("before_row", err)
	err = db.Callback().Row().After("gorm:row").Register("after_row", after)
	e.add("after_row", err)

	err = db.Callback().Raw().Before("gorm:raw").Register("before_raw", before)
	e.add("before_raw", err)
	err = db.Callback().Raw().After("gorm:raw").Register("after_raw", after)
	e.add("after_raw", err)

	return e.toError()
}

func before(db *gorm.DB) {
	if db == nil {
		return
	}

	sp, _ := opentracing.StartSpanFromContextWithTracer(db.Statement.Context, opentracing.GlobalTracer(), "mysql")

	db.InstanceSet("db_span", sp)
}

func after(db *gorm.DB) {
	if db == nil {
		return
	}

	v, ok := db.InstanceGet("db_span")
	if !ok || v == nil {
		return
	}

	sp, ok := v.(opentracing.Span)
	if !ok || sp == nil {
		return
	}

	defer sp.Finish()

	sp.LogFields(log.String("SQL", db.Statement.SQL.String()))
	sp.LogFields(log.Object("param", db.Statement.Vars))
}

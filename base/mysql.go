package base

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/sirupsen/logrus"
	"time"
)

const TraceContextKey string = "traceContext"

const SqlTraceContextKey = "traceContext"

type MysqlConf struct {
	User            string
	Password        string
	Addr            string
	DataBase        string
	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxLifeTime time.Duration
	LogMode         bool
}

type GormLogger struct {
	logger gorm.Logger
}

// 设置为debug ?
func (logger GormLogger) Print(v ...interface{}) {
	if logrus.IsLevelEnabled(logrus.DebugLevel) {
		logger.logger.Print(v...)
	}
}

// InitMysqlClient 一定要在bootstrap之后
func InitMysqlClient(conf MysqlConf) (client *gorm.DB, err error) {
	client, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Asia%%2FShanghai",
		conf.User,
		conf.Password,
		conf.Addr,
		conf.DataBase))

	if err != nil {
		return client, err
	}

	client.DB().SetMaxIdleConns(conf.MaxIdleConns)
	client.DB().SetMaxOpenConns(conf.MaxOpenConns)
	client.DB().SetConnMaxLifetime(conf.ConnMaxLifeTime)
	client.LogMode(conf.LogMode)

	// 暂时用不着
	//logger := GormLogger{logger: gorm.DefaultLogger}
	//client.SetLogger(logger)

	// register tracer callback
	setCallback(client, "create")
	setCallback(client, "delete")
	setCallback(client, "update")
	setCallback(client, "query")
	setCallback(client, "row_query")

	return client, nil
}

func setCallback(client *gorm.DB, callbackName string) {
	beforeName := fmt.Sprintf("tracer:%v_before", callbackName)
	afterName := fmt.Sprintf("tracer:%v_after", callbackName)
	gormCallbackName := fmt.Sprintf("gorm:%v", callbackName)
	switch callbackName {
	case "create":
		client.Callback().Create().Before(gormCallbackName).Register(beforeName, func(scope *gorm.Scope) {
			tracerBefore(scope, callbackName)
		})
		client.Callback().Create().After(gormCallbackName).Register(afterName, func(scope *gorm.Scope) {
			tracerAfter(scope, callbackName)
		})
	case "query":
		client.Callback().Query().Before(gormCallbackName).Register(beforeName, func(scope *gorm.Scope) {
			tracerBefore(scope, callbackName)
		})
		client.Callback().Query().After(gormCallbackName).Register(afterName, func(scope *gorm.Scope) {
			tracerAfter(scope, callbackName)
		})
	case "update":
		client.Callback().Update().Before(gormCallbackName).Register(beforeName, func(scope *gorm.Scope) {
			tracerBefore(scope, callbackName)
		})
		client.Callback().Update().After(gormCallbackName).Register(afterName, func(scope *gorm.Scope) {
			tracerAfter(scope, callbackName)
		})
	case "delete":
		client.Callback().Delete().Before(gormCallbackName).Register(beforeName, func(scope *gorm.Scope) {
			tracerBefore(scope, callbackName)
		})
		client.Callback().Delete().After(gormCallbackName).Register(afterName, func(scope *gorm.Scope) {
			tracerAfter(scope, callbackName)
		})
	case "row_query":
		client.Callback().RowQuery().Before(gormCallbackName).Register(beforeName, func(scope *gorm.Scope) {
			tracerBefore(scope, callbackName)
		})
		client.Callback().RowQuery().After(gormCallbackName).Register(afterName, func(scope *gorm.Scope) {
			tracerAfter(scope, callbackName)
		})
	}
}

func tracerBefore(scope *gorm.Scope, callbackName string) {
	//ctx, ok := scope.Search.GetCtx().(*gin.Context)
	//if !ok {
	//	return
	//}

	var a = gin.Context{}
	var parentSpanContext opentracing.SpanContext
	//if ctx != nil {
	tempSpan, exist := a.Get(TraceContextKey)
	if exist {
		parentSpanContext = tempSpan.(opentracing.Span).Context()
		//}
	}

	span := TracerStartSpan(parentSpanContext, "DB_"+callbackName, map[string]interface{}{
		ext.SpanKindRPCClient.Key: ext.SpanKindRPCClient,
		string(ext.Component):     "gorm",
		string(ext.DBInstance):    scope.InstanceID(),
		string(ext.DBType):        "mysql",
	})
	scope.Set(SqlTraceContextKey, span)
}

func tracerAfter(scope *gorm.Scope, callbackName string) {
	tempSpan, exist := scope.Get(SqlTraceContextKey)
	if !exist {
		return
	}
	span := tempSpan.(opentracing.Span)
	// todo sql 的长度限制
	span.SetTag(string(ext.DBStatement), scope.SQL)
	span.Finish()
}

func TracerStartSpan(parent opentracing.SpanContext, operationName string, tags map[string]interface{}) opentracing.Span {
	var options []opentracing.StartSpanOption
	if parent != nil {
		options = append(options, opentracing.ChildOf(parent))
	}

	span := opentracing.StartSpan(operationName, options...)
	for k, v := range tags {
		span.SetTag(k, v)
	}
	return span
}

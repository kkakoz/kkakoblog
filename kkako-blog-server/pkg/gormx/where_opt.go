package gormx

import (
	"gorm.io/gorm"
	"reflect"
)

type Option func(db *gorm.DB)

var nullOption = func(db *gorm.DB) {}

func Cond(b bool, option Option) Option {
	if b {
		return option
	}
	return nullOption
}

func DBOpt(db *gorm.DB, opts []Option)  {
	for _, opt := range opts {
		opt(db)
	}
}

func WithEQ(key string, value interface{}) Option {
	return func(db *gorm.DB) {
		if isBlank(reflect.ValueOf(value)) {
			return
		}
		s := key + " = ?"
		db = db.Where(s, value)
	}
}

func WithEQForce(key string, value interface{}) Option {
	return func(db *gorm.DB) {
		s := key + " = ?"
		db = db.Where(s, value)
	}
}

func WithWhere(key string, value ...interface{}) Option {
	return func(db *gorm.DB) {
		if isBlank(reflect.ValueOf(value)) {
			return
		}
		db = db.Where(key, value...)
	}
}

func WithWhereForce(key string, value interface{}) Option {
	return func(db *gorm.DB) {
		db = db.Where(key, value)
	}
}

func WithOr(key string, value ...interface{}) Option {
	return func(db *gorm.DB) {
		if isBlank(reflect.ValueOf(value)) {
			return
		}
		db = db.Where(key, value...)
	}
}

func WithIn(key string, value interface{}) Option {
	return func(db *gorm.DB) {
		if isBlank(reflect.ValueOf(value)) {
			return
		}
		s := key + " in (?)"
		db = db.Where(s, value)
	}
}

func WithInForce(key string, value interface{}) Option {
	return func(db *gorm.DB) {
		s := key + " in (?)"
		db = db.Where(s, value)
	}
}

func WithNull(key string) Option {
	return func(db *gorm.DB) {
		s := key + " is null "
		db = db.Where(s)
	}
}

func WithNotNull(key string) Option {
	return func(db *gorm.DB) {
		s := key + " is not null "
		db = db.Where(s)
	}
}

func WithLike(key, value string) Option  {
	return func(db *gorm.DB) {
		if value == "" {
			return
		}
		key = key + " like ? "
		value = "%" + value + "%"
		db = db.Where(key, value)
	}
}

func isBlank(value reflect.Value) bool {
	switch value.Kind() {
	case reflect.String:
		return value.Len() == 0
	case reflect.Bool:
		return !value.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return value.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return value.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return value.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return value.IsNil()
	case reflect.Slice, reflect.Map:
		return value.Len() == 0 || value.IsNil()
	}
	return reflect.DeepEqual(value.Interface(), reflect.Zero(value.Type()).Interface())
}

package main

import (
	"gorm.io/gorm"
)

// ScopeOption 定义scope选项的函数类型
type ScopeOption func(*gorm.DB) *gorm.DB

// ScopeBuilder 创建一个组合多个选项的scope函数
func NewScope(options ...ScopeOption) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		for _, option := range options {
			db = option(db)
		}
		return db
	}
}

// Select 设置查询字段
func Select(fields ...string) ScopeOption {
	return func(db *gorm.DB) *gorm.DB {
		return db.Select(fields)
	}
}

// Where 添加where条件
func Where(query interface{}, args ...interface{}) ScopeOption {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where(query, args...)
	}
}

// Limit 设置查询限制数量
func Limit(limit int) ScopeOption {
	return func(db *gorm.DB) *gorm.DB {
		return db.Limit(limit)
	}
}

// Offset 设置查询偏移量
func Offset(offset int) ScopeOption {
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(offset)
	}
}

// Order 设置排序
func Order(order string) ScopeOption {
	return func(db *gorm.DB) *gorm.DB {
		return db.Order(order)
	}
}

// Group 设置分组
func Group(group string) ScopeOption {
	return func(db *gorm.DB) *gorm.DB {
		return db.Group(group)
	}
}

// Having 设置having条件
func Having(query interface{}, args ...interface{}) ScopeOption {
	return func(db *gorm.DB) *gorm.DB {
		return db.Having(query, args...)
	}
}

// Joins 设置join连接
func Joins(query string, args ...interface{}) ScopeOption {
	return func(db *gorm.DB) *gorm.DB {
		return db.Joins(query, args...)
	}
}

// Preload 设置预加载关联，支持嵌套选项
func Preload(association string, options ...ScopeOption) ScopeOption {
	return func(db *gorm.DB) *gorm.DB {
		if len(options) > 0 {
			// 如果有嵌套选项，创建一个组合的scope函数
			nestedScope := NewScope(options...)
			return db.Preload(association, nestedScope)
		}
		return db.Preload(association)
	}
}

// Distinct 设置去重
func Distinct(columns ...string) ScopeOption {
	return func(db *gorm.DB) *gorm.DB {
		if len(columns) > 0 {
			return db.Distinct(columns)
		}
		return db.Distinct()
	}
}

// Omit 忽略指定字段
func Omit(columns ...string) ScopeOption {
	return func(db *gorm.DB) *gorm.DB {
		return db.Omit(columns...)
	}
}

// Scopes 应用其他自定义scope
func Scopes(scopes ...func(*gorm.DB) *gorm.DB) ScopeOption {
	return func(db *gorm.DB) *gorm.DB {
		return db.Scopes(scopes...)
	}
}

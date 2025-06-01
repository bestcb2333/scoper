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

// WithSelect 设置查询字段
func WithSelect(fields ...string) ScopeOption {
	return func(db *gorm.DB) *gorm.DB {
		return db.Select(fields)
	}
}

// WithWhere 添加where条件
func WithWhere(query interface{}, args ...interface{}) ScopeOption {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where(query, args...)
	}
}

// WithLimit 设置查询限制数量
func WithLimit(limit int) ScopeOption {
	return func(db *gorm.DB) *gorm.DB {
		return db.Limit(limit)
	}
}

// WithOffset 设置查询偏移量
func WithOffset(offset int) ScopeOption {
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(offset)
	}
}

// WithOrder 设置排序
func WithOrder(order string) ScopeOption {
	return func(db *gorm.DB) *gorm.DB {
		return db.Order(order)
	}
}

// WithGroup 设置分组
func WithGroup(group string) ScopeOption {
	return func(db *gorm.DB) *gorm.DB {
		return db.Group(group)
	}
}

// WithHaving 设置having条件
func WithHaving(query interface{}, args ...interface{}) ScopeOption {
	return func(db *gorm.DB) *gorm.DB {
		return db.Having(query, args...)
	}
}

// WithJoins 设置join连接
func WithJoins(query string, args ...interface{}) ScopeOption {
	return func(db *gorm.DB) *gorm.DB {
		return db.Joins(query, args...)
	}
}

// WithPreload 设置预加载关联，支持嵌套选项
func WithPreload(association string, options ...ScopeOption) ScopeOption {
	return func(db *gorm.DB) *gorm.DB {
		if len(options) > 0 {
			// 如果有嵌套选项，创建一个组合的scope函数
			nestedScope := NewScope(options...)
			return db.Preload(association, nestedScope)
		}
		return db.Preload(association)
	}
}

// WithDistinct 设置去重
func WithDistinct(columns ...string) ScopeOption {
	return func(db *gorm.DB) *gorm.DB {
		if len(columns) > 0 {
			return db.Distinct(columns)
		}
		return db.Distinct()
	}
}

// WithOmit 忽略指定字段
func WithOmit(columns ...string) ScopeOption {
	return func(db *gorm.DB) *gorm.DB {
		return db.Omit(columns...)
	}
}

// WithScopes 应用其他自定义scope
func WithScopes(scopes ...func(*gorm.DB) *gorm.DB) ScopeOption {
	return func(db *gorm.DB) *gorm.DB {
		return db.Scopes(scopes...)
	}
}

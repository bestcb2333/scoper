That's the code simplified by scoper:
```go
var companies []Company
db.Preload("Users", scoper.NewScope(
  scoper.WithSelect("id", "name", "avatar", "company_id"),
  scoper.WithWhere("gender = ?", "male"),
  scoper.WithPreload("Comments",
    scoper.WithSelect("id", "title", "user_id"),
    scoper.WithWhere("topic = ?", "environment"),
  ),
)).Find(&companies).Error
```

That's the original code:
```go
var companies []Company
db.Preload("Users", func(db *gorm.DB) *gorm.DB {
  return db.
    Select("id", "name", "avatar", "company_id").
    Where("gender = ?", "male").
    Preload("Comments", func(db *gorm.DB) *gorm.DB {
      return db.
        Select("id", "title", "user_id").
        Where("topic = ?", "environment")
    })
}).Find(&companies).Error
```


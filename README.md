That's the code simplified by scoper:
```go
import s "github.com/bestcb2333/scoper"

var companies []Company
db.Preload("Users", s.NewScope(
  s.Select("id", "name", "avatar", "company_id"),
  s.Where("gender = ?", "male"),
  s.Preload("Comments",
    s.Select("id", "title", "user_id"),
    s.Where("topic = ?", "environment"),
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


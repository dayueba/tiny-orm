## 效果
database/sql
```go
func FindUsers(ctx context.Context) ([]*User, error) {
    rows, err := db.QueryContext(ctx, "SELECT `id`,`name`,`age`,`ctime`,`mtime` FROM user WHERE `age`<? ORDER BY `id` LIMIT 20 ", 20)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    result := []*User{}
    for rows.Next() {
        a := &User{}
        if err := rows.Scan(&a.Id, &a.Name, &a.Age, &a.Ctime, &a.Mtime); err != nil {
            return nil, err
        }
        result = append(result, a)
    }
    if rows.Err() != nil {
        return nil, rows.Err()
    }
    return result, nil
}
```
orm
```go
func FindUsers(ctx context.Context) ([]*User, error) {
	users := make([]*User, 20)
	err := db.Table(&User{}).Select("id", "name", "age", "ctime","mtime`").Where("age < ", 20).Limit(20).Find(&users).Err
	return users, err
}
```

## orm的核心组成
1. SQLBuilder：SQL语句要非硬编码，通过某种链式调用构造器帮助我构建SQL语句。
2. Scanner：从数据库返回的数据可以自动映射赋值到结构体中

### SQL SelectBuilder

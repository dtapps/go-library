# fly

A simple go db library

- data hook, Easy transition db data
- sql builder, Not need handwritten SQL
- hasOne/hasMany, Convenient access to associated data
- validator, Flexible verification policies
- extensible, Easy extend custom hook/sql/validator
- cacheEnable, Support for custom cache implementations

# usages

more example please check out [model_test.go](./model_test.go)

```go
package main

import (
    "fmt"

    "github.com/daodao97/fly"
    _ "github.com/go-sql-driver/mysql"
)

func init() {
    err := fly.Init(map[string]*fly.Config{
        "default": {DSN: "root@tcp(127.0.0.1:3306)/fly_test?&parseTime=true"},
    })
    if err != nil {
        panic(err)
    }
}

func main() {
    m := fly.New(
        "user",
        fly.ColumnHook(fly.CommaInt("role_id"), fly.Json("profile")),
    )

    var err error

    _, err = m.Insert(map[string]interface{}{
        "name": "Seiya",
        "profile": map[string]interface{}{
            "hobby": "Pegasus Ryuseiken",
        },
    })

    var result []*User
    err = m.Select(fly.WhereGt("id", 1)).Binding(&result)

    var result1 *User
    err = m.SelectOne(fly.WhereEq("id", 1)).Binding(&result1)

    count, err := m.Count()
    fmt.Println("count", count)

    _, err = m.Update(User{
        ID:   1,
        Name: "星矢",
        Profile: &Profile{
            Hobby: "天马流行拳",
        },
        RoleIds: []int{2, 3},
    })

    _, err = m.Delete(fly.WhereEq("id", 1))

    fmt.Println(err)
}

type User struct {
    ID        int64    `json:"id"`
    Name      string   `json:"name"`
    Status    int64    `json:"status"`
    Profile   *Profile `json:"profile"`
    IsDeleted int      `json:"is_deleted"`
    RoleIds   []int    `json:"role_ids"`
    Score     int      `json:"score"`
}

type Profile struct {
    Hobby string `json:"hobby"`
}
```
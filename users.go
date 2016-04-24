package main

import (
    "net/http"

    "github.com/labstack/echo"
    "github.com/labstack/echo/engine/standard"
    "github.com/rs/cors"
)

// UserのJSONの形を指定します。
type (
    user struct {
        ID   string `json:"id"`
        Name string `json:"name"`
        Age  int    `json:"age"`
    }
)

var (
    users map[string]user
)

func init() {
    // 初期データを登録してみます
    users = map[string]user{
        "1": user{
            ID:   "1",
            Name: "ジョナサン・ジョースター",
            Age:  20,
        },
        "2": user{
            ID:   "2",
            Name: "ディオ・ブランドー",
            Age:  21,
        },
    }

    g := e.Group("/users")
    g.Use(standard.WrapMiddleware(cors.Default().Handler))

    g.POST("", createUser)
    g.GET("", getUsers)
    g.GET("/:id", getUser)
}

// Userの作成、一覧、詳細です
func createUser(c echo.Context) error {
    u := new(user)
    if err := c.Bind(u); err != nil {
        return err
    }
    users[u.ID] = *u
    return c.JSON(http.StatusCreated, u)
}

func getUsers(c echo.Context) error {
    return c.JSON(http.StatusOK, users)
}

func getUser(c echo.Context) error {
    return c.JSON(http.StatusOK, users[c.P(0)])
}
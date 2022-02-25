package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"log"
	"github.com/Sowiriro/Phione/database"
	"github.com/Sowiriro/Phione/greeting"
	"fmt"
	//"database/sql"
	_"github.com/go-sql-driver/mysql"
)

type Hero struct {
	Id int
	Name string
}

	var err error

func main() {
	//他のパッケージからインポートできているのか？の確認
	message := greeting.Hello("Sowiriro")
	fmt.Println(message)

	// hero := Hero{Name: "alpha"}
	// log.Printf("Starting")
	// hero.Add()
	// log.Printf("ending")

	//インスタンスの生成
	e := echo.New()

	// //ルーティング
	e.GET("/", hello)
	e.GET("/heros", all)
	e.POST("/hero/create", Create)
//	e.POST("/hero/:id", detail)
//	e.POST("/hero/:id/update", update)
//	e.DELETE("hero/delete", delete)	


	//サーバーのスタート
	e.Logger.Fatal(e.Start(":1323"))
}


func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello Sowiriro")
}

func all(c echo.Context) error {
	

	return c.String(http.StatusOK, "All")
}

func detail(c echo.Context) error {
	return c.String(http.StatusOK, "Detail")
}


func Create(c echo.Context) error {

	name := c.FormValue("name")

	db := database.Connect()

	_, err := db.Exec("INSERT INTO heros(name) VALUES (?)", name)
	if err != nil {
		log.Println("うまく値が入れられない")
		log.Fatal(err)
	}

	log.Printf("libraを作成しました")


	 return c.String(http.StatusOK, name + "を取得しました")
}


func update(c echo.Context) error {
	return c.String(http.StatusOK, "Update")
}


func delete(c echo.Context) error {
	return c.String(http.StatusOK, "Delete")
}
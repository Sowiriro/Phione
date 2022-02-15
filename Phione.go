package main

import (
	"database/sql"
	"github.com/labstack/echo/v4"
	"net/http"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"phione/database"
)

type Hero struct {
	Id int
	Name string
}

	var err error
	var db *sql.DB


func main() {

	Insert()

	// hero := Hero{Name: "alpha"}
	// log.Printf("Starting")
	// hero.Add()
	// log.Printf("ending")

	//インスタンスの生成
	// e := echo.New()

	// //ルーティング
	// e.GET("/", hello)
	// e.GET("/heros", all)
	// e.POST("/hero/create", create)
//	e.POST("/hero/:id", detail)
//	e.POST("/hero/:id/update", update)
//	e.DELETE("hero/delete", delete)	


	//サーバーのスタート
	// e.Logger.Fatal(e.Start(":1323"))
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


func create(c echo.Context) error {

	// _, err := Db.Exec("INSERT INTO heros (name) VALUES ('アルファ')")
	// if err != nil {
	// 	return err 
	// }

	name := c.FormValue("name")

	 return c.String(http.StatusOK, name + "をCreateしました")
}


func update(c echo.Context) error {
	return c.String(http.StatusOK, "Update")
}


func delete(c echo.Context) error {
	return c.String(http.StatusOK, "Delete")
}


func Insert() {
	db := database.Connect()
	ins, err := db.Prepare("INSERT INTO heros(name) VALUES(?)")
	if err != nil {
		log.Fatal(err)
	}
	defer ins.Close()

	//sql 実行
	ins.Exec("alpha")


	log.Printf("アルファを作成しました")
}



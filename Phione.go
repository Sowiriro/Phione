package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"log"
	"github.com/Sowiriro/Phione/database"
	"github.com/Sowiriro/Phione/greeting"
	"fmt"
	"database/sql"
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

	//インスタンスの生成
	e := echo.New()

	// //ルーティング
	e.GET("/", hello)
	e.GET("/heros", all)
	e.POST("/hero/create", Create)
	e.GET("/hero/:id", detail)
	e.POST("/hero/:id/update", update)
	e.DELETE("hero/:id/delete", delete)	


	//サーバーのスタート
	e.Logger.Fatal(e.Start(":1323"))
}


func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello Sowiriro")
}

func all(c echo.Context) error {
	db := database.Connect()

	rows, err := db.Query(`SELECT * FROM heros`)
	if err != nil {
		return err
	}

	defer rows.Close()

	var heros []Hero

	for rows.Next() {
		var hero Hero
		if err := rows.Scan(&hero.Id ,&hero.Name); err != nil {
			return err
		}
		log.Println(hero.Id, hero.Name)
		heros = append(heros, hero)
	}

	return c.String(http.StatusOK, "全てを取得できました")
}

func detail(c echo.Context) error {
	id := c.Param("id")
	db := database.Connect()

	row :=db.QueryRow("SELECT * FROM heros where id = ?", id)
	
	var hero Hero
	err := row.Scan(&hero.Id, &hero.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("そんな行はありませんでした")
		} else {
			log.Println(err)
		}
	}
	log.Println(hero.Id, hero.Name + "を取得しました")
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
	id := c.Param("id")
	name := c.FormValue("name")
	db := database.Connect()

	//prepare文の追加するのがベストプラクティスらしいので追記しますここに↓

	_, err := db.Exec("UPDATE heros SET name = ? WHERE id = ?", name, id)
	if err != nil {
		return err
	}

	log.Println(name + "にUpdateしました")
	return c.String(http.StatusOK, "Update")
}


func delete(c echo.Context) error {
	id := c.Param("id")
	db := database.Connect()
	//prepare文の追加するのがベストプラクティスらしいので追記しますここに↓

	_, err := db.Exec("DELETE FROM heros WHERE id = ?", id)
	if err != nil {
		log.Fatalln(err)
	}

	return c.String(http.StatusOK, id + "番をDeleteしました")
}
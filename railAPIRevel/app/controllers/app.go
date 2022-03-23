package controllers

import (
	"log"
	"net/http"
	handler "railAPIRevel/handler"
	model "railAPIRevel/model"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
}

func (c App) GetCelanaById() revel.Result {
	db := handler.Connect()
	defer db.Close()

	id := c.Params.Route.Get("celana_id")
	query := "SELECT id,nama,harga,stok,deskripsi FROM celana WHERE id = " + id

	rows, err := db.Query(query)
	if err != nil {
		log.Println(err)
	}

	var celana model.DataCelana
	rows.Next()
	if err := rows.Scan(&celana.ID, &celana.Nama, &celana.Harga, &celana.Stok, &celana.Deskripsi); err != nil {
		log.Println(err)
		var response model.Response
		response.Message = "Error"
		response.Status = 404
		return c.Render(response)
	}
	var response model.ResponseCelana
	response.Message = "Success"
	response.Status = 200
	response.Data = celana
	return c.RenderJSON(response)
}

func (c App) UpdateCelanaById() revel.Result {
	db := handler.Connect()
	defer db.Close()

	var celana model.DataCelana
	celana.ID, _ = strconv.Atoi(c.Params.Get("id"))
	celana.Nama = c.Params.Get("nama")
	celana.Harga, _ = strconv.Atoi(c.Params.Get("harga"))
	celana.Stok, _ = strconv.Atoi(c.Params.Get("stok"))
	celana.Deskripsi = c.Params.Get("deskripsi")

	var response model.Response
	_, err := db.Exec("UPDATE celana SET nama = ?, harga = ?, stok = ?, deskripsi = ? WHERE id = ?", celana.Nama, celana.Harga, celana.Stok, celana.Deskripsi, celana.ID)
	if err != nil {
		log.Println(err)
		c.Response.Status = http.StatusBadGateway
		response.Message = "Error"
		response.Status = 404
	} else {
		response.Message = "Success"
		response.Status = 200
	}
	c.Response.Status = http.StatusOK
	return c.RenderJSON(response)
}

func (c App) InsertNewCelana() revel.Result {
	db := handler.Connect()
	defer db.Close()

	var celana model.DataCelana
	celana.Nama = c.Params.Get("nama")
	celana.Harga, _ = strconv.Atoi(c.Params.Get("harga"))
	celana.Stok, _ = strconv.Atoi(c.Params.Get("stok"))
	celana.Deskripsi = c.Params.Get("deskripsi")

	var response model.Response

	_, err := db.Exec("INSERT INTO celana(nama, harga, stok, deskripsi) values (?,?,?,?)", celana.Nama, celana.Harga, celana.Stok, celana.Deskripsi)
	if err != nil {
		log.Println(err)
		c.Response.Status = http.StatusBadGateway
		response.Message = "Error"
		response.Status = 404
		return c.RenderJSON(response)
	}
	response.Message = "Success"
	response.Status = 200

	return c.RenderJSON(response)
}

func (c App) DeleteCelanaById() revel.Result {
	db := handler.Connect()
	defer db.Close()

	id := c.Params.Route.Get("celana_id")

	var response model.Response
	_, err := db.Exec("DELETE FROM celana WHERE id = ?", id)
	if err != nil {
		log.Println(err)
		c.Response.Status = http.StatusBadGateway
		response.Message = "Error"
		response.Status = 404
		return c.RenderJSON(response)
	}

	response.Message = "Success"
	response.Status = 200
	return c.RenderJSON(response)
}

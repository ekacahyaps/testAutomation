package handler

import (
	"api/features/book"
	"api/helper"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type bookHandle struct {
	srv book.BookService
}

func New(bs book.BookService) book.BookHandler {
	return &bookHandle{
		srv: bs,
	}
}

func (bh *bookHandle) Add() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := AddBookRequest{}
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, "format inputan salah")
		}

		cnv := ToCore(input)

		res, err := bh.srv.Add(c.Get("user"), *cnv)
		if err != nil {
			log.Println("trouble :  ", err.Error())
			return c.JSON(helper.PrintErrorResponse(err.Error()))
		}

		return c.JSON(helper.PrintSuccessReponse(http.StatusCreated, "sukses menambahkan buku", res))
	}
}
func (bh *bookHandle) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		var token interface{}

		bookID := c.Param("id")
		cnv, err := strconv.Atoi(bookID)
		if err != nil {
			return c.JSON(http.StatusBadRequest, "id salah")
		}

		input := UpdateBookRequest{}
		err2 := c.Bind(&input)
		if err2 != nil {
			return c.JSON(http.StatusBadRequest, "format inputan salah")
		}

		cnvInput := ToCore(input)

		res, err := bh.srv.Update(token, cnv, *cnvInput)
		if err != nil {
			log.Println("trouble : ", err.Error())
			return c.JSON(helper.PrintErrorResponse(err.Error()))
		}
		return c.JSON(helper.PrintSuccessReponse(http.StatusCreated, "sukses mengupdate buku", res))
	}
}

func (bh *bookHandle) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		var token interface{}
		input := c.Param("id")
		cnv, err := strconv.Atoi(input)
		if err != nil {
			return c.JSON(http.StatusBadRequest, "id salah")
		}

		err2 := bh.srv.Delete(token, cnv)
		if err2 != nil {
			return c.JSON(helper.PrintErrorResponse(err2.Error()))
		}

		return c.JSON(http.StatusOK, "berhasil hapus data buku")
	}
}

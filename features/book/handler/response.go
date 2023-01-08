package handler

import "api/features/book"

type BookResponse struct {
	ID          uint   `json:"id"`
	Judul       string `json:"judul"`
	TahunTerbit int    `json:"tahun_terbit"`
	Penulis     string `json:"penulis"`
	Pemilik     string `json:"pemilik"`
}
type AddBookResponse struct {
	ID          uint   `json:"id"`
	Judul       string `json:"judul"`
	TahunTerbit int    `json:"tahun_terbit"`
	Penulis     string `json:"penulis"`
}
type MyBookResponse struct {
	ID          uint   `json:"id"`
	Judul       string `json:"judul"`
	TahunTerbit int    `json:"tahun_terbit"`
	Penulis     string `json:"penulis"`
}

func ToResponse(feature string, book book.Core) interface{} {
	switch feature {
	case "add":
		return AddBookResponse{
			Judul:       book.Judul,
			TahunTerbit: book.TahunTerbit,
			Penulis:     book.Penulis,
		}
	case "my":
		return MyBookResponse{
			ID:          book.ID,
			Judul:       book.Judul,
			TahunTerbit: book.TahunTerbit,
			Penulis:     book.Penulis,
		}
	default:
		return BookResponse{
			ID:          book.ID,
			Judul:       book.Judul,
			TahunTerbit: book.TahunTerbit,
			Penulis:     book.Penulis,
			Pemilik:     book.Pemilik,
		}
	}
}

func AddBookToResponse(dataCore book.Core) AddBookResponse {
	return AddBookResponse{
		ID:          dataCore.ID,
		Judul:       dataCore.Judul,
		TahunTerbit: dataCore.TahunTerbit,
		Penulis:     dataCore.Penulis,
	}
}

//For MyBooks
func MyBookToResponse(dataCore book.Core) MyBookResponse {
	return MyBookResponse{
		ID:          dataCore.ID,
		Judul:       dataCore.Judul,
		TahunTerbit: dataCore.TahunTerbit,
		Penulis:     dataCore.Penulis,
	}
}
func ListMyBookToResponse(dataCore []book.Core) []MyBookResponse {
	var DataResponse []MyBookResponse

	for _, value := range dataCore {
		DataResponse = append(DataResponse, MyBookToResponse(value))
	}
	return DataResponse
}

//For AllBooks (view all books collection)
func AllBooksToResponse(dataCore book.Core) BookResponse {
	return BookResponse{
		ID:          dataCore.ID,
		Judul:       dataCore.Judul,
		TahunTerbit: dataCore.TahunTerbit,
		Penulis:     dataCore.Penulis,
		Pemilik:     dataCore.Pemilik,
	}
}
func ListAllBooksToResponse(dataCore []book.Core) []BookResponse {
	var DataResponse []BookResponse

	for _, value := range dataCore {
		DataResponse = append(DataResponse, AllBooksToResponse(value))
	}
	return DataResponse
}

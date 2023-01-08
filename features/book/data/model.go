package data

import (
	"api/features/book"

	"gorm.io/gorm"
)

type Books struct {
	gorm.Model
	Judul       string
	Penulis     string
	TahunTerbit int
	UserID      uint
}

type BooksOwner struct {
	ID          uint
	Judul       string
	Penulis     string
	TahunTerbit int
	Pemilik     string
}

func ToCore(data Books) book.Core {
	return book.Core{
		ID:          data.ID,
		Judul:       data.Judul,
		TahunTerbit: data.TahunTerbit,
		Penulis:     data.Penulis,
	}
}

func CoreToData(data book.Core) Books {
	return Books{
		Model:       gorm.Model{ID: data.ID},
		Judul:       data.Judul,
		Penulis:     data.Penulis,
		TahunTerbit: data.TahunTerbit,
	}
}

// For MyBooks
func (dataModel *Books) ModelsToCore() book.Core {
	return book.Core{
		ID:          dataModel.ID,
		Judul:       dataModel.Judul,
		Penulis:     dataModel.Penulis,
		TahunTerbit: dataModel.TahunTerbit,
	}
}

func ListModelToCore(dataModel []Books) []book.Core {
	var dataCore []book.Core
	for _, value := range dataModel {
		dataCore = append(dataCore, value.ModelsToCore())
	}
	return dataCore
}

// For AllBooks (view all books collection)
func (dataModel *BooksOwner) AllModelsToCore() book.Core {
	return book.Core{
		ID:          dataModel.ID,
		Judul:       dataModel.Judul,
		Penulis:     dataModel.Penulis,
		TahunTerbit: dataModel.TahunTerbit,
		Pemilik:     dataModel.Pemilik,
	}
}

func ListAllModelToCore(dataModel []BooksOwner) []book.Core {
	var dataCore []book.Core
	for _, value := range dataModel {
		dataCore = append(dataCore, value.AllModelsToCore())
	}
	return dataCore
}

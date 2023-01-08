package data

import (
	"api/features/book"
	"errors"
	"log"

	"gorm.io/gorm"
)

type bookData struct {
	db *gorm.DB
}

func New(db *gorm.DB) book.BookData {
	return &bookData{
		db: db,
	}
}

func (bd *bookData) Add(userID int, newBook book.Core) (book.Core, error) {
	cnv := CoreToData(newBook)
	cnv.UserID = uint(userID)
	err := bd.db.Create(&cnv).Error
	if err != nil {
		return book.Core{}, err
	}

	newBook.ID = cnv.ID

	return newBook, nil
}

func (bd *bookData) Update(bookID int, userID int, updatedData book.Core) (book.Core, error) {
	cnv := CoreToData(updatedData)
	qry := bd.db.Where("id = ? AND user_id = ?", bookID, userID).Updates(&cnv)
	if qry.RowsAffected <= 0 {
		log.Println("update book query error : data not found")
		return book.Core{}, errors.New("not found")
	}

	if err := qry.Error; err != nil {
		log.Println("update book query error :", err.Error())
		return book.Core{}, err
	}

	return ToCore(cnv), nil
}

func (bd *bookData) Delete(bookID int, userID int) error {
	qry := bd.db.Where("user_id = ?", userID).Delete(&Books{}, bookID)

	affrows := qry.RowsAffected

	if affrows == 0 {
		log.Println("no rows affected")
		return errors.New("tidak ada data buku yang dihapus")
	}

	err := qry.Error

	if err != nil {
		log.Println("delete query error")
		return errors.New("tidak bisa menghapus data buku")
	}

	return nil

}

func (bd *bookData) MyBook(userID int) ([]book.Core, error) {
	myBooks := []Books{}
	err := bd.db.Where("user_id = ?", userID).Find(&myBooks).Error
	if err != nil {
		return []book.Core{}, err
	}

	listMyBooks := ListModelToCore(myBooks)

	return listMyBooks, nil
}

func (bd *bookData) AllBooks() ([]book.Core, error) {
	myBooks := []BooksOwner{}
	err := bd.db.Raw("SELECT books.id, books.judul, books.penulis, books.tahun_terbit, users.nama as pemilik FROM books JOIN users ON users.id = books.user_id WHERE books.deleted_at is NULL").Find(&myBooks).Error
	if err != nil {
		return []book.Core{}, err
	}

	listMyBooks := ListAllModelToCore(myBooks)

	return listMyBooks, nil
}

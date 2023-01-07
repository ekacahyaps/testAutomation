package handler

import "api/features/book"

type AddBookRequest struct {
	Judul       string `json:"judul"`
	Penulis     string `json:"penulis"`
	TahunTerbit int    `json:"tahun_terbit"`
}

type UpdateBookRequest struct {
	Judul       string `json:"judul"`
	Penulis     string `json:"penulis"`
	TahunTerbit int    `json:"tahun_terbit"`
}

func ToCore(data interface{}) *book.Core {
	res := book.Core{}

	switch data.(type) {
	case AddBookRequest:
		cnv := data.(AddBookRequest)
		res.Judul = cnv.Judul
		res.Penulis = cnv.Penulis
		res.TahunTerbit = cnv.TahunTerbit
	case UpdateBookRequest:
		cnv := data.(UpdateBookRequest)
		res.Judul = cnv.Judul
		res.Penulis = cnv.Penulis
		res.TahunTerbit = cnv.TahunTerbit
	default:
		return nil
	}

	return &res
}

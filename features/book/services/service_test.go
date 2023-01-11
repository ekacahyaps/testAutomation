package services

import (
	"api/features/book"
	"api/helper"
	"api/mocks"
	"testing"

	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	repo := mocks.NewBookData(t)

	t.Run("berhasil add buku", func(t *testing.T) {
		userID := 2
		inputData := book.Core{Judul: "kalkulus 1", Penulis: "Alfian", TahunTerbit: 2019}

		resData := book.Core{ID: 3, Judul: "kalkulus 1", Penulis: "Alfian", TahunTerbit: 2019}

		repo.On("Add", userID, inputData).Return(resData, nil)

		srv := New(repo)

		_, token := helper.GenerateJWT(2)

		pToken := token.(*jwt.Token)
		pToken.Valid = true

		res, err := srv.Add(pToken, inputData)
		assert.Nil(t, err)
		assert.Equal(t, resData.ID, res.ID)
		repo.AssertExpectations(t)
	})
}

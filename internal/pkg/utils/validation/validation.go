package utils

import (
	"errors"
	"strings"

	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/models"
)

const (
	maxShortFieldLength = 20
	maxAddressLength    = 200
	maxCommentLength    = 300
	minFieldLength      = 1
)

// Общий набор допустимых символов
const allowedSymbols = "абвгдеёжзийклмнопрстуфхцчшщъыьэюя" +
	"АБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯ" +
	"abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789 -_#*,."

func isValidShortField(s string) bool {
	if len(s) < minFieldLength || len(s) > maxShortFieldLength {
		return false
	}
	for _, r := range s {
		if !strings.ContainsRune(allowedSymbols, r) {
			return false
		}
	}
	return true
}

func isValidAddress(s string) bool {
	if len(s) < minFieldLength || len(s) > maxAddressLength {
		return false
	}
	for _, r := range s {
		if !strings.ContainsRune(allowedSymbols, r) {
			return false
		}
	}
	return true
}

func isValidComment(s string) bool {
	if len(s) > maxCommentLength {
		return false
	}
	for _, r := range s {
		if !strings.ContainsRune(allowedSymbols, r) {
			return false
		}
	}
	return true
}

func ValidateOrderInput(req *models.OrderInReq) error {
	if !isValidShortField(req.Status) {
		return errors.New("некорректный статус (макс 20 символов)")
	}
	if !isValidAddress(req.Address) {
		return errors.New("некорректный адрес (макс 200 символов)")
	}
	if !isValidShortField(req.ApartmentOrOffice) {
		return errors.New("некорректная квартира/офис (макс 20 символов)")
	}
	if !isValidShortField(req.Intercom) {
		return errors.New("некорректный домофон (макс 20 символов)")
	}
	if !isValidShortField(req.Entrance) {
		return errors.New("некорректный подъезд (макс 20 символов)")
	}
	if !isValidShortField(req.Floor) {
		return errors.New("некорректный этаж (макс 20 символов)")
	}
	if req.CourierComment != "" && !isValidComment(req.CourierComment) {
		return errors.New("некорректный комментарий (макс 300 символов)")
	}
	if req.FinalPrice < 0 {
		return errors.New("цена не может быть отрицательной")
	}
	return nil
}

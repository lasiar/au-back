package auth

import (
	"errors"
	"sync"
)

// создание переменных для ошибок
var (
	errWrongPassword         = errors.New("wrong username or password")
	errAuthForbidden         = errors.New("authorization is forbidden")
	errSessionIsNotCompleted = errors.New("session is not completed")
	errSession               = errors.New("session error")
)

// IsWrongPassword возврашает истину если ошибка это неверный логин пароль
func IsWrongPassword(err error) bool { return err == errWrongPassword }

// IsAuthForbidden возвращает истину если ошибка это запрет авторизации
func IsAuthForbidden(err error) bool { return err == errAuthForbidden }

// IsSessionNotCompleted возвращает истину если ошибка это незавершеная сессия
func IsSessionNotCompleted(err error) bool { return err == errSessionIsNotCompleted }

// IsSessionError возвращает истину если ошибка в сесии
func IsSessionError(err error) bool { return err == errSession }

var _once sync.Once
var _auth *Auth

// GetAuth возвращает текущий объект авторизации
// вызовет панику если config и log не были установлены сетерами
func GetAuth() *Auth {
	_once.Do(func() {
		db := &database{}
		db.connect()
		_auth = &Auth{db: db}
	})
	return _auth
}

package file

import (
	"sync"
	"time"
)

type Row struct {
	sync.Mutex

	// ID номер строки
	ID int64

	// Колонки
	Coefficient     CoefficientColumn
	Name            NameColumn
	Volume          VolumeColumn
	StartDateColumn StartDateColumn
	EndDateColumn   EndDateColumn

	// ErrorValue значение которое необходимо поместить в столбец ошибка
	ErrorValue []string
}

func (r *Row) AddErrorValue(errorValue string) {
	r.Lock()
	r.ErrorValue = append(r.ErrorValue, errorValue)
	r.Unlock()
}

type ErrorInfo struct {
	sync.Mutex
	// ErrorComment комментарий при ошибочном значении
	ErrorComment []string
}

func (ei *ErrorInfo) AddComment(comment string) {
	ei.Lock()
	ei.ErrorComment = append(ei.ErrorComment, comment)
	ei.Unlock()
}

type CoefficientColumn struct {
	// ID номер столбца (нужен для указания комента)
	ID int64
	// Value значение
	Value int
	ErrorInfo
}

type NameColumn struct {
	// ID номер столбца (нужен для указания комента)
	ID int64
	// Value значение
	Value string
	ErrorInfo
}

type VolumeColumn struct {
	// ID номер столбца (нужен для указания комента)
	ID int64
	// Value значение
	Value int64
	ErrorInfo
}

type StartDateColumn struct {
	// ID номер столбца (нужен для указания комента)
	ID int64
	// Value значение
	Value time.Time
	ErrorInfo
}

type EndDateColumn struct {
	// ID номер столбца (нужен для указания комента)
	ID int64
	// Value значение
	Value time.Time
	ErrorInfo
}

type ParsedFile struct {
	FileRows []Row
}

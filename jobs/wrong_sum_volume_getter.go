package jobs

import (
	"awesomeProject/file"
	"sync"
	"time"
)

type WrongSumVolumeGetter struct {
	// Общие данные для всех валидаций, тут он передаетс для того чтобы с данными из файла куда либо сбегать
	// Например за весовыми товарами, он пробежится по файлу, соберет идишки и сгоняет за ними
	ParsedFile *file.ParsedFile
	// Канал для передачи внештатной ситуации
	errChan chan error

	wrongSumVolume WrongSumVolume

	wg *sync.WaitGroup

	// Канал, который закроется, когда валидации закончатся
	closeChannel  chan interface{}
	resultChannel chan interface{}
}

func NewWrongVolumeGetter(wg *sync.WaitGroup, parsedFile *file.ParsedFile, errChan chan error) *WrongSumVolumeGetter {
	return &WrongSumVolumeGetter{
		wg:            wg,
		errChan:       errChan,
		ParsedFile:    parsedFile,
		closeChannel:  make(chan interface{}),
		resultChannel: make(chan interface{}),
	}
}

func (e *WrongSumVolumeGetter) Start() {
	// Имитируем поход за данными
	time.Sleep(50 * time.Second)
	e.wrongSumVolume = WrongSumVolume{120}
	for {
		select {
		case <-e.closeChannel:
			close(e.resultChannel)
			return
		case e.resultChannel <- e.wrongSumVolume:
		}
	}
}

func (e *WrongSumVolumeGetter) GetResult() chan interface{} {
	return e.resultChannel
}

func (e *WrongSumVolumeGetter) Stop() {
	close(e.closeChannel)
}

func (e *WrongSumVolumeGetter) GetName() string {
	return "wrong_sum_volume_getter"
}

func (e *WrongSumVolumeGetter) SetDependencies(jobs []Job) {

}

func (e *WrongSumVolumeGetter) GetDependenciesNames() []string {
	return []string{}
}

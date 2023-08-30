package jobs

import (
	"awesomeProject/file"
	"sync"
	"time"
)

type ExtraVolumeGetter struct {
	// Общие данные для всех валидаций, тут он передаетс для того чтобы с данными из файла куда либо сбегать
	// Например за весовыми товарами, он пробежится по файлу, соберет идишки и сгоняет за ними
	ParsedFile *file.ParsedFile
	// Канал для передачи внештатной ситуации
	errChan chan error

	extraVolume ExtraVolume

	wg *sync.WaitGroup

	// Канал, который закроется, когда валидации закончатся
	closeChannel  chan interface{}
	resultChannel chan interface{}
}

func NewExtraVolumeGetter(wg *sync.WaitGroup, parsedFile *file.ParsedFile, errChan chan error) *ExtraVolumeGetter {
	return &ExtraVolumeGetter{
		wg:            wg,
		errChan:       errChan,
		ParsedFile:    parsedFile,
		closeChannel:  make(chan interface{}),
		resultChannel: make(chan interface{}),
	}
}

func (e *ExtraVolumeGetter) Start() {
	// Имитируем поход за данными
	time.Sleep(30 * time.Second)
	e.extraVolume = ExtraVolume{100}
	for {
		select {
		case <-e.closeChannel:
			close(e.resultChannel)
			return
		case e.resultChannel <- e.extraVolume:
		}
	}
}

func (e *ExtraVolumeGetter) GetResult() chan interface{} {
	return e.resultChannel
}

func (e *ExtraVolumeGetter) Stop() {
	close(e.closeChannel)
}

func (e *ExtraVolumeGetter) GetName() string {
	return "extra_volume_getter"
}

func (e *ExtraVolumeGetter) SetDependencies(jobs []Job) {

}

func (e *ExtraVolumeGetter) GetDependenciesNames() []string {
	return []string{}
}

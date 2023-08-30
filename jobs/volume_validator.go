package jobs

import (
	"awesomeProject/file"
	"errors"
	"fmt"
	"sync"
)

type VolumeValidator struct {
	// Общие данные для всех валидаций
	ParsedFile   *file.ParsedFile
	wg           *sync.WaitGroup
	dependencies []Job
	// Канал для передачи внештатной ситуации
	errChan chan error

	// Дополнительные данные, которые необходимы для обработки валидации
	extraVolume    ExtraVolume
	wrongSumVolume WrongSumVolume
}

func NewVolumeValidator(wg *sync.WaitGroup, parsedFile *file.ParsedFile, errChan chan error) *VolumeValidator {
	return &VolumeValidator{
		ParsedFile: parsedFile,
		wg:         wg,
		errChan:    errChan,
	}
}

func (v *VolumeValidator) GetResult() chan interface{} {
	ch := make(chan interface{})
	close(ch)
	return ch
}

func (v *VolumeValidator) Start() {
	// Говорим, что закончили валидацию
	defer v.wg.Done()

	var sumVolume int64 = 0

	// Тут можно распаралелить как хотим
	for i := range v.ParsedFile.FileRows {
		sumVolume += v.ParsedFile.FileRows[i].Volume.Value

		// Если объем четный, считаем за ошибку для комента
		if v.ParsedFile.FileRows[i].Volume.Value%2 == 0 {
			v.ParsedFile.FileRows[i].Volume.ErrorInfo.AddComment("О-оу объем четный!")
		}
	}

	// Получаем зависимые данные, так как все равно нам нужны все данные, то можем подождать
	// Так же мы знаем какие именно данные нужны для этой валидации
	for _, job := range v.dependencies {
		result := <-job.GetResult()
		err := v.checkResult(result)
		if err != nil {
			v.errChan <- err
			return
		}
	}

	// Предполагаемая ситуация, когда ошибку помещаем в колоннку - "ошибка"
	if sumVolume+v.extraVolume.Volume == v.wrongSumVolume.Volume {
		for i := range v.ParsedFile.FileRows {
			v.ParsedFile.FileRows[i].AddErrorValue(
				fmt.Sprintf("Неверный суммарный объем на файл, суммарный объем не должен быть равен - %d. "+
					"Изначальный объем: %d, Объем с екстра: %d", v.wrongSumVolume.Volume, sumVolume, sumVolume+v.extraVolume.Volume),
			)
		}
	}
}

func (v *VolumeValidator) Stop() {
	return
}

func (v *VolumeValidator) GetName() string {
	return "volume_validator"
}

func (v *VolumeValidator) SetDependencies(jobs []Job) {
	v.dependencies = jobs
}

func (v *VolumeValidator) GetDependenciesNames() []string {
	return []string{"extra_volume_getter", "wrong_sum_volume_getter"}
}

func (v *VolumeValidator) checkResult(result interface{}) error {
	switch t := result.(type) {
	case WrongSumVolume:
		v.wrongSumVolume = t
		return nil
	case ExtraVolume:
		v.extraVolume = t
		return nil
	default:
		return errors.New("неожиданные данные")
	}
}

package jobs

import (
	"awesomeProject/file"
	"sync"
)

type CoefValidator struct {
	// Общие данные для всех валидаций
	ParsedFile   *file.ParsedFile
	wg           *sync.WaitGroup
	dependencies []Job
	// Канал для передачи внештатной ситуации
	errChan chan error
}

func NewCoefValidator(wg *sync.WaitGroup, parsedFile *file.ParsedFile, errChan chan error) *CoefValidator {
	return &CoefValidator{
		ParsedFile: parsedFile,
		wg:         wg,
		errChan:    errChan,
	}
}

func (v *CoefValidator) GetResult() chan interface{} {
	ch := make(chan interface{})
	close(ch)
	return ch
}

func (v *CoefValidator) Start() {
	// Говорим, что закончили валидацию
	defer v.wg.Done()

	// Тут можно распаралелить как хотим
	for i := range v.ParsedFile.FileRows {
		if v.ParsedFile.FileRows[i].Coefficient.Value < 1 || v.ParsedFile.FileRows[i].Coefficient.Value > 3 {
			v.ParsedFile.FileRows[i].Coefficient.AddComment("Коэффициент должен быть 1 <= k <= 3")
		}
	}
}

func (v *CoefValidator) Stop() {
	return
}

func (v *CoefValidator) GetName() string {
	return "coef_validator"
}

func (v *CoefValidator) SetDependencies(jobs []Job) {
	v.dependencies = jobs
}

func (v *CoefValidator) GetDependenciesNames() []string {
	return []string{}
}

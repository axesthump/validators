package builder

import (
	"awesomeProject/file"
	"awesomeProject/jobs"
	"sync"
)

type coefValidatorJobBuilder struct {
}

func (coefValidatorJobBuilder) build(wg *sync.WaitGroup, parsedFile *file.ParsedFile, errChan chan error) jobs.Job {
	return jobs.NewCoefValidator(wg, parsedFile, errChan)
}

package builder

import (
	"awesomeProject/file"
	"awesomeProject/jobs"
	"sync"
)

type volumeValidatorJobBuilder struct{}

func (volumeValidatorJobBuilder) build(wg *sync.WaitGroup, parsedFile *file.ParsedFile, errChan chan error) jobs.Job {
	return jobs.NewVolumeValidator(wg, parsedFile, errChan)
}

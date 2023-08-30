package builder

import (
	"awesomeProject/file"
	"awesomeProject/jobs"
	"sync"
)

type extraVolumeGetterJobBuilder struct{}

func (extraVolumeGetterJobBuilder) build(wg *sync.WaitGroup, parsedFile *file.ParsedFile, errChan chan error) jobs.Job {
	return jobs.NewExtraVolumeGetter(wg, parsedFile, errChan)
}

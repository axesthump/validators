package builder

import (
	"awesomeProject/file"
	"awesomeProject/jobs"
	"sync"
)

type wrongSumVolumeGetterJobBuilder struct{}

func (wrongSumVolumeGetterJobBuilder) build(wg *sync.WaitGroup, parsedFile *file.ParsedFile, errChan chan error) jobs.Job {
	return jobs.NewWrongVolumeGetter(wg, parsedFile, errChan)
}

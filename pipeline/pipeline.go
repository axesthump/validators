package pipeline

import (
	"awesomeProject/builder"
	"awesomeProject/file"
	"fmt"
	"sync"
)

type PipelineStarter interface {
	StartJobs() error
}

type PipelineImpl struct {
	parsedFile *file.ParsedFile
	// Имена джоб которые нужно будет стартануть
	jobNames []string

	wg      *sync.WaitGroup
	errChan chan error
}

func NewPipeline(parsedFile *file.ParsedFile, jobNames []string) *PipelineImpl {
	return &PipelineImpl{
		parsedFile: parsedFile,
		jobNames:   jobNames,
		wg:         &sync.WaitGroup{},
		errChan:    make(chan error, len(jobNames)),
	}
}

func (p *PipelineImpl) StartJobs() error {
	// Сортируем имена
	// p.jobName.sort()

	jobs, err := builder.BuildJobs(p.jobNames, p.wg, p.parsedFile, p.errChan)
	if err != nil {
		return fmt.Errorf("cant create jobs: %w", err)
	}

	// Тут можно посчитать для скольких джоб нам нужно апать wg
	// Для джоб сборки данных это не нужно, так как они бесконечно пытаются
	// отдать свои данные, пока примем за веру, что надо отдать только по 2
	p.wg.Add(2)

	for _, job := range jobs {
		go job.Start()
	}

	p.wg.Wait()
	select {
	case err = <-p.errChan:
		for _, job := range jobs {
			job.Stop()
		}
		return err
	default:
		for _, job := range jobs {
			job.Stop()
		}
		return nil
	}
}

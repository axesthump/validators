package builder

import (
	"awesomeProject/file"
	"awesomeProject/jobs"
	"errors"
	"fmt"
	"sync"
)

// build создает джобы
type build interface {
	build(wg *sync.WaitGroup, parsedFile *file.ParsedFile, errChan chan error) jobs.Job
}

// builders Тут будут все билдеры для наших джоб
var builders = map[string]build{
	"extra_volume_getter":     extraVolumeGetterJobBuilder{},
	"volume_validator":        volumeValidatorJobBuilder{},
	"wrong_sum_volume_getter": wrongSumVolumeGetterJobBuilder{},
	"coef_validator":          coefValidatorJobBuilder{},
}

// BuildJobs создает джобы, имена должны придти уже в отсортированном виде
func BuildJobs(names []string, wg *sync.WaitGroup, parsedFile *file.ParsedFile, errChan chan error) ([]jobs.Job, error) {
	createdJobsMap := make(map[string]jobs.Job, len(names))
	createdJobsList := make([]jobs.Job, 0, len(names))

	for _, name := range names {
		jobBuilder, ok := builders[name]
		if !ok {
			return nil, errors.New(fmt.Sprintf("cant create job with name: %s", name))
		}
		createdJobsMap[name] = jobBuilder.build(wg, parsedFile, errChan)
		jobDependenciesNames := createdJobsMap[name].GetDependenciesNames()
		if len(jobDependenciesNames) != 0 {
			jobDependencies := make([]jobs.Job, 0, len(jobDependenciesNames))
			for _, dependenciesName := range jobDependenciesNames {
				existDependencie, ok := createdJobsMap[dependenciesName]
				if !ok {
					return nil, errors.New(fmt.Sprintf("cant create dependencie job with name: %s for job: %s", dependenciesName, name))
				}
				jobDependencies = append(jobDependencies, existDependencie)
			}
			createdJobsMap[name].SetDependencies(jobDependencies)
		}
		createdJobsList = append(createdJobsList, createdJobsMap[name])
	}

	return createdJobsList, nil
}

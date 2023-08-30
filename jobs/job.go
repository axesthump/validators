package jobs

type Job interface {
	// GetResult возвращает результат выполнения
	GetResult() chan interface{} // можнет возвращать ид при прихронке в редис и тип

	// Start запускаем джобу
	Start()

	// Stop останавливаем джобу, больше нужно для джоб сборщиков данных
	Stop()

	// GetName возвращает имя джобы
	GetName() string

	// SetDependencies инитит зависимости
	SetDependencies(jobs []Job)

	// GetDependenciesNames возвращает имена зависимостей
	GetDependenciesNames() []string
}

// ExtraVolume какие-то данные которые понадобятся другим джобам
type ExtraVolume struct {
	Volume int64
}

// WrongSumVolume какие-то данные которые понадобятся другим джобам
type WrongSumVolume struct {
	Volume int64
}

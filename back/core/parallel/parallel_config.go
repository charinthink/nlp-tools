package parallelconfig

const (
	cpu             = 4
	errorBufferSize = 1000
	embedBufferSize = 100000
	scanBufferSize  = 256 * 1024
	workerPoolSize  = 12
)

type config struct {
	Enabled       bool
	RoutineConfig routineConfig
}
type routineConfig struct {
	Cpu             int
	ErrorBufferSize int
	EmbedBufferSize int
	ScanBufferSize  int
	WorkerPoolSize  int
}

type Config interface {
	IsEnabled() bool
	Get() config
}

func Default() Config {
	return &config{
		true,
		routineConfig{
			cpu,
			errorBufferSize,
			embedBufferSize,
			scanBufferSize,
			workerPoolSize,
		},
	}
}

func Set(core, errorBufferSize, embedBufferSize, scanBufferSize, workerPoolSize int) Config {
	return &config{
		true,
		routineConfig{
	      core,
			 errorBufferSize,
			 embedBufferSize,
			 scanBufferSize,
			  workerPoolSize,
		},
	}
}

func (c *config) Get() config {
	return *c
}

func (c *config) IsEnabled() bool {
	return c.Enabled
}

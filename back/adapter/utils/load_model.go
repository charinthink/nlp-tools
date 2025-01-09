package utils

import (
	"bufio"
	"io"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync"

	"github.com/charinthink/thai-sentence/back/core/embedding"
	parallelconfig "github.com/charinthink/thai-sentence/back/core/parallel"
	"github.com/pkg/errors"
)

func openFile(fileName string) (*os.File, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func LoadModel(modelName string, parallelConfig parallelconfig.Config) (embedding.Embeddings, error) {
	if parallelConfig == nil {
		return loadModelVec(modelName)
	}
	return loadlParallelModelVec(modelName, parallelConfig)
}

func loadlParallelModelVec(modelName string, parallelConfig parallelconfig.Config) (embedding.Embeddings, error) {
	file, err := openFile(modelName)
	if err != nil {
		return nil, err
	}
	embs, err := loadFileParallel(file, parallelConfig)
	if err != nil {
		return nil, err
	}
	file.Close()
	return embs, nil
}

func loadModelVec(modelName string) (embedding.Embeddings, error) {
	file, err := openFile(modelName)
	if err != nil {
		return nil, err
	}
	embs, err := loadFile(file)
	if err != nil {
		return nil, err
	}
	file.Close()
	return embs, nil
}

func loadFileParallel(r io.Reader, p parallelconfig.Config) (embedding.Embeddings, error) {
	configs := p.Get()
	runtime.GOMAXPROCS(configs.RoutineConfig.Cpu)
	s := bufio.NewScanner(r)
	buf := make([]byte, configs.RoutineConfig.ScanBufferSize)
	s.Buffer(buf, configs.RoutineConfig.ScanBufferSize)

	errChan := make(chan error, configs.RoutineConfig.ErrorBufferSize)
	embChan := make(chan embedding.Embedding, configs.RoutineConfig.EmbedBufferSize)
	lineChan := make(chan string, configs.RoutineConfig.EmbedBufferSize)

	var processWg sync.WaitGroup

	processWg.Add(configs.RoutineConfig.WorkerPoolSize)
	for i := 0; i < configs.RoutineConfig.WorkerPoolSize; i++ {
		go func() {
			defer processWg.Done()
			for line := range lineChan {
				emb, err := parseLine(line)
				if err != nil {
					errChan <- err
					break
				}
				embChan <- emb
			}
		}()
	}

	go func() {
		defer close(lineChan)
		for s.Scan() {
			line := s.Text()
			if strings.HasPrefix(line, " ") {
				continue
			}
			lineChan <- line
		}
		if err := s.Err(); err != nil {
			errChan <- err
		}
	}()

	go func() {
		processWg.Wait()
		close(embChan)
		close(errChan)
	}()

	var embs []embedding.Embedding
	for {
		select {
		case err := <-errChan:
			if err != nil {
				return nil, err
			}
		case emb, ok := <-embChan:
			if !ok {
				return embs, nil
			}
			embs = append(embs, emb)
		}
	}
}

func loadFile(r io.Reader) (embedding.Embeddings, error) {
	s := bufio.NewScanner(r)
	var embs []embedding.Embedding

	for s.Scan() {
		line := s.Text()
		if strings.HasPrefix(line, " ") {
			continue
		}
		emb, err := parseLine(line)
		if err != nil {
			return nil, err
		}
		embs = append(embs, emb)
	}
	if err := s.Err(); err != nil {
		return nil, err
	}
	return embs, nil
}

func parseLine(line string) (embedding.Embedding, error) {
	slice := strings.Fields(line)
	if len(slice) < 2 {
		return embedding.Embedding{}, errors.New("Must be over 2 lenghth for word and vector elems")
	}

	word := slice[0]
	vector := slice[1:]
	dim := len(vector)

	vec := make([]float64, dim)
	for k, elem := range vector {
		val, err := strconv.ParseFloat(elem, 64)
		if err != nil {
			return embedding.Embedding{}, err
		}
		vec[k] = val
	}
	return embedding.Embedding{Word: word, Vector: vec, Dim: dim}, nil
}

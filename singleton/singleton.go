package singleton

import (
	"bufio"
	"os"
	"path/filepath"
	"strconv"
	"sync"
)

type Repository interface {
	GetPopulation(name string) int
}

type repository struct {
	capitals map[string]int
}

func (r *repository) GetPopulation(name string) int {
	return r.capitals[name]
}

var repoInstance Repository

func readData(filename string) (map[string]int, error) {
	wdPath, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	file, err := os.Open(filepath.Join(wdPath, filename))
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	result := map[string]int{}

	for scanner.Scan() {
		k := scanner.Text()
		scanner.Scan()
		v, _ := strconv.Atoi(scanner.Text())
		result[k] = v
	}

	return result, nil
}

// both init and sync.Once are thread-safe
// but only sync.Once is lazy
// init() — we could, but it's not lazy
var once sync.Once

func GetRepoInstance() Repository {
	once.Do(func() {
		repo := repository{}
		caps, err := readData("capitals.db")
		if err == nil {
			repo.capitals = caps
		}
		repoInstance = &repo
	})
	return repoInstance
}

func CalculateTotalPopulation(repo Repository, cities ...string) int {
	result := 0
	for _, city := range cities {
		result += repo.GetPopulation(city)
	}
	return result
}

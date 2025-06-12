package singleton

import (
	"testing"
)

func TestGetPopulation(t *testing.T) {
	repo := GetRepoInstance()
	if pop := repo.GetPopulation("Tokyo"); pop != 33200000 {
		t.Errorf("expected 33200000, got %d", pop)
	}
}

func TestCalculateTotalPopulation(t *testing.T) {
	repo := GetRepoInstance()
	if total := CalculateTotalPopulation(repo, "Tokyo", "New York"); total != 51000000 {
		t.Errorf("expected 51000000, got %d", total)
	}
}

type FakeDatabase struct {
	dummyData map[string]int
}

func (d *FakeDatabase) GetPopulation(name string) int {
	if len(d.dummyData) == 0 {
		d.dummyData = map[string]int{
			"alpha": 1,
			"beta":  2,
			"gamma": 3}
	}
	return d.dummyData[name]
}

func TestGetPopulation_WithFakeDatabase(t *testing.T) {
	repo := &FakeDatabase{}
	if repo.GetPopulation("alpha") != 1 || repo.GetPopulation("gamma") != 3 {
		t.Errorf("expected fake data to be correct, got %+v", repo.dummyData)
	}
}

func TestCalculateTotalPopulation_WithFakeDatabase(t *testing.T) {
	repo := &FakeDatabase{}
	if total := CalculateTotalPopulation(repo, "alpha", "gamma"); total != 4 {
		t.Errorf("expected 4, got %d", total)
	}
}

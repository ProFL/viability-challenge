package viability

import (
	"testing"

	"github.com/ProFL/viability-challenge/config"
)

func TestFind(t *testing.T) {
	svc := Repository{
		Database:   "test",
		Connection: config.GetConnection(),
	}

	splitters, err := svc.Find(150, Point{-3.724348, -38.482041})

	if err != nil {
		t.Error(err)
	}

	t.Logf("The following splitters were found:")
	t.Log(splitters)

	if len(splitters) < 1 {
		t.Errorf("No splitters were found")
	}

	if len(splitters[0].ID.String()) < 1 {
		t.Errorf("Splitter is missing its id")
	}

	if len(splitters[0].Location.ID.String()) < 1 {
		t.Errorf("Splitter's location is missing its id")
	}

	if len(splitters[0].Location.Coordinates) != 2 {
		t.Errorf("There are more than two coordinates on splitter.Location.Coordinates")
	}
}

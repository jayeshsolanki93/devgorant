package devrant_test

import (
	"testing"

	"github.com/jayeshsolanki93/devgorant/devrant"
)

func TestProfile_jayeshs(t *testing.T) {
	devrant := devrant.New()
	user, _, err := devrant.Profile("jayeshs")
	if err != nil {
		t.Fatalf(`err != nil, got %v`, err)
	}

	if got, want := user.Username, "jayeshs"; got != want {
		t.Errorf(`user.Username = %s, want %s`, got, want)
	}

	if got, want := user.CreatedTime, 1463402493; got != want {
		t.Errorf(`user.CreatedTime = %d, want %d`, got, want)
	}
}

func TestRant_27317(t *testing.T) {
	devrant := devrant.New()
	rant, _, err := devrant.Rant(27317)
	if err != nil {
		t.Fatalf(`err != nil, got %v`, err)
	}

	if got, want := rant.UserId, 27292; got != want {
		t.Errorf(`rant.UserId = %d, want %d`, got, want)
	}
}

func TestRants(t *testing.T) {
	devrant := devrant.New()
	rants, err := devrant.Rants("algo", 50, 0)
	if err != nil {
		t.Fatalf(`err != nil, got %v`, err)
	}

	if got, want := len(rants), 50; got != want {
		t.Errorf(`len(rants) = %d, want > 0`, got)
	}
}

// TODO: More tests

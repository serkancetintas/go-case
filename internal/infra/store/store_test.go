package store

import (
	"testing"
)

func TestGet(t *testing.T) {
	store := NewStore()

	store.Set("Yas", "30")
	val := store.Get("Yas")

	if val != "30" {
		t.Errorf("Get value should be equal set value ")
	}
}

func TestSet(t *testing.T) {
	store := NewStore()

	store.Set("Yas", "30")

	if  store.data["Yas"] != "30" {
		t.Errorf("Value could not be set")
	}
}

func TestFlush(t *testing.T) {
	store := NewStore()

	store.Set("Yas", "30")
	store.Set("Adi", "Ahmet")

	store.Flush()
	if len(store.data) > 0 {
		t.Errorf("store length should be zero")
	}

}

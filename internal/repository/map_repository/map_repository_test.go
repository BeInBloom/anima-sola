package maprepository

import (
	"errors"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
)

const checks = 1000

func TestBaseSet(t *testing.T) {
	repo := New()
	for range checks {
		var key, value = gofakeit.UUID(), gofakeit.LoremIpsumSentence(10)

		if err := repo.Set(key, value); err != nil {
			t.Fatalf("случилась какая-то неведомая хуйня kekw: %s", err)
		}

		val, err := repo.Get(key)
		if err != nil {
			t.Fatalf("cant get result: %s", val)
		}
		if val != value {
			t.Fatalf("wrong value: %s", val)
		}
	}
}

func TestRebase(t *testing.T) {
	repo := New()

	for range checks {
		keys := getData(gofakeit.UUID)
		vals := getData(getWords)

		if err := setData(repo, keys, vals); err != nil {
			t.Fatalf("cant set data: %s", err)
		}

		if err := setCheck(repo, keys, vals); err != nil {
			t.Fatalf("wrong set: %s", err)
		}

		newVlas := getData(getWords)
		if err := setData(repo, keys, newVlas); err != nil {
			t.Fatalf("cant set data: %s", err)
		}

		if err := setCheck(repo, keys, newVlas); err != nil {
			t.Fatalf("wrong set: %s", err)
		}
	}
}

func TestNoKey(t *testing.T) {
	repo := New()

	for range checks {
		keys := getData(gofakeit.UUID)
		for _, key := range keys {
			if _, err := repo.Get(key); err == nil {
				t.Errorf("key in repo kekw: %s", key)
			}
		}
	}
}

func getData(gen func() string) []string {
	const numberOfData = 100

	data := make([]string, 0, numberOfData)
	for i := range data {
		data[i] = gen()
	}

	return data
}

func getWords() string {
	const numberOfWords = 10
	return gofakeit.LoremIpsumSentence(numberOfWords)
}

func setData(repo *Repo, keys, value []string) error {
	l := len(keys)
	if l-len(value) < 0 {
		l = len(value)
	}

	for i := range l {
		if err := repo.Set(keys[i], value[i]); err != nil {
			return err
		}
	}

	return nil
}

func setCheck(repo *Repo, keys, vals []string) error {
	if len(keys) != len(vals) {
		return errors.New("wrong values")
	}

	for i := range len(keys) {
		repoVal, err := repo.Get(keys[i])
		if err != nil {
			return err
		}
		if vals[i] != repoVal {
			return errors.New("wrong set")
		}
	}

	return nil
}

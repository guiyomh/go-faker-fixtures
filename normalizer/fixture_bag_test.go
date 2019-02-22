package normalizer

import (
	"fmt"
	"testing"

	"github.com/guiyomh/charlatan/contract"
	"github.com/stretchr/testify/assert"
)

func TestFixtureBag(t *testing.T) {
	t.Parallel()

	var bag contract.Bager = make(FixtureBag, 0)
	assert.Equal(t, 0, len(bag.(FixtureBag)))
	assert.False(t, bag.Has("foo"))

	testObj := new(contract.MockFixture)
	testObj.On("Id").Return("foo")
	bag = bag.Add(testObj)
	testObj.AssertExpectations(t)

	assert.Equal(t, 1, len(bag.(FixtureBag)))
	assert.True(t, bag.Has("foo"))

	got, err := bag.Get("foo")
	assert.NoError(t, err)
	assert.Equal(t, testObj, got)

	_, err = bag.Get("unknown")
	assert.Error(t, err)
}

func TestFixtureBagWithout(t *testing.T) {
	t.Parallel()

	fixtures := []contract.Fixture{}
	var bag contract.Bager = make(FixtureBag, 0)
	nb := 10
	for i := 1; i <= nb; i++ {
		f := new(contract.MockFixture)
		f.On("Id").Return(fmt.Sprintf("fix_%d", i))
		fixtures = append(fixtures, f)
		bag = bag.Add(f)
	}
	assert.Equal(t, nb, len(bag.(FixtureBag)))

	for _, f := range fixtures {
		m, _ := f.(*contract.MockFixture)
		m.AssertExpectations(t)
	}
	assert.True(t, bag.Has("fix_3"))

	bag = bag.Without(fixtures[2])
	assert.False(t, bag.Has("fix_3"))
	assert.Equal(t, nb-1, len(bag.(FixtureBag)))
}

func TestFixtureBagMerge(t *testing.T) {
	t.Parallel()
	fixtures := []contract.Fixture{}

	var bag contract.Bager = make(FixtureBag, 0)
	nb := 10
	for i := 1; i <= nb; i++ {
		f := new(contract.MockFixture)
		f.On("Id").Return(fmt.Sprintf("fix_%d", i))
		fixtures = append(fixtures, f)
		bag = bag.Add(f)
	}
	assert.Equal(t, nb, len(bag.(FixtureBag)))

	var bag2 contract.Bager = make(FixtureBag, 0)
	nb2 := 10
	for i := 1; i <= nb2; i++ {
		f := new(contract.MockFixture)
		f.On("Id").Return(fmt.Sprintf("mock_%d", i))
		bag2 = bag2.Add(f)
	}
	assert.Equal(t, nb2, len(bag2.(FixtureBag)))
	for _, f := range fixtures[0:3] {
		bag2 = bag2.Add(f)
	}
	assert.Equal(t, nb2+3, len(bag2.(FixtureBag)))
	bag3, err := bag.MergeWith(bag2)
	assert.Equal(t, nb+nb2, len(bag3.(FixtureBag)))
	assert.NoError(t, err)
}

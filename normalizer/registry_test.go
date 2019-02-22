package normalizer

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/guiyomh/charlatan/contract"
)

func TestRegistry_CanDenormalize(t *testing.T) {
	fields := map[string]interface{}{
		"username": "<Username()>",
	}
	data := map[string]map[string]interface{}{
		"user_{bob,alice}": fields,
	}
	expectedFixture := &contract.MockBager{}
	denorm1 := &contract.MockChainabler{}
	denorm1.On("CanDenormalize", "user_{bob,alice}").Return(true)
	denorm1.On("Denormalize", "user_{bob,alice}", fields).Return(expectedFixture, nil)
	denorm2 := &contract.MockChainabler{}

	registry := NewRegistry([]contract.Chainabler{
		denorm1,
		denorm2,
	})
	registry.Denormalize("user", data)
	denorm1.AssertExpectations(t)
	denorm2.AssertNotCalled(t, "CanDenormalize", "user_{bob,alice}")
}

func TestRegistry_CanDenormalizer2(t *testing.T) {
	fields := map[string]interface{}{
		"username": "<Username()>",
	}
	data := map[string]map[string]interface{}{
		"user_{bob,alice}": fields,
	}
	expectedFixture := &contract.MockBager{}
	denorm1 := &contract.MockChainabler{}
	denorm1.On("CanDenormalize", "user_{bob,alice}").Return(false)
	denorm2 := &contract.MockChainabler{}
	denorm2.On("CanDenormalize", "user_{bob,alice}").Return(true)
	denorm2.On("Denormalize", "user_{bob,alice}", fields).Return(expectedFixture, nil)

	registry := NewRegistry([]contract.Chainabler{
		denorm1,
		denorm2,
	})
	registry.Denormalize("user_{bob,alice}", data)
	denorm1.AssertExpectations(t)
	denorm1.AssertNotCalled(t, "Denormalize", "user_{bob,alice}", data)
	denorm2.AssertExpectations(t)
}

func TestRegistry_NonormalizerMatch(t *testing.T) {
	fields := map[string]interface{}{
		"username": "<Username()>",
	}
	data := map[string]map[string]interface{}{
		"user_{bob,alice}": fields,
	}
	denorm1 := &contract.MockChainabler{}
	denorm1.On("CanDenormalize", "user_{bob,alice}").Return(false)
	denorm2 := &contract.MockChainabler{}
	denorm2.On("CanDenormalize", "user_{bob,alice}").Return(false)

	registry := NewRegistry([]contract.Chainabler{
		denorm1,
		denorm2,
	})
	bag, err := registry.Denormalize("user_{bob,alice}", data)
	denorm1.AssertExpectations(t)
	denorm1.AssertNotCalled(t, "Denormalize", "user_{bob,alice}", fields)
	denorm2.AssertExpectations(t)
	denorm2.AssertNotCalled(t, "Denormalize", "user_{bob,alice}", fields)
	assert.NoError(t, err)
	assert.Len(t, bag, 0)
}

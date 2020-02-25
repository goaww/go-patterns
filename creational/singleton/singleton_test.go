package singleton

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetSingleton(t *testing.T) {
	instance := GetInstance()

	assert.NotNil(t, instance)

}

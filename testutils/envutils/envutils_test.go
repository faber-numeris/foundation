package envutils_test

import (
	"os"
	"testing"

	"github.com/faber-numeris/foundation/testutils/envutils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPushClearsEnvironmentAndPopRestoresIt(t *testing.T) {
	const key = "ENVUTILS_TEST_VAR"
	require.NoError(t, os.Setenv(key, "original"))
	t.Cleanup(func() { _ = os.Unsetenv(key) })

	require.NoError(t, envutils.Push())

	_, ok := os.LookupEnv(key)
	assert.False(t, ok, "Push should clear all environment variables")

	require.NoError(t, os.Setenv(key, "mutated"))

	require.NoError(t, envutils.Pop())

	got, ok := os.LookupEnv(key)
	require.True(t, ok, "Pop should restore the snapshotted variable")
	assert.Equal(t, "original", got, "Pop should discard changes made after Push")
}

func TestPushPopNesting(t *testing.T) {
	const key = "ENVUTILS_NESTED_VAR"
	require.NoError(t, os.Setenv(key, "outer"))
	t.Cleanup(func() { _ = os.Unsetenv(key) })

	require.NoError(t, envutils.Push())
	require.NoError(t, os.Setenv(key, "inner"))
	require.NoError(t, envutils.Push())

	require.NoError(t, envutils.Pop())
	got, _ := os.LookupEnv(key)
	assert.Equal(t, "inner", got, "first Pop restores the inner snapshot")

	require.NoError(t, envutils.Pop())
	got, _ = os.LookupEnv(key)
	assert.Equal(t, "outer", got, "second Pop restores the outer snapshot")
}

func TestPopOnEmptyStackReturnsError(t *testing.T) {
	assert.ErrorIs(t, envutils.Pop(), envutils.ErrEmptyStack)
}

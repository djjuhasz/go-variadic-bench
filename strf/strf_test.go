package strf_test

import (
	"fmt"
	"testing"

	"github.com/djjuhasz/go-variadic-bench/strf"

	"gotest.tools/v3/assert"
)

var (
	name     string   = "David"
	weekdays []string = []string{
		"Monday",
		"Tuesday",
		"Wednesday",
		"Thursday",
		"Friday",
		"Saturday",
		"Sunday",
	}
)

func TestNoop(t *testing.T) {
	t.Parallel()

	assert.Equal(t, strf.Noop("hello"), "hello")
}

func TestF(t *testing.T) {
	t.Parallel()

	assert.Equal(t,
		strf.F("Hello %s, happy %s!", "Frank", "Tuesday"),
		"Hello Frank, happy Tuesday!",
	)
}

func TestF2(t *testing.T) {
	t.Parallel()

	assert.Equal(t,
		strf.F2("Hello %s, happy %s!", "Frank", "Tuesday"),
		"Hello Frank, happy Tuesday!",
	)
}

func BenchmarkNoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strf.Noop(fmt.Sprintf("Hello %s, happy: %s!", name, weekdays[i%7]))
	}
}

func BenchmarkF(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strf.F("Hello %s, happy: %s!", name, weekdays[i%7])
	}
}

func BenchmarkF2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strf.F2("Hello %s, happy: %s!", name, weekdays[i%7])
	}
}

package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func BenchmarkHelloWorldTable(b *testing.B) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Hello World(Made)",
			request:  "Made",
			expected: "Hello Made",
		}, {
			name:     "Hello World(Garda)",
			request:  "Garda",
			expected: "Hello Garda",
		},
	}

	for _, test := range tests {
		b.Run(test.name, func(b *testing.B) {
			HelloWorld(test.request)
		})
	}

}

func BenchmarkSub(b *testing.B) {

	b.Run("Pertama", func(b *testing.B) {
		HelloWorld("Elang")
	})

	b.Run("Groovy", func(b *testing.B) {
		HelloWorld("Groovy")
	})

}

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Made")
	}
}

func BenchmarkHelloWorldMa(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Ma")
	}
}

func TestHelloWorldTable(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Hello World(Made)",
			request:  "Made",
			expected: "Hello Made",
		}, {
			name:     "Hello World(Garda)",
			request:  "Garda",
			expected: "Hello Garda",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			r := HelloWorld(test.request)
			require.Equal(t, test.expected, r, "must be Hello Elang")
		})
	}

}

func TestSubTest(t *testing.T) {

	t.Run("Pertama", func(t *testing.T) {
		r := HelloWorld("Elang")
		require.Equal(t, "Hello Elang", r, "must be Hello Elang")
	})

	t.Run("Groovy", func(t *testing.T) {
		r := HelloWorld("Groovy")
		assert.Equal(t, "Hi Groovy", r)
	})

}

func TestMain(m *testing.M) {
	fmt.Println("sebelum unit test")
	m.Run()
	fmt.Println("setelah")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("can't run on windows")
	}
	r := HelloWorld("Made")
	require.Equal(t, "Hello Made", r, "result must be hello eko")
}

func TestHelloWorldAssert(t *testing.T) {
	r := HelloWorld("Made")
	assert.Equal(t, "Hello Made", r, "result must be hello eko")
	fmt.Println("test Hello word assert done")

}

func TestHelloWorld(t *testing.T) {
	r := HelloWorld("Made")

	if r != "Hello Made" {
		t.Error("result must be hello made")
	}
	fmt.Println("test hello world done")

}

func TestHelloWorldSanskara(t *testing.T) {
	r := HelloWorld("Sanskara")

	if r != "Hello Sanskara" {
		t.Fatal("result must be hello sanskara")
	}
	fmt.Println("test hello done")
}

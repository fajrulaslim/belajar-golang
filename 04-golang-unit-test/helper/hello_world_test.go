package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func BenchmarkTable(b *testing.B) {
	benchmarks := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Eko",
			request:  "Eko",
			expected: "Hello Eko",
		},
		{
			name:     "Kurniawan",
			request:  "Kurniawan",
			expected: "Hello Kurniawan",
		},
		{
			name:     "Khannedy",
			request:  "Khannedy",
			expected: "Hello Khannedy",
		},
		{
			name:     "Budi",
			request:  "Budi",
			expected: "Hello Budi",
		},
		{
			name:     "Joko",
			request:  "Joko",
			expected: "Hello Joko",
		},
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(benchmark.request)
			}
		})
	}
}

func BenchmarkSubTest(b *testing.B) {
	b.Run("Benchmark 1", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Fajrul")
		}
	})
	b.Run("Benchmark 2", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Aslim")
		}
	})
}

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Fajrul")
	}
}

func TestHelloWorldTable(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Eko",
			request:  "Eko",
			expected: "Hello Eko",
		},
		{
			name:     "Kurniawan",
			request:  "Kurniawan",
			expected: "Hello Kurniawan",
		},
		{
			name:     "Khannedy",
			request:  "Khannedy",
			expected: "Hello Khannedy",
		},
		{
			name:     "Budi",
			request:  "Budi",
			expected: "Hello Budi",
		},
		{
			name:     "Joko",
			request:  "Joko",
			expected: "Hello Joko",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}

func TestSubTest(t *testing.T) {
	t.Run("Subtest 1", func(t *testing.T) {
		result := HelloWorld("Fajrul")
		require.Equal(t, "Hello Fajrul", result, "Pesan Error Require")
	})

	t.Run("Subtest 2", func(t *testing.T) {
		result := HelloWorld("Aslim")
		require.Equal(t, "Hello Aslim", result, "Pesan Error Require")
	})
}

func TestMain(m *testing.M) {
	// before
	fmt.Println("-----------")
	fmt.Println("Sebelum Tes")

	m.Run()

	// after
	fmt.Println("Sesudah Tes")
	fmt.Println("-----------")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Tidak bisa berjalan di windows")
	}
	fmt.Println("Tes Skip Windows")
}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Fajrul")
	assert.Equal(t, "Hello Fajrul", result, "Pesan Error Assert") // assert = fail
	fmt.Println("Test Assert Done")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Fajrul")
	require.Equal(t, "Hello Fajrul", result, "Pesan Error Require") // require = failnow
	fmt.Println("Test Require Done")
}

func TestHelloWorld1(t *testing.T) {
	result := HelloWorld("Fajrul")
	if result != "Hello Fajrul" {
		// unit test failed,
		// t.Fail() // gagal tetapi dilanjutkan ke kode berikutnya
		t.Error("Harusnya Hello Fajrul") // seperti fail tapi berisi pesan
	}
	fmt.Println("Test Hello World 1 Done")
}

func TestHelloWorld2(t *testing.T) {
	result := HelloWorld("Fajrul")
	if result != "Hello Fajrul" {
		// unit test failed,
		// t.FailNow() // berenti langsung
		t.Fatal("Harusnya Hello Fajrul") // seperti failnow tetapi berisi pesan
	}
	fmt.Println("Test Hello World 2 Done")
}

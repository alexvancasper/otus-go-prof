package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCopy(t *testing.T) {
	FullFileSize := int64(200)
	cases := []struct {
		name         string
		offset       int64
		limit        int64
		expectedSize int64
	}{
		{
			name:         "full file",
			offset:       0,
			limit:        0,
			expectedSize: FullFileSize,
		}, {
			name:         "first half file",
			offset:       0,
			limit:        100,
			expectedSize: 100,
		}, {
			name:         "second half file",
			offset:       100,
			limit:        0,
			expectedSize: 100,
		}, {
			name:         "offset equal file size",
			offset:       200,
			limit:        0,
			expectedSize: 0,
		}, {
			name:         "offset + limit",
			offset:       150,
			limit:        100,
			expectedSize: 50,
		}, {
			name:         "middle of file",
			offset:       50,
			limit:        100,
			expectedSize: 100,
		},
	}

	hFile, err := os.OpenFile("./input.dat", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0o666)
	if err != nil {
		t.Errorf("error %v", err)
	}

	buf := make([]byte, FullFileSize)
	_, err = hFile.Write(buf)
	defer hFile.Close()

	if err != nil {
		t.Errorf("error write initial file %v", err)
	}
	defer os.Remove("./input.dat")
	defer os.Remove("./out.dat")

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			err := Copy("./input.dat", "./out.dat", c.offset, c.limit)
			if err != nil {
				t.Errorf("error copy file %v", err)
			}
			dstFile, err := os.OpenFile("./out.dat", os.O_RDONLY, 0o644)
			if err != nil {
				t.Errorf("dst copy file %v", err)
			}
			defer dstFile.Close()
			dstFileStat, err := dstFile.Stat()
			if err != nil {
				t.Errorf("dst stat file %v", err)
			}
			require.Equal(t, c.expectedSize, dstFileStat.Size())
		})
	}

	t.Run("offset more than filesize", func(t *testing.T) {
		err := Copy("./input.dat", "./out.dat", 250, 0)
		require.Equal(t, ErrOffsetExceedsFileSize, err)
	})

	t.Run("unsupported file", func(t *testing.T) {
		err := Copy("/dev/random", "./out.dat", 0, 0)
		require.Equal(t, ErrUnsupportedFile, err)
	})
}

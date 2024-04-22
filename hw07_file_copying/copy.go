package main

import (
	"errors"
	"io"
	"os"
	"time"

	"github.com/cheggaaa/pb/v3"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds filesize")
)

func Copy(fromPath, toPath string, offset, limit int64) error {
	hInFile, err := os.OpenFile(fromPath, os.O_RDONLY, 0o644)
	if err != nil {
		return err
	}
	defer hInFile.Close()

	inFileInfo, err := hInFile.Stat()
	if err != nil {
		return err
	}

	if !inFileInfo.Mode().IsRegular() {
		return ErrUnsupportedFile
	}

	if offset > inFileInfo.Size() {
		return ErrOffsetExceedsFileSize
	}
	_, err = hInFile.Seek(offset, 0)
	if err != nil {
		return ErrOffsetExceedsFileSize
	}

	limit = checkLimit(inFileInfo.Size(), limit, offset)

	hOutFile, err := os.OpenFile(toPath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0o666)
	if err != nil {
		return err
	}
	defer hOutFile.Close()

	return lcopy(hInFile, hOutFile, limit)
}

func checkLimit(fileSize, limit, offset int64) int64 {
	if limit == 0 {
		limit = fileSize
	}

	if limit > fileSize-offset {
		limit = fileSize - offset
	}
	return limit
}

func lcopy(src io.Reader, dst io.Writer, size int64) error {
	bar := pb.New64(size)
	for i := int64(0); i < size; i++ {
		_, err := io.CopyN(dst, src, 1)
		bar.Increment()
		if errors.Is(err, io.EOF) {
			break
		}

		time.Sleep(time.Millisecond)
	}
	bar.Finish()
	return nil
}

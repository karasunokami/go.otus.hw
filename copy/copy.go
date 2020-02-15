package copy

import (
	"errors"
	"github.com/cheggaaa/pb/v3"
	"io"
	"os"
)

var bufSize = 1024

func Copy(from, to string, limit, offset int) error {
	fi, err := os.Stat(from)
	if err != nil {
		return err
	}

	fileSize := int(fi.Size())

	err = validate(offset, limit, fileSize)
	if err != nil {
		return err
	}

	var size int

	if limit > 0 {
		size = limit
	} else {
		size = fileSize - offset
	}

	fileFrom, err := os.Open(from)
	if err != nil {
		return err
	}

	fileTo, err := os.Create(to)
	if err != nil {
		return err
	}

	bar := pb.StartNew(size)
	buf := make([]byte, bufSize)
	toFileOffset := 0

	for {
		if toFileOffset > size {
			break
		}

		if size-toFileOffset < bufSize {
			buf = make([]byte, size-toFileOffset)
		}

		err := read(fileFrom, offset, bufSize, &buf)
		if err != nil {
			return err
		}

		err = write(fileTo, toFileOffset, &buf)
		if err != nil {
			return err
		}

		toFileOffset += bufSize
		offset += bufSize

		bar.Add(bufSize)
	}

	err = fileTo.Close()
	if err != nil {
		return err
	}

	err = fileFrom.Close()
	if err != nil {
		return err
	}

	bar.Finish()

	return nil
}

func read(file *os.File, offset, size int, buf *[]byte) error {
	_, err := file.Seek(int64(offset), 0)
	if err != nil {
		return err
	}

	offset = 0

	for offset < size {
		read, err := file.Read((*buf)[offset:])
		if err != nil {
			return err
		}

		offset += read

		if err == io.EOF || read == 0 {
			break
		}
	}

	return nil
}

func write(file *os.File, offset int, buf *[]byte) error {
	_, err := file.Seek(int64(offset), 0)
	if err != nil {
		return err
	}

	_, err = file.Write(*buf)
	if err != nil {
		return err
	}

	return nil
}

func validate(offset, limit, fileSize int) error {
	if offset < 0 || offset > fileSize {
		return errors.New("offset cant be more then file size or less then 0")
	} else if limit < 0 || limit > fileSize {
		return errors.New("limit cant be more then file size or less then 0")
	}

	if (offset + limit) > fileSize {
		return errors.New("failed to copy file - offset + limit more then file size")
	}

	return nil
}

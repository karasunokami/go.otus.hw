package hw6

import (
	"errors"
	"github.com/cheggaaa/pb/v3"
	"io"
	"os"
)

func Copy(from, to string, limit, offset int) error {
	var size int

	fi, err := os.Stat(from)
	if err != nil {
		return err
	}

	fileSize := int(fi.Size())

	err = validate(offset, limit, fileSize)
	if err != nil {
		return err
	}

	if limit == 0 {
		size = fileSize - offset
	} else {
		size = limit
	}

	bar := pb.StartNew(size)

	buf, err := read(from, offset, size, bar)
	if err != nil {
		return err
	}

	test := string(buf)
	_ = test

	err = write(to, buf)
	if err != nil {
		return err
	}

	bar.Finish()

	return nil
}

func read(from string, offset, size int, bar  *pb.ProgressBar) ([]byte, error) {
	buf := make([]byte, size)

	file, err := os.Open(from)
	if err != nil {
		return buf, err
	}

	_, err = file.Seek(int64(offset), 0)
	if err != nil {
		return buf, err
	}

	for size > 0 {
		read, err := file.Read(buf)
		if err != nil {
			return buf, err
		}

		size -= read
		bar.Add(read)

		if err == io.EOF || read == 0 {
			break
		}
	}

	err = file.Close()
	if err != nil {
		return buf, err
	}

	return buf, nil
}

func write(to string, buf []byte) error {
	file, _ := os.Create(to)
	_, err := file.Write(buf)
	if err != nil {
		return err
	}

	err = file.Close()
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
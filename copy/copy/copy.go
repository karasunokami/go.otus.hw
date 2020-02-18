package copy

import (
	"errors"
	"github.com/cheggaaa/pb/v3"
	"io"
	"os"
)

const copySize = 1024 * 32

func Copy(from, to string, limit, offset int) error {
	file, err := os.Open(from)
	if err != nil {
		return err
	}

	defer file.Close()

	info, err := file.Stat()
	if err != nil {
		return err
	}

	fileSize := info.Size()

	err = validate(offset, limit, int(fileSize))
	if err != nil {
		return err
	}

	input := io.Reader(file)
	if limit != 0 {
		input = io.LimitReader(file, int64(limit))
		fileSize = int64(limit)
	}

	if offset > 0 {
		offset64 := int64(offset)
		if pos, err := file.Seek(offset64, 0); err != nil || pos != offset64 {
			return err
		}

		if offset64+fileSize >= info.Size() {
			fileSize = info.Size() - offset64
		}
	}

	bar := pb.StartNew(int(fileSize))
	bar.SetWriter(os.Stdout)

	output, err := os.Create(to)
	if err != nil {
		return err
	}

	defer output.Close()

	for totalWritten := int64(0); totalWritten < fileSize; {
		written, err := io.CopyN(output, input, copySize)
		if err != nil {
			if err == io.EOF {
				bar.Add64(written)
				break
			}
			return err
		}

		totalWritten += written
		bar.Add64(written)
	}

	bar.Finish()

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

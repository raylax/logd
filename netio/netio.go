package netio

import (
	"encoding/json"
	"github.com/raylax/logd/model"
	"io"
)

func WriteLog(writer io.Writer, logLine *model.LogLine) error {
	data, _ := json.Marshal(logLine)
	err := WriteInt16(writer, len(data))
	if err != nil {
		return err
	}
	return WriteData(writer, data)
}

func ReadLog(reader io.Reader) (*model.LogLine, error) {
	dataLen, err := ReadInt16(reader)
	if err != nil {
		return nil, err
	}
	data, err := ReadData(reader, dataLen)
	if err != nil {
		return nil, err
	}
	logLine := &model.LogLine{}
	_ = json.Unmarshal(data, logLine)
	return logLine, nil
}

func WriteInt16(writer io.Writer, n int) error {
	return WriteData(writer, []byte{
		byte(n),
		byte(n >> 8),
	})
}

func ReadInt16(reader io.Reader) (int, error) {
	data, err := ReadData(reader, 2)
	if err != nil {
		return 0, err
	}
	var n int
	n |= int(data[0])
	n |= int(data[1]) << 8
	return n, nil
}

func WriteData(writer io.Writer, data []byte) error {
	_, err := writer.Write(data)
	if err != nil {
		return err
	}
	return nil
}

func ReadData(reader io.Reader, n int) ([]byte, error) {
	data := make([]byte, n)
	_, err := reader.Read(data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

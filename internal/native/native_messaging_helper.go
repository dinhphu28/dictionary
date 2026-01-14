package native

import (
	"encoding/binary"
	"encoding/json"
	"io"
	"os"
)

func ReadMessage() ([]byte, error) {
	var length uint32
	err := binary.Read(os.Stdin, binary.LittleEndian, &length)
	if err != nil {
		return nil, err
	}

	buf := make([]byte, length)
	_, err = io.ReadFull(os.Stdin, buf)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func WriteMessage(v any) error {
	payload, err := json.Marshal(v)
	if err != nil {
		return err
	}

	// length prefix
	err = binary.Write(os.Stdout, binary.LittleEndian, uint32(len(payload)))
	if err != nil {
		return err
	}

	_, err = os.Stdout.Write(payload)
	if err != nil {
		return err
	}

	os.Stdout.Sync()
	return nil
}

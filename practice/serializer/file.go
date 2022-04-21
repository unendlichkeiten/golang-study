package serializer

import (
	"fmt"
	"io/ioutil"

	"github.com/golang/protobuf/proto"
)

// WriteProtobufToJSONFile writes protocol buffer to JSON file
func WriteProtobufToJSONFile(message proto.Message, fileName string) error {
	data, err := ProtobufToJSON(message)
	if err != nil {
		return fmt.Errorf("cannot marshal proto message to JSON: %s", err)
	}

	err = ioutil.WriteFile(fileName, []byte(data), 0644)
	if err != nil {
		return fmt.Errorf("cannot write JSON data to file: %s", err)
	}
	return nil
}

// WriteProtobufToBinaryFile writes protocol buffer to binary file
func WriteProtobufToBinaryFile(message proto.Message, fileName string) error {
	data, err := proto.Marshal(message)
	if err != nil {
		return fmt.Errorf("cannot marshal proto message to binary: %w", err)
	}

	err = ioutil.WriteFile(fileName, []byte(data), 0644)
	if err != nil {
		return fmt.Errorf("cannot write binary data to file: %s", err)
	}

	return nil
}

// ReadProtobufFromBinaryFile reads protocol buffer message from binary file
func ReadProtobufFromBinaryFile(fileName string, message proto.Message) error {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return fmt.Errorf("cannot read binary data from file: %s", err)
	}

	err = proto.Unmarshal(data, message)
	if err != nil {
		return fmt.Errorf("cannot unmarshal binary to proto message: %s", err)
	}
	return nil
}

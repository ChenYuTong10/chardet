package chardet

import (
	"bufio"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBOMUTF8Detector(t *testing.T) {

	for i := 1; i <= 3; i++ {
		f, err := os.Open(fmt.Sprintf(".\\raw\\BOM UTF8 %d.txt", i))
		if err != nil {
			t.Errorf("open file error: %v", err)
		}

		buf := make([]byte, 1024)
		reader := bufio.NewReader(f)
		_, err = reader.Read(buf)
		if err != nil {
			t.Errorf("read buf error: %v", err)
		}

		d := new(Detector)
		d.Feed(buf)
		assert.Equal(t, d.Encoding, BOM_UTF8)

		defer f.Close()
	}
}

func TestBOMUTF16BEDetector(t *testing.T) {

	for i := 1; i <= 3; i++ {
		f, err := os.Open(fmt.Sprintf(".\\raw\\BOM UTF16 BE %d.txt", i))
		if err != nil {
			t.Errorf("open file error: %v", err)
		}

		buf := make([]byte, 1024)
		reader := bufio.NewReader(f)
		_, err = reader.Read(buf)
		if err != nil {
			t.Errorf("read buf error: %v", err)
		}

		d := new(Detector)
		d.Feed(buf)
		assert.Equal(t, d.Encoding, BOM_UTF16_BE)

		defer f.Close()
	}
}

func TestBOMUTF16LEDetector(t *testing.T) {

	for i := 1; i <= 3; i++ {
		f, err := os.Open(fmt.Sprintf(".\\raw\\BOM UTF16 LE %d.txt", i))
		if err != nil {
			t.Errorf("open file error: %v", err)
		}

		buf := make([]byte, 1024)
		reader := bufio.NewReader(f)
		_, err = reader.Read(buf)
		if err != nil {
			t.Errorf("read buf error: %v", err)
		}

		d := new(Detector)
		d.Feed(buf)
		assert.Equal(t, d.Encoding, BOM_UTF16_LE)

		defer f.Close()
	}
}

func TestUTF8Detector(t *testing.T) {

	for i := 1; i <= 3; i++ {
		f, err := os.Open(fmt.Sprintf(".\\raw\\UTF8 %d.txt", i))
		if err != nil {
			t.Errorf("open file error: %v", err)
		}

		buf := make([]byte, 1024)
		reader := bufio.NewReader(f)
		_, err = reader.Read(buf)
		if err != nil {
			t.Errorf("read buf error: %v", err)
		}

		d := new(Detector)
		d.Feed(buf)
		assert.Equal(t, d.Encoding, UTF8)

		defer f.Close()
	}
}

func TestANSIBEDetector(t *testing.T) {

	for i := 1; i <= 3; i++ {
		f, err := os.Open(fmt.Sprintf(".\\raw\\ANSI %d.txt", i))
		if err != nil {
			t.Errorf("open file error: %v", err)
		}

		buf := make([]byte, 1024)
		reader := bufio.NewReader(f)
		_, err = reader.Read(buf)
		if err != nil {
			t.Errorf("read buf error: %v", err)
		}

		d := new(Detector)
		d.Feed(buf)
		assert.Equal(t, d.Encoding, ANSI)

		defer f.Close()
	}
}

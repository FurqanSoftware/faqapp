package cfg

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var (
	Port     int
	Secret   string
	MongoURL string
)

func Load() error {
	err := ParseFile(".env")
	if err != nil {
		return fmt.Errorf("parse file: %s", err)
	}

	errs := []error{}
	errs = append(errs, loadInt(&Port, "PORT", 5000))
	errs = append(errs, loadString(&Secret, "SECRET", ""))
	errs = append(errs, loadString(&MongoURL, "MONGO_URL", ""))
	if len(errs) != 0 {
		return errs[0]
	}
	return nil
}

func loadString(dst *string, key string, value string) error {
	str := os.Getenv(key)
	if str != "" {
		*dst = str
	} else {
		*dst = value
	}
	return nil
}

func loadInt(dst *int, key string, value int) error {
	str := os.Getenv(key)
	if str != "" {
		v, err := strconv.ParseInt(str, 10, 32)
		if err != nil {
			return fmt.Errorf("parse int: %s", err)
		}
		*dst = int(v)
	} else {
		*dst = value
	}
	return nil
}

// ParseFile reads an environment file and loads environment variables that are not already set.
func ParseFile(name string) error {
	f, err := os.Open(name)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}

	r := bufio.NewReader(f)
	for {
		line, err := r.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}

		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			return errors.New("malformed environment file")
		}
		key := parts[0]
		value := parts[1]
		if os.Getenv(key) == "" {
			os.Setenv(key, value)
		}
	}
	return nil
}

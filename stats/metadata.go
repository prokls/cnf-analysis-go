package stats

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/prokls/cnf-analysis-go/output"
	"github.com/prokls/cnf-hash-go/cnfhash"
)

func getTimestamp() string {
	now := time.Now().UTC()
	str := now.Format(time.RFC3339Nano)
	for i := 0; i < len(str); i++ {
		if str[i] == 'Z' {
			str = str[:i]
			break
		}
	}
	return str
}

func computeHashes(path string) (string, string, error) {
	sha1 := sha1.New()
	md5 := md5.New()

	f, err := os.Open(path)
	if err != nil {
		return "", "", err
	}
	defer f.Close()

	buf := make([]byte, os.Getpagesize())
	for {
		n, err := f.Read(buf)
		if err != nil && err != io.EOF {
			return "", "", err
		}
		sha1.Write(buf[:n])
		md5.Write(buf[:n])

		if err == io.EOF {
			break
		}
	}

	return hex.EncodeToString(sha1.Sum(nil)), hex.EncodeToString(md5.Sum(nil)), nil
}

func Metadata(s *output.Stats, path string, conf *FeatureConfig) error {
	var err error

	// cnfhash
	fd, err := os.Open(path)
	if err != nil {
		return err
	}
	s.CNFHash, err = cnfhash.HashDIMACS(fd, cnfhash.Config{})
	if err != nil {
		return err
	}
	fd.Close()

	// filename
	if conf.FullPath {
		s.Filename = path
	} else {
		s.Filename = filepath.Base(path)
	}

	// hashes
	s.SHA1Sum, s.MD5Sum, err = computeHashes(path)
	if err != nil {
		return err
	}

	// timestamp
	s.Timestamp = getTimestamp()

	// version
	s.Version = "1.0.0"

	return nil
}

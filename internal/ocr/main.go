package ocr

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/rs/xid"
	log "github.com/sirupsen/logrus"
	"gitlab.com/swinkelhofer/paperless-office/internal/config"
	"gitlab.com/swinkelhofer/paperless-office/internal/parsers"
)

func NewOCR(rawPDFFile string) (*OCR, error) {
	var (
		ocr = &OCR{
			Hash: xid.New().String(),
		}
		err       error
		appConfig = config.NewConfig()
		hash      xid.ID
	)

	if hash, err = xid.FromString(strings.ReplaceAll(filepath.Base(rawPDFFile), filepath.Ext(rawPDFFile), "")); err == nil {
		ocr.Hash = hash.String()
		ocr.RawPDFFile = filepath.Join(filepath.Dir(rawPDFFile), fmt.Sprintf("%s.pdf", ocr.Hash))
	} else {
		ocr.RawPDFFile = filepath.Join(filepath.Dir(rawPDFFile), fmt.Sprintf("%s.pdf", ocr.Hash))

		if err = os.Rename(rawPDFFile, ocr.RawPDFFile); err != nil {
			log.Errorf("Move error: %w", err)
			return nil, err
		}
		return nil, nil
	}
	ocr.PreviewFile = filepath.Join(appConfig.ProcessedPDFDir, fmt.Sprintf("%s.png", ocr.Hash))
	ocr.ProcessedPDFFile = filepath.Join(appConfig.ProcessedPDFDir, fmt.Sprintf("%s.pdf", ocr.Hash))

	return ocr, nil
}

func (ocr *OCR) RunOCR() error {
	var (
		err     error
		cmd     = exec.Command("ocrmypdf", "-l", "eng+deu", "--deskew", "--rotate-pages", "--force-ocr", "--remove-background", ocr.RawPDFFile, ocr.ProcessedPDFFile)
		copyBuf []byte
	)
	cmd.Env = os.Environ()
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	if err = cmd.Run(); err != nil && cmd.ProcessState.ExitCode() != 6 {
		log.Error(err)
		return err
	}
	if cmd.ProcessState.ExitCode() == 6 {
		if copyBuf, err = ioutil.ReadFile(ocr.RawPDFFile); err != nil {
			log.Error(err)
			return err
		}
		if err = ioutil.WriteFile(ocr.ProcessedPDFFile, copyBuf, 0644); err != nil {
			log.Error(err)
			return err
		}
	}
	return nil
}

func (ocr *OCR) RunInvasiveOCR() error {
	var (
		err     error
		cmd     = exec.Command("ocrmypdf", "-l", "eng+deu", "--deskew", "-i", "-c", "--rotate-pages", "--force-ocr", "--remove-background", "--unpaper-args", "--blackfilter-scan-direction v,h --deskew-scan-direction top,bottom,left,right", ocr.RawPDFFile, ocr.ProcessedPDFFile)
		copyBuf []byte
	)

	cmd.Env = os.Environ()
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	if err = cmd.Run(); err != nil && cmd.ProcessState.ExitCode() != 6 {
		log.Error(err)
		return err
	}
	if cmd.ProcessState.ExitCode() == 6 {
		if copyBuf, err = ioutil.ReadFile(ocr.RawPDFFile); err != nil {
			log.Error(err)
			return err
		}
		if err = ioutil.WriteFile(ocr.ProcessedPDFFile, copyBuf, 0644); err != nil {
			log.Error(err)
			return err
		}
	}
	return nil
}

func (ocr *OCR) GeneratePreview() error {
	var (
		err    error
		cmd    = exec.Command("pdftoppm", "-scale-to", "800", "-png", "-singlefile", "-gray", ocr.ProcessedPDFFile)
		buffer = bytes.NewBuffer([]byte(""))
	)
	cmd.Env = os.Environ()
	cmd.Stderr = os.Stderr
	cmd.Stdout = buffer
	if err = cmd.Run(); err != nil {
		log.Error(err)
		return err
	}
	if err = ioutil.WriteFile(ocr.PreviewFile, buffer.Bytes(), 0644); err != nil {
		return err
	}
	return nil
}

func (ocr *OCR) ExtractText() (string, error) {
	var (
		err    error
		cmd    = exec.Command("pdftotext", "-layout", ocr.ProcessedPDFFile, "-")
		buffer = bytes.NewBuffer([]byte(""))
	)
	cmd.Env = os.Environ()
	cmd.Stderr = os.Stderr
	cmd.Stdout = buffer
	if err = cmd.Run(); err != nil {
		log.Error(err)
		return "", err
	}
	return buffer.String(), nil
}

func (ocr *OCR) HasTextLayer() bool {
	var (
		err    error
		cmd    = exec.Command("pdftotext", "-layout", ocr.RawPDFFile, "-")
		buffer = bytes.NewBuffer([]byte(""))
		reg    *regexp.Regexp
	)
	cmd.Env = os.Environ()
	cmd.Stderr = os.Stderr
	cmd.Stdout = buffer
	if err = cmd.Run(); err != nil {
		log.Error(err)
		return false
	}

	reg, err = regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		log.Error(err)
		return false
	}
	return len(reg.ReplaceAllString(buffer.String(), "")) > 0
}

func (ocr *OCR) ExtractURLs(content string) string {
	var (
		regex = regexp.MustCompile(`\b(http:\/\/www\.|https:\/\/www\.|ftp:\/\/www\.|http:\/\/|https:\/\/|ftp:\/\/)?[a-zA-Z0-9]+([\-\.]{1}[a-zA-Z0-9]+)*\.[a-zA-Z]{2,5}(\/.*)?($|\s+)`)
		err   error
		raw   []byte
		index int
		found = regex.FindAllString(content, -1)
	)

	for index = range found {
		found[index] = strings.TrimSpace(found[index])
	}
	if raw, err = json.Marshal(found); err != nil {
		return ""
	}
	return string(raw)
}

func (ocr *OCR) ExtractMailAdresses(content string) string {
	var (
		regex = `[a-zA-Z0-9.!#$%&*+/=?^_{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*`
		err   error
		found []byte
	)

	if found, err = json.Marshal(regexp.MustCompile(regex).FindAllString(content, -1)); err != nil {
		return ""
	}
	return string(found)
}

func (ocr *OCR) ExtractPhoneNumbers(content string) string {
	var (
		regex = `((\+|00)[1-9]\d{0,3}|0 ?[1-9]|\(00? ?[1-9][\d]*\))[\d\-/]{5,}`
		err   error
		found []byte
	)

	if found, err = json.Marshal(regexp.MustCompile(regex).FindAllString(content, -1)); err != nil {
		return ""
	}
	return string(found)
}

func (ocr *OCR) ExtractDate(content string) time.Time {
	var (
		regex = `[0-3]?[0-9][/.][0-3]?[0-9][/.](?:[0-9]{2})?[0-9]{2}`
		found []string
		date  = time.Now()
	)

	found = regexp.MustCompile(regex).FindAllString(content, -1)
	if len(found) > 0 {
		date, _ = parsers.ParseDate(found[0])
	}
	return date
}

package paperlessoffice

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	log "github.com/sirupsen/logrus"
	"gitlab.com/swinkelhofer/paperless-office/internal/config"
	"gitlab.com/swinkelhofer/paperless-office/internal/ocr"
	"gorm.io/gorm"
)

func Respond(w http.ResponseWriter, response interface{}, statusCode int) {
	var (
		responseBody []byte
		err          error
	)

	if responseBody, err = json.MarshalIndent(response, "    ", ""); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(statusCode)
	w.Write(responseBody)
}

func HandleUpload(w http.ResponseWriter, r *http.Request) {
	var (
		err         error
		requestBody []byte
		fileUpload  UploadRequest
		content     []byte
	)

	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Cache-Control", "no-cache")
	w.Header().Add("Access-Control-Allow-Headers", "*")
	w.Header().Add("Access-Control-Allow-Methods", "*")

	if r.Method == http.MethodPost {
		if requestBody, err = ioutil.ReadAll(r.Body); err != nil {
			Respond(w, err.Error(), http.StatusBadRequest)
			return
		}
		if err = json.Unmarshal(requestBody, &fileUpload); err != nil {
			Respond(w, err.Error(), http.StatusBadRequest)
			return
		}
		if len(fileUpload.FileData) == 0 || len(fileUpload.Name) == 0 {
			Respond(w, "Request body doesn't seem to contain a valid file upload", http.StatusBadRequest)
		}
		if strings.HasPrefix(fileUpload.FileData, "data:") {
			if content, err = base64.StdEncoding.DecodeString(strings.Join(strings.Split(fileUpload.FileData, ",")[1:], ",")); err != nil {
				Respond(w, err.Error(), http.StatusBadRequest)
				return
			}
		} else {
			if content, err = base64.StdEncoding.DecodeString(fileUpload.FileData); err != nil {
				Respond(w, err.Error(), http.StatusBadRequest)
				return
			}
		}

		if err = ioutil.WriteFile(filepath.Join(globalConfig.RawPDFDir, fileUpload.Name), content, 0644); err != nil {
			Respond(w, err.Error(), http.StatusBadRequest)
			return
		}

		Respond(w, "Upload successful", http.StatusCreated)
	} else if r.Method == http.MethodOptions {
		Respond(w, "OK", http.StatusOK)
	} else {
		Respond(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func HandleDocuments(w http.ResponseWriter, r *http.Request) {
	var (
		pdfEntries  = make([]config.PDFEntry, 0)
		pdfEntry    = config.PDFEntry{}
		err         error
		requestBody []byte
	)

	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Cache-Control", "no-cache")
	w.Header().Add("Access-Control-Allow-Headers", "*")
	w.Header().Add("Access-Control-Allow-Methods", "*")

	if r.Method == http.MethodGet {
		if err = globalConfig.DB.Preload("Tags").Order("date DESC").Where("confirmed = ? AND in_trash = ?", true, false).Find(&pdfEntries).Error; err != nil {
			Respond(w, err.Error(), http.StatusBadRequest)
		} else {
			Respond(w, pdfEntries, http.StatusOK)
		}
	} else if r.Method == http.MethodPost {
		if requestBody, err = ioutil.ReadAll(r.Body); err != nil {
			Respond(w, err.Error(), http.StatusBadRequest)
			return
		}
		if err = json.Unmarshal(requestBody, &pdfEntry); err != nil {
			Respond(w, err.Error(), http.StatusBadRequest)
			return
		}
		log.Infof("%+v", pdfEntry)
		if err = globalConfig.DB.Save(&pdfEntry).Error; err != nil {
			Respond(w, err.Error(), http.StatusBadRequest)
			return
		}
		if err = globalConfig.DB.Model(&pdfEntry).Association("Tags").Replace(pdfEntry.Tags); err != nil {
			Respond(w, err.Error(), http.StatusBadRequest)
			return
		}
		Respond(w, "Updated", http.StatusOK)
	} else if r.Method == http.MethodDelete {
		if requestBody, err = ioutil.ReadAll(r.Body); err != nil {
			Respond(w, err.Error(), http.StatusBadRequest)
			return
		}
		if err = json.Unmarshal(requestBody, &pdfEntry); err != nil {
			Respond(w, err.Error(), http.StatusBadRequest)
			return
		}
		log.Infof("%+v", pdfEntry)
		if err = globalConfig.DB.Delete(&pdfEntry).Error; err != nil {
			Respond(w, err.Error(), http.StatusBadRequest)
			return
		}
		Respond(w, "Deleted", http.StatusOK)

	} else if r.Method == http.MethodOptions {
		Respond(w, "OK", http.StatusOK)
	} else {
		Respond(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func HandleTags(w http.ResponseWriter, r *http.Request) {
	var (
		tags        = make([]config.Tag, 0)
		tag         = config.Tag{}
		err         error
		requestBody []byte
	)
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Cache-Control", "no-cache")
	w.Header().Add("Access-Control-Allow-Headers", "*")
	w.Header().Add("Access-Control-Allow-Methods", "*")
	if r.Method == http.MethodGet {
		if err = globalConfig.DB.Find(&tags).Error; err != nil {
			Respond(w, err.Error(), http.StatusBadRequest)
		} else {
			Respond(w, tags, http.StatusOK)
		}
	} else if r.Method == http.MethodPut {
		if requestBody, err = ioutil.ReadAll(r.Body); err != nil {
			Respond(w, err.Error(), http.StatusBadRequest)
			return
		}
		if err = json.Unmarshal(requestBody, &tag); err != nil {
			Respond(w, err.Error(), http.StatusBadRequest)
			return
		}
		log.Infof("%+v", tag)
		if err = globalConfig.DB.Create(&tag).Error; err != nil {
			Respond(w, err.Error(), http.StatusBadRequest)
			return
		}
		Respond(w, tag, http.StatusCreated)
	} else if r.Method == http.MethodPost {
		if requestBody, err = ioutil.ReadAll(r.Body); err != nil {
			Respond(w, err.Error(), http.StatusBadRequest)
			return
		}
		if err = json.Unmarshal(requestBody, &tag); err != nil {
			Respond(w, err.Error(), http.StatusBadRequest)
			return
		}
		log.Infof("%+v", tag)
		if err = globalConfig.DB.Updates(&tag).Error; err != nil {
			Respond(w, err.Error(), http.StatusBadRequest)
			return
		}
		Respond(w, "Updated", http.StatusCreated)
	} else if r.Method == http.MethodDelete {
		if requestBody, err = ioutil.ReadAll(r.Body); err != nil {
			Respond(w, err.Error(), http.StatusBadRequest)
			return
		}
		if err = json.Unmarshal(requestBody, &tag); err != nil {
			Respond(w, err.Error(), http.StatusBadRequest)
			return
		}
		log.Infof("%+v", tag)
		if err = globalConfig.DB.Delete(&tag).Error; err != nil {
			Respond(w, err.Error(), http.StatusBadRequest)
			return
		}
		Respond(w, "Deleted", http.StatusOK)
	} else if r.Method == http.MethodOptions {
		Respond(w, "OK", http.StatusOK)
	} else {
		Respond(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// HandleRescan re-runs scan with more invasive unpaper options
func HandleRescan(w http.ResponseWriter, r *http.Request) {
	var (
		requestBody []byte
		err         error
		ocrClient   *ocr.OCR
		pdfEntry    = config.PDFEntry{}
	)

	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Cache-Control", "no-cache")
	w.Header().Add("Access-Control-Allow-Headers", "*")
	w.Header().Add("Access-Control-Allow-Methods", "*")
	if r.Method == http.MethodPost {
		if requestBody, err = ioutil.ReadAll(r.Body); err != nil {
			Respond(w, err.Error(), http.StatusBadRequest)
			return
		}
		if err = json.Unmarshal(requestBody, &pdfEntry); err != nil {
			Respond(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err = globalConfig.DB.Preload("Tags").Find(&pdfEntry).Error; err != nil {
			Respond(w, err.Error(), http.StatusBadRequest)
			return
		}

		if len(pdfEntry.Filename) == 0 {
			Respond(w, "Failed to find entry by ID", http.StatusBadRequest)
			return
		}

		pdfEntry.RescanFilename = fmt.Sprintf("rescanned_%s", pdfEntry.Filename)

		ocrClient = &ocr.OCR{
			RawPDFFile:       filepath.Join(globalConfig.ProcessedPDFDir, pdfEntry.Filename),
			ProcessedPDFFile: filepath.Join(globalConfig.ProcessedPDFDir, pdfEntry.RescanFilename),
		}

		if err = ocrClient.RunInvasiveOCR(); err != nil {
			Respond(w, err.Error(), http.StatusBadRequest)
			return
		}
		globalConfig.DB.Model(&pdfEntry).Update("rescan_filename", pdfEntry.RescanFilename)

		Respond(w, pdfEntry, http.StatusOK)
	} else if r.Method == http.MethodOptions {
		Respond(w, "OK", http.StatusOK)
	} else {
		Respond(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func HandleRescanAbort(w http.ResponseWriter, r *http.Request) {
	var (
		requestBody []byte
		err         error
		pdfEntry    = config.PDFEntry{}
	)

	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Cache-Control", "no-cache")
	w.Header().Add("Access-Control-Allow-Headers", "*")
	w.Header().Add("Access-Control-Allow-Methods", "*")
	if r.Method == http.MethodPost {
		if requestBody, err = ioutil.ReadAll(r.Body); err != nil {
			Respond(w, err.Error(), http.StatusBadRequest)
			return
		}
		if err = json.Unmarshal(requestBody, &pdfEntry); err != nil {
			Respond(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err = globalConfig.DB.Find(&pdfEntry).Error; err != nil {
			Respond(w, err.Error(), http.StatusBadRequest)
			return
		}

		if len(pdfEntry.Filename) == 0 {
			Respond(w, "Failed to find entry by ID", http.StatusBadRequest)
			return
		}
		os.Remove(filepath.Join(globalConfig.ProcessedPDFDir, pdfEntry.RescanFilename))

		if err = globalConfig.DB.Model(&pdfEntry).Update("rescan_filename", "").Error; err != nil {
			Respond(w, err.Error(), http.StatusBadRequest)
			return
		}
		Respond(w, "Deleted", http.StatusOK)
	} else if r.Method == http.MethodOptions {
		Respond(w, "OK", http.StatusOK)
	} else {
		Respond(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func HandleRescanFinalize(w http.ResponseWriter, r *http.Request) {
	var (
		requestBody []byte
		err         error
		pdfEntry    = config.PDFEntry{}
		ocrClient   ocr.OCR
		content     string
	)

	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Cache-Control", "no-cache")
	w.Header().Add("Access-Control-Allow-Headers", "*")
	w.Header().Add("Access-Control-Allow-Methods", "*")
	if r.Method == http.MethodPost {
		if requestBody, err = ioutil.ReadAll(r.Body); err != nil {
			Respond(w, err, http.StatusBadRequest)
			return
		}
		if err = json.Unmarshal(requestBody, &pdfEntry); err != nil {
			Respond(w, err, http.StatusBadRequest)
			return
		}

		if err = globalConfig.DB.Find(&pdfEntry).Error; err != nil {
			Respond(w, err, http.StatusBadRequest)
			return
		}

		if len(pdfEntry.Filename) == 0 {
			Respond(w, "Failed to find entry by ID", http.StatusBadRequest)
			return
		}

		os.Remove(filepath.Join(globalConfig.ProcessedPDFDir, pdfEntry.Filename))

		if err = os.Rename(
			filepath.Join(globalConfig.ProcessedPDFDir, pdfEntry.RescanFilename),
			filepath.Join(globalConfig.ProcessedPDFDir, pdfEntry.Filename),
		); err != nil {
			Respond(w, err, http.StatusBadRequest)
			return
		}

		ocrClient = ocr.OCR{
			ProcessedPDFFile: filepath.Join(globalConfig.ProcessedPDFDir, pdfEntry.Filename),
			PreviewFile:      filepath.Join(globalConfig.ProcessedPDFDir, strings.Replace(pdfEntry.Filename, ".pdf", ".png", 1)),
		}

		if err = ocrClient.GeneratePreview(); err != nil {
			Respond(w, err.Error(), http.StatusBadRequest)
			return
		}

		if content, err = ocrClient.ExtractText(); err != nil {
			Respond(w, err.Error(), http.StatusBadRequest)
			return
		}

		pdfEntry.Content = content
		pdfEntry.Date = config.Date(ocrClient.ExtractDate(content))
		pdfEntry.PhoneNumbers = ocrClient.ExtractPhoneNumbers(content)
		pdfEntry.URLs = ocrClient.ExtractURLs(content)
		pdfEntry.EMailAdresses = ocrClient.ExtractMailAdresses(content)
		pdfEntry.RescanFilename = ""

		if err = globalConfig.DB.Updates(&pdfEntry).Error; err != nil {
			Respond(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err = globalConfig.DB.Model(&pdfEntry).Update("rescan_filename", "").Error; err != nil {
			Respond(w, err.Error(), http.StatusBadRequest)
			return
		}

		Respond(w, "Success", http.StatusOK)
	} else if r.Method == http.MethodOptions {
		Respond(w, "OK", http.StatusOK)
	} else {
		Respond(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func HandleSearch(w http.ResponseWriter, r *http.Request) {
	var (
		requestBody []byte
		err         error
		search      Search
		tx          *gorm.DB
		pdfEntries  = []config.PDFEntry{}
		//tag                       config.Tag
		doSearch = false
	)
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Cache-Control", "no-cache")
	w.Header().Add("Access-Control-Allow-Headers", "*")
	w.Header().Add("Access-Control-Allow-Methods", "*")
	if r.Method == http.MethodPost {
		if requestBody, err = ioutil.ReadAll(r.Body); err != nil {
			Respond(w, err.Error(), http.StatusBadRequest)
			return
		}
		if err = json.Unmarshal(requestBody, &search); err != nil {
			Respond(w, err.Error(), http.StatusBadRequest)
			return
		}
		log.Infof("%+v", search)

		tx = globalConfig.DB.Preload("Tags")

		if len(search.ID) > 0 {
			tx = tx.Where("id LIKE ?", search.ID)
			doSearch = true
		}
		if len(search.Title) > 0 {
			tx = tx.Where("title LIKE ?", fmt.Sprintf("%%%s%%", search.Title))
			doSearch = true
		}
		if len(search.Content) > 0 {
			tx = tx.Where("content LIKE ?", fmt.Sprintf("%%%s%%", search.Content))
			doSearch = true
		}
		if len(search.EmailAdress) > 0 {
			tx = tx.Where("e_mail_adresses LIKE ?", fmt.Sprintf("%%%s%%", search.EmailAdress))
			doSearch = true
		}
		if len(search.URL) > 0 {
			tx = tx.Where("urls LIKE ?", fmt.Sprintf("%%%s%%", search.URL))
			doSearch = true
		}
		if len(search.URL) > 0 {
			tx = tx.Where("urls LIKE ?", fmt.Sprintf("%%%s%%", search.URL))
			doSearch = true
		}
		if search.Unconfirmed {
			tx = tx.Where("confirmed <> ? and in_trash = ?", search.Unconfirmed, false)
			doSearch = true
		}
		if search.InTrash {
			tx = tx.Where("in_trash = ?", search.InTrash)
			doSearch = true
		}
		// for _, tag = range search.Tags {
		// 	log.Info(tag)
		// 	tx = tx.Joins("pdf_tags")
		// 	doSearch = true
		// }
		if doSearch {
			if err = tx.Find(&pdfEntries).Error; err != nil {
				log.Error()
			}
		}
		Respond(w, pdfEntries, http.StatusOK)
	} else if r.Method == http.MethodOptions {
		Respond(w, "OK", http.StatusOK)
	} else {
		Respond(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

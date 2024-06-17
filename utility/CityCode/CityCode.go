package CityCode

import (
	"archive/zip"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Thenecromance/OurStories/utility/log"
	"github.com/xuri/excelize/v2"
	"io"
	"net/http"
	"os"
)

const (
	downloadUrl = "https://a.amap.com/lbs/static/code_resource/AMap_adcode_citycode.zip"
)

type Code struct {
	Name     string `json:"name"`
	AdCode   string `json:"adcode"`
	CityCode string `json:"citycode"`
}

func writeToFile(codeList []Code) {
	//check dir exist or not
	_, err := os.Stat("setting")
	if os.IsNotExist(err) {
		err := os.Mkdir("setting", os.ModePerm)
		if err != nil {
			log.Errorf("fail to create setting dir .... %s", err.Error())
			return
		}
	}

	file, err := os.Create("setting/cityCode.json")
	if err != nil {
		log.Errorf("fail to create cityCode.json .... %s", err.Error())
		return
	}
	marshal, err := json.Marshal(codeList[2:]) // for skip the first two line of the excel file. first line is the title, second line is the China country code
	if err != nil {
		log.Errorf("fail to Marshal City code .... %s", err.Error())
		return
	}
	_, err = file.Write(marshal)
	if err != nil {
		log.Errorf("fail to Write City code .... %s", err.Error())
		return
	}
}

func requestCityCodeUpdate() []byte {

	response, err := http.Get(downloadUrl)
	if err != nil {
		log.Errorf("fail to Request City code .... %s", err.Error())
		return nil
	}
	fmt.Println(response.StatusCode)
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Errorf("fail to Read City code .... %s", err.Error())
		return nil
	}
	return body
}

func ReadExcel(xlsxStream io.ReadCloser) (codeList []Code) {
	f, err := excelize.OpenReader(xlsxStream)
	if err != nil {
		log.Errorf("fail to Xml Read City code .... %s", err.Error())
		return
	}
	rows, err := f.GetRows("Sheet1")
	for _, row := range rows {
		code := Code{}
		for i, colCell := range row {
			/*if i == 0 {
				code.Name = colCell
				continue
			}else if */
			switch i {
			case 0:
				code.Name = colCell
				continue
			case 1:
				code.AdCode = colCell
				continue
			case 2:
				code.CityCode = colCell
				continue
			default:
				continue
			}
		}
		codeList = append(codeList, code)
	}
	return
}

func unZipCityCode(buffer []byte) io.ReadCloser {
	zipFile, err := zip.NewReader(bytes.NewReader(buffer), int64(len(buffer)))
	if err != nil {
		log.Errorf("fail to Unzip City code .... %s", err.Error())
		return nil
	}
	for _, f := range zipFile.File {
		if f.FileInfo().Name() == "AMap_adcode_citycode.xlsx" {
			//read file
			raw, err := f.Open()
			if err != nil {
				return nil
			}
			return raw
		}
	}
	return nil
}

// Update will request the latest City code from AMap and update the local file
func Update() {
	buffer := requestCityCodeUpdate()
	if buffer == nil {
		log.Error("fail to request City code from AMap, please check your network")
		return
	}
	stream := unZipCityCode(buffer)
	if stream == nil {
		log.Error("fail to unzip City code from AMap")
		return
	}
	codeList := ReadExcel(stream)
	if codeList == nil {
		log.Error("fail to read City code from AMap")
		return
	}
	writeToFile(codeList)
	log.Info("City code update success")
}

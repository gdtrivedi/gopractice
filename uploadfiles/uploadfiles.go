package uploadfiles

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gdtrivedi/gopractice/fileutil"
)

func UploadFilesTest() {
	// Read property file for secrets.
	props, e := fileutil.ReadPropertiesFile("/Users/gautam.trivedi/Documents/Work/Projects/OuterEdge/env.properties")
	if e != nil {
		fmt.Println("Error reading properties", e)
		return
	}
	fmt.Println(time.Now())
	uploadMultipleFilesToPlatform(props)
	fmt.Println(time.Now())
}

func uploadMultipleFilesToPlatform(props fileutil.AppConfigProperties) {

	url := props["prod.platformapi1.baseurl"] + "/ssl/index.php?method=upload&account_name=uploadcert02"
	method := "POST"

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)

	file, errFile1 := os.Open("/Users/gautam.trivedi/Downloads/upload_cert/uploadcert02.gtrivedi.xyz.crt")
	defer file.Close()
	part1, errFile1 := writer.CreateFormFile("upload[]", filepath.Base("/Users/gautam.trivedi/Downloads/upload_cert/uploadcert02.gtrivedi.xyz.crt"))
	_, errFile1 = io.Copy(part1, file)
	if errFile1 != nil {
		fmt.Println(errFile1)
		return
	}

	file, errFile2 := os.Open("/Users/gautam.trivedi/Downloads/upload_cert/uploadcert02.gtrivedi.xyz.key")
	defer file.Close()
	part2, errFile2 := writer.CreateFormFile("upload[]", filepath.Base("/Users/gautam.trivedi/Downloads/upload_cert/uploadcert02.gtrivedi.xyz.key"))
	_, errFile2 = io.Copy(part2, file)
	if errFile2 != nil {
		fmt.Println(errFile2)
		return
	}

	_ = writer.WriteField("wpe_apikey", props["prod.platformapi1.wpe_apikey"])

	err := writer.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}

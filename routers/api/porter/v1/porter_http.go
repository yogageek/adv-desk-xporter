package v1

import (
	"io"
	"net/http"
	"os"
	client "porter/pkg/logic/client"
	logic "porter/pkg/logic/client"
	controller "porter/pkg/logic/controller"

	vars "porter/pkg/logic/vars"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
)

/*
使用 Query 與 DefaultQuery 來取得 request 參數
DefaultQuery 的話如果沒有 firstname 這參數，就會給預設值第二個參數(None)
firstname := c.DefaultQuery("firstname", "None")
lastname := c.Query("lastname")
*/

func Export(c *gin.Context) {
	client.PrepareGQLClient()

	vars.Update_PublicRes_Start(vars.ModeExport)
	defer vars.Update_PublicRes_Done()

	controller.Export()
	//下載檔案
	c.FileAttachment("./exportingData.json", "exportingData.json") //注意都要加.json 否則會找不到轉很久

	// 下面再加會報錯
	// c.JSON(http.StatusOK, gin.H{
	// 	"message": "OK",
	// })
}

func Import(c *gin.Context) {
	logic.PrepareGQLClient()
	//查看是否正在做，如果是則值接返回錯誤
	if !vars.Get_PublicRes_State() {
		c.JSON(http.StatusLocked, gin.H{
			"error": "work is in process",
		})
		return
	}

	vars.Update_PublicRes_Start(vars.ModeImport)
	defer vars.Update_PublicRes_Done()

	//step1讀取客戶端傳來的formdata
	// method1
	file, err := c.FormFile("file") //FormFile()可以取得数据本体、标头，而标头本身有纪录该数据的文件名称，接着只是单纯的I/O处理，这样就会将文件复制一份放到当前目录中。
	// method2
	// rfile, handler, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
	}

	//step2開啟檔案
	sourceFile, err := file.Open()
	if err != nil {
		glog.Error(err)
	}

	//step3
	//***注意以下io copy都只能做一次不然後續會讀不到***

	//---->檔案讀取方式一，通過os.Open返回一個檔案控制代碼，然後利用它進行讀取
	// body, err := ioutil.ReadAll(sourceFile)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(string(body))

	//---->檔案讀取方式二，只讀取檔案
	// buf := bytes.NewBuffer(nil)
	// if _, err := io.Copy(buf, sourceFile); err != nil {
	// 	glog.Error(err)
	// }
	//sth to do

	//---->檔案讀取方式三，存檔(把剛剛開啟的檔案存入到另一個檔案)
	filename := "importingData.json"
	out, err := os.Create(filename)
	if err != nil {
		glog.Error(err)
	}
	if _, err = io.Copy(out, sourceFile); err != nil { //會在目錄裡存檔
		glog.Error(err)
	}
	defer out.Close()

	// 快速方式
	// c.SaveUploadedFile()

	// step4 business logic
	// logic.Import() //old
	controller.Import() //new

	c.JSON(http.StatusOK, gin.H{
		"fileName": file.Filename,
		"size":     file.Size,
		"mimeType": file.Header,
	})

}

//上傳範例
/*
//上传多文件
c.Request.MultipartForm
//上传单文件
c.Request.FormFile("file") //file 是文件名

//把文件保存下来
	fileName:="../import.xlsx"
	saveFile, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0755)
	defer saveFile.Close()
	if err != nil {
		log.Error("error(%v)(%v)", err, fileName)
		return err
	}
	_, _ = io.Copy(saveFile, file)
	//defer os.Remove(fileName) //这句话是删除保存的文件
*/

func DownloadSample(c *gin.Context) {

	// //step1 Read json file...
	// b, err := ioutil.ReadFile("./exportingData.json")
	// if err != nil {
	// 	glog.Error(err)
	// }

	// c.Writer.Header().Add("Content-type", "application/octet-stream")
	// _, err = io.Copy(c.Writer, b)
	// if err != nil {
	// 	c.JSON(http.StatusNotFound, HTTPGenericResponse{
	// 		Code:    http.StatusInternalServerError,
	// 		Message: "文件加载失败:" + err.Error(),
	// 	})
	// 	return
	// }
}

//application/octet-stream ： 二進制流數據（如常見的文檔下載）
//multipart/form-data ： 需要在表單中進行文檔上傳時

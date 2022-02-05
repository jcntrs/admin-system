package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/JeanCntrs/admin-system/dal"
	"github.com/JeanCntrs/admin-system/models"
	"github.com/JeanCntrs/admin-system/utils"
	"github.com/gorilla/mux"
)

func Page(w http.ResponseWriter, r *http.Request) {
	utils.RenderTemplate(w, "page", nil)
}

func GetPages(w http.ResponseWriter, r *http.Request) {
	pages := dal.GetPages()
	pagesByte, _ := json.Marshal(pages)
	fmt.Fprint(w, string(pagesByte))
}

func GetPageById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pageId := vars["pageId"]
	pageIdConv, _ := strconv.Atoi(pageId)

	pageFound := dal.GetPageById(pageIdConv)
	pageByte, _ := json.Marshal(pageFound)
	fmt.Fprint(w, string(pageByte))
}

func CreatePage(w http.ResponseWriter, r *http.Request) {
	page := models.Page{}
	data := json.NewDecoder(r.Body)
	err := data.Decode(&page)

	if err != nil {
		panic("An error occurred while decoding page")
	}

	if page.PageId == 0 {
		_, err := dal.InsertPage(page)
		if err != nil {
			fmt.Fprintf(w, "0")
		}
	} else {
		_, err := dal.UpdatePage(page)
		if err != nil {
			fmt.Fprintf(w, "0")
		}
	}

	fmt.Fprintf(w, "1")
}
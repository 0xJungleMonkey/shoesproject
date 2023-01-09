package api

import (
	"net/http"

	"shoes/dao"
	"shoes/model"

	"github.com/gin-gonic/gin"
	"github.com/guregu/null"
	"github.com/julienschmidt/httprouter"
)

var (
	_ = null.Bool{}
)

func configShoesRouter(router *httprouter.Router) {
	router.GET("/shoes", GetAllShoes)
	router.POST("/shoes", AddShoes)
	router.GET("/shoes/:argID", GetShoes)
	router.PUT("/shoes/:argID", UpdateShoes)
	router.DELETE("/shoes/:argID", DeleteShoes)
}

func configGinShoesRouter(router gin.IRoutes) {
	router.GET("/shoes", ConverHttprouterToGin(GetAllShoes))
	router.POST("/shoes", ConverHttprouterToGin(AddShoes))
	router.GET("/shoes/:argID", ConverHttprouterToGin(GetShoes))
	router.PUT("/shoes/:argID", ConverHttprouterToGin(UpdateShoes))
	router.DELETE("/shoes/:argID", ConverHttprouterToGin(DeleteShoes))
}

// GetAllShoes is a function to get a slice of record(s) from shoes table in the shoes_development database
// @Summary Get list of Shoes
// @Tags Shoes
// @Description GetAllShoes is a handler to get a slice of record(s) from shoes table in the shoes_development database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.Shoes}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /shoes [get]
// http "http://localhost:8080/shoes?page=0&pagesize=20" X-Api-User:user123
func GetAllShoes(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	page, err := readInt(r, "page", 0)
	if err != nil || page < 0 {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	pagesize, err := readInt(r, "pagesize", 20)
	if err != nil || pagesize <= 0 {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	order := r.FormValue("order")

	if err := ValidateRequest(ctx, r, "shoes", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllShoes(ctx, page, pagesize, order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: totalRows}
	writeJSON(ctx, w, result)
}

// GetShoes is a function to get a single record from the shoes table in the shoes_development database
// @Summary Get record from table Shoes by  argID
// @Tags Shoes
// @ID argID
// @Description GetShoes is a function to get a single record from the shoes table in the shoes_development database
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Success 200 {object} model.Shoes
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /shoes/{argID} [get]
// http "http://localhost:8080/shoes/1" X-Api-User:user123
func GetShoes(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "shoes", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetShoes(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddShoes add to add a single record to shoes table in the shoes_development database
// @Summary Add an record to shoes table
// @Description add to add a single record to shoes table in the shoes_development database
// @Tags Shoes
// @Accept  json
// @Produce  json
// @Param Shoes body model.Shoes true "Add Shoes"
// @Success 200 {object} model.Shoes
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /shoes [post]
// echo '{"id": 3,"brand": "NbrbtERxigVhMZdwDhCXfqlxp","model": "ldPplaIdqvMsaHlkjAiffImld","occasion": "nevueHgAeEnHWlXUljQYbYkMo","favorite": 47,"pic": "AKZyxpUOoYaqliuGuWwhueEyb","picture": "RCRWHiBQDlkMK2FEJhJKM15BXxseHRco","created_at": "2177-09-12T19:31:57.746575425-04:00","updated_at": "2161-04-02T18:21:46.101816943-04:00"}' | http POST "http://localhost:8080/shoes" X-Api-User:user123
func AddShoes(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	shoes := &model.Shoes{}

	if err := readJSON(r, shoes); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := shoes.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	shoes.Prepare()

	if err := shoes.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "shoes", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	shoes, _, err = dao.AddShoes(ctx, shoes)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, shoes)
}

// UpdateShoes Update a single record from shoes table in the shoes_development database
// @Summary Update an record in table shoes
// @Description Update a single record from shoes table in the shoes_development database
// @Tags Shoes
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Param  Shoes body model.Shoes true "Update Shoes record"
// @Success 200 {object} model.Shoes
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /shoes/{argID} [put]
// echo '{"id": 3,"brand": "NbrbtERxigVhMZdwDhCXfqlxp","model": "ldPplaIdqvMsaHlkjAiffImld","occasion": "nevueHgAeEnHWlXUljQYbYkMo","favorite": 47,"pic": "AKZyxpUOoYaqliuGuWwhueEyb","picture": "RCRWHiBQDlkMK2FEJhJKM15BXxseHRco","created_at": "2177-09-12T19:31:57.746575425-04:00","updated_at": "2161-04-02T18:21:46.101816943-04:00"}' | http PUT "http://localhost:8080/shoes/1"  X-Api-User:user123
func UpdateShoes(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	shoes := &model.Shoes{}
	if err := readJSON(r, shoes); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := shoes.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	shoes.Prepare()

	if err := shoes.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "shoes", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	shoes, _, err = dao.UpdateShoes(ctx,
		argID,
		shoes)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, shoes)
}

// DeleteShoes Delete a single record from shoes table in the shoes_development database
// @Summary Delete a record from shoes
// @Description Delete a single record from shoes table in the shoes_development database
// @Tags Shoes
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Success 204 {object} model.Shoes
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /shoes/{argID} [delete]
// http DELETE "http://localhost:8080/shoes/1" X-Api-User:user123
func DeleteShoes(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "shoes", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteShoes(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}

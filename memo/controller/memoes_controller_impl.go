package controller

import (
	"memoAPI/helper"
	"memoAPI/model/web"
	"memoAPI/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type MemoesControllerImpl struct {
	memoesService service.MemoesService
}

func NewMemoesController(service service.MemoesService)MemoesController{
	return &MemoesControllerImpl{
		memoesService: service,
	}
}

func(memoesController *MemoesControllerImpl)Create(writer http.ResponseWriter,request *http.Request,params httprouter.Params){
	memoesCreateRequest := web.MemoesCreateRequest{}
	helper.ReadFromRequestBody(request, &memoesCreateRequest)

	memoesResponse := memoesController.memoesService.Create(request.Context(),memoesCreateRequest)
	webResponse := web.WebResponse{
		Code: http.StatusCreated,
		Status: "Success Create Memo",
		Data: memoesResponse,
	}

	helper.WriteToResponseBody(writer,webResponse)
}

func(memoesController *MemoesControllerImpl)Update(writer http.ResponseWriter,request *http.Request,params httprouter.Params){
	memoesUpdateRequest := web.MemoesUpdateRequest{}
	helper.ReadFromRequestBody(request, &memoesUpdateRequest)

	memoesId := params.ByName("memoesId")
	id,err := strconv.Atoi(memoesId)
	helper.PanicIfError(err)

	memoesUpdateRequest.Id = id

	memoesResponse := memoesController.memoesService.Update(request.Context(),memoesUpdateRequest)
	webResponse := web.WebResponse{
		Code: http.StatusOK,
		Status: "Success Update Memo",
		Data: memoesResponse,
	}

	helper.WriteToResponseBody(writer,webResponse)
}

func(memoesController *MemoesControllerImpl)Delete(writer http.ResponseWriter,request *http.Request,params httprouter.Params){
	memoesId := params.ByName("memoesId")
	id,err := strconv.Atoi(memoesId)
	helper.PanicIfError(err)

	memoesController.memoesService.Delete(request.Context(),id)
	webResponse := web.WebResponse{
		Code: http.StatusNoContent,
		Status: "Success Delete Memo",
	}

	helper.WriteToResponseBody(writer,webResponse)
}

func(memoesController *MemoesControllerImpl)FindById(writer http.ResponseWriter,request *http.Request,params httprouter.Params){
	memoesId := params.ByName("memoesId")
	id,err := strconv.Atoi(memoesId)
	helper.PanicIfError(err)

	memoesResponse := memoesController.memoesService.FindById(request.Context(),id)
	webResponse := web.WebResponse{
		Code: http.StatusOK,
		Status: "Success Get Memo By Id",
		Data: memoesResponse,
	}

	helper.WriteToResponseBody(writer,webResponse)
}

func(memoesController *MemoesControllerImpl)FindByTitle(writer http.ResponseWriter,request *http.Request,params httprouter.Params){
	memoesTitle := params.ByName("memoesTitle")

	memoesResponse := memoesController.memoesService.FindByTitle(request.Context(),memoesTitle)
	webResponse := web.WebResponse{
		Code: http.StatusOK,
		Status: "Success Get Memo By Title ",
		Data: memoesResponse,
	}

	helper.WriteToResponseBody(writer,webResponse)
}

func(memoesController *MemoesControllerImpl)OrderByTitleAsc(writer http.ResponseWriter,request *http.Request,params httprouter.Params){
	memoesResponses := memoesController.memoesService.OrderByTitleAsc(request.Context())

	webResponse := web.WebResponse{
		Code: http.StatusOK,
		Status: "Success Get All Memo Order By Title Ascending",
		Data: memoesResponses,
	}

	helper.WriteToResponseBody(writer,webResponse)
}

func(memoesController *MemoesControllerImpl)OrderByIdDesc(writer http.ResponseWriter,request *http.Request,params httprouter.Params){
	memoesResponses := memoesController.memoesService.OrderByIdDesc(request.Context())

	webResponse := web.WebResponse{
		Code: http.StatusOK,
		Status: "Success Get All Memo Order By Id Descending",
		Data: memoesResponses,
	}

	helper.WriteToResponseBody(writer,webResponse)
}

func(memoesController *MemoesControllerImpl)FindAll(writer http.ResponseWriter,request *http.Request,params httprouter.Params){
	memoesResponses := memoesController.memoesService.FindAll(request.Context())

	webResponse := web.WebResponse{
		Code: http.StatusOK,
		Status: "Success Get All Memo",
		Data: memoesResponses,
	}

	helper.WriteToResponseBody(writer,webResponse)
}

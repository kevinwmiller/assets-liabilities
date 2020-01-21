package finances

import (
	"assets-liabilities/entities"
	"assets-liabilities/errors"
	"assets-liabilities/models/record"
	"assets-liabilities/server/routes"
	"assets-liabilities/types"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Router manages routes related to authentication of user objects
type Router struct{}

// List returns a list of authentication route handlers
func (r Router) List() routes.Routes {
	return routes.Routes{
		"/finances/records": routes.Methods{
			routes.Get:  listRecords,
			routes.Post: createRecord,
		},
		"/finances/records/{id}": routes.Methods{
			routes.Get:    getRecord,
			routes.Put:    updateRecord,
			routes.Delete: deleteRecord,
		},
	}
}

// URLParameters:
//	limit: int [1-500]
//	offset: int
//	type: RecordType
func listRecords(w http.ResponseWriter, r *http.Request) {
	limit := types.CreateIntFromString(r.URL.Query().Get("limit"))
	offset := types.CreateIntFromString(r.URL.Query().Get("offset"))
	recordType, err := entities.ConvStrToRecordType(r.URL.Query().Get("type"))
	if r.URL.Query().Get("type") != "" && err != nil {
		routes.RespondWithError(w, errors.Error(err))
		return
	}

	ctx := r.Context()
	records, err := record.CtxModel(ctx).List(ctx, &entities.Record{
		Type: recordType,
	}, &entities.QueryParams{
		Limit:  limit,
		Offset: offset,
	})

	if err != nil {
		routes.RespondWithError(w, errors.Error(err))
		return
	}

	responseJSON, err := json.Marshal(&records)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	routes.Respond(w, http.StatusOK, responseJSON)
}

func createRecord(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	decoder := json.NewDecoder(r.Body)

	var data entities.Record
	err := decoder.Decode(&data)
	if err != nil {
		fmt.Printf("Failed to decode %+v\n", err.Error())
		routes.RespondWithError(w, errors.NewErrorWithCode(http.StatusUnprocessableEntity, err.Error()))
		return
	}

	newRecord, err := record.CtxModel(ctx).Create(ctx, data)
	if err != nil {
		routes.RespondWithError(w, errors.Error(err))
		return
	}

	responseJSON, err := json.Marshal(&newRecord)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	routes.Respond(w, http.StatusOK, responseJSON)
}

func getRecord(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	ctx := r.Context()
	record, err := record.CtxModel(ctx).FindByID(ctx, id)

	if err != nil {
		routes.RespondWithError(w, errors.Error(err))
		return
	}

	responseJSON, err := json.Marshal(&record)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	routes.Respond(w, http.StatusOK, responseJSON)
}

func updateRecord(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	decoder := json.NewDecoder(r.Body)
	var data entities.Record
	err := decoder.Decode(&data)
	if err != nil {
		routes.RespondWithError(w, errors.NewErrorWithCode(http.StatusUnprocessableEntity, err.Error()))
		return
	}

	data.ID = id
	ctx := r.Context()
	record, err := record.CtxModel(ctx).Update(ctx, data)

	if err != nil {
		routes.RespondWithError(w, errors.Error(err))
		return
	}

	responseJSON, err := json.Marshal(&record)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	routes.Respond(w, http.StatusOK, responseJSON)

}

func deleteRecord(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	ctx := r.Context()
	err := record.CtxModel(ctx).Delete(ctx, id)

	if err != nil {
		routes.RespondWithError(w, errors.Error(err))
		return
	}

	routes.Respond(w, http.StatusOK, []byte("{}"))
}

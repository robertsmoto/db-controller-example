package handlers

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

type MemberValues struct {
	ID   string
	Path []string
}

// Helper function to construct redigo.Do varidac args.
func ConstrArgs(key string, values []string) []interface{} {
	sv := []interface{}{}
	sv = append(sv, key)
	if values[0] != "" {
		for _, x := range values {
			sv = append(sv, x)
		}
	} else {
		sv = append(sv, ".")
	}
	return sv
}

// Handler for Member GET requests
func (h *BaseHandler) GetMemberHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	startTime := ctx.Value("startTime").(time.Time)
	// get the request query values
	queryValues := r.URL.Query()
	fmt.Printf("## qvalues %[1]s %[1]T ", queryValues)
	mv := new(MemberValues)

	mv.ID = queryValues.Get("ID")
	fmt.Printf("\n## id --> %[1]s %[1]T", mv.ID)
	if &mv.ID == nil {
		// handle error
	}
	mv.Path = strings.Split(queryValues.Get("path"), ",")
	if &mv.Path == nil {
		// handle error
	}
	args := ConstrArgs(mv.ID, mv.Path)
	result, err := h.memberDataLayer.Get(args)
	if err != nil {
		// send error
		log.Printf("Reading request body. %s", err)
		response := &http.Response{
			Status:        "400 Bad Request",
			StatusCode:    400,
			Proto:         "HTTP/1.1",
			ProtoMajor:    1,
			ProtoMinor:    1,
			Body:          ioutil.NopCloser(bytes.NewBufferString(fmt.Sprintf("%v", err))),
			ContentLength: int64(len(fmt.Sprintf("%v", err))),
			Request:       r,
			Header:        make(http.Header, 0),
		}
		response.Write(w)
	} else {
		// send body
		fmt.Printf("\n## result %[1]v type %[1]T", result)
		strResult := string(result.([]byte)[:])
		response := &http.Response{
			Status:        "200 OK",
			StatusCode:    200,
			Proto:         "HTTP/1.1",
			ProtoMajor:    1,
			ProtoMinor:    1,
			Body:          ioutil.NopCloser(bytes.NewBufferString(strResult)),
			ContentLength: int64(len(strResult)),
			Request:       r,
			Header:        make(http.Header, 0),
		}
		response.Header.Add("Response-Time", fmt.Sprintf("%s", time.Since(startTime)))
		w.Header().Set("Content-Type", "application/json")
		response.Write(w)
	}
}

func (h *BaseHandler) GetCollectionHandler(w http.ResponseWriter, r *http.Request) {
}

// for use with full text search
func (h *BaseHandler) GetSearchHandler(w http.ResponseWriter, r *http.Request) {
}

// creates or updates member(s) and collection(s)
func (h *BaseHandler) PutMemberHandler(w http.ResponseWriter, r *http.Request) {
}

package api

import (
  
  "net/http"
  
)

// ServerInterface represents all server handlers.
type RootStub interface {
Get(w http.ResponseWriter, r *http.Request)

}

type RootWrapper struct {
  RootDelegate RootStub
}


func(stub *RootWrapper) Get(w http.ResponseWriter, r *http.Request) {


  
  stub.RootDelegate.Get(w, r)
}


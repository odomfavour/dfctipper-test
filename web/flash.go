package web

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"time"
)

  func NotifySuccess(w http.ResponseWriter, message string) {
	  SetFlash(w, "success", []byte(message))
  }

  func NotifyError(w http.ResponseWriter, message string) {
	  fmt.Println(message)
	SetFlash(w, "error", []byte(message))
  }

  func SetFlash(w http.ResponseWriter, name string, value []byte) {
	  // TODO: use javascript for current request display
	c := &http.Cookie{Name: name, Value: encode(value)}
	http.SetCookie(w, c)
  }
  
  func GetFlash(w http.ResponseWriter, r *http.Request, name string) ([]byte, error) {
	c, err := r.Cookie(name)
	if err != nil {
	  switch err {
	  case http.ErrNoCookie:
		return nil, nil
	  default:
		return nil, err
	  }
	}
	value, err := decode(c.Value)
	if err != nil {
	  return nil, err
	}
	dc := &http.Cookie{Name: name, MaxAge: -1, Expires: time.Unix(1, 0)}
	http.SetCookie(w, dc)
	return value, nil
  }
  
  // -------------------------
  
  func encode(src []byte) string {
	return base64.URLEncoding.EncodeToString(src)
  }
  
  func decode(src string) ([]byte, error) {
	return base64.URLEncoding.DecodeString(src)
  }
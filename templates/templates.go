package templates

import (
	"bytes"
	"errors"
	"html/template"
	"io"
	"net/http"
	"strconv"
	"sync"
)

type BinTemplate struct {
	Readfunc func(string) ([]byte, error)
	Filename string
	Cascade  bool
	Funcmap  template.FuncMap

	template *template.Template
	mutex    sync.Mutex
}

func MustInit(bt *BinTemplate) {
	bt.mutex.Lock()
	defer bt.mutex.Unlock()
	template.Must(bt.initUnlocked())
}

func (bt *BinTemplate) initUnlocked() (*template.Template, error) {
	if bt.Readfunc == nil {
		return nil, errors.New("BinTemplate must have non-nil .Readfunc")
	}
	text, err := bt.Readfunc(bt.Filename)
	if err != nil {
		return nil, err
	}
	if !bt.Cascade { // the simple case
		t, err := template.New(bt.Filename).Parse(string(text))
		if err == nil {
			bt.template = t
		}
		return t, err
	}
	t, err := template.New("./" + bt.Filename).Parse(`Empty`)
	if err != nil {
		return nil, err
	}
	if bt.Funcmap != nil {
		t.Funcs(bt.Funcmap)
	}
	if _, err := t.New(bt.Filename).Parse(string(text)); err != nil {
		return nil, err
	}
	bt.template = t
	return bt.template, nil
}

func (bt *BinTemplate) CloneExecute(data interface{}) (*bytes.Buffer, error) {
	var (
		filename string
		clone    *template.Template
		err      error
	)
	func() {
		bt.mutex.Lock()
		defer bt.mutex.Unlock()
		if bt.template == nil {
			if _, err = bt.initUnlocked(); err != nil {
				return
			}
		}
		clone, err = bt.template.Clone()
		filename = bt.Filename
	}()
	if err != nil {
		return nil, err
	}
	buf := new(bytes.Buffer)
	if err := clone.ExecuteTemplate(buf, filename, data); err != nil {
		return nil, err
	}
	return buf, nil
}

type templateWriter struct {
	writer http.ResponseWriter
	buf    *bytes.Buffer
	err    error
}

func (bt *BinTemplate) Response(w http.ResponseWriter, data interface{}) templateWriter {
	tw := templateWriter{writer: w}
	tw.buf, tw.err = bt.CloneExecute(data)
	return tw
}

func (tw *templateWriter) SetHeader(name, value string) {
	if tw.err != nil {
		return
	}
	tw.writer.Header().Set(name, value)
}

func (tw *templateWriter) SetContentLength() {
	tw.SetHeader("Content-Length", strconv.Itoa(tw.buf.Len()))
}

func (tw *templateWriter) Send() {
	if tw.err != nil {
		http.Error(tw.writer, tw.err.Error(), http.StatusInternalServerError)
	} else {
		io.Copy(tw.writer, tw.buf) // or w.Write(buf.Bytes())
	}
}
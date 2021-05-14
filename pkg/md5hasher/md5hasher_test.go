package md5hasher

import (
	"bytes"
	"io/ioutil"
	"net/http"
	urlpkg "net/url"
	"testing"
)

const (
	examplePage = `<!doctype html><html lang="en"><head><meta charset="utf-8"><title>The HTML5 empty page</title><meta name="description" content="The HTML5 empty page"><meta name="author" content="SitePoint"><link rel="stylesheet" href="css/styles.css?v=1.0"></head><body><p>Hello world</p><script src="js/scripts.js"></script></body></html>`
)

type httpStub struct {
	reqResp map[string]string
}

func (hm *httpStub) Get(url string) (*http.Response, error) {
	//verify url
	_, err := urlpkg.Parse(url)
	if err != nil {
		return nil, err
	}
	// fill response from map
	if response, ok := hm.reqResp[url]; ok {
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewReader([]byte(response))),
		}, nil
	}
	return &http.Response{
		StatusCode: http.StatusNotFound,
		Body:       ioutil.NopCloser(bytes.NewReader([]byte("Page not found"))),
	}, nil
}

func Test_hasher_Hash(t *testing.T) {
	type fields struct {
		Client client
	}
	type args struct {
		url string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Happy Path example",
			fields: fields{
				Client: &httpStub{reqResp: map[string]string{"http://html5template.com": examplePage}},
			},
			args: args{
				url: "http://html5template.com",
			},
			want:    "50a2d2ad02124b2c70cfb6ccc24d83f4",
			wantErr: false,
		},
		{
			name: "Bad URL",
			fields: fields{
				Client: &httpStub{reqResp: map[string]string{"http://html5template.com": examplePage}},
			},
			args: args{
				url: "http://html5template com ",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "Page not found",
			fields: fields{
				Client: &httpStub{reqResp: map[string]string{"http://html5template.net": examplePage}},
			},
			args: args{
				url: "http://html5template.com",
			},
			want:    "d0fbda9855d118740f1105334305c126",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &hasher{
				client: tt.fields.Client,
			}
			got, err := h.Hash(tt.args.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("hasher.Hash() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("hasher.Hash() = %v, want %v", got, tt.want)
			}
		})
	}
}

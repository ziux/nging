/*
   Nging is a toolbox for webmasters
   Copyright (C) 2018-present  Wenhui Shen <swh@admpub.com>

   This program is free software: you can redistribute it and/or modify
   it under the terms of the GNU Affero General Public License as published
   by the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.

   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU Affero General Public License for more details.

   You should have received a copy of the GNU Affero General Public License
   along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

package upload

import (
	"io"
	"mime/multipart"
	"net/url"

	"github.com/webx-top/echo"
)

func BatchUpload(
	ctx echo.Context,
	fieldName string,
	dstNamer func(*multipart.FileHeader) (dst string, err error),
	storer Storer,
) ([]string, error) {
	req := ctx.Request()
	if req == nil {
		return nil, ctx.E(`Invalid upload content`)
	}
	m := req.MultipartForm()
	if m == nil {
		return nil, ctx.E(`Invalid upload content`)
	}
	files, ok := m.File[fieldName]
	if !ok {
		return nil, echo.ErrNotFoundFileInput
	}
	var dstFile string
	var viewURLs []string
	for _, fileHdr := range files {
		//for each fileheader, get a handle to the actual file
		file, err := fileHdr.Open()
		if err != nil {
			file.Close()
			return viewURLs, err
		}

		dstFile, err = dstNamer(fileHdr)
		if err != nil {
			file.Close()
			return viewURLs, err
		}
		if len(dstFile) == 0 {
			file.Close()
			continue
		}
		viewURL, err := storer.Put(dstFile, file, fileHdr.Size)
		if err != nil {
			file.Close()
			return viewURLs, err
		}
		file.Close()
		viewURLs = append(viewURLs, viewURL)
	}
	return viewURLs, nil
}

type Sizer interface {
	Size() int64
}

type Storer interface {
	Engine() string
	Put(dst string, src io.Reader, size int64) (string, error)
	Get(file string) (io.ReadCloser, error)
	Delete(file string) error
	DeleteDir(dir string) error
	PublicURL(dst string) string
	FixURL(content string, embedded ...bool) string
	FixURLWithParams(content string, values url.Values, embedded ...bool) string
}

type Constructor func(typ string) Storer

var storers = map[string]Constructor{}

var DefaultConstructor Constructor

func StorerRegister(engine string, constructor Constructor) {
	storers[engine] = constructor
}

func StorerGet(engine string) Constructor {
	constructor, ok := storers[engine]
	if !ok {
		return DefaultConstructor
	}
	return constructor
}

func StorerAll(engine string) map[string]Constructor {
	return storers
}

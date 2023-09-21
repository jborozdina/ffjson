/**
 *  Copyright 2014 Paul Querna
 *
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 *
 */

package ffjsoninception

import (
	"bytes"
	"go/format"
	"text/template"
)

const ffjsonTemplate = `
// Code generated by ffjson <https://github.com/jborozdina/ffjson>. DO NOT EDIT.
// source: {{.InputPath}}

package {{.PackageName}}

import (
{{range $k, $v := .OutputImports}}{{$k}}
{{end}}
)

{{range .OutputFuncs}}
{{.}}
{{end}}

`

func RenderTemplate(ic *Inception) ([]byte, error) {
	t := template.Must(template.New("ffjson.go").Parse(ffjsonTemplate))
	buf := new(bytes.Buffer)
	err := t.Execute(buf, ic)
	if err != nil {
		return nil, err
	}
	return format.Source(buf.Bytes())
}

func tplStr(t *template.Template, data interface{}) string {
	buf := bytes.Buffer{}
	err := t.Execute(&buf, data)
	if err != nil {
		panic(err)
	}
	return buf.String()
}

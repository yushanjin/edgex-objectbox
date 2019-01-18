/*
 * Copyright 2019 ObjectBox Ltd. All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package templates

import (
	"text/template"
)

var InitDb = template.Must(template.New("initdb").Parse(`// Code generated by EdgeX; DO NOT EDIT.

package {{.Service.Package}}

import (
	"github.com/edgexfoundry/edgex-go/internal/pkg/db"
	"github.com/edgexfoundry/edgex-go/{{.Service.InterfacePath}}"
	impl "{{.ProviderImport}}"
)

// Return the dbClient interface
func newDBClient(config db.Configuration) ({{.Service.Interface}}.DBClient, error) {
	return impl.NewClient(config)
}
`))
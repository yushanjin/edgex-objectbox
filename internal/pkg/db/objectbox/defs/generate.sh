#!/usr/bin/env bash

set -euo pipefail

obxmodels=internal/pkg/db/objectbox/defs #source
# obxmodels=internal/pkg/db/objectbox/defs/correlation #source
obxbindings=internal/pkg/db/objectbox/obx  #target

generator="go run github.com/objectbox/objectbox-go/cmd/objectbox-gogen -byValue -persist ${obxbindings}/objectbox-model.json"

# cleanup if there was a failed generation previously
rm ${obxmodels}/*.obx.go -f

# generate
for f in "${obxmodels}"/*.go; do if [[ ${f: -8} != ".skip.go" ]]; then ${generator} "${f}"; fi; done

# fix import path
for f in "${obxmodels}"/*.obx.go; do sed -i 's/import (/import (\n\t. "github.com\/edgexfoundry\/go-mod-core-contracts\/models"/g' "$f"; done

# move to the output folder
mv "${obxmodels}"/*.obx.go "${obxbindings}/"

# fix package name on generated files and objectbox-model.go
for f in "${obxbindings}"/*.go; do sed -i 's/package defs/package obx/g' "$f"; done

gofmt -l -w "${obxbindings}"
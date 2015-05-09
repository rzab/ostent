#!/usr/bin/env make -f

# This repo clone location (final subdirectories) defines package name thus
# it should be */github.com/[ostrost]/ostent to make package=github.com/[ostrost]/ostent
package=$(shell echo $$PWD | awk -F/ '{ OFS="/"; print $$(NF-2), $$(NF-1), $$NF }')
testpackage?=./...
singletestpackage=$(testpackage)
ifeq ($(testpackage), ./...)
singletestpackage=$(package)
endif

acepp.go=acepp/acepp.go

binassets_develgo         = share/assets/bindata.devel.go
binassets_productiongo    = share/assets/bindata.production.go
bintemplates_develgo      = share/templates/bindata.devel.go
bintemplates_productiongo = share/templates/bindata.production.go

PATH=$(shell printf %s: $$PATH; echo $$GOPATH | awk -F: 'BEGIN { OFS="/bin:"; } { print $$1,$$2,$$3,$$4,$$5,$$6,$$7,$$8,$$9 "/bin"}')

xargs=xargs
ifeq (Linux, $(shell uname -s))
xargs=xargs --no-run-if-empty
endif
go-bindata=go-bindata -ignore '.*\.go'# Go regexp syntax for -ignore

.PHONY: all al init test covertest coverfunc coverhtml bindata bindata-devel bindata-production
ifneq (init, $(MAKECMDGOALS))
# before init:
# - go list would fail => unknown $(destbin)
# - go test fails without dependencies installed
# - go-bindata is not installed yet

cmdname=$(notdir $(PWD))
destbin=$(shell echo $(GOPATH) | awk -F: '{ print $$1 "/bin" }')
# destbin=$(abspath $(dir $(shell go list -f '{{.Target}}' $(package))))
packagefiles=$(shell \
go list -tags production -f '{{.ImportPath}}{{"\n"}}{{join .Deps "\n"}}' $(package) | xargs \
go list -tags production -f '{{if and (not .Standard) (not .Goroot)}}\
{{$$dir := .Dir}}\
{{range .GoFiles }}{{$$dir}}/{{.}}{{"\n"}}{{end}}\
{{range .CgoFiles}}{{$$dir}}/{{.}}{{"\n"}}{{end}}{{end}}' | \
sed -n "s,^ *,,g; s,$(PWD)/,,p" | sort) # | tee /dev/stderr

all: $(destbin)/$(cmdname)
endif
init:
	go get -u -v \
github.com/jteeuwen/go-bindata/go-bindata \
github.com/skelterjohn/rerun \
github.com/campoy/jsonenums \
github.com/clipperhouse/gen \
github.com/yosssi/ace
# TODO rm {src,pkg/*}/github.com/clipperhouse/slice{,.a}
	cd system/operating && gen add github.com/rzab/slice
	git remote set-url origin https://$(package) # travis & tip & https://code.google.com/p/go/issues/detail?id=8850
	go get -v -tags production $(package)
	go list -f '{{.Target}}' $(package) | $(xargs) rm # clean the library archive
	go get -v -a $(package)
	go list -f '{{.Target}}' $(package) | $(xargs) rm # clean the library archive

%: %.sh # clear the implicit *.sh rule
# print-* rule for debugging. http://blog.jgc.org/2015/04/the-one-line-you-should-add-to-every.html :
print-%: ; @echo $*=$($*)

ifneq (init, $(MAKECMDGOALS))
test:
	go vet $(testpackage)
	go test -v $(testpackage)
covertest:           ; go test -coverprofile=coverage.out -covermode=count -v $(singletestpackage)
coverfunc: covertest ; go tool  cover  -func=coverage.out
coverhtml: covertest ; go tool  cover  -html=coverage.out

system/operating/%_slice.go: system/operating/operating.go ; cd $(dir $@) && go generate
client/uint%_jsonenums.go:   client/tabs.go                ; cd $(dir $@) && go generate

al: $(packagefiles)
# al: like `all' but without final go build $(package). For when rerun does the build

$(destbin)/$(cmdname): $(packagefiles)
	go build -ldflags -w -a -tags production -o $@ $(package)

share/assets/css/index.css: share/style/index.scss
	type sass   >/dev/null || exit 0; sass $< $@
share/assets/js/src/gen/jscript.js: share/tmp/jscript.jsx
	type jsx    >/dev/null || exit 0; jsx <$^ >/dev/null && jsx <$^ 2>/dev/null >$@
share/assets/js/src/milk/index.js: share/coffee/index.coffee
	type coffee >/dev/null || exit 0; coffee -p $^ >/dev/null && coffee -o $(@D)/ $^
share/assets/js/min/index.min.js: $(shell find share/assets/js/src/ -type f)
	type r.js   >/dev/null || exit 0; cd share/assets/js/src/milk && r.js -o build.js

share/templates/index.html: share/ace.templates/index.ace share/ace.templates/defines.ace $(acepp.go)
	go run $(acepp.go) -defines share/ace.templates/defines.ace -output $@ $<
share/templates/defines.html: share/ace.templates/defines.ace $(acepp.go)
	go run $(acepp.go) -defines share/ace.templates/defines.ace -output $@ -savedefines
share/tmp/jscript.jsx: share/ace.templates/jscript.txt share/ace.templates/defines.ace $(acepp.go)
	go run $(acepp.go) -defines share/ace.templates/defines.ace -output $@ -javascript $<

$(bintemplates_productiongo) $(bintemplates_develgo): $(shell find share/templates/ -type f \! -name \*.go)

$(bintemplates_productiongo):
	cd $(@D) && $(go-bindata) -pkg $(notdir $(@D)) -o $(@F) -tags production -mode 0600 -modtime 1400000000 ./...
$(bintemplates_develgo):
	cd $(@D) && $(go-bindata) -pkg $(notdir $(@D)) -o $(@F) -tags '!production' -dev ./...

$(binassets_productiongo):
	cd $(@D) && $(go-bindata) -pkg $(notdir $(@D)) -o $(@F) -tags production -mode 0600 -modtime 1400000000 -ignore js/src/ ./...
$(binassets_develgo):
	cd $(@D) && $(go-bindata) -pkg $(notdir $(@D)) -o $(@F) -tags '!production' -dev -ignore js/min/ ./...

$(binassets_productiongo): $(shell find \
                           share/assets/ -type f \! -name '*.go' \! -path \
                          'share/assets/js/src/*')
$(binassets_productiongo): share/assets/css/index.css
$(binassets_productiongo): share/assets/js/min/index.min.js

$(binassets_develgo): $(shell find \
                      share/assets/ -type f \! -name '*.go' \! -path \
                     'share/assets/js/min/*')
$(binassets_develgo): share/assets/css/index.css
$(binassets_develgo): share/assets/js/src/gen/jscript.js

# spare shortcuts
bindata-production: $(binassets_productiongo) $(bintemplates_productiongo)
bindata-devel: $(binassets_develgo) $(bintemplates_develgo)
bindata: bindata-devel bindata-production

endif # END OF ifneq (init, $(MAKECMDGOALS))

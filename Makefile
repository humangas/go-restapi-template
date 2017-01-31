# Make REST-API module with goa
# goa: https://goa.design/ 

# For the design path, specify relative path from $GOPATH/src/
BASE_PATH=$(subst $(GOPATH)/src/,,$(PWD))
DESIGN_PATH=$(BASE_PATH)/design

.PHONY: all
all:
	@make clean
	@make bootstrap 
	@make serve 

.PHONY: test
test:
	@make cleanall
	@make bootstrap 

.PHONY: bootstrap
bootstrap:
	goagen bootstrap -d $(DESIGN_PATH)

.PHONY: app
app:
	goagen app -d $(DESIGN_PATH)
	
.PHONY: client 
client:
	goagen client -d $(DESIGN_PATH)

.PHONY: main 
main:
	goagen main -d $(DESIGN_PATH) --force

.PHONY: serve
serve:
	go run *.go

.PHONY: clean 
clean:
	rm -rf $(PWD)/app
	rm -rf $(PWD)/tool
	rm -rf $(PWD)/client
	rm -rf $(PWD)/swagger

.PHONY: cleanall
cleanall:
	@make clean
	rm -rf $(PWD)/*.go

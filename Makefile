# Make REST-API module with goa
# goa: https://goa.design/ 

# For the design path, specify relative path from $GOPATH/src/
DESIGN_PATH=github.com/humangas/template-restapi

.PHONY: all
all:
	@make clean
	@make bootstrap 
	@make serve 

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
	rm -rf app
	rm -rf tool
	rm -rf client
	rm -rf swagger
	#rm -rf *.go

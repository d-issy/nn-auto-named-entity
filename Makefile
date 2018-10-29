data/wiki.bit:
	@echo started parser of xml file
	@bzcat $(xmlfile) | go run cmd/xml_to_proto/main.go > logs/xml_to_proto.log
	@echo end parser

data/stat.bin:
	@echo create statistic file
	@go run cmd/statistics/main.go
	@echo done creating statistic file

xml2proto: data/wiki.bit
stat: xml2proto data/stat.bin

clean:
	@rm -rfv data/wiki.bit data/stat.bit logs/%.log

deps:
	@go get -v github.com/golang/protobuf/proto
	@pip install -U pip protobuf

data/wiki.bin:
	@echo started parser of xml file
	@bzcat $(xmlfile) | go run cmd/xml_to_proto/main.go > logs/xml_to_proto.log
	@echo end parser

data/stat.bin:
	@echo create statistic file
	@go run cmd/statistics/main.go
	@echo done creating statistic file

data/tf/title_filters.index:
	@python cmd/create_filters/main.py

data/npy:
	@mkdir data/npy

xml2proto: data/wiki.bin
stat: xml2proto data/stat.bin
create/filter: data/tf/title_filters.index
create/vector: data/npy
	@PYTHONPATH=. python cmd/create_vector/main.py

clean:
	@rm -rfv data/wiki.bin data/stat.bin data/tf data/np logs/%.log

deps:
	@go get -v github.com/golang/protobuf/proto
	@pip install -U pip protobuf

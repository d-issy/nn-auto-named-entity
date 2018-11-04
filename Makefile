data/wiki.bin:
	@echo started parser of xml file
	@bzcat $(xmlfile) | go run cmd/xml_to_proto/main.go > logs/xml_to_proto.log
	@echo end parser

data/stat.bin:
	@echo create statistic file
	@go run cmd/statistics/main.go
	@echo done creating statistic file

data/nn:
	@mkdir data/nn

data/nn/filters.bin: data/nn
	@PYTHONPATH=. python cmd/generate_filters/main.py

data/nn/cnv_table.bin: data/nn
	@PYTHONPATH=. python cmd/cnv_table/main.py

xml2proto: data/wiki.bin
stat: xml2proto data/stat.bin
generate/filter: data/nn/filters.bin
generate/cnv_table: data/nn/cnv_table.bin
create/vector: data/npy
	@PYTHONPATH=. python cmd/create_vector/main.py

clean:
	@rm -rfv data/wiki.bin data/stat.bin data/tf data/np logs/%.log

deps:
	@go get -v github.com/golang/protobuf/proto
	@pip install -U pip protobuf

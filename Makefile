create:
	@echo started parser of xml file
	@bzcat $(xmlfile) | go run cmd/xml_to_proto/main.go > logs/xml_to_proto.log
	@echo end parser

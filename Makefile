create-topic:
	@echo "Creating topic $(TOPIC)"
	@kafka-topics --create --topic $(TOPIC) --bootstrap-server localhost:29092

read-topic:
	@echo "Reading topic $(TOPIC)"
	@kafka-console-consumer --topic $(TOPIC) --bootstrap-server localhost:29092

write-topic:
	@echo "Writing to topic $(TOPIC)"
	@kafka-console-producer --topic $(TOPIC) --bootstrap-server localhost:29092

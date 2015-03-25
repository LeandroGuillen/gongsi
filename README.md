# Gongsi
NGSI utilities for Orion in Go. Low level API to do NGSI queries to an Orion Context Broker plus a command-line utility to test such queries.

## ngsi command-line utility

You can use the command-line utility in the ngsi folder.

To query an entity do:

		ngsi query <entityID>

or

		ngsi q <entity>

To query a specific attribute do:

		ngsi query <entityID> <attributeID>

or

		ngsi q <entity> <attributeID>

## TODO

* Queries
	* Types

* Updates
* Subscriptions
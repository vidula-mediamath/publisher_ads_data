# publisher_ads_data

The src folder contains 3 modules:
main, storage and server

main() calls a function to read data from ads.txt page for a publisher,
performs some validations on the data

It initiates a wait group and send the work to go routine per ads.txt page.
To validate the data, ParseHttpResp() function checks if each record contains
one supply source url domain, and it has "DIRECT" or "RESELLER" as relation
 
After the data is validated, its sent to database
primary key/unique key constraint is publisher_name, supply_source_domain, id, relationship

While inserting the record for a publisher, if duplicate keys are encountered, program will omit those.
Storage module contains DBInsert and DBQuery functions

The server module contains code to start an http server and handler for the GET endpoint it defines.
The output of this program will be json array of records

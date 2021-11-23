const fs = require('fs');
const kafka = require("../.errors/errors_kafka.json")

console.log("Generating kafka errors file")

if (!kafka || !kafka.items  ){
  console.log("invalid error file detected", kafka)
  exit(1)
}

stringBuffer=`
package kafkamgmt

// KafkaManagementErrorCode represents error code returned by Kafka Management API
type KafkaManagementErrorCode string

const (
`

kafka.items.forEach(function(errorType) {
  codeNr = errorType.code.split("-")[2];
  stringBuffer += `  // ${errorType.reason}\n`
  stringBuffer += `  ERROR_${codeNr} KafkaManagementErrorCode = "${errorType.code}"`
  stringBuffer += `\n\n`
})

stringBuffer += `)`

fs.writeFileSync(__dirname+"/../kafkamgmt/apiv1/kafka_errors.go", stringBuffer)

console.log("Sucessfully generated kafka errors file")

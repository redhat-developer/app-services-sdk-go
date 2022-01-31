


module.exports ={
    kafkamgmt: {
        definition: require("./errors_kafka_mgmt.json"),
        file: "kafkamgmt/apiv1/errors/errors.go"
    },
    registrymgmt:  {
        definition: require("./errors_srs_mgmt.json"),
        file: "registrymgmt/apiv1/errors/errors.go"
    },
    connectormgmt: {
        definition: require("./errors_connector_mgmt.json"),
        file: "connectormgmt/apiv1/errors/errors.go"
    }, 
}


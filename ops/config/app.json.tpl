{
  "grpcGateway": {
    "httpPort": 80,
    "grpcPort": 6080,
    "exitTimeout": "20s",
    "validators": [
      "Default"
    ],
    "usePascalNameLogKey": false,
    "usePascalNameErrKey": false,
    "marshalUseProtoNames": true,
    "marshalEmitUnpopulated": false,
    "unmarshalDiscardUnknown": true,
    "enablePing": true,
    "enableTrace": false,
    "enableMetric": false,
    "enablePprof": false,
    "jaeger": {
      "serviceName": "rpc-ops",
      "sampler": {
        "type": "const",
        "param": 1,
        "samplingServerURL": "${JAEGER_SAMPLING_SERVER_URL}"
      },
      "reporter": {
        "logSpans": false,
        "localAgentHostPort": "${JAEGER_REPORTER_LOCAL_AGENT_HOST_PORT}"
      }
    },
    "enableCors": true,
    "cors": {
      "allowAll": true,
      "allowMethod": ["GET, HEAD, POST, PUT, DELETE"],
    }
  },
  "service": {
  },
  "logger": {
    "grpc": {
      "level": "Info",
      "writers": [{
        "type": "RotateFile",
        "options": {
          "filename": "log/app.rpc",
          "maxAge": "24h",
          "formatter": {
            "type": "Json",
            "options": {
              "flatMap": true,
              "pascalNameKey": true
            }
          }
        }
      }, {
        "type": "ElasticSearch",
        "options": {
          "index": "ops-grpc",
          "idField": "requestID",
          "timeout": "200ms",
          "msgChanLen": 200,
          "workerNum": 2,
          "es": {
            "es": {
              "uri": "${ELASTICSEARCH_ENDPOINT}",
              "username": "elastic",
              "password": "${ELASTICSEARCH_PASSWORD}"
            },
            "retry": {
              "attempt": 3,
              "delay": "1s",
              "lastErrorOnly": true,
              "delayType": "BackOff"
            }
          }
        }
      }]
    },
    "info": {
      "level": "Info",
      "writers": [{
        "type": "RotateFile",
        "options": {
          "filename": "log/app.log",
          "maxAge": "24h",
          "formatter": {
            "type": "Json",
            "options": {
              "pascalNameKey": true
            }
          }
        }
      }, {
        "type": "ElasticSearch",
        "options": {
          "index": "ops-log",
          "idField": "requestID",
          "timeout": "200ms",
          "msgChanLen": 200,
          "workerNum": 2,
          "es": {
            "es": {
              "uri": "${ELASTICSEARCH_ENDPOINT}",
              "username": "elastic",
              "password": "${ELASTICSEARCH_PASSWORD}"
            },
            "retry": {
              "attempt": 3,
              "delay": "1s",
              "lastErrorOnly": true,
              "delayType": "BackOff"
            }
          }
        }
      }]
    }
  }
}
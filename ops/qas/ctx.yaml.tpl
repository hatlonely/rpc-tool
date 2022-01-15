name: rpc-tool

ctx:
  rpcTool:
    type: http
    args:
      endpoint: "http://${INGRESS_HOST}"
      headers:
        Origin: "http://${INGRESS_HOST}"
      method: POST
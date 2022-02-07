name: rpc-tool

ctx:
  rpcTool:
    type: http
    args:
      endpoint: "http://${INGRESS_HOST}"
    dft:
      req:
        headers:
          Origin: "http://${INGRESS_HOST}"
        method: POST
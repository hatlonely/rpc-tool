# rpc-tool

## 运维

```shell
ops --variable .cfg/dev.yaml -a run --env dev --task image
ops --variable .cfg/dev.yaml -a run --env dev --task helm --cmd=diff
ops --variable .cfg/dev.yaml -a run --env dev --task helm --cmd=install
```
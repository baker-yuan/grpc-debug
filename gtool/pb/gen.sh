echo "生成 rpc 代码"

# 输出目录
OUT=./

protoc \
--proto_path=${OUT} \
--go_out=${OUT} \
--go-grpc_out=${OUT} \
--openapiv2_out ${OUT} \
--grpc-gateway_out=${OUT} \
--validate_out="lang=go:./" \
--openapiv2_out ./docs/ \
--openapiv2_opt logtostderr=true \
--openapiv2_opt json_names_for_fields=false \
gtool.proto

go-bindata --nocompress -pkg swagger -o swagger/datafile.go docs/...
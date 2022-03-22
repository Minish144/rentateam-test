generate:
	rm -rf gen/pb && mkdir -p gen/pb
	prototool generate --config-data "$$(envsubst < prototool.yaml | yq -o=json)"
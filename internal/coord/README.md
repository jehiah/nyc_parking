brew install go-swagger
curl -O coordprodsearchcurbs.yaml "https://jsapi.apiary.io/apis/coordprodsearchcurbs.source"
mkdir -p internal/coord; cd internal/coord
swagger generate client -f data/coordprodsearchcurbs.yaml 
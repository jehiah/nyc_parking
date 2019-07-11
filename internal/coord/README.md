brew install go-swagger
curl -O coordprodsearchcurbs.yaml "https://jsapi.apiary.io/apis/coordprodsearchcurbs.source"
* manually remove references to common_model so all models are generated
mkdir -p internal/coord; cd internal/coord
swagger generate client -f data/coordprodsearchcurbs.yaml 
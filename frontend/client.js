const {TransformTextRequest} = require('../proto/js/text_tools_pb');
const {TextToolsServiceClient} = require('../proto/js/text_tools_grpc_web_pb');
const grpc = {};

grpc.web = require('grpc-web');

var client = new TextToolsServiceClient('http://localhost:8080');

var request = new TransformTextRequest();
request.setTransformer(Transformer.TRANSFORMER_UPPERCASE);
request.setText('abcdef');

client.transformText(request, {}, function(err, response) {
    console.log(response);
});
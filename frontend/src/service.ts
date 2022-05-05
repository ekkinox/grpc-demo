import {TextToolsServiceClient} from "../js/text_tools_grpc_web_pb";
import {Transformer, TransformTextRequest} from "../js/text_tools_pb";

export class Service {
    private readonly request: TransformTextRequest;
    private client: TextToolsServiceClient;
    constructor(public c: TextToolsServiceClient,public  r: TransformTextRequest) {
        this.request = r;
        this.client = c;
    }
    transform(text: string) {
        this.request.setTransformer(Transformer.TRANSFORMER_UPPERCASE);
        this.request.setText(text);
        this.client.transformText(this.request, {}, function(err, response) {
            console.log(response);
        });
    }
}
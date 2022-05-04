import { Service } from './service';
import {TextToolsServiceClient} from "../js/text_tools_grpc_web_pb";
import {TransformTextRequest} from "../js/text_tools_pb";

const client = new TextToolsServiceClient(process.env.GRPC_SERVER);
const request = new TransformTextRequest();
const ob = new Service(client, request);

const urlParams = new URLSearchParams(window.location.search);
const text = urlParams.get('text');
ob.transform(text);
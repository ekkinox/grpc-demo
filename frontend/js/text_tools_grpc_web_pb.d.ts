import * as grpcWeb from 'grpc-web';

import * as text_tools_pb from './text_tools_pb';


export class TextToolsServiceClient {
  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; });

  transformText(
    request: text_tools_pb.TransformTextRequest,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.RpcError,
               response: text_tools_pb.TransformTextResponse) => void
  ): grpcWeb.ClientReadableStream<text_tools_pb.TransformTextResponse>;

}

export class TextToolsServicePromiseClient {
  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; });

  transformText(
    request: text_tools_pb.TransformTextRequest,
    metadata?: grpcWeb.Metadata
  ): Promise<text_tools_pb.TransformTextResponse>;

}


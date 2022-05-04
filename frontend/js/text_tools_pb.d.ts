import * as jspb from 'google-protobuf'



export class TransformTextRequest extends jspb.Message {
  getTransformer(): Transformer;
  setTransformer(value: Transformer): TransformTextRequest;

  getText(): string;
  setText(value: string): TransformTextRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): TransformTextRequest.AsObject;
  static toObject(includeInstance: boolean, msg: TransformTextRequest): TransformTextRequest.AsObject;
  static serializeBinaryToWriter(message: TransformTextRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): TransformTextRequest;
  static deserializeBinaryFromReader(message: TransformTextRequest, reader: jspb.BinaryReader): TransformTextRequest;
}

export namespace TransformTextRequest {
  export type AsObject = {
    transformer: Transformer,
    text: string,
  }
}

export class TransformTextResponse extends jspb.Message {
  getResult(): string;
  setResult(value: string): TransformTextResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): TransformTextResponse.AsObject;
  static toObject(includeInstance: boolean, msg: TransformTextResponse): TransformTextResponse.AsObject;
  static serializeBinaryToWriter(message: TransformTextResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): TransformTextResponse;
  static deserializeBinaryFromReader(message: TransformTextResponse, reader: jspb.BinaryReader): TransformTextResponse;
}

export namespace TransformTextResponse {
  export type AsObject = {
    result: string,
  }
}

export enum Transformer { 
  TRANSFORMER_UNSPECIFIED = 0,
  TRANSFORMER_UPPERCASE = 1,
  TRANSFORMER_LOWERCASE = 2,
}

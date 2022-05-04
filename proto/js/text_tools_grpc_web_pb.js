/**
 * @fileoverview gRPC-Web generated client stub for text_tools
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck



const grpc = {};
grpc.web = require('grpc-web');


var google_api_annotations_pb = require('./google/api/annotations_pb.js')
const proto = {};
proto.text_tools = require('./text_tools_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.text_tools.TextToolsServiceClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options.format = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.text_tools.TextToolsServicePromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options.format = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.text_tools.TransformTextRequest,
 *   !proto.text_tools.TransformTextResponse>}
 */
const methodDescriptor_TextToolsService_TransformText = new grpc.web.MethodDescriptor(
  '/text_tools.TextToolsService/TransformText',
  grpc.web.MethodType.UNARY,
  proto.text_tools.TransformTextRequest,
  proto.text_tools.TransformTextResponse,
  /**
   * @param {!proto.text_tools.TransformTextRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.text_tools.TransformTextResponse.deserializeBinary
);


/**
 * @param {!proto.text_tools.TransformTextRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.text_tools.TransformTextResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.text_tools.TransformTextResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.text_tools.TextToolsServiceClient.prototype.transformText =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/text_tools.TextToolsService/TransformText',
      request,
      metadata || {},
      methodDescriptor_TextToolsService_TransformText,
      callback);
};


/**
 * @param {!proto.text_tools.TransformTextRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.text_tools.TransformTextResponse>}
 *     Promise that resolves to the response
 */
proto.text_tools.TextToolsServicePromiseClient.prototype.transformText =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/text_tools.TextToolsService/TransformText',
      request,
      metadata || {},
      methodDescriptor_TextToolsService_TransformText);
};


module.exports = proto.text_tools;


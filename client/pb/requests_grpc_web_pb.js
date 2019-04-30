/**
 * @fileoverview gRPC-Web generated client stub for vida
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!



const grpc = {};
grpc.web = require('grpc-web');


var google_protobuf_empty_pb = require('google-protobuf/google/protobuf/empty_pb.js')
const proto = {};
proto.vida = require('./requests_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.vida.MoviesRequestsClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

  /**
   * @private @const {?Object} The credentials to be used to connect
   *    to the server
   */
  this.credentials_ = credentials;

  /**
   * @private @const {?Object} Options for the client
   */
  this.options_ = options;
};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.vida.MoviesRequestsPromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

  /**
   * @private @const {?Object} The credentials to be used to connect
   *    to the server
   */
  this.credentials_ = credentials;

  /**
   * @private @const {?Object} Options for the client
   */
  this.options_ = options;
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.vida.SearchMovieRequest,
 *   !proto.vida.Movie>}
 */
const methodInfo_MoviesRequests_SearchMovies = new grpc.web.AbstractClientBase.MethodInfo(
  proto.vida.Movie,
  /** @param {!proto.vida.SearchMovieRequest} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.vida.Movie.deserializeBinary
);


/**
 * @param {!proto.vida.SearchMovieRequest} request The request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.vida.Movie>}
 *     The XHR Node Readable Stream
 */
proto.vida.MoviesRequestsClient.prototype.searchMovies =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/vida.MoviesRequests/SearchMovies',
      request,
      metadata || {},
      methodInfo_MoviesRequests_SearchMovies);
};


/**
 * @param {!proto.vida.SearchMovieRequest} request The request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.vida.Movie>}
 *     The XHR Node Readable Stream
 */
proto.vida.MoviesRequestsPromiseClient.prototype.searchMovies =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/vida.MoviesRequests/SearchMovies',
      request,
      metadata || {},
      methodInfo_MoviesRequests_SearchMovies);
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.google.protobuf.Empty,
 *   !proto.vida.SearchMovieResponse>}
 */
const methodInfo_MoviesRequests_ListMovies = new grpc.web.AbstractClientBase.MethodInfo(
  proto.vida.SearchMovieResponse,
  /** @param {!proto.google.protobuf.Empty} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.vida.SearchMovieResponse.deserializeBinary
);


/**
 * @param {!proto.google.protobuf.Empty} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.vida.SearchMovieResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.vida.SearchMovieResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.vida.MoviesRequestsClient.prototype.listMovies =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/vida.MoviesRequests/ListMovies',
      request,
      metadata || {},
      methodInfo_MoviesRequests_ListMovies,
      callback);
};


/**
 * @param {!proto.google.protobuf.Empty} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.vida.SearchMovieResponse>}
 *     A native promise that resolves to the response
 */
proto.vida.MoviesRequestsPromiseClient.prototype.listMovies =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/vida.MoviesRequests/ListMovies',
      request,
      metadata || {},
      methodInfo_MoviesRequests_ListMovies);
};


module.exports = proto.vida;


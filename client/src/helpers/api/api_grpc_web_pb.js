/**
 * @fileoverview gRPC-Web generated client stub for v1
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!



const grpc = {};
grpc.web = require('grpc-web');

const proto = {};
proto.v1 = require('./api_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.v1.UsersClient =
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
proto.v1.UsersPromiseClient =
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
 *   !proto.v1.NewUser,
 *   !proto.v1.ID>}
 */
const methodInfo_Users_Create = new grpc.web.AbstractClientBase.MethodInfo(
  proto.v1.ID,
  /** @param {!proto.v1.NewUser} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.v1.ID.deserializeBinary
);


/**
 * @param {!proto.v1.NewUser} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.v1.ID)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.v1.ID>|undefined}
 *     The XHR Node Readable Stream
 */
proto.v1.UsersClient.prototype.create =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/v1.Users/Create',
      request,
      metadata || {},
      methodInfo_Users_Create,
      callback);
};


/**
 * @param {!proto.v1.NewUser} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.v1.ID>}
 *     A native promise that resolves to the response
 */
proto.v1.UsersPromiseClient.prototype.create =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/v1.Users/Create',
      request,
      metadata || {},
      methodInfo_Users_Create);
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.v1.ID,
 *   !proto.v1.ReadUser>}
 */
const methodInfo_Users_Read = new grpc.web.AbstractClientBase.MethodInfo(
  proto.v1.ReadUser,
  /** @param {!proto.v1.ID} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.v1.ReadUser.deserializeBinary
);


/**
 * @param {!proto.v1.ID} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.v1.ReadUser)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.v1.ReadUser>|undefined}
 *     The XHR Node Readable Stream
 */
proto.v1.UsersClient.prototype.read =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/v1.Users/Read',
      request,
      metadata || {},
      methodInfo_Users_Read,
      callback);
};


/**
 * @param {!proto.v1.ID} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.v1.ReadUser>}
 *     A native promise that resolves to the response
 */
proto.v1.UsersPromiseClient.prototype.read =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/v1.Users/Read',
      request,
      metadata || {},
      methodInfo_Users_Read);
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.v1.ID,
 *   !proto.v1.ReadUserLog>}
 */
const methodInfo_Users_ReadLog = new grpc.web.AbstractClientBase.MethodInfo(
  proto.v1.ReadUserLog,
  /** @param {!proto.v1.ID} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.v1.ReadUserLog.deserializeBinary
);


/**
 * @param {!proto.v1.ID} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.v1.ReadUserLog)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.v1.ReadUserLog>|undefined}
 *     The XHR Node Readable Stream
 */
proto.v1.UsersClient.prototype.readLog =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/v1.Users/ReadLog',
      request,
      metadata || {},
      methodInfo_Users_ReadLog,
      callback);
};


/**
 * @param {!proto.v1.ID} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.v1.ReadUserLog>}
 *     A native promise that resolves to the response
 */
proto.v1.UsersPromiseClient.prototype.readLog =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/v1.Users/ReadLog',
      request,
      metadata || {},
      methodInfo_Users_ReadLog);
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.v1.UpdateUser,
 *   !proto.v1.Result>}
 */
const methodInfo_Users_Update = new grpc.web.AbstractClientBase.MethodInfo(
  proto.v1.Result,
  /** @param {!proto.v1.UpdateUser} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.v1.Result.deserializeBinary
);


/**
 * @param {!proto.v1.UpdateUser} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.v1.Result)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.v1.Result>|undefined}
 *     The XHR Node Readable Stream
 */
proto.v1.UsersClient.prototype.update =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/v1.Users/Update',
      request,
      metadata || {},
      methodInfo_Users_Update,
      callback);
};


/**
 * @param {!proto.v1.UpdateUser} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.v1.Result>}
 *     A native promise that resolves to the response
 */
proto.v1.UsersPromiseClient.prototype.update =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/v1.Users/Update',
      request,
      metadata || {},
      methodInfo_Users_Update);
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.v1.ID,
 *   !proto.v1.Result>}
 */
const methodInfo_Users_Delete = new grpc.web.AbstractClientBase.MethodInfo(
  proto.v1.Result,
  /** @param {!proto.v1.ID} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.v1.Result.deserializeBinary
);


/**
 * @param {!proto.v1.ID} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.v1.Result)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.v1.Result>|undefined}
 *     The XHR Node Readable Stream
 */
proto.v1.UsersClient.prototype.delete =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/v1.Users/Delete',
      request,
      metadata || {},
      methodInfo_Users_Delete,
      callback);
};


/**
 * @param {!proto.v1.ID} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.v1.Result>}
 *     A native promise that resolves to the response
 */
proto.v1.UsersPromiseClient.prototype.delete =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/v1.Users/Delete',
      request,
      metadata || {},
      methodInfo_Users_Delete);
};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.v1.GroupsClient =
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
proto.v1.GroupsPromiseClient =
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
 *   !proto.v1.NewGroup,
 *   !proto.v1.ID>}
 */
const methodInfo_Groups_Create = new grpc.web.AbstractClientBase.MethodInfo(
  proto.v1.ID,
  /** @param {!proto.v1.NewGroup} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.v1.ID.deserializeBinary
);


/**
 * @param {!proto.v1.NewGroup} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.v1.ID)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.v1.ID>|undefined}
 *     The XHR Node Readable Stream
 */
proto.v1.GroupsClient.prototype.create =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/v1.Groups/Create',
      request,
      metadata || {},
      methodInfo_Groups_Create,
      callback);
};


/**
 * @param {!proto.v1.NewGroup} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.v1.ID>}
 *     A native promise that resolves to the response
 */
proto.v1.GroupsPromiseClient.prototype.create =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/v1.Groups/Create',
      request,
      metadata || {},
      methodInfo_Groups_Create);
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.v1.ID,
 *   !proto.v1.Group>}
 */
const methodInfo_Groups_Read = new grpc.web.AbstractClientBase.MethodInfo(
  proto.v1.Group,
  /** @param {!proto.v1.ID} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.v1.Group.deserializeBinary
);


/**
 * @param {!proto.v1.ID} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.v1.Group)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.v1.Group>|undefined}
 *     The XHR Node Readable Stream
 */
proto.v1.GroupsClient.prototype.read =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/v1.Groups/Read',
      request,
      metadata || {},
      methodInfo_Groups_Read,
      callback);
};


/**
 * @param {!proto.v1.ID} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.v1.Group>}
 *     A native promise that resolves to the response
 */
proto.v1.GroupsPromiseClient.prototype.read =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/v1.Groups/Read',
      request,
      metadata || {},
      methodInfo_Groups_Read);
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.v1.ID,
 *   !proto.v1.Members>}
 */
const methodInfo_Groups_ReadGroupMembers = new grpc.web.AbstractClientBase.MethodInfo(
  proto.v1.Members,
  /** @param {!proto.v1.ID} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.v1.Members.deserializeBinary
);


/**
 * @param {!proto.v1.ID} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.v1.Members)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.v1.Members>|undefined}
 *     The XHR Node Readable Stream
 */
proto.v1.GroupsClient.prototype.readGroupMembers =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/v1.Groups/ReadGroupMembers',
      request,
      metadata || {},
      methodInfo_Groups_ReadGroupMembers,
      callback);
};


/**
 * @param {!proto.v1.ID} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.v1.Members>}
 *     A native promise that resolves to the response
 */
proto.v1.GroupsPromiseClient.prototype.readGroupMembers =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/v1.Groups/ReadGroupMembers',
      request,
      metadata || {},
      methodInfo_Groups_ReadGroupMembers);
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.v1.ID,
 *   !proto.v1.ReadUser>}
 */
const methodInfo_Groups_ReadGroupOrganiser = new grpc.web.AbstractClientBase.MethodInfo(
  proto.v1.ReadUser,
  /** @param {!proto.v1.ID} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.v1.ReadUser.deserializeBinary
);


/**
 * @param {!proto.v1.ID} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.v1.ReadUser)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.v1.ReadUser>|undefined}
 *     The XHR Node Readable Stream
 */
proto.v1.GroupsClient.prototype.readGroupOrganiser =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/v1.Groups/ReadGroupOrganiser',
      request,
      metadata || {},
      methodInfo_Groups_ReadGroupOrganiser,
      callback);
};


/**
 * @param {!proto.v1.ID} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.v1.ReadUser>}
 *     A native promise that resolves to the response
 */
proto.v1.GroupsPromiseClient.prototype.readGroupOrganiser =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/v1.Groups/ReadGroupOrganiser',
      request,
      metadata || {},
      methodInfo_Groups_ReadGroupOrganiser);
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.v1.UpdateGroup,
 *   !proto.v1.Result>}
 */
const methodInfo_Groups_Update = new grpc.web.AbstractClientBase.MethodInfo(
  proto.v1.Result,
  /** @param {!proto.v1.UpdateGroup} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.v1.Result.deserializeBinary
);


/**
 * @param {!proto.v1.UpdateGroup} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.v1.Result)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.v1.Result>|undefined}
 *     The XHR Node Readable Stream
 */
proto.v1.GroupsClient.prototype.update =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/v1.Groups/Update',
      request,
      metadata || {},
      methodInfo_Groups_Update,
      callback);
};


/**
 * @param {!proto.v1.UpdateGroup} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.v1.Result>}
 *     A native promise that resolves to the response
 */
proto.v1.GroupsPromiseClient.prototype.update =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/v1.Groups/Update',
      request,
      metadata || {},
      methodInfo_Groups_Update);
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.v1.GroupMember,
 *   !proto.v1.Result>}
 */
const methodInfo_Groups_UpdateMember = new grpc.web.AbstractClientBase.MethodInfo(
  proto.v1.Result,
  /** @param {!proto.v1.GroupMember} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.v1.Result.deserializeBinary
);


/**
 * @param {!proto.v1.GroupMember} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.v1.Result)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.v1.Result>|undefined}
 *     The XHR Node Readable Stream
 */
proto.v1.GroupsClient.prototype.updateMember =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/v1.Groups/UpdateMember',
      request,
      metadata || {},
      methodInfo_Groups_UpdateMember,
      callback);
};


/**
 * @param {!proto.v1.GroupMember} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.v1.Result>}
 *     A native promise that resolves to the response
 */
proto.v1.GroupsPromiseClient.prototype.updateMember =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/v1.Groups/UpdateMember',
      request,
      metadata || {},
      methodInfo_Groups_UpdateMember);
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.v1.ID,
 *   !proto.v1.Result>}
 */
const methodInfo_Groups_Delete = new grpc.web.AbstractClientBase.MethodInfo(
  proto.v1.Result,
  /** @param {!proto.v1.ID} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.v1.Result.deserializeBinary
);


/**
 * @param {!proto.v1.ID} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.v1.Result)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.v1.Result>|undefined}
 *     The XHR Node Readable Stream
 */
proto.v1.GroupsClient.prototype.delete =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/v1.Groups/Delete',
      request,
      metadata || {},
      methodInfo_Groups_Delete,
      callback);
};


/**
 * @param {!proto.v1.ID} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.v1.Result>}
 *     A native promise that resolves to the response
 */
proto.v1.GroupsPromiseClient.prototype.delete =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/v1.Groups/Delete',
      request,
      metadata || {},
      methodInfo_Groups_Delete);
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.v1.GroupMember,
 *   !proto.v1.Result>}
 */
const methodInfo_Groups_DeleteMember = new grpc.web.AbstractClientBase.MethodInfo(
  proto.v1.Result,
  /** @param {!proto.v1.GroupMember} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.v1.Result.deserializeBinary
);


/**
 * @param {!proto.v1.GroupMember} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.v1.Result)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.v1.Result>|undefined}
 *     The XHR Node Readable Stream
 */
proto.v1.GroupsClient.prototype.deleteMember =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/v1.Groups/DeleteMember',
      request,
      metadata || {},
      methodInfo_Groups_DeleteMember,
      callback);
};


/**
 * @param {!proto.v1.GroupMember} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.v1.Result>}
 *     A native promise that resolves to the response
 */
proto.v1.GroupsPromiseClient.prototype.deleteMember =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/v1.Groups/DeleteMember',
      request,
      metadata || {},
      methodInfo_Groups_DeleteMember);
};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.v1.TransactionsClient =
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
proto.v1.TransactionsPromiseClient =
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
 *   !proto.v1.NewTransaction,
 *   !proto.v1.Result>}
 */
const methodInfo_Transactions_Create = new grpc.web.AbstractClientBase.MethodInfo(
  proto.v1.Result,
  /** @param {!proto.v1.NewTransaction} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.v1.Result.deserializeBinary
);


/**
 * @param {!proto.v1.NewTransaction} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.v1.Result)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.v1.Result>|undefined}
 *     The XHR Node Readable Stream
 */
proto.v1.TransactionsClient.prototype.create =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/v1.Transactions/Create',
      request,
      metadata || {},
      methodInfo_Transactions_Create,
      callback);
};


/**
 * @param {!proto.v1.NewTransaction} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.v1.Result>}
 *     A native promise that resolves to the response
 */
proto.v1.TransactionsPromiseClient.prototype.create =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/v1.Transactions/Create',
      request,
      metadata || {},
      methodInfo_Transactions_Create);
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.v1.TID,
 *   !proto.v1.Transaction>}
 */
const methodInfo_Transactions_ReadTransaction = new grpc.web.AbstractClientBase.MethodInfo(
  proto.v1.Transaction,
  /** @param {!proto.v1.TID} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.v1.Transaction.deserializeBinary
);


/**
 * @param {!proto.v1.TID} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.v1.Transaction)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.v1.Transaction>|undefined}
 *     The XHR Node Readable Stream
 */
proto.v1.TransactionsClient.prototype.readTransaction =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/v1.Transactions/ReadTransaction',
      request,
      metadata || {},
      methodInfo_Transactions_ReadTransaction,
      callback);
};


/**
 * @param {!proto.v1.TID} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.v1.Transaction>}
 *     A native promise that resolves to the response
 */
proto.v1.TransactionsPromiseClient.prototype.readTransaction =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/v1.Transactions/ReadTransaction',
      request,
      metadata || {},
      methodInfo_Transactions_ReadTransaction);
};


module.exports = proto.v1;


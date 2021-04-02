// source: errors.proto
/**
 * @fileoverview
 * @enhanceable
 * @suppress {missingRequire} reports error on implicit type usages.
 * @suppress {messageConventions} JS Compiler reports an error if a variable or
 *     field starts with 'MSG_' and isn't a translatable message.
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!
/* eslint-disable */
// @ts-nocheck

var jspb = require('google-protobuf');
var goog = jspb;
var global = Function('return this')();

goog.exportSymbol('proto.ItemRequestError', null, global);
goog.exportSymbol('proto.ItemRequestError.ErrorType', null, global);
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.ItemRequestError = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.ItemRequestError, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.ItemRequestError.displayName = 'proto.ItemRequestError';
}



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.ItemRequestError.prototype.toObject = function(opt_includeInstance) {
  return proto.ItemRequestError.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.ItemRequestError} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.ItemRequestError.toObject = function(includeInstance, msg) {
  var f, obj = {
    errortype: jspb.Message.getFieldWithDefault(msg, 2, 0),
    errorstring: jspb.Message.getFieldWithDefault(msg, 3, ""),
    context: jspb.Message.getFieldWithDefault(msg, 4, "")
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.ItemRequestError}
 */
proto.ItemRequestError.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.ItemRequestError;
  return proto.ItemRequestError.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.ItemRequestError} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.ItemRequestError}
 */
proto.ItemRequestError.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 2:
      var value = /** @type {!proto.ItemRequestError.ErrorType} */ (reader.readEnum());
      msg.setErrortype(value);
      break;
    case 3:
      var value = /** @type {string} */ (reader.readString());
      msg.setErrorstring(value);
      break;
    case 4:
      var value = /** @type {string} */ (reader.readString());
      msg.setContext(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.ItemRequestError.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.ItemRequestError.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.ItemRequestError} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.ItemRequestError.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getErrortype();
  if (f !== 0.0) {
    writer.writeEnum(
      2,
      f
    );
  }
  f = message.getErrorstring();
  if (f.length > 0) {
    writer.writeString(
      3,
      f
    );
  }
  f = message.getContext();
  if (f.length > 0) {
    writer.writeString(
      4,
      f
    );
  }
};


/**
 * @enum {number}
 */
proto.ItemRequestError.ErrorType = {
  OTHER: 0,
  NOTFOUND: 1,
  NOCONTEXT: 2
};

/**
 * optional ErrorType errorType = 2;
 * @return {!proto.ItemRequestError.ErrorType}
 */
proto.ItemRequestError.prototype.getErrortype = function() {
  return /** @type {!proto.ItemRequestError.ErrorType} */ (jspb.Message.getFieldWithDefault(this, 2, 0));
};


/**
 * @param {!proto.ItemRequestError.ErrorType} value
 * @return {!proto.ItemRequestError} returns this
 */
proto.ItemRequestError.prototype.setErrortype = function(value) {
  return jspb.Message.setProto3EnumField(this, 2, value);
};


/**
 * optional string errorString = 3;
 * @return {string}
 */
proto.ItemRequestError.prototype.getErrorstring = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 3, ""));
};


/**
 * @param {string} value
 * @return {!proto.ItemRequestError} returns this
 */
proto.ItemRequestError.prototype.setErrorstring = function(value) {
  return jspb.Message.setProto3StringField(this, 3, value);
};


/**
 * optional string context = 4;
 * @return {string}
 */
proto.ItemRequestError.prototype.getContext = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 4, ""));
};


/**
 * @param {string} value
 * @return {!proto.ItemRequestError} returns this
 */
proto.ItemRequestError.prototype.setContext = function(value) {
  return jspb.Message.setProto3StringField(this, 4, value);
};


goog.object.extend(exports, proto);

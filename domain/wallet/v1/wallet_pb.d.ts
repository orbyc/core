// @generated by protoc-gen-es v1.1.0
// @generated from file wallet/v1/wallet.proto (package wallet.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";
import type { Transaction } from "../../transaction/v1/transaction_pb.js";

/**
 * @generated from message wallet.v1.Wallet
 */
export declare class Wallet extends Message<Wallet> {
  /**
   * @generated from field: string account = 1;
   */
  account: string;

  /**
   * @generated from field: map<string, wallet.v1.Wallet.Collection> collections = 2;
   */
  collections: { [key: string]: Wallet_Collection };

  /**
   * @generated from field: repeated wallet.v1.Wallet.Contact contacts = 3;
   */
  contacts: Wallet_Contact[];

  constructor(data?: PartialMessage<Wallet>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "wallet.v1.Wallet";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Wallet;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Wallet;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Wallet;

  static equals(a: Wallet | PlainMessage<Wallet> | undefined, b: Wallet | PlainMessage<Wallet> | undefined): boolean;
}

/**
 * @generated from message wallet.v1.Wallet.Contact
 */
export declare class Wallet_Contact extends Message<Wallet_Contact> {
  constructor(data?: PartialMessage<Wallet_Contact>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "wallet.v1.Wallet.Contact";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Wallet_Contact;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Wallet_Contact;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Wallet_Contact;

  static equals(a: Wallet_Contact | PlainMessage<Wallet_Contact> | undefined, b: Wallet_Contact | PlainMessage<Wallet_Contact> | undefined): boolean;
}

/**
 * @generated from message wallet.v1.Wallet.Collection
 */
export declare class Wallet_Collection extends Message<Wallet_Collection> {
  /**
   * @generated from field: repeated uint64 assets = 1;
   */
  assets: bigint[];

  constructor(data?: PartialMessage<Wallet_Collection>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "wallet.v1.Wallet.Collection";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Wallet_Collection;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Wallet_Collection;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Wallet_Collection;

  static equals(a: Wallet_Collection | PlainMessage<Wallet_Collection> | undefined, b: Wallet_Collection | PlainMessage<Wallet_Collection> | undefined): boolean;
}

/**
 * @generated from message wallet.v1.GetAddressResponse
 */
export declare class GetAddressResponse extends Message<GetAddressResponse> {
  /**
   * @generated from field: string address = 1;
   */
  address: string;

  constructor(data?: PartialMessage<GetAddressResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "wallet.v1.GetAddressResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetAddressResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetAddressResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetAddressResponse;

  static equals(a: GetAddressResponse | PlainMessage<GetAddressResponse> | undefined, b: GetAddressResponse | PlainMessage<GetAddressResponse> | undefined): boolean;
}

/**
 * @generated from message wallet.v1.GetMetadataRequest
 */
export declare class GetMetadataRequest extends Message<GetMetadataRequest> {
  /**
   * @generated from field: string id = 1;
   */
  id: string;

  constructor(data?: PartialMessage<GetMetadataRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "wallet.v1.GetMetadataRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetMetadataRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetMetadataRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetMetadataRequest;

  static equals(a: GetMetadataRequest | PlainMessage<GetMetadataRequest> | undefined, b: GetMetadataRequest | PlainMessage<GetMetadataRequest> | undefined): boolean;
}

/**
 * @generated from message wallet.v1.GetSignatureRequest
 */
export declare class GetSignatureRequest extends Message<GetSignatureRequest> {
  /**
   * @generated from field: string message = 1;
   */
  message: string;

  constructor(data?: PartialMessage<GetSignatureRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "wallet.v1.GetSignatureRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetSignatureRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetSignatureRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetSignatureRequest;

  static equals(a: GetSignatureRequest | PlainMessage<GetSignatureRequest> | undefined, b: GetSignatureRequest | PlainMessage<GetSignatureRequest> | undefined): boolean;
}

/**
 * @generated from message wallet.v1.GetSignatureResponse
 */
export declare class GetSignatureResponse extends Message<GetSignatureResponse> {
  /**
   * @generated from field: string signature = 1;
   */
  signature: string;

  constructor(data?: PartialMessage<GetSignatureResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "wallet.v1.GetSignatureResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetSignatureResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetSignatureResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetSignatureResponse;

  static equals(a: GetSignatureResponse | PlainMessage<GetSignatureResponse> | undefined, b: GetSignatureResponse | PlainMessage<GetSignatureResponse> | undefined): boolean;
}

/**
 * @generated from message wallet.v1.SendTransactionRequest
 */
export declare class SendTransactionRequest extends Message<SendTransactionRequest> {
  /**
   * @generated from field: transaction.v1.Transaction transaction = 1;
   */
  transaction?: Transaction;

  /**
   * @generated from field: repeated string authorizations = 2;
   */
  authorizations: string[];

  constructor(data?: PartialMessage<SendTransactionRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "wallet.v1.SendTransactionRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): SendTransactionRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): SendTransactionRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): SendTransactionRequest;

  static equals(a: SendTransactionRequest | PlainMessage<SendTransactionRequest> | undefined, b: SendTransactionRequest | PlainMessage<SendTransactionRequest> | undefined): boolean;
}

/**
 * @generated from message wallet.v1.SendTransactionResponse
 */
export declare class SendTransactionResponse extends Message<SendTransactionResponse> {
  /**
   * @generated from field: bool success = 1;
   */
  success: boolean;

  constructor(data?: PartialMessage<SendTransactionResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "wallet.v1.SendTransactionResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): SendTransactionResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): SendTransactionResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): SendTransactionResponse;

  static equals(a: SendTransactionResponse | PlainMessage<SendTransactionResponse> | undefined, b: SendTransactionResponse | PlainMessage<SendTransactionResponse> | undefined): boolean;
}


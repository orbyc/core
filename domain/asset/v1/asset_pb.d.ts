// @generated by protoc-gen-es v1.0.0
// @generated from file asset/v1/asset.proto (package asset.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage, Timestamp } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * METADATA 
 *
 * @generated from message asset.v1.Asset
 */
export declare class Asset extends Message<Asset> {
  /**
   * @generated from field: string name = 1;
   */
  name: string;

  /**
   * @generated from field: string description = 2;
   */
  description: string;

  /**
   * @generated from field: string image = 3;
   */
  image: string;

  /**
   * @generated from field: string external_url = 4;
   */
  externalUrl: string;

  /**
   * @generated from field: repeated asset.v1.Asset.Attribute attributes = 5;
   */
  attributes: Asset_Attribute[];

  /**
   * @generated from oneof asset.v1.Asset.footprint
   */
  footprint: {
    /**
     * carbon emissions
     *
     * @generated from field: uint64 co2e = 6;
     */
    value: bigint;
    case: "co2e";
  } | {
    /**
     * carbon offset
     *
     * @generated from field: uint64 co2o = 7;
     */
    value: bigint;
    case: "co2o";
  } | { case: undefined; value?: undefined };

  /**
   * @generated from field: repeated string gallery = 8;
   */
  gallery: string[];

  /**
   * @generated from field: optional string video = 9;
   */
  video?: string;

  /**
   * @generated from field: google.protobuf.Timestamp created_at = 10;
   */
  createdAt?: Timestamp;

  /**
   * @generated from field: google.protobuf.Timestamp updated_at = 11;
   */
  updatedAt?: Timestamp;

  constructor(data?: PartialMessage<Asset>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "asset.v1.Asset";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Asset;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Asset;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Asset;

  static equals(a: Asset | PlainMessage<Asset> | undefined, b: Asset | PlainMessage<Asset> | undefined): boolean;
}

/**
 * @generated from message asset.v1.Asset.Attribute
 */
export declare class Asset_Attribute extends Message<Asset_Attribute> {
  /**
   * @generated from field: string name = 1;
   */
  name: string;

  /**
   * @generated from field: string value = 2;
   */
  value: string;

  constructor(data?: PartialMessage<Asset_Attribute>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "asset.v1.Asset.Attribute";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Asset_Attribute;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Asset_Attribute;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Asset_Attribute;

  static equals(a: Asset_Attribute | PlainMessage<Asset_Attribute> | undefined, b: Asset_Attribute | PlainMessage<Asset_Attribute> | undefined): boolean;
}


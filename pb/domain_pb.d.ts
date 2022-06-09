// package: core
// file: domain.proto

import * as jspb from "google-protobuf";

export class Asset extends jspb.Message {
  getId(): number;
  setId(value: number): void;

  getOwner(): string;
  setOwner(value: string): void;

  getIssuer(): string;
  setIssuer(value: string): void;

  getCo2e(): number;
  setCo2e(value: number): void;

  getCertid(): number;
  setCertid(value: number): void;

  getMetadata(): string;
  setMetadata(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Asset.AsObject;
  static toObject(includeInstance: boolean, msg: Asset): Asset.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Asset, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Asset;
  static deserializeBinaryFromReader(message: Asset, reader: jspb.BinaryReader): Asset;
}

export namespace Asset {
  export type AsObject = {
    id: number,
    owner: string,
    issuer: string,
    co2e: number,
    certid: number,
    metadata: string,
  }
}

export class AssetCertificates extends jspb.Message {
  clearCertificatesList(): void;
  getCertificatesList(): Array<number>;
  setCertificatesList(value: Array<number>): void;
  addCertificates(value: number, index?: number): number;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): AssetCertificates.AsObject;
  static toObject(includeInstance: boolean, msg: AssetCertificates): AssetCertificates.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: AssetCertificates, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): AssetCertificates;
  static deserializeBinaryFromReader(message: AssetCertificates, reader: jspb.BinaryReader): AssetCertificates;
}

export namespace AssetCertificates {
  export type AsObject = {
    certificatesList: Array<number>,
  }
}

export class AssetComposition extends jspb.Message {
  clearAssetsList(): void;
  getAssetsList(): Array<number>;
  setAssetsList(value: Array<number>): void;
  addAssets(value: number, index?: number): number;

  clearPortionsList(): void;
  getPortionsList(): Array<number>;
  setPortionsList(value: Array<number>): void;
  addPortions(value: number, index?: number): number;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): AssetComposition.AsObject;
  static toObject(includeInstance: boolean, msg: AssetComposition): AssetComposition.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: AssetComposition, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): AssetComposition;
  static deserializeBinaryFromReader(message: AssetComposition, reader: jspb.BinaryReader): AssetComposition;
}

export namespace AssetComposition {
  export type AsObject = {
    assetsList: Array<number>,
    portionsList: Array<number>,
  }
}

export class AssetTraceability extends jspb.Message {
  clearMovementsList(): void;
  getMovementsList(): Array<number>;
  setMovementsList(value: Array<number>): void;
  addMovements(value: number, index?: number): number;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): AssetTraceability.AsObject;
  static toObject(includeInstance: boolean, msg: AssetTraceability): AssetTraceability.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: AssetTraceability, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): AssetTraceability;
  static deserializeBinaryFromReader(message: AssetTraceability, reader: jspb.BinaryReader): AssetTraceability;
}

export namespace AssetTraceability {
  export type AsObject = {
    movementsList: Array<number>,
  }
}

export class Certificate extends jspb.Message {
  getId(): number;
  setId(value: number): void;

  getIssuer(): string;
  setIssuer(value: string): void;

  getMetadata(): string;
  setMetadata(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Certificate.AsObject;
  static toObject(includeInstance: boolean, msg: Certificate): Certificate.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Certificate, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Certificate;
  static deserializeBinaryFromReader(message: Certificate, reader: jspb.BinaryReader): Certificate;
}

export namespace Certificate {
  export type AsObject = {
    id: number,
    issuer: string,
    metadata: string,
  }
}

export class Movement extends jspb.Message {
  getId(): number;
  setId(value: number): void;

  getIssuer(): string;
  setIssuer(value: string): void;

  getCo2e(): number;
  setCo2e(value: number): void;

  getCertid(): number;
  setCertid(value: number): void;

  getMetadata(): string;
  setMetadata(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Movement.AsObject;
  static toObject(includeInstance: boolean, msg: Movement): Movement.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Movement, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Movement;
  static deserializeBinaryFromReader(message: Movement, reader: jspb.BinaryReader): Movement;
}

export namespace Movement {
  export type AsObject = {
    id: number,
    issuer: string,
    co2e: number,
    certid: number,
    metadata: string,
  }
}


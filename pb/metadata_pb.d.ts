// package: core
// file: metadata.proto

import * as jspb from "google-protobuf";

export class Image extends jspb.Message {
  getName(): string;
  setName(value: string): void;

  getAttachment(): string;
  setAttachment(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Image.AsObject;
  static toObject(includeInstance: boolean, msg: Image): Image.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Image, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Image;
  static deserializeBinaryFromReader(message: Image, reader: jspb.BinaryReader): Image;
}

export namespace Image {
  export type AsObject = {
    name: string,
    attachment: string,
  }
}

export class Link extends jspb.Message {
  getName(): string;
  setName(value: string): void;

  getIcon(): string;
  setIcon(value: string): void;

  getUrl(): string;
  setUrl(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Link.AsObject;
  static toObject(includeInstance: boolean, msg: Link): Link.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Link, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Link;
  static deserializeBinaryFromReader(message: Link, reader: jspb.BinaryReader): Link;
}

export namespace Link {
  export type AsObject = {
    name: string,
    icon: string,
    url: string,
  }
}

export class AssetMetadata extends jspb.Message {
  getName(): string;
  setName(value: string): void;

  getHeader(): string;
  setHeader(value: string): void;

  getDescription(): string;
  setDescription(value: string): void;

  hasBackground(): boolean;
  clearBackground(): void;
  getBackground(): Image | undefined;
  setBackground(value?: Image): void;

  clearImagesList(): void;
  getImagesList(): Array<Image>;
  setImagesList(value: Array<Image>): void;
  addImages(value?: Image, index?: number): Image;

  clearLinksList(): void;
  getLinksList(): Array<Link>;
  setLinksList(value: Array<Link>): void;
  addLinks(value?: Link, index?: number): Link;

  clearPropertiesList(): void;
  getPropertiesList(): Array<AssetMetadata.Property>;
  setPropertiesList(value: Array<AssetMetadata.Property>): void;
  addProperties(value?: AssetMetadata.Property, index?: number): AssetMetadata.Property;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): AssetMetadata.AsObject;
  static toObject(includeInstance: boolean, msg: AssetMetadata): AssetMetadata.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: AssetMetadata, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): AssetMetadata;
  static deserializeBinaryFromReader(message: AssetMetadata, reader: jspb.BinaryReader): AssetMetadata;
}

export namespace AssetMetadata {
  export type AsObject = {
    name: string,
    header: string,
    description: string,
    background?: Image.AsObject,
    imagesList: Array<Image.AsObject>,
    linksList: Array<Link.AsObject>,
    propertiesList: Array<AssetMetadata.Property.AsObject>,
  }

  export class Property extends jspb.Message {
    getName(): string;
    setName(value: string): void;

    getValue(): string;
    setValue(value: string): void;

    getIcon(): string;
    setIcon(value: string): void;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): Property.AsObject;
    static toObject(includeInstance: boolean, msg: Property): Property.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: Property, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): Property;
    static deserializeBinaryFromReader(message: Property, reader: jspb.BinaryReader): Property;
  }

  export namespace Property {
    export type AsObject = {
      name: string,
      value: string,
      icon: string,
    }
  }
}

export class CertificateMetadata extends jspb.Message {
  getDate(): number;
  setDate(value: number): void;

  getUrl(): string;
  setUrl(value: string): void;

  getAttachment(): string;
  setAttachment(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CertificateMetadata.AsObject;
  static toObject(includeInstance: boolean, msg: CertificateMetadata): CertificateMetadata.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: CertificateMetadata, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CertificateMetadata;
  static deserializeBinaryFromReader(message: CertificateMetadata, reader: jspb.BinaryReader): CertificateMetadata;
}

export namespace CertificateMetadata {
  export type AsObject = {
    date: number,
    url: string,
    attachment: string,
  }
}

export class MovementMetadata extends jspb.Message {
  getCity(): string;
  setCity(value: string): void;

  getStartDate(): string;
  setStartDate(value: string): void;

  getEndDate(): string;
  setEndDate(value: string): void;

  getDistance(): number;
  setDistance(value: number): void;

  getType(): MovementMetadata.TransportMap[keyof MovementMetadata.TransportMap];
  setType(value: MovementMetadata.TransportMap[keyof MovementMetadata.TransportMap]): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): MovementMetadata.AsObject;
  static toObject(includeInstance: boolean, msg: MovementMetadata): MovementMetadata.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: MovementMetadata, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): MovementMetadata;
  static deserializeBinaryFromReader(message: MovementMetadata, reader: jspb.BinaryReader): MovementMetadata;
}

export namespace MovementMetadata {
  export type AsObject = {
    city: string,
    startDate: string,
    endDate: string,
    distance: number,
    type: MovementMetadata.TransportMap[keyof MovementMetadata.TransportMap],
  }

  export interface TransportMap {
    CREATE: 0;
    AIR: 1;
    SEA: 2;
    LAND: 3;
  }

  export const Transport: TransportMap;
}

export class AccountMetadata extends jspb.Message {
  getName(): string;
  setName(value: string): void;

  getDescription(): string;
  setDescription(value: string): void;

  hasLogo(): boolean;
  clearLogo(): void;
  getLogo(): Image | undefined;
  setLogo(value?: Image): void;

  clearLinksList(): void;
  getLinksList(): Array<Link>;
  setLinksList(value: Array<Link>): void;
  addLinks(value?: Link, index?: number): Link;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): AccountMetadata.AsObject;
  static toObject(includeInstance: boolean, msg: AccountMetadata): AccountMetadata.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: AccountMetadata, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): AccountMetadata;
  static deserializeBinaryFromReader(message: AccountMetadata, reader: jspb.BinaryReader): AccountMetadata;
}

export namespace AccountMetadata {
  export type AsObject = {
    name: string,
    description: string,
    logo?: Image.AsObject,
    linksList: Array<Link.AsObject>,
  }
}

export class RoleMetadata extends jspb.Message {
  getInfo(): string;
  setInfo(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): RoleMetadata.AsObject;
  static toObject(includeInstance: boolean, msg: RoleMetadata): RoleMetadata.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: RoleMetadata, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): RoleMetadata;
  static deserializeBinaryFromReader(message: RoleMetadata, reader: jspb.BinaryReader): RoleMetadata;
}

export namespace RoleMetadata {
  export type AsObject = {
    info: string,
  }
}


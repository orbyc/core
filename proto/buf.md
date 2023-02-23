# Orbyc Supply Chain Core Documentation

This package contains the core protocol buffers for the Orbyc Supply Chain Suite. The protocol buffers in this package define the data structures and formats used throughout the suite.

## Module: Asset

The Asset module defines the protocol buffer for representing assets in the `Orbyc Supply Chain Suite`. An asset represents a physical object or product that is tracked through the supply chain.

The Asset protocol buffer has the following fields:

- `name` (string): The name of the asset.
- `description` (string): A description of the asset.
- `image` (string): A URL to an image of the asset.
- `external_url` (string, optional): A URL to an external webpage with more information about the asset.
- `attributes` (repeated Attribute): A list of attributes associated with the asset.
- `footprint` (oneof): Either `co2e` or `co2o`. `co2e` represents the carbon emissions associated with the asset, while `co2o` represents the carbon offset associated with the asset.
- `gallery` (repeated string): A list of URLs to images of the asset.
- `video` (string, optional): A URL to a video of the asset.
- `created_at` (google.protobuf.Timestamp): The timestamp when the asset was created.
- `updated_at` (google.protobuf.Timestamp): The timestamp when the asset was last updated.

## Module: Certificate

The Certificate module defines the protocol buffer for representing certificates in the Orbyc Supply Chain Suite. A certificate is a document that certifies a specific attribute of an asset, such as its sustainability or origin.

The Certificate protocol buffer has the following fields:

- `name` (string): The name of the certificate.
- `description` (string): A description of the certificate.
- `issuer` (string): The name of the organization that issued the certificate.
- `external_reference` (string, optional): An external reference number associated with the certificate.
- `issued_at` (google.protobuf.Timestamp): The timestamp when the certificate was issued.

## Module: Trace

The Trace module defines the protocol buffer for representing traces in the Orbyc Supply Chain Suite. A trace represents the movement of an asset through the supply chain, from its origin to its destination.

The Trace protocol buffer has the following fields:

- `type` (Type): The type of transportation used for the trace.
- `departure` (Location, optional): The location where the trace started.
- `destination` (Location): The location where the trace ended.
- `description` (string, optional): A description of the trace.
- `footprint` (oneof): Either `co2e` or `co2o`. `co2e` represents the carbon emissions associated with the trace, while `co2o` represents the carbon offset associated with the trace.
- `started_at` (google.protobuf.Timestamp): The timestamp when the trace started.
- `duration` (google.protobuf.Duration): The duration of the trace.
- `Location`:
    - `place` (string): The name of the location.
    - `lat` (string, optional): The latitude of the location.
    - `lng` (string, optional): The longitude of the location.

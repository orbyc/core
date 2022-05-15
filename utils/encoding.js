import { Buffer } from "buffer";

export function encodeHex(data) {
    return Buffer.from(data).toString("hex")
}

export function decodeHex(encoded) {
    return Uint8Array.from(Buffer.from(encoded, "hex"))
}
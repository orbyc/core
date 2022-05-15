import { Buffer } from "buffer";

export function encodeHex(data: Uint8Array): string {
    return Buffer.from(data).toString("hex")
}

export function decodeHex(encoded: string): Uint8Array {
    return Uint8Array.from(Buffer.from(encoded, "hex"))
}
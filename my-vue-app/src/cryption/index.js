import {JSEncrypt} from "jsencrypt";
var crypt = new JSEncrypt();
var decrpt = new JSEncrypt()

const PublicKey  = "-----BEGIN PUBLIC KEY-----\nMFwwDQYJKoZIhvcNAQEBBQADSwAwSAJBAMz4u0O0AI4TfvilTlJxQedoZgA8NuEC\nM/N16k3Pb2mObSpPGDIYdkqm1w3mjUGqK+Sld0Za2JEGAHrN1w+VvVECAwEAAQ==\n-----END PUBLIC KEY-----\n"
const encoder = new TextEncoder();
crypt.setPublicKey(PublicKey)
// 导出函数

export function encrypt(sting){
    const enc = crypt.encrypt(sting)
    const dataBytes = encoder.encode(enc);
    const hexString = Array.from(new Uint8Array(dataBytes))
        .map(b => b.toString(16).padStart(2, '0'))
        .join('');
    return hexString
}
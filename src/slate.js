/**
 */
//
import { lexer } from "./lexer.js";
import { parser } from "./parser.js";
import { converter } from "./converter.js";
import "https://cdn.rawgit.com/WebAssembly/wabt/e5b3830/demo/libwabt.js";

//
export {
    lexer,
    parser,
    converter,
}

//
const wabt_module = WabtModule();

//
export async function convert_to_wasm(text) {
    return Promise.resolve(text)
    .then(x => wabt_module.parseWat("test.wat", x))
    .then(x => x.toBinary({}))
    .then(x => x.buffer)
    .then(x => new Blob([x], { type:"application/wasm" }))
    .then(x => new Response(x));
}

//
export async function parse(text) {
    return Promise.resolve(text)
    .then(x => lexer.parse(x))
    .then(x => parser.convert_tokens_to_expressions(x))
    .then(x => parser.parse(x))
    .then(x => converter.normalize(x))
    .then(x => x.convert_to_wast())
    .then(x => convert_to_wasm(x))
}

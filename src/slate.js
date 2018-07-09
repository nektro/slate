/**
 */
//
import { lexer } from "./lexer.js";
import { parser } from "./parser.js";
import { converter } from "./converter.js";

//
export {
    lexer,
    parser,
    converter,
}

//
export async function parse(text) {
    return Promise.resolve(text)
    .then(x => lexer.parse(x))
    .then(x => parser.convert_tokens_to_expressions(x))
    .then(x => parser.parse(x))
    .then(x => converter.normalize(x))
    .then(x => x.convert_to_wast())
}

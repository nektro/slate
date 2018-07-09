/**
 */
//
import { lexer } from "./lexer.js";
import { parser } from "./parser.js";

//
export {
    lexer,
    parser,
}

//
export async function parse(text) {
    return Promise.resolve(text)
    .then(x => lexer.parse(x))
    .then(x => parser.convert_tokens_to_expressions(x))
    .then(x => parser.parse(x))
}

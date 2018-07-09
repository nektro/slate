/**
 */
//
import { lexer } from "./lexer.js";

//
export {
    lexer,
}

//
export async function parse(text) {
    return Promise.resolve(text)
    .then(x => lexer.parse(x))
}

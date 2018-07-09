/**
 */
//
import { Lexer } from "./imports.js";

//
export const lexer = new (class SlateLexer extends Lexer {
    constructor() {
        super(
            ["export","const","return","function"],
            ["=",";","(",")",":","{","}","+"],
            ["\""],
            true,
            true
        );
    }
    isValidVarChar(c) {
        return (c >= "a" && c <= "z") || (c >= "A" && c <= "Z") || (c === "_") || (c >= "0" && c <= "9");
    }
})();

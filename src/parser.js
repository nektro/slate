/**
 */
//
import { Parser, ExpressionSimple, ExpressionContainer } from "./imports.js";

//
const nES = (...args) => new ExpressionSimple(...args);
const nEC = (...args) => new ExpressionContainer(...args);

//
export const parser = new (class SlateParser extends Parser {
    constructor() {
        super();
        //
        this.addRule(["K_const", "Id", "S_="], (list, i) => nEC("ConstDef", [list[i + 1]]), true);
        this.addRule(["Int", "S_+", "Int"], (list, i) => nEC("Add", [list[i], list[i + 2]]));
        this.addRule(["ConstDef", "Literal"], (list, i) => nEC("Const", [list[i].value[0], list[i + 1]]));
        this.addRule(["Declaration", "S_;"], (list, i) => nEC("Initialization", [list[i]]));
        this.addRule(["K_export", "Exportable"], (list, i) => nEC("Export", [list[i + 1]]));
        this.addRule(["K_return", "Returnable"], (list, i) => nEC("Return", [list[i + 1]]));
        this.addRule(["CanBeStatement", "S_;"], (list, i) => nEC("Statement", [list[i]]));
        this.addRule(["S_{", "Statement"], (list, i) => nEC("BlockPart", [list[i + 1]]));
        this.addRule(["BlockPart", "S_}"], (list, i) => nEC("Block", [...list[i].value]));
        this.addRule(["S_(", "S_)"], (list, i) => nEC("EmptyParen", []));
        this.addRule(["K_function", "Id", "EmptyParen"], (list, i) => nEC("FunctionPart_1", [list[i + 1], list[i + 2]]));
        this.addRule(["FunctionPart_1", "S_:", "Id"], (list, i) => nEC("FunctionPart_2", [...list[i].value, list[i + 2]]));
        this.addRule(["FunctionPart_2", "Block"], (list, i) => nEC("Function", [...list[i].value, list[i + 1]]));
        this.addRule(["TopLevel"], (list, i) => nEC("File", [list[i]]));
        this.addRule(["File", "File"], (list, i) => nEC("File", [...list[i].value, ...list[i + 1].value]));

        this.addAlias("Literal", ["Int", "String"]);
        this.addAlias("Declaration", ["Const"]);
        this.addAlias("Returnable", ["Add"]);
        this.addAlias("Exportable", ["Initialization", "Function"]);
        this.addAlias("CanBeStatement", ["Return"]);
        this.addAlias("TopLevel", ["Export"]);
    }
    is_valid_identifier(word) {
        return new RegExp(/[a-zA-Z][a-zA-Z0-9_]*/, "g").test(word);
    }
})();

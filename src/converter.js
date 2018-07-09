/**
 */
//
import * as ob from "./objects.js";

//
export const converter = new (class SlateConverter {
    constructor() {
    }
    normalize(e) {
        switch (e.type) {
            case "Block":
                return e.value.map(x => this.normalize(x));
            case "File":
            case "Export":
            case "Const":
            case "Function":
            case "Return":
            case "Add":
                return new (this._normalize_to_class(e))(...e.value.map(this.normalize.bind(this)));
            case "Id":
            case "String":
            case "Int":
                return e.value;
            case "Initialization":
            case "Statement":
                return this.normalize(e.value[0]);
            case "EmptyParen":
                return [];
            default: {
                console.debug(e);
                throw new Error(`ConversionError: No conversion found for expression type: ${e.type}`);
            }
        }
    }
    _normalize_to_class(e) {
        switch (e.type) {
            case "File": return ob.S_Module;
            case "Export": return ob.S_Export;
            case "Const": return ob.S_Const;
            case "Function": return ob.S_Function;
            case "Return": return ob.S_Return;
            case "Add": return ob.S_Add;

            default: {
                console.debug(e);
                throw new Error(`ConversionError: No class found for expression type: ${e.type}`);
            }
        }
    }
})();

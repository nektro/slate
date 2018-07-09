/**
 */
//
export class S_Object {
    constructor() {
    }
    convert_to_wast() {
        console.debug(this);
        console.warn(`GeneratorError: No WAST conversion found for type: ${this.__proto__.constructor.name}`);
        return "";
    }
}
export class S_Module extends S_Object {
    constructor(...x) {
        super();
        this.top = x;
    }
    convert_to_wast() {
        return `(module (memory $0 1)${this.top.map(v => v.convert_to_wast(), "").join("")})`;
    }
}
export class S_Export extends S_Object {
    constructor(v) {
        super();
        this.value = v;
    }
    convert_to_wast() {
        if (this.value instanceof S_Const) return ` ${this._b()} (export "${this.value.name}" (global 0))`;
        if (this.value instanceof S_Function) return this.value.convert_to_wast(true);
    }
    _b() {
        switch (typeof this.value.value) {
            case "bigint": return `(global i32 (i32.const ${this.value.value}))`;
            default: throw new Error("SlateError: no export wasm for type of " + typeof this.value.value);
        }
    }
}
export class S_Const extends S_Object {
    constructor(n, v) {
        super();
        this.name = n;
        this.value = v;
    }
}
export class S_Function extends S_Object {
    constructor(n, p, r, b) {
        super();
        this.name = n;
        this.parameters = p;
        this.returnType = r;
        this.body = b;
    }
    convert_to_wast(to_export) {
        return ` (func (export "${this.name}") (result ${this.returnType})${this.body.map(x => x.convert_to_wast()).join('')})`;
    }
}
export class S_Return extends S_Object {
    constructor(v) {
        super();
        this.value = v;
    }
    convert_to_wast() {
        return this.value.convert_to_wast();
    }
}
export class S_Add extends S_Object {
    constructor(x, y) {
        super();
        this.a = x;
        this.b = y;
    }
    convert_to_wast() {
        return ` (i32.add (i32.const ${this.a}) (i32.const ${this.b}))`;
    }
}

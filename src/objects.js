/**
 */
//
export class S_Object {
    constructor() {
    }
}
export class S_Module extends S_Object {
    constructor(...x) {
        super();
        this.top = x;
    }
}
export class S_Export extends S_Object {
    constructor(v) {
        super();
        this.value = v;
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
}
export class S_Return extends S_Object {
    constructor(v) {
        super();
        this.value = v;
    }
}
export class S_Add extends S_Object {
    constructor(x, y) {
        super();
        this.a = x;
        this.b = y;
    }
}

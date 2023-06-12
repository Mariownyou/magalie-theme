function someFunction(a, b, c) {
    let x = 1;
    const y = 2;
    var z = 3;

    console.log('someFunction called')

    return a + b + c;
}

class Some {
    constructor() {
        console.log('Some constructor called')
    }
}

class SomeOther extends Some {
    constructor() {
        super();
        console.log('SomeOther constructor called')
    }
}

export { someFunction, Some }

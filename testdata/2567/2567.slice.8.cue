// 2567.slice.8.cue
e: _
d: a[3:]        // slice in a top level field
c: {
    z: a[2:]    // slice in a struct field
}
b:[
    a[1:]      // slice in a list
]
a: [1,2,3,4,5]

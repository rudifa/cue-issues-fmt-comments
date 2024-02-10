// 2567.8.comprehension.in.X.BAD.cue

x: [
    if true // about this condition MISPLACED
    {}
    if false // about that condition BADLY MISPLACED
    {}]

y: {

    for k, v in #a {
        "\( v )": {         // about the key
            nameLen: len(v) // about the value
        }}
}

#a: ["Barcelona", "Shanghai", "Munich"]

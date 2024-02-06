// 2567.8.comprehension.in.X.BAD.cue

x: [
    if true // about this condition
    {}
    if false // about that condition
    {}]

y: {

    for k, v in #a {
        "\( v )": {         // about the key
            nameLen: len(v) // about the value
        }}
}

#a: ["Barcelona", "Shanghai", "Munich"]

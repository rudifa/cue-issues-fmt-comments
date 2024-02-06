// 2567.8.indexexpr.in.X.3.BAD.cue

foo: [
	bar["baz"], // in list: misplaced comment
]

goo: {
	bar["baz"] // in struct: correctly placed
}


hoo: bar["baz"], // top level: correctly placed

bar: baz: 41

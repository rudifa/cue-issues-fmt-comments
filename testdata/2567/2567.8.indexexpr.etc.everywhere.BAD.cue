// 2567.8.indexexpr.etc.everywhere.BAD.cue

foo: [
	bar["baz"], // indexexpr in list: comment misplaced
	11,              // an integer
	{three:   33},   // a struct
	[44],            // a list
]

goo: {
	a: bar["baz"],      // in struct: comment ok
	b: 41,              // an integer
	c: {three:   33},   // a struct
	d: [44],            // a list
}

hoo: bar["baz"], // top level: comment ok

bar: baz: 41



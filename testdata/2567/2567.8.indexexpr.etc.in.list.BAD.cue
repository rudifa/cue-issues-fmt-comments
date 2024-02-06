// 2567.8.indexexpr.etc.in.list.BAD.cue
foo: [
	11,              // an integer
	bar["baz"],      // an index expr
	{three:   33},   // a struct
	[44],            // a list
]
bar: baz: 41

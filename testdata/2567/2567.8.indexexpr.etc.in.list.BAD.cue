// 2567.8.indexexpr.etc.in.list.BAD.cue
foo: [
	bar["baz"],      // an index expr
	11,              // an integer
	{three:   33},   // a struct
	[44],            // a list
]
bar: baz: 41

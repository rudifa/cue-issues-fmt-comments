// 2567.n.fbb.s.cue
foo: [
	1              // a number
	bar["baz"]     // an index expr
	{               // a struct opens
	    three:      // the label
	            3   // the value
	}              // a struct closes
]
bar: baz: 41

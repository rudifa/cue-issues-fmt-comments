// 2567.indexexpr.list.number.comprehension.cue
foo: [
	bar["baz"],     // an index expression
	["four"],       // a list
	1,              // an integer
	if true         // a condition in a comprehension
    {}]             // fmt => invalid cue
bar: baz: string

// 2423/2423.comment.in.empty.struct.BAD.cue
ex0: {
	v: 10 // ex0: comment after a value
} & {
	// ex0: comment in an empty struct
}

ex1: [
	11, // ex1: comment after a value
] & [
	// ex1: comment in an empty list
	...
]

ex2: {
	// ex2: comment in an empty struct
}

ex3: [
	// ex3: comment in an empty list
]

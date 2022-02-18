let readline = require('readline');

let rl = readline.createInterface({
	input: process.stdin,
	output: process.stdout
});

let prog;
let input;

rl.question('Prog: ', programString => {
	rl.question('Input: ', inputString => {
		prog = programString.split('');
		input = inputString.split('');
		rl.close();
	});
});

rl.on('close', function () {
	console.log('\nYour input:');
	console.log(prog);
	console.log(input);
	process.exit(0);
});
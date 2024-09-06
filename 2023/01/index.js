const input = require('fs').readFileSync('./input.txt').toString().split('\n');

console.log('Part 1:', input.reduce((acc, value) => {
    const cleantValue = value.trim().replace(/\D/g, '');

    if (!cleantValue) return acc;

    return acc + +(cleantValue[0] + cleantValue[cleantValue.length - 1]);
}, 0))

// pt2

console.log('Part 2:', input.reduce((acc, value) => {
    let residualValue = value.trim();
    let i = 0;

    while (i<residualValue.length) {
        const initial = residualValue.slice(0, i);

        const final = residualValue.slice(i)
            .replace(/^one/, '1')
            .replace(/^two/, '2')
            .replace(/^three/, '3')
            .replace(/^four/, '4')
            .replace(/^five/, '5')
            .replace(/^six/, '6')
            .replace(/^seven/, '7')
            .replace(/^eight/, '8')
            .replace(/^nine/, '9');

        i++;

        residualValue = initial + final
    }

    const cleantValue = residualValue.replace(/\D/g, '');

    if (!cleantValue) return acc;

    const computed = +(cleantValue[0] + cleantValue[cleantValue.length - 1]);

    // console.log(value, cleantValue, computed)

    return acc + computed;

}, 0));

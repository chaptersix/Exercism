var resistance = [
    "black",
    "brown",
    "red",
    "orange",
    "yellow",
    "green",
    "blue",
    "violet",
    "grey",
    "white"
]

export function value(colors) {
    if (!Array.isArray(colors)) {
        throw "invalid argument";
    }
    let numStr = "";
    for (let i = 0; i < colors.length; i++) {
        numStr += resistance.indexOf(colors[i]);
    }
    return parseInt(numStr, 10);
}
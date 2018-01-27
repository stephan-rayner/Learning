const a = {
    A: 1,
    B: 2,
    C: 3
}

console.log(a)
delete a.B
console.log(a)
delete a.E

if (!a.B) {
    console.log("Couldn't find B")
}

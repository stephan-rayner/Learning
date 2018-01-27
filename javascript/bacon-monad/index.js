const Bacon = require('baconjs')
const stream = new Bacon.Bus()

stream
  .filter(element => typeof(element) === 'string')
  .filter(word => word.length > 3)
  .map(word => word.toUpperCase())
  .onValue(word => console.log(word))

stream.push('cheese')
stream.push([])
stream.push('dog')
stream.push('doodle')
stream.push('waldo')

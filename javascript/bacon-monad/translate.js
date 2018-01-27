const fetch = require('node-fetch')
const Bacon = require('baconjs')

function getInFrench(word) {
    const apiKey = 'xxxxxx'
    const url = 'https://www.googleapis.com' +
        '/language/translate/v2' +
        '?key=' + apiKey +
        '&source=en' +
        '&target=fr' +
        '&q=' + encodeURIComponent(word)

    const promise = fetch(url)
        .then(response => response.json())
        .then(parsedResponse => {
            parsedResponse
                .data
                .translations[0]
                .translatedText
        })
        .catch(reason => console.error("[ERROR]:\n\t", reason))

    return Bacon.fromPromise(promise)
}

const stream = new Bacon.Bus()

stream
    .flatMap(word => getInFrench(word))
    .map(word => word.toUpperCase())
    .onValue(word => console.log(word))

stream.push('cheese')
stream.push('dog')
stream.push('doodle')
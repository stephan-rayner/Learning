const Redux = require('redux')

const counter = (state = 0, action) => {
    switch (action.type) {
        case "INCREMENT":
            return state + 1
        case "DECREMENT":
            return state - 1
        default:
            return state
    }
}

const render = () =>  console.log("Stuff has changed!!!")

// Creating the Store
const { createStore } = Redux
const store = createStore(counter)

// Setting up a subscription
store.subscribe(render)

// Poking the Bear
console.log("BEFORE INCREMENT", store.getState())
store.dispatch({type: "INCREMENT"})
console.log("AFTER INCREMENT", store.getState())

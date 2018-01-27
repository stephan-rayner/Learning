/* ===================================
    Redux from Scratch 
 =================================== */
const createStore(reducer) => {
    let state
    let listeners = []

    const getState() => state

    const dispatch = (action) => {
        state = reducer(state,  action)
        listeners.forEach(listener => listener())
    }

    const subscribe = (listener) => {
        listeners.push(listener)
        return () => {
            // Return a fuction that the user can call to unsubscribe
            listeners = listeners.filter(l => l !=== listeners)
        }
    }

    dispatch({})

    return {getState, dispatch,  subscribe}
}

/* ===================================
    The Application
 =================================== */
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

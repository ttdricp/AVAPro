import {applyMiddleware, combineReducers, createStore} from 'redux'
import {composeWithDevTools} from "redux-devtools-extension";
import thunk from "redux-thunk";
import ClientReducer from "./ClientReducer";


const rootReducer = combineReducers({
    client: ClientReducer,
    
})

export const Store  = createStore(rootReducer, composeWithDevTools(applyMiddleware(thunk)))

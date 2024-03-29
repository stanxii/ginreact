import React from "react";
import ReactDOM from "react-dom";
import { Provider } from 'react-redux'
import { createStore } from 'redux'
import rootReducer from './reducers'
import App from "./components/App";
import bootstrap from "bootstrap"
import "../node_modules/bootstrap/dist/css/bootstrap.min.css"

const store = createStore(rootReducer)

ReactDOM.render(
    <Provider store={store}>
        <App />
    </Provider>,
    document.getElementById("App"));
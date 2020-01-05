import React from "react";
import ReactDOM from "react-dom";
import { Provider } from 'react-redux'
import { createStore } from 'redux'
import App from "./components/App";
import bootstrap from "bootstrap"
import "../node_modules/bootstrap/dist/css/bootstrap.min.css"

ReactDOM.render(
    <App />
    , document.getElementById("App"));
import React from 'react';
import {createRoot} from 'react-dom/client';
import App from './App';

import {Store} from "./Reducers";
import {Provider} from "react-redux";
import {BrowserRouter} from "react-router-dom";

const container = document.getElementById('root')
const root = createRoot(container)
root.render(
    <Provider store={Store}>
        <BrowserRouter>
            <App />
        </BrowserRouter>
    </Provider>
)
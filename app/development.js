import React from 'react'
import ReactDOM from 'react-dom'
import { createStore, combineReducers } from 'redux'
import { Provider } from 'react-redux'
import { Router, Route } from 'react-router'
import { createHistory } from 'history'
import { syncReduxAndRouter, routeReducer } from 'redux-simple-router'

import {Main} from './main'
import {Routes} from './routes'
const reducers = require('./reducers');

const reducer = combineReducers(Object.assign({}, reducers, {
    routing: routeReducer
}));
const store = createStore(reducer);
const history = createHistory();

syncReduxAndRouter(history, store);

Main({assets_path: '/assets'});


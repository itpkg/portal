import { createStore, combineReducers } from 'redux'
import { syncReduxAndRouter, routeReducer } from 'redux-simple-router'
import createBrowserHistory from 'history/lib/createBrowserHistory'

import currentUser from './reducers/current_user'


const reducer = combineReducers(Object.assign({}, currentUser, {
    routing: routeReducer
}));

const store = createStore(reducer);
const history = createBrowserHistory();

syncReduxAndRouter(history, store);

module.exports = {
    Store: store,
    History: history
};

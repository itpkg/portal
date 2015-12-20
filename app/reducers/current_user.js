import {SIGN_IN, SIGN_OUT} from '../actions'

function currentUser(state = {}, action) {
    switch (action.type) {
        case SIGN_IN:
            console.log("reducer sign in");
            return Object.assign({}, state, {current_user: action.user});
        case SIGN_OUT:
            console.log("reducer sign out");
            return Object.assign({}, state);
        default:
            return state;
    }
}

export default currentUser
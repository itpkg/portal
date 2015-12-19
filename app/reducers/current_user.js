const constants = require('../constants');

const initialState = {
    current_user: null
};

function update(state = initialState, action) {
    if(action.type === constants.SIGN_IN) {
        return { current_user:action.user };
    }
    else if(action.type === constants.SIGN_OUT) {
        return { current_user:null };
    }
    return state;
}
module.exports = update;
var Reflux = require('reflux');

const current_user = "current_user";

const actions = Reflux.createActions(['signIn', 'signOut']);

const store = Reflux.createStore({
    listenables: actions,
    init: function () {
        this.data = {
            current_user: JSON.parse(localStorage.getItem(current_user))
        };
    },
    onSignIn: function (user) {
        localStorage.setItem(current_user, JSON.stringify(user));
        this.data.current_user = user;
    },
    onSingOut: function () {
        localStorage.removeItem(current_user);
        this.data.current_user = null;
    },
    getInitialState: function () {
        return this.data;
    }
});

module.exports = {
    Store: store,
    Actions: actions
};
import React from 'react';
import {Router, Route, IndexRoute} from 'react-router';

import createBrowserHistory from 'history/lib/createBrowserHistory'

import {NoMatch, Version} from './Widgets'
import {Application, Home, AboutUs, Notice} from  './Application'
import {Users, SignIn, SignUp, Confirm, Unlock, ForgotPassword, ResetPassword} from './Users'
import {Personal, Profile} from './Personal'
import {Message} from "./Message"


const routes = (<Router history={createBrowserHistory()}>
    <Route path="/" component={Application}>
        <IndexRoute component={Home}/>
        <Route path="message" component={Message}/>
        <Route path="about-us" component={AboutUs}/>
        <Route path="notices/:id" component={Notice}/>

        <Route path="users" component={Users}>
            <Route path="sign-in" component={SignIn}/>
            <Route path="sign-up" component={SignUp}/>
            <Route path="confirm" component={Confirm}/>
            <Route path="unlock" component={Unlock}/>
            <Route path="forgot-password" component={ForgotPassword}/>
            <Route path="reset-password" component={ResetPassword}/>
        </Route>

        <Route path="personal" component={Personal}>
            <Route path="profile" component={Profile}/>
        </Route>

        <Route path="*" component={NoMatch}/>
    </Route>
</Router>);

module.exports = {
    Routes:routes
};
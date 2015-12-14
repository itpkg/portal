"use strict";

import $ from 'jquery'

import React from 'react'
import  ReactDOM  from 'react-dom'
import { Router, Route, IndexRoute} from 'react-router'

import {NoMatch, Version} from './Widgets'
import {Application, Home, AboutUs, Notice} from  './Application'
import {Users, SignIn, SignUp, Confirm, Unlock, ForgotPassword, ResetPassword, Profile} from './Users'


ReactDOM.render(
    (<Router>
        <Route path="/users" component={Users}>
            <Route path="sign-in" component={SignIn}/>
            <Route path="sign-up" component={SignUp}/>
            <Route path="confirm" component={Confirm}/>
            <Route path="unlock" component={Unlock}/>
            <Route path="forgot-password" component={ForgotPassword}/>
            <Route path="reset-password" component={ResetPassword}/>
        </Route>

        <Route path="/" component={Application}>
            <IndexRoute component={Home}/>
            <Route path="about_us" component={AboutUs}/>
            <Route path="notices/:id" component={Notice}/>
        </Route>

        <Route path="*" component={NoMatch}/>
    </Router>),
    document.getElementById('root')
);

$(function(){
    console.log('init');
});
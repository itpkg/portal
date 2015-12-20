import React from 'react';
import {Route, IndexRoute} from 'react-router';

import {Application} from './Layout'
import {Home, AboutUs, NoMatch, Message} from './Pages'
import {Users, SignIn, SignUp, Confirm, Unlock, ForgotPassword,ResetPassword} from './Users'
import {Personal,Profile} from './Personal'

const routes = (
    <Route path="/" component={Application}>
        <IndexRoute component={Home}/>
        <Route path="message" component={Message}/>
        <Route path="about-us" component={AboutUs}/>

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
);

export default routes
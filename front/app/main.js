"use strict";

require('bootstrap/dist/css/bootstrap.css');

import React from 'react';
import ReactDOM from 'react-dom';
import {Router, Route, IndexRoute} from 'react-router';

import i18next from 'i18next/lib';
import LanguageDetector from 'i18next-browser-languagedetector/lib';

import {NoMatch, Version} from './Widgets'
import {Application, Home, AboutUs, Notice} from  './Application'
import {Users, SignIn, SignUp, Confirm, Unlock, ForgotPassword, ChangePassword, Profile} from './Users'
import {enUS, zhCN} from './locales'

i18next.use(LanguageDetector).init({
    fallbackLng: 'en',
    resources: {
        'en': enUS,
        'zh-CN': zhCN
    },
    detection: {
        order: ['querystring', 'cookie', 'localStorage', 'navigator'],
        lookupQuerystring: 'locale',
        lookupCookie: 'locale',
        lookupLocalStorage: 'locale',
        caches: ['localStorage', 'cookie']
    }
}, (err, t)=> {
    ReactDOM.render(
        (<Router>
            <Route path="/users" component={Users}>
                <Route path="sign-in" component={SignIn}/>
                <Route path="sign-up" component={SignUp}/>
                <Route path="confirm" component={Confirm}/>
                <Route path="unlock" component={Unlock}/>
                <Route path="forgot-password" component={ForgotPassword}/>
                <Route path="change-password" component={ChangePassword}/>
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

});

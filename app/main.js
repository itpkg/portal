"use strict";

require('bootstrap/dist/css/bootstrap.css');

import React from 'react';
import ReactDOM from 'react-dom';
import {Router, Route, IndexRoute} from 'react-router';

import i18next from 'i18next/lib';
import XHR from 'i18next-xhr-backend/lib';
import LanguageDetector from 'i18next-browser-languagedetector/lib';

import {NoMatch, Version} from './Widgets'
import {Application, Home, AboutUs, Notice} from  './Application'
import {Users, SignIn, SignUp, Confirm, Unlock, ForgotPassword, ResetPassword, Profile} from './Users'
import {Message} from "./Message"

function main(options) {
    i18next
        .use(XHR)
        .use(LanguageDetector)
        .init({
            fallbackLng: 'en-US',
            backend: {
                loadPath: options.assets_path + '/locales/{{lng}}.json'
            },
            detection: {
                order: ['querystring', 'localStorage', 'cookie', 'navigator'],
                lookupQuerystring: 'locale',
                lookupCookie: 'locale',
                lookupLocalStorage: 'locale',
                caches: ['localStorage', 'cookie']
            }
        }, (err, t)=> {
            ReactDOM.render(
                (<Router>

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

                        <Route path="*" component={NoMatch}/>
                    </Route>
                </Router>),
                document.getElementById('root')
            );

        });
}

module.exports = {
    Main: main
};
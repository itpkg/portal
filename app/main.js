require('./main.css');

import React from 'react'
import ReactDOM from 'react-dom'
import {Router} from 'react-router';
import { Provider } from 'react-redux'
import i18next from 'i18next/lib';
import XHR from 'i18next-xhr-backend/lib';
import LanguageDetector from 'i18next-browser-languagedetector/lib';

import routes from './components/routes'
import {History,Store} from './store'

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
                (<Provider store={Store}>
                    <Router history={History} routes={routes}/>
                </Provider>),
                document.getElementById('root')
            );
        });
}

export default main
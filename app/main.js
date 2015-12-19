"use strict";

require('bootstrap/dist/css/bootstrap.css');

import React from 'react';
import ReactDOM from 'react-dom';

import i18next from 'i18next/lib';
import XHR from 'i18next-xhr-backend/lib';
import LanguageDetector from 'i18next-browser-languagedetector/lib';

import {Routes} from './routes'

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
            ReactDOM.render(Routes, document.getElementById('root'));
        });
}

module.exports = {
    Main: main
};
"use strict";

import React from 'react'
import  ReactDOM  from 'react-dom'
import { Router, Route, IndexRoute} from 'react-router'

import {NoMatch, Version} from './Widgets'
import {Application, Home, AboutUs, Notice} from  './Application'


ReactDOM.render(
    (<Router>
        <Route path="/" component={Application}>
            <IndexRoute component={Home}/>
            <Route path="about_us" component={AboutUs}/>
            <Route path="notices/:id" component={Notice}/>
        </Route>
        <Route path="*" component={NoMatch}/>
    </Router>),
    document.getElementById('root')
);
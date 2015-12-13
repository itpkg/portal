"use strict";

import React from 'react'
import  ReactDOM  from 'react-dom'
import { Router, Route, IndexRoute} from 'react-router'

import {App, AboutUsPage, HomePage, NoMatchPage, Version} from './Widgets'


ReactDOM.render(
    (<Router>
        <Route path="/" component={App}>
            <IndexRoute component={HomePage}/>
            <Route path="about-us" component={AboutUsPage}/>
        </Route>
        <Route path="*" component={NoMatchPage}/>
    </Router>),
    document.getElementById('root')
);
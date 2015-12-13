import React from 'react'
import  ReactDOM  from 'react-dom'
import { Router, Route, Link } from 'react-router'

import {aboutUs, Home, noMatch, Version} from './Widgets'


ReactDOM.render(
    (<Router>
        <Route path="about-us" component={aboutUs}/>
        <Route path="home" component={Home}/>
        <Route path="*" component={noMatch}/>
    </Router>),
    document.getElementById('root')
);
